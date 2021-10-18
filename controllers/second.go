package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SecondController(c *gin.Context) {
	// var res []int
	var err_ string
	var final string
	fmt.Println(c.Request.Method)
	if c.Request.Method == "POST" {
		val, err := strconv.Atoi(c.PostForm("value"))
		if err != nil {
			err_ = "Only can procced number"
			c.HTML(http.StatusOK, "second.html", gin.H{
				"title": "Excercise",
				"res":   final,
				"err":   err_,
			})
			return
		}
		for i := 0; i < val; i++ {
			num := i*(i+1)/2 + 1
			final += strconv.Itoa(num)
			if i < (val - 1) {
				final += "-"
			}
			// res = append(res[:i], num)
		}
		log.Println(final)

	}
	c.HTML(http.StatusOK, "second.html", gin.H{
		"title": "Excercise",
		// "res":   res,
		"final": final,
	})
}
