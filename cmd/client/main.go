package main

import (
	"context"
	"log"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := content.NewCoursePageServiceClient(conn)

	res, err := client.CreateCoursePage(context.Background(), &content.CoursePage{
		Id:       uuid.New().String(),
		Title:    "test",
		Sections: sections,
		Question: questions,
	})
	if err != nil {
		log.Fatalf("could not create course page: %v", err)
	}
	log.Printf("CreateCoursePage: %s", res)
	resp, err := client.GetCoursePage(context.Background(), &content.ReqID{
		ID: "5592b166-71ad-490e-a76f-c5cd3d28db8c",
	})
	if err != nil {
		log.Fatalf("could not get course page: %v", err)
	}

	log.Printf("GetCoursePage: %s", resp)
}

var questions *content.Question = &content.Question{

	Content: "test",
	Options: []string{"test1", "test2"},
}

var sections []*content.Section = []*content.Section{

	{
		Title:       "test",
		Description: "test",
		Resources: []*content.Resource{
			{
				Id:          uuid.New().String(),
				Name:        "test",
				Description: "test",
				Url:         "test",
			},
		},
		Paragraph: "test",
	},

	{
		Title:       "test",
		Description: "test",
		Resources: []*content.Resource{
			{
				Id:          uuid.New().String(),
				Name:        "test",
				Description: "test",
				Url:         "test",
			},
		},
		Paragraph: "test",
	},
}
