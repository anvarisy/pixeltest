package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type NodeF struct {
	Key    int
	Data   string
	Grade  int
	Left   *NodeF
	Middle *NodeF
	Right  *NodeF
}

var room []string

func (n *NodeF) exists(value, grade int) bool {
	if n == nil {
		return false
	}
	room = append(room, n.Data)
	// fmt.Println(grade)
	if n.Key == value && n.Grade == grade {
		return true
	}
	// println("value ", n.Key)
	if value < n.Key {
		return n.Left.exists(value, grade)
	} else if value > n.Key {
		return n.Right.exists(value, grade)
	} else {
		return n.Middle.exists(value, grade)
	}
}

func ForthController(c *gin.Context) {
	room = nil
	root := NodeF{7, "A", 0, nil, nil, nil}
	root.Left = &NodeF{5, "B", 1, nil, nil, nil}
	root.Right = &NodeF{11, "C", 1, nil, nil, nil}
	root.Left.Left = &NodeF{1, "D", 2, nil, nil, nil}
	root.Left.Right = &NodeF{3, "E", 2, nil, nil, nil}
	root.Right.Left = &NodeF{9, "F", 2, nil, nil, nil}
	root.Right.Middle = &NodeF{11, "G", 2, nil, nil, nil}
	root.Right.Right = &NodeF{13, "H", 2, nil, nil, nil}
	var val string
	var status bool
	var temp_room []string
	if c.Request.Method == "POST" {
		post := strings.ToUpper(c.PostForm("value"))
		val = post
		var grade int = 0
		for _, item := range post {
			room = nil
			a := fmt.Sprintf("%c", item)
			key, _ := CharToKey(a)
			status = root.exists(key, grade)
			grade += 1
			if !status {
				break
			}
			temp_room = append(temp_room, a)

		}
		/*
			if len(room) > 3 {
				room = room[:len(room)-1]
			}
		*/
		fmt.Println(room)
		fmt.Println(temp_room)
		status = reflect.DeepEqual(room, temp_room)
	}
	c.HTML(http.StatusOK, "forth.html", gin.H{
		"title":  "Excercise",
		"input":  val,
		"res":    room,
		"status": status,
	})
}

func CharToKey(Ch string) (int, error) {
	switch Ch {
	case "A":
		return 7, nil
	case "B":
		return 5, nil
	case "C":
		return 11, nil
	case "D":
		return 1, nil
	case "E":
		return 3, nil
	case "F":
		return 9, nil
	case "G":
		return 11, nil
	case "H":
		return 13, nil
	default:
		return 0, errors.New("char not registered")
	}
}
