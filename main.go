package main

import (
	"github.com/gin-gonic/gin"
	metrics "github.com/nbrink91/GoPlayground/metrics"
)

func main() {
	router := gin.Default()
	router.POST("/transcript", PostTranscript)
	router.Run()
}

func PostTranscript(c *gin.Context) {
	var transcript metrics.Transcript

	err := c.ShouldBindJSON(&transcript)

	if err != nil {
		panic(err)
	}

	metrics := metrics.GetMetrics(transcript)

	c.JSON(200, metrics)
}
