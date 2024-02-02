package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhail34/pdf-compressor/producer-service/internal/databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func DownloadHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fileId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("Not a proper file ID: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	client, err := databases.MongoConnect()
	if err != nil {
		log.Print("Failed to connect to MongoDB : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	db := client.Database("myFiles")
	fs, err := gridfs.NewBucket(db)
	if err != nil {
		log.Print("Gridfs Bucket error : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	downloadStream, err := fs.OpenDownloadStream(primitive.ObjectID(fileId))
	if err != nil {
		log.Print("Error while downloading : ", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	defer downloadStream.Close()
	var fileName string
	fileInfo := bson.M{}
	err = fs.GetFilesCollection().FindOne(ctx, bson.M{"_id": fileId}).Decode(&fileInfo)
	if err != nil {
		log.Print("Failed to retrieve file Info : ", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve file info " + err.Error()})
		return
	}
	fileNameRaw, ok := fileInfo["filename"].(bson.Raw)
	if !ok {
		fileName = "compressed.pdf"
	} else {
		fileName = fileNameRaw.String()
	}
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/pdf")
	_, err = io.Copy(ctx.Writer, downloadStream)
	if err != nil {
		log.Print("Error while downloading : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download file"})
		return
	}
	ctx.Status(http.StatusOK)
}
