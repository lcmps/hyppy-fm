package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// func initializeRoutes() {

// 	// Handle the index route
// 	router.GET("/", showIndexPage)
// }

// Host any
func Host() {
	router := gin.Default()
	router.LoadHTMLGlob("./html/pages/*")
	router.Static("/assets", "./html/assets/")

	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	// initializeRoutes()

	router.Run()
}
