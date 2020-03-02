
package main // import "github.com/caffeine-driven-developers/modi-backend-go"
import (
	"fmt"
	"github.com/caffeine-driven-developers/modi-backend-go/database"
	"github.com/caffeine-driven-developers/modi-backend-go/search"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	fmt.Println("DB is connected!")
	defer db.Close()

	r := gin.Default() // create gin app
	r.Use(database.Inject(db)) // db to gin context
	r.GET("/search", search.OMDB)

	_ = r.Run(":3001") //listen on port :3001

}