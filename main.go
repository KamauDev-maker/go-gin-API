package main

import "github.com/gin-gonic/gin"

func main(){
	// create a default Gin router
	r := gin.Default()

	// Define a GET route
	r.GET("/", func(c *gin.Context) {
		//Return JSON response
		c.JSON(200, gin.H{
			"message": "Hello, World! Oscar",
		})
	})

	//Start the server on port 8080
	r.Run(":8080") //default is ":8080"
}