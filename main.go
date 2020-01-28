
package main // import "github.com/caffeine-driven-developers/modi-backend-go"

import (
	"github.com/caffeine-driven-developers/modi-backend-go/search"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/search", search.OMDB)

	_ = r.Run(":3001") //listen on port :3001
}