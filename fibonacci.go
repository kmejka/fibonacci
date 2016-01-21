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

	if err != nil || n == 0 {
		c.JSON(http.StatusBadRequest, errorResponse{
			ErrorCode: 1000,
			Message: "Received inputParameter is not a correct limit value",
			InputParameters: nParam,
		})
		return
	}

	fibelems, err := fibSrv.CountNValues(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{
			ErrorCode: 1000,
			Message: "Unexpected error occurred during fibonacci calculation",
			InputParameters: nParam,
		})
		return
	}

	c.JSON(http.StatusOK, response{N: n, Elems: fibelems.ToString()})
}

type response struct {
	N int `json:"n"`
	Elems string `json:"elems"`
}

type errorResponse struct {
	ErrorCode int `json:"errorCode"`
	Message string `json:"message"`
	InputParameters string `json:"inputParameters"`
}
