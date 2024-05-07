package test

import (
	"context"
	"testing"

	content "github.com/Soyaka/content/api/protogen/golang"
	"google.golang.org/grpc"
)

func TestCreateTutorial(t *testing.T) {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	mainClient := NewClientRegistry(conn)

	res, err := mainClient.Tutorial.CreateTutorial(context.Background(), createMockTutorial())

	if err != nil {
		t.Fatalf("could not create tutorial: %v", err)
	}

	if !res.OK {
		t.Fatalf("tutorial not created ")
	}
}

func TestGetTutorial(t *testing.T) {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	mainClient := NewClientRegistry(conn)

	res, err := mainClient.Tutorial.GetTutorial(context.Background(), &content.ReqID{ID: "tutorial_id_1"})

	if err != nil {
		t.Fatalf("could not get tutorial: %v", err)
	}

	if res == nil {
		t.Fatalf("tutorial not found ")
	}
}

func TestCreateTutorialPage(t *testing.T) {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	nainClient := NewClientRegistry(conn)

	res, err := nainClient.TutorialPage.CreateTutorialPage(context.Background(), createMockTutorialPage())

	if err != nil {
		t.Fatalf("could not create tutorial page: %v", err)
	}

	if !res.OK {
		t.Fatalf("tutorial page not created ")
	}
}

func TestGetTutorialPage(t *testing.T) {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	mainClient := NewClientRegistry(conn)

	res, err := mainClient.TutorialPage.GetTutorialPage(context.Background(), &content.ReqID{ID: "tutorial_page_id_1"})

	if err != nil {
		t.Fatalf("could not get tutorial page: %v", err)
	}

	if res == nil {
		t.Fatalf("tutorial page not found ")
	}
}
func TestCreateCoursePage(t *testing.T) {

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	mainClient := NewClientRegistry(conn)

	res, err := mainClient.CoursePage.CreateCoursePage(context.Background(), createMockCoursePage())
	if err!=nil{
		t.Fatal(err)
	}
	if res == nil  || !res.OK {
		t.Fatalf("course page not created ")
	} 
}

func TestGetCoursePage(t *testing.T) {	

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	mainClient := NewClientRegistry(conn)

	res, err := mainClient.CoursePage.GetCoursePage(context.Background(), &content.ReqID{ID: "course_page_id_1"})

	if err != nil {
		t.Fatalf("could not get course page: %v", err)
	}

	if res == nil {
		t.Fatalf("course page not found ")
	}

}





type MockClient struct {
	CoursePage   content.CoursePageServiceClient
	Course       content.CourseSeriesServiceClient
	Tutorial     content.TutorialServiceClient
	TutorialPage content.TutorialPageServiceClient
	Podcast      content.PodcastServiceClient
	LearningPath content.LearningPathServiceClient
	VidSerries   content.VideoSeriesServiceClient
}

func NewClientRegistry(service *grpc.ClientConn) *MockClient {
	var m MockClient
	m.CoursePage = content.NewCoursePageServiceClient(service)
	m.Course = content.NewCourseSeriesServiceClient(service)
	m.Tutorial = content.NewTutorialServiceClient(service)
	m.TutorialPage = content.NewTutorialPageServiceClient(service)
	m.Podcast = content.NewPodcastServiceClient(service)
	m.LearningPath = content.NewLearningPathServiceClient(service)
	m.VidSerries = content.NewVideoSeriesServiceClient(service)
	return &m
}

func createMockTutorial() *content.Tutorial {
	return &content.Tutorial{
		Id: "tutorial_id_1",
		ContentInfo: &content.ContentInfo{
			Title:       "Mock Tutorial Title",
			Description: "Mock Tutorial Description",
			ContentType: "Tutorial",
		},
		Duration: "Mock Tutorial Duration",
		Pages: []*content.TutorialPage{
			createMockTutorialPage(), 
		},
	}
}

func createMockTutorialPage() *content.TutorialPage {
	return &content.TutorialPage{
		Id:          "tutorial_page_id_1",
		Title:       "Mock Tutorial Page Title",
		Description: "Mock Tutorial Page Description",
		Sections: []*content.Section{
			{
				Title:       "Mock Section Title 1",
				Description: "Mock Section Description 1",
			},
			{
				Title:       "Mock Section Title 2",
				Description: "Mock Section Description 2",
			},
		},
	}
}

func createMockCoursePage() *content.CoursePage {
	return &content.CoursePage{
		Id:          "course_page_id_1",
		Title:       "Mock Course Page Title",
		Sections: []*content.Section{
			{
				Title:       "Mock Section Title 1",
				Description: "Mock Section Description 1",
				// Add more fields as needed
			},
			{
				Title:       "Mock Section Title 2",
				Description: "Mock Section Description 2",
				// Add more fields as needed
			},
			// Add more sections if needed
		},
		Question: &content.Question{
			Content: "Mock Question Content",
			Options: []string{"Option 1", "Option 2"},
			// Add more fields as needed
		},
	}
}
