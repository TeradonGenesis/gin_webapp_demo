package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set router as the default one by gin
	router = gin.Default()

	// process templates at the start so dont have to be loaded form the disk again, serves the HTMP pages faster
	router.LoadHTMLGlob("templates/*")

	//define the route for index page abd display index.html

	// router.GET("/", func(c *gin.Context) {
	// 	//call html method of the context to render a template
	// 	c.HTML(
	// 		//set the http status to 200 ok
	// 		http.StatusOK,
	// 		// use the template
	// 		"index.html",
	// 		// pass the data the page uses
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	router.GET("/", showIndexPage)

	//get single article
	router.GET("/article/view/:article_id", getArticle)

	//start serving application
	router.Run()
}
