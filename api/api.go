package api

import (
	"github.com/caffeine-driven-developers/modi-backend-go/api/v1"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes ( r *gin.Engine){
		api := r.Group("/api")
		{
			v1.ApplyRoutes(api)
		}
}
