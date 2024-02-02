package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suhail34/pdf-compressor/producer-service/internal/middlewares"
	"github.com/suhail34/pdf-compressor/producer-service/internal/routes"
)

func NewServer() *gin.Engine {
	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.49.2:30002"}
	app.Use(cors.New(config))
	app.Use(middlewares.RateLimiterMiddleware)
	routes.SetupRoutes(app)
	return app
}
