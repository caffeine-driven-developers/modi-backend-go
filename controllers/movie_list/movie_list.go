package movie_list

import (
	"github.com/caffeine-driven-developers/modi-backend-go/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type MovieList = models.MovieList


func Create(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	title := c.Query("title")
	movieIdList := c.Query("movieIdList")

	var movieList MovieList

	if title == "" {
		title = "untitle"
	}
	movieList.Title = title
	movieList.Movie_Id_List = movieIdList
	movieList.Created_At = time.Now()
	movieList.Updated_At = time.Now()
	db.Create(&movieList)
}
/*
func Read(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
}

func Update(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
}

func Remove(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
}
 */