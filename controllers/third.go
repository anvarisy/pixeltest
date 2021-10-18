package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Node struct {
	Data   string
	Left   *Node
	Middle *Node
	Right  *Node
}

func (root *Node) PreOrderTraversal() {
	if root != nil {
		fmt.Println("* ", root.Data)
		root.Left.PreOrderTraversal()
		root.Middle.PreOrderTraversal()
		root.Right.PreOrderTraversal()
	}
}

func ThirdController(c *gin.Context) {
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
	c.HTML(http.StatusOK, "third.html", gin.H{
		"title": "Excercise",
	})
}
