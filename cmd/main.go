package main

import (
	"fmt"
	"log"
	"net"
	"os"

	content "github.com/Soyaka/content/api/protogen/golang"
	"github.com/Soyaka/content/internal"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	Addr := NewLocalConfig().serverAddr
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	ContentService := internal.NewContentService()

	RegisterServerServices(server, &ContentService)

	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


type LocalConfig struct {
	serverAddr string
}

func NewLocalConfig() *LocalConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	return &LocalConfig{
		serverAddr: Addr,
	}
}

func RegisterServerServices(server *grpc.Server, service *internal.ContentService) {
	content.RegisterTutorialPageServiceServer(server, service)
	content.RegisterCourseSeriesServiceServer(server, service)
	content.RegisterCoursePageServiceServer(server, service)
	content.RegisterPodcastServiceServer(server, service)
	content.RegisterVideoSeriesServiceServer(server, service)
	content.RegisterTutorialServiceServer(server, service)
	content.RegisterLearningPathServiceServer(server, service)

}
