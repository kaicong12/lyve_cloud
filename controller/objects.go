package controller

import (
	"net/http"
	"seagate-hackathon/db"
	"seagate-hackathon/models"

	"github.com/gin-gonic/gin"
)

type ObjectsResponse struct {
	ID            uint          `json:"id"`
	Name          string        `json:"name"`
	Status        db.StatusEnum `json:"status"`
	Progress      int           `json:"progress"`
	Total_Objects int           `json:"total_objects"`
	Objects       []db.Object   `json:"objects"`
}

func GetObjectsUnderAMigration(c *gin.Context) {
	var objects []db.Object
	var migration db.Migration
	var doneObjects int

	migration_pk := c.Params.ByName("id")
	err := models.GetMigrationObjects(migration_pk, &migration, &objects)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"Message": err.Error()})
	} else {
		// Objects with "DONE" status are counted to measure migration progress
		total_objects := len(objects)
		for _, obj := range objects {
			if obj.Status == db.Done {
				doneObjects++
			}
		}

		// initialize response struct
		r := ObjectsResponse{
			ID:            migration.ID,
			Status:        migration.Status,
			Progress:      doneObjects,
			Total_Objects: total_objects,
			Objects:       objects,
		}

		c.JSON(http.StatusOK, r)
	}

}
