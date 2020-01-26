
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/modi-backend-go/search"
)
func main() {
	r := gin.Default()
	r.GET("/search", search.OMDB)

	_ = r.Run(":3001") // listen and serve on 0.0.0.0:8080
}