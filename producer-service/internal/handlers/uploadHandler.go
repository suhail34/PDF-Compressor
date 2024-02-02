package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/suhail34/pdf-compressor/producer-service/internal/databases"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UploadFileHandler(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Print("MultipartForm Error : ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}
	files := form.File["files"]
	log.Print(files)
	if len(ctx.Request.MultipartForm.File) == 0 {
		log.Print("No files Uploaded : ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No files Uploaded"})
    return
	}
	producer, err := databases.KafkaConnect()
	if err != nil {
		log.Print("Failed to make producer : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
	topic := "my-topic"
	client, err := databases.MongoConnect()
	if err != nil {
		log.Print("Failed to connect to mongodb: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
	db := client.Database("myFiles")
	fs, err := gridfs.NewBucket(db)
	if err != nil {
		log.Print("Gridfs Bucket Error : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
  fileIds := make(map[primitive.ObjectID]string)
	for _, file := range files {
		fileHandle, err := file.Open()
		if err != nil {
			log.Print("Cannot open file : ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
		}
		defer fileHandle.Close()
		opts := options.GridFSUpload().SetMetadata(map[string]string{"originalFileName": file.Filename})
		fileID := primitive.NewObjectID()
    fileIds[fileID] = file.Filename
		uploadStream, err := fs.OpenUploadStreamWithID(fileID, file.Filename, opts)
		if err != nil {
			log.Print("Cannot open uploadStream : ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
		}
		defer uploadStream.Close()
		_, err = io.Copy(uploadStream, fileHandle)
		if err != nil {
			log.Print("Cannot Copy to destination : ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
		}
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(file.Filename),
			Value:          []byte(fileID.Hex()),
		}
		err = producer.Produce(msg, nil)
		if err != nil {
			log.Print("Failed to produce message : ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"resp": fileIds})
}
