package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (root *Node) Find(val string) (string, bool) {
	if root == nil {
		return "", false
	}
	switch {
	case val == root.Data:
		return root.Data, true
	case val < root.Data:
		return root.Left.Find(val)
	default:
		return val, false
	}

}
func ForthController(c *gin.Context) {
	root := Node{"A", nil, nil, nil}
	root.Left = &Node{"-B", nil, nil, nil}
	root.Left.Left = &Node{"--D", nil, nil, nil}
	root.Left.Right = &Node{"--E", nil, nil, nil}
	root.Right = &Node{"-C", nil, nil, nil}
	root.Right.Left = &Node{"--F", nil, nil, nil}
	root.Right.Middle = &Node{"--G", nil, nil, nil}
	root.Right.Right = &Node{"--H", nil, nil, nil}
	fmt.Println("/**")
	// fmt.Printf("Pre Order Traversal of the given tree is: ")
	root.PreOrderTraversal()
	fmt.Println("**/")
	var is_found bool
	if c.Request.Method == "POST" {
		val := strings.ToUpper(c.PostForm("value"))
		_, is_found = root.Find(val)
		/*
			for _, item := range val {

			}
		*/

	}
	c.HTML(http.StatusOK, "forth.html", gin.H{
		"title": "Excercise",
		"res":   is_found,
	})
}
