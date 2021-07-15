// Golang REST API program
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MathMessage struct {
	Operation string
	Result    int
}

type Message struct {
	Method  string
	Message string
}

func SampleGet(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Method:  http.MethodGet,
		Message: "GetMethod called",
	})
}

func SamplePost(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Method:  http.MethodPost,
		Message: "PostMethod called",
	})
}

func SamplePut(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Method:  http.MethodPut,
		Message: "PutMethod called",
	})
}

func SampleDelete(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Method:  http.MethodDelete,
		Message: "DeleteMethod called",
	})
}

func AddMethod(c *gin.Context) {
	x, _ := strconv.Atoi(c.Param("x"))
	y, _ := strconv.Atoi(c.Param("y"))

	c.JSON(http.StatusOK, MathMessage{
		Operation: fmt.Sprintf("%d + %d", x, y),
		Result:    x + y,
	})
}

func SubtractMethod(c *gin.Context) {
	x, _ := strconv.Atoi(c.Param("x"))
	y, _ := strconv.Atoi(c.Param("y"))

	c.JSON(http.StatusOK, MathMessage{
		Operation: fmt.Sprintf("%d - %d", x, y),
		Result:    x - y,
	})
}

func main() {
	router := gin.Default()

	// Sample routes
	router.GET("/", SampleGet)
	router.POST("/", SamplePost)
	router.PUT("/", SamplePut)
	router.DELETE("/", SampleDelete)

	// Math operations
	router.GET("/add/:x/:y", AddMethod)
	router.GET("/subtract/:x/:y", SubtractMethod)

	listenPort := "4000"
	router.Run(":" + listenPort)
}
