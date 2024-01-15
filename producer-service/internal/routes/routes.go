package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhail34/pdf-compressor/producer-service/internal/handlers"
)

func SetupRoutes(app *gin.Engine) {
  app.GET("/api/files", handlers.UploadFileHandler)
}
