package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/transcript", PostTranscript)
	router.Run()
}

func PostTranscript(c *gin.Context) {
	var transcript Transcript

	err := c.ShouldBindJSON(&transcript)

	if err != nil {
		panic(err)
	}

	metrics := GetMetrics(transcript)

	c.JSON(200, metrics)
}
