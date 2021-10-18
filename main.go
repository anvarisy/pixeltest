package main

import (
	"net/http"
	"pixeltest/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./statics")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Excercise",
		})
	})
	r.GET("/first", controllers.FirstController)
	r.POST("/post-first", controllers.FirstController)
	r.GET("/second", controllers.SecondController)
	r.POST("/post-second", controllers.SecondController)
	r.GET("/third", controllers.ThirdController)
	r.GET("/forth", controllers.ForthController)
	r.POST("/post-forth", controllers.ForthController)
	r.Run()
}
