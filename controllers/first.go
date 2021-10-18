package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FirstController(c *gin.Context) {
	var res int
	var err_ string

	fmt.Println(c.Request.Method)
	if c.Request.Method == "POST" {
		val, err := strconv.Atoi(c.PostForm("value"))
		if err != nil {
			res = 0
			err_ = "Only can procced number"
			c.HTML(http.StatusOK, "first.html", gin.H{
				"title": "Excercise",
				"res":   res,
				"err":   err_,
			})
			return
		}
		res = val*(val+1)/2 + 1

	}
	c.HTML(http.StatusOK, "first.html", gin.H{
		"title": "Excercise",
		"res":   res,
	})
}
