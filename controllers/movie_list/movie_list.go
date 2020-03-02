package movie_list

import (
	"github.com/caffeine-driven-developers/modi-backend-go/database/models"
	"github.com/caffeine-driven-developers/modi-backend-go/lib"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"

	//"strconv"
	"time"
)
// MovieList type alias
type MovieList = models.MovieList
// JSON type alias
type JSON = lib.JSON
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

func Read(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	id,_ := strconv.Atoi(c.Param("id"))
	var movieList MovieList

	if err := db.Where("movie_list_id = ?", id).Find(&movieList); err.Error !=nil{
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, movieList.Seriallize())
}
/*
func Update(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
}
*/
func Remove(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	id,_ := strconv.Atoi(c.Param("id"))

	var movieList MovieList

	if err := db.Where("movie_list_id = ?", id).Find(&movieList); err.Error !=nil{
		c.AbortWithStatus(404)
		return
	}
	// authorization 후 creator 와 지금 유저 확인 필

	db.Delete(&movieList)
	c.Status(204)
}
