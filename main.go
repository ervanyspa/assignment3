package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	assignment3()
}

func assignment3() {
	g := gin.Default()

	data := map[string]interface{}{}

	go func() {
		for {
			data["water"] = rand.Int31n(100)
			data["water_status"] = getStatus(data["water"].(int32))
			data["wind"] = rand.Int31n(100)
			data["wind_status"] = getStatus(data["wind"].(int32))
			data["fire"] = rand.Int31n(100)
			data["fire_status"] = getStatus(data["fire"].(int32))
			data["earth"] = rand.Int31n(100)
			data["earth_status"] = getStatus(data["earth"].(int32))
			time.Sleep(15 * time.Second)
		}
	}()

	g.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": data,
		})
	})
	g.Run(":3000")
}

func getStatus(value int32) string {
	if value < 6 {
		return "Aman"
	} else if value >= 7 && value <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
