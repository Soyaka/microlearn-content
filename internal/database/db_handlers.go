package database

/* ---------------------Database Handlers--------------------------------------- */

import (
	"context"
	"errors"
	"time"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	TimeOut         = 5 * time.Second
	DATA            = "content"
	TutorialCol     = "tutorials"
	TutorialPageCol = "tutorial_pages"
	CourseCol       = "courses"
	CoursePageCol   = "course_pages"
	VideoSeriesCol  = "videosSeries"
	PodcastCol      = "podcasts"
	LearningCol     = "learningPath"
)

/* CoursePage Database Handlers */

func (c *Service) CreateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("course page cannot be nil")
	}
	result, err := c.Client.Database(DATA).Collection(CoursePageCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(CoursePageCol).UpdateOne(context, bson.M{"_id": in.GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetId()}, nil
}

func (c *Service) DeleteCoursePage(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(CoursePageCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) GetCoursePage(ctx context.Context, in *content.ReqID) (*content.CoursePage, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result := c.Client.Database(DATA).Collection(CoursePageCol).FindOne(context, bson.M{"_id": in.ID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var coursePage *content.CoursePage
	raw, err := result.Raw()
	if err != nil {
		return &content.CoursePage{}, err
	}
	err = bson.Unmarshal(raw, &coursePage)
	if err != nil {
		return &content.CoursePage{}, err
	}
	return coursePage, nil
}

// Tutorial pages handlers
func (c *Service) CreateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("tutorial page cannot be nil")
	}
	result, err := c.Client.Database(DATA).Collection(TutorialPageCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) GetTutorialPage(ctx context.Context, in *content.ReqID) (*content.TutorialPage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result := c.Client.Database(DATA).Collection(TutorialPageCol).FindOne(ctox, bson.M{"_id": in.ID})
	if result.Err() != nil {
		return &content.TutorialPage{}, result.Err()
	}

	var TutorialPage *content.TutorialPage
	raw, err := result.Raw()
	if err != nil {
		return &content.TutorialPage{}, err
	}
	err = bson.Unmarshal(raw, &TutorialPage)
	return TutorialPage, err
}

func (c *Service) DeleteTutorialPage(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(TutorialPageCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) UpdateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(TutorialPageCol).UpdateOne(context, bson.M{"_id": in.GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetId()}, nil
}

/* CourseSeries Database Handlers */

func (c *Service) CreateCourseSeries(ctx context.Context, in *content.Course) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(CourseCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdateCourseSeries(ctx context.Context, in *content.Course) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(CourseCol).UpdateOne(context, bson.M{"_id": in.GetContentInfo().GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetContentInfo().GetId()}, nil

}

func (c *Service) DeleteCourseSeries(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(CourseCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) GetCourseSeries(ctx context.Context, in *content.ReqID) (*content.Course, error) {
	var res *content.Course
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result := c.Client.Database(DATA).Collection(CourseCol).FindOne(ctox, bson.M{"_id": in.GetID()})
	if result.Err() != nil {
		return nil, result.Err()
	}
	resp, err := result.Raw()
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(resp, &res)
	if err != nil || res == nil {
		return nil, err
	}
	return res, nil
}

/* VideoSeries Database Handlers */

func (c *Service) CreateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(VideoSeriesCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(VideoSeriesCol).UpdateOne(context, bson.M{"_id": in.GetContentInfo().GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetContentInfo().GetId()}, nil
}

func (c *Service) DeleteVideoSeries(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(VideoSeriesCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) GetVideoSeries(ctx context.Context, in *content.ReqID) (*content.VideoSeries, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result := c.Client.Database(DATA).Collection(VideoSeriesCol).FindOne(context, bson.M{"_id": in.ID})
	if result.Err() != nil {
		return &content.VideoSeries{}, result.Err()
	}
	var vidSeries *content.VideoSeries
	bsonRes, err := result.Raw()
	if err != nil {
		return &content.VideoSeries{}, err
	}
	err = bson.Unmarshal(bsonRes, &vidSeries)
	if err != nil || vidSeries == nil {
		return &content.VideoSeries{}, err
	}
	return vidSeries, nil
}

/* Event Database Handlers */

/* Exam Database Handlers */

/* Learning Path Database Handlers */

func (c *Service) CreateLearningPath(ctx context.Context, in *content.LearningPath) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(LearningCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdateLearningPath(ctx context.Context, in *content.LearningPath) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(LearningCol).UpdateOne(context, bson.M{"_id": in.GetContentInfo().GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetContentInfo().GetId()}, nil
}

func (c *Service) DeleteLearningPath(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(LearningCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) GetLearningPath(ctx context.Context, in *content.ReqID) (*content.LearningPath, error) {

	var res *content.LearningPath
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result := c.Client.Database(DATA).Collection(LearningCol).FindOne(context, bson.M{"_id": in.ID})

	if result.Err() != nil {
		return nil, result.Err()

	}
	resp, err := result.Raw()

	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(resp, &res)

	if err != nil || res == nil {
		return nil, err
	}

	return res, nil
}

/* Podcast Database Handlers */

func (c *Service) CreatePodcast(ctx context.Context, in *content.Podcast) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(PodcastCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdatePodcast(ctx context.Context, in *content.Podcast) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(PodcastCol).UpdateOne(context, bson.M{"_id": in.GetContentInfo().GetId()}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetContentInfo().GetId()}, nil
}

func (c *Service) DeletePodcast(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(PodcastCol).DeleteOne(context, bson.M{"_id": in.GetID()})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.GetID()}, nil
}

func (c *Service) GetPodcast(ctx context.Context, in *content.ReqID) (*content.Podcast, error) {

	var res *content.Podcast
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result := c.Client.Database(DATA).Collection(PodcastCol).FindOne(context, bson.M{"_id": in.ID})

	if result.Err() != nil {
		return nil, result.Err()
	}

	resp, err := result.Raw()

	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(resp, &res)

	if err != nil || res == nil {
		return nil, err
	}

	return res, nil
}

/* Tutorials*/

func (c *Service) CreateTutorial(ctx context.Context, in *content.Tutorial) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("tutorial cannot be nil")
	}
	result, err := c.Client.Database(DATA).Collection(TutorialCol).InsertOne(context, in)
	if err != nil || result.InsertedID == nil {
		return nil, err
	}
	return &content.ReqID{ID: result.InsertedID.(string)}, nil
}

func (c *Service) UpdateTutorial(ctx context.Context, in *content.Tutorial) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(TutorialCol).UpdateOne(context, bson.M{"_id": in.ContentInfo.Id}, in)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.ContentInfo.Id}, nil
}

func (c *Service) DeleteTutorial(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	result, err := c.Client.Database(DATA).Collection(TutorialCol).DeleteOne(context, bson.M{"_id": in.ID})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &content.ReqID{ID: in.ID}, nil
}

func (c *Service) GetTutorial(ctx context.Context, in *content.ReqID) (*content.Tutorial, error) {

	var res *content.Tutorial
	context, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result := c.Client.Database(DATA).Collection(TutorialCol).FindOne(context, bson.M{"_id": in.ID})

	if result.Err() != nil {
		return nil, result.Err()
	}

	resp, err := result.Raw()

	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(resp, &res)

	if err != nil || res == nil {
		return nil, err
	}

	return res, nil
}
