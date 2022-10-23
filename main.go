package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-robot-api/game"
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
	X           int
	Y           int
	Orientation string
}

func main() {
	r := gin.Default()

	r.POST("/game", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")

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

	})

	r.GET("/game", func(c *gin.Context) {
		var queryStr Req
		a := c.Request.URL.Query()

		queryStr.Command = a.Get("command")
		queryStr.Deep, _ = strconv.Atoi(a.Get("deep"))
		queryStr.Width, _ = strconv.Atoi(a.Get("width"))
		queryStr.X, _ = strconv.Atoi(a.Get("x"))
		queryStr.Y, _ = strconv.Atoi(a.Get("y"))
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

	})

	r.Run()
}
