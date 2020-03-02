package v1

import (
	"github.com/caffeine-driven-developers/modi-backend-go/api/v1/movie_list"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		// make api each compo
		movie_list.ApplyRoutes(v1)
	}
}