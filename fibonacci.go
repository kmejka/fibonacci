package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var fibSrv FibService

func main() {
	fibSrv = &FibServiceImpl{}

	server := gin.Default()
	server.GET("/", info)
	server.GET("/fibonacci/:n", fibonacci)

	server.Run(":8080")
}

func info(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to fib service, call /fibonacci/:n")
}

func fibonacci(c *gin.Context) {

	nParam := c.Param("n")
	n, err := strconv.Atoi(nParam)
	if err != nil {
		//todo create error type
		c.JSON(http.StatusBadRequest, "'"+nParam+"' is not a correct limit value")
	}

	fibelems, err := fibSrv.CountNValues(n)
	if err != nil {
		//todo use error type
		c.Error(err)
	}

	c.String(http.StatusOK, fibelems.ToString())
}
