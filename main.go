package main // import "github.com/caffeine-driven-developers/modi-backend-go"
import (
	"fmt"
	"github.com/caffeine-driven-developers/modi-backend-go/api"
	"github.com/caffeine-driven-developers/modi-backend-go/database"
	"github.com/caffeine-driven-developers/modi-backend-go/search"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	db, err := gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	} else {
		fmt.Println("DB is connected!")
	}
	defer db.Close()

	r := gin.Default() // create gin app

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(database.Inject(db)) // db to gin context
	r.GET("/search", search.OMDB)
	api.ApplyRoutes(r)

	_ = r.Run(":3001") //listen on port :3001

}
