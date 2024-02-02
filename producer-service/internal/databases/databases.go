package databases

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoOnce       sync.Once
	kafkaOnce       sync.Once
	mongoConnection *mongo.Client
	kafkaProducer   *kafka.Producer
)

func MongoConnect() (*mongo.Client, error) {
  
	if mongoConnection == nil {
		mongoOnce.Do(
			func() {
        usernameBase64 := os.Getenv("MONGO_USERNAME")
        passwordBase64 := os.Getenv("MONGO_PASSWORD")
        usernameByte, _ := base64.StdEncoding.DecodeString(usernameBase64)
        passwordByte, _ := base64.StdEncoding.DecodeString(passwordBase64)
        username := string(usernameByte)
        password := string(passwordByte)
        connectionString := fmt.Sprintf("mongodb://%v:%v@mongodb.default.svc.cluster.local:27017/myFiles?authSource=admin&authMechanism=SCRAM-SHA-1", username, password)
				clientOptions := options.Client().ApplyURI(connectionString)

				ctx := context.Background()
				mongoConnect, err := mongo.Connect(ctx, clientOptions)
				if err != nil {
					log.Fatal(err, " can't connect to mongodb altas")
				}
				err = mongoConnect.Ping(ctx, readpref.Primary())
				if err != nil {
					log.Fatal(err, " error pinging mongodb Atlas")
				}
				mongoConnection = mongoConnect
				log.Print("MongoDB Connected")
			})
	}
	return mongoConnection, nil
}

func KafkaConnect() (*kafka.Producer, error) {
	if kafkaProducer == nil {
		kafkaOnce.Do(
			func() {
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka-release.default.svc.cluster.local:9092"})
				if err != nil {
					log.Print("Failed to Create producer : ", err)
				}
				kafkaProducer = p
				log.Print("Kafka Connected")
			})
	}
	return kafkaProducer, nil
}
