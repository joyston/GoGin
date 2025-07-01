package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

// Data in the memory //In real world we use
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) { //We Context
	c.IndentedJSON(http.StatusOK, albums) // IndentedJSON Serializes the struct to Json format
} // http.StatusOK returns success Http code 200

func addAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)            //We append the new data
	c.IndentedJSON(http.StatusCreated, newAlbum) //Serialize album struct to JSON & return Created Http code 201
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id") //Use Context.Param to retrieve the id path parameter from the URL

	for _, a := range albums { //looping throught the data to find the id
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"}) // return 404 not found http code
	}
}

func main() {
	router := gin.Default() //Initialize a Gin router using Default
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbums)

	router.Run("localhost:8080")
}
