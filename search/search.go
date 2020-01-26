package search

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func OMDB(c *gin.Context){
	query := c.Request.URL.Query()
	if query["i"] == nil && query["s"] == nil {
		c.String(http.StatusBadRequest, "both i and t are missing")
	} else if query["i"] != nil && query["s"] != nil {
		c.String(http.StatusConflict, "please specify either i(id) or t(title)")
	} else if query["i"] != nil {
		res := searchOmdbById(query["i"])
		c.String(http.StatusOK,string(res[:]))
	} else if query["s"] != nil {
		res := searchOmdbByTitle(query["s"])
		c.String(http.StatusOK,string(res[:]))
	}
}

func searchOmdbById(id []string) []byte{
	apiKey := goDotEnvVariable("OMDB_API_KEY")
	resp, _ := http.Get(`https://www.omdbapi.com/?i=`+id[0]+`&apikey=`+apiKey)
	defer resp.Body.Close()
	body ,_ := ioutil.ReadAll(resp.Body)
	return body
}
func searchOmdbByTitle(title []string) []byte{
	apiKey := goDotEnvVariable("OMDB_API_KEY")
	resp, _ := http.Get(`https://www.omdbapi.com/?s=`+title[0]+`&apikey=`+apiKey)
	defer resp.Body.Close()
	body ,_ := ioutil.ReadAll(resp.Body)
	return body
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
