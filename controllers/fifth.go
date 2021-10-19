package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
var x = make(map[string][]string)

func (n *NodeF) exists_II(value, grade int) bool {

	if n == nil {
		return false
	}
	x["Key"] = append(x["key"], strconv.Itoa(n.Key))
	x["Data"] = append(x["key"], n.Data)
	x["Grade"] = append(x["key"], strconv.Itoa(n.Grade))
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
*/
func FifthController(c *gin.Context) {
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
	var result string
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
		if status {
			if len(room) >= 3 {
				result = " Tidak ada Child"
			} else if len(room) == 2 {
				if room[1] == "B" {
					result = "[D E]"
				} else {
					result = "[F G H]"
				}
			} else {
				result = "[B C]"
			}
		} else {
			result = "Path tidak valid"
		}
	}
	c.HTML(http.StatusOK, "fifth.html", gin.H{
		"title": "Excercise",
		"input": val,
		"res":   result,
	})
}
