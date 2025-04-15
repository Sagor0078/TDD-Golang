package main

import (
    "math"
    "net/http"

    "github.com/gin-gonic/gin"
)

func cpuIntensiveTask() {
    // Perform a CPU-intensive calculation
    for i := 0; i < 1000000; i++ {
        _ = math.Sqrt(float64(i))
    }
}

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        cpuIntensiveTask() // Add CPU load
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    r.Run() // listen and serve on 0.0.0.0:8080 (default)
}