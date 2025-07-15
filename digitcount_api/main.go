package main

import (
	"digitcount/digitcount"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type findInSeriesRequest struct {
	Start     int `json:"series_start"`
	End       int `json:"series_end"`
	Increase  int `json:"series_increment"`
	Type      int `json:"series_type"`
	SearchKey int `json:"specified_digit"`
}

func main() {
	router := gin.Default()
	router.POST("/findInSeries", findInSeries)
	router.GET("/health", isHealthy)

	if err := router.Run(":" + getPort()); err != nil {
		panic(err)
	}
}

func getPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

func isHealthy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func findInSeries(context *gin.Context) {
	var request findInSeriesRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	count, err := digitcount.CountDigitOccurrencesInSeries(request.Start, request.End, request.Increase, request.Type, request.SearchKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err, "count": count})
		return
	}

	context.JSON(http.StatusOK, gin.H{"count": count})

}
