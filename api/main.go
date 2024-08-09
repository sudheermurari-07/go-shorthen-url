package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sudheermurari-07/shorten-url/routes"
)

func setupRoutes(app *gin.Engine) {
	app.GET("/:url", routes.ResolveURL)
	app.POST("/api/url", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := gin.Default()

	setupRoutes(app)

	log.Fatal(app.Run(os.Getenv("APP_PORT")))
}
