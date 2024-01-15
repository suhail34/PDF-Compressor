package handlers

import (
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/suhail34/pdf-compressor/producer-service/internal/databases"
)

func UploadFileHandler(ctx *gin.Context) {
	producer, err := databases.KafkaConnect()
	if err != nil {
		log.Print("Failed to make producer : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	topic := "my-topic"
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello World"),
	}
	err = producer.Produce(message, nil)
	if err != nil {
		log.Print("Failed to produce message : ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
  _, err = databases.MongoConnect()
  if err != nil {
    log.Print("Failed to connect to mongodb: ", err)
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  }
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfull"})
}
