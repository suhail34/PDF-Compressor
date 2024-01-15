package databases

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once

var (
	MongoConnection *mongo.Client
	KafkaProducer   *kafka.Producer
)

func MongoConnect() (*mongo.Client, error) {
	if MongoConnection == nil {
		once.Do(
			func() {
				connectionString := fmt.Sprintf("mongodb://root:root@mongodb.default.svc.cluster.local:27017/compressor")
				clientOptions := options.Client().ApplyURI(connectionString)

				ctx := context.Background()
				MongoConnection, err := mongo.Connect(ctx, clientOptions)
				if err != nil {
					log.Fatal(err, " can't connect to mongodb altas")
				}
				err = MongoConnection.Ping(ctx, readpref.Primary())
				if err != nil {
					log.Fatal(err, " error pinging mongodb Atlas")
				}
				fmt.Println("MongoDB Connected")
			})
	}
	return MongoConnection, nil
}

func KafkaConnect() (*kafka.Producer, error) {
	if KafkaProducer == nil {
		once.Do(
			func() {
        p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka.default.svc.cluster.local:9092"})
				if err != nil {
					log.Print("Failed to Create producer : ", err)
				}
				KafkaProducer = p
			})
	}
	return KafkaProducer, nil
}