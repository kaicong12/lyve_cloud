package routes

import (
	"../Controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("migrations", Controller.GetMigrations)
	}
	return r
}