package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the lsit of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
    router := gin.Default()

    // Configure CORS middleware to allow reqs from vite app
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://127.0.0.1:5173"}
    // this should allow my vite app to make requests
    router.Use(cors.New(config))

    // When hitting localhost:8085/albums we get the return value of getAlbums function
    router.GET("/albums", getAlbums)
    
    // Hosts the server on port 8085
    router.Run("localhost:8085")
}
