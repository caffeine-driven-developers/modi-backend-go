package movie_list

import (
	"github.com/caffeine-driven-developers/modi-backend-go/controllers/movie_list"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	movieList := r.Group("/movie-list")
	{
		movieList.POST("/",movie_list.Create)
		//movieList.GET("/:id", movie_list.Read)
		//movieList.DELETE("/:id", movie_list.Remove)
		//movieList.PATCH("/:id", movie_list.Update)

	}
}
