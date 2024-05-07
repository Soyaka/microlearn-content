package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.vallahaye.net/protobson"
)

type Service struct {
	Client *mongo.Client
}

type DataConfig struct {
	Username string
	Password string
	AppName  string
}

func GetConfig() *DataConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_ROOT_PASSWORD")
	APPNAME := os.Getenv("DB_APPNAME")
	return &DataConfig{Username: USERNAME, Password: PASSWORD, AppName: APPNAME}
}

func New( dt *DataConfig) *Service {

	encodedUsername := url.QueryEscape(dt.Username)
	encodedPassword := url.QueryEscape(dt.Password)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	URI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.v8lyn78.mongodb.net/?retryWrites=true&w=majority&appName=%s", encodedUsername, encodedPassword, dt.AppName)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	opts := options.Client().SetRegistry(protobson.DefaultRegistry).ApplyURI(URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return &Service{
		Client: client,
	}
}
