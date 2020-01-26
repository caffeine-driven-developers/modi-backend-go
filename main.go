
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junpyo/modi-backend-go/search"
)
func main() {
	r := gin.Default()
	r.GET("/search", search.OMDB)

	_ = r.Run(":3001") //listen on port :3001
}