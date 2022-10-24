package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-robot-api/game"
	"github.com/tarcea/go-robot-api/middlewares"
)

type Req struct {
	Width       int    `json:"width,string,omitempty"`
	Deep        int    `json:"deep,string,omitempty"`
	Orientation string `json:"orientation,omitempty"`
	X           int    `json:"x,string,omitempty"`
	Y           int    `json:"y,string,omitempty"`
	Command     string `json:"command,omitempty"`
}

type Res struct {
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Orientation string `json:"orientation"`
}

func PostHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var reqBody Req

	if err := c.BindJSON(&reqBody); err != nil {
		fmt.Println("Error binding body")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	g := game.NewGame(reqBody.Deep, reqBody.Width, reqBody.X, reqBody.Y)
	g.Orientation = reqBody.Orientation
	g.Command = reqBody.Command
	game.RunCommand(g)
	var res Res
	res.X = g.PositionX
	res.Y = g.PositionY
	res.Orientation = g.Orientation

	c.JSON(http.StatusOK, gin.H{
		"Report": res,
	})
}

func GetHandler(c *gin.Context) {
	var queryStr Req
	a := c.Request.URL.Query()

	if a.Get("command") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'command' not present or wrong query provided",
		})
		return
	}
	queryStr.Command = a.Get("command")

	if a.Get("deep") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'deep' not present or wrong query provided",
		})
		return
	}
	queryStr.Deep, _ = strconv.Atoi(a.Get("deep"))

	if a.Get("width") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'width' not present or wrong query provided",
		})
		return
	}
	queryStr.Width, _ = strconv.Atoi(a.Get("width"))

	if a.Get("x") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'x' not present or wrong query provided",
		})
		return
	}
	queryStr.X, _ = strconv.Atoi(a.Get("x"))

	if a.Get("y") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'y' not present or wrong query provided",
		})
		return
	}
	queryStr.Y, _ = strconv.Atoi(a.Get("y"))

	if a.Get("orientation") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "value for 'orientation' not present or wrong query provided",
		})
		return
	}
	queryStr.Orientation = a.Get("orientation")

	g := game.NewGame(queryStr.Deep, queryStr.Width, queryStr.X, queryStr.Y)
	g.Orientation = queryStr.Orientation
	g.Command = queryStr.Command
	game.RunCommand(g)

	var res Res
	res.X = g.PositionX
	res.Y = g.PositionY
	res.Orientation = g.Orientation

	c.JSON(http.StatusOK, gin.H{
		"Report": res,
	})

}

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.POST("/game", PostHandler)
	r.GET("/game", GetHandler)

	r.Run()
}
