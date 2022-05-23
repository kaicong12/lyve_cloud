package controller

import (
	"net/http"
	"seagate-hackathon/db"
	"seagate-hackathon/models"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MigrationResponse struct {
	ID               uint          `json:"id"`
	Name             string        `json:"name"`
	Status           db.StatusEnum `json:"status"`
	Objects_Migrated int           `json:"objects_migrated"`
	Failed_Objects   int           `json:"failed_objects"`
	Created_Date     time.Time     `json:"created_date"`
}

func GetMigrations(c *gin.Context) {
	var migrations []db.Migration

	// querystring handling
	order_by := c.DefaultQuery("ordering", "name")
	search := c.DefaultQuery("search", "")

	page_size, ps_err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	page, p_err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if ps_err != nil || p_err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"Message": "Invalid page or page size"})
	}

	if order_by == "-name" {
		order_by = "name desc"
	}

	err := models.GetAllMigrations(&migrations, order_by, search, page, page_size)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var responses []MigrationResponse
		for _, m := range migrations {
			r, m_err := GetResponse(m)

			if m_err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
			}

			responses = append(responses, r)
		}

		// other orderby querystring
		sortResponse(responses, order_by)

		c.JSON(http.StatusOK, responses)
	}
}

func GetAMigration(c *gin.Context) {
	pk := c.Params.ByName("id")
	var migration db.Migration

	err := models.GetAMigration(&migration, pk)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"Message": err.Error()})
	} else {
		response, m_err := GetResponse(migration)

		if m_err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		c.JSON(http.StatusOK, response)
	}
}

func CreateAMigration(c *gin.Context) {
	var newMigration db.Migration
	c.ShouldBindJSON(&newMigration)

	err := models.CreateAMigration(&newMigration)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}

func GetResponse(m db.Migration) (MigrationResponse, error) {
	// query all objects under this Migration
	status_count, err := models.GroupByStatus(m.ID)

	if err != nil {
		return MigrationResponse{}, err
	}

	migrated, failed := 0, 0
	for _, s := range status_count {
		if s.Status == db.Failed {
			failed += s.Count
		} else {
			migrated += s.Count
		}

	}
	response := MigrationResponse{
		ID:               m.ID,
		Name:             m.Name,
		Objects_Migrated: migrated,
		Failed_Objects:   failed,
		Status:           m.Status,
		Created_Date:     m.CreatedAt,
	}

	return response, nil
}

func sortResponse(responses []MigrationResponse, order_by string) {
	if order_by == "failed" {
		sort.Slice(responses, func(i, j int) bool { return responses[i].Failed_Objects < responses[j].Failed_Objects })
	} else if order_by == "-failed" {
		sort.Slice(responses, func(i, j int) bool { return responses[i].Failed_Objects > responses[j].Failed_Objects })
	} else if order_by == "migrated" {
		sort.Slice(responses, func(i, j int) bool { return responses[i].Objects_Migrated < responses[j].Objects_Migrated })
	} else if order_by == "-migrated" {
		sort.Slice(responses, func(i, j int) bool { return responses[i].Objects_Migrated > responses[j].Objects_Migrated })
	}
}
