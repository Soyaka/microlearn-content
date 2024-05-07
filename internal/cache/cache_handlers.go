package cache

import (
	"context"
	"time"

	content "github.com/Soyaka/content/api/protogen/golang"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	TimeOut  = 10 * time.Millisecond
	CacheOut = 10 * time.Minute
)

func (c *Cache) DeleteFromCache(ctx context.Context, ID string) error {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	return c.Client.Del(ctox, ID).Err()
}

// CoursePage Cache Handlers
func (c *Cache) SetCoursePageToCache(ctx context.Context, coursePage *content.CoursePage) error {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(coursePage)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, coursePage.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedCoursePage(ctx context.Context, ID string) (*content.CoursePage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.CoursePage{}, err
	}
	var coursePage *content.CoursePage

	err = bson.Unmarshal(result, &coursePage)

	if err != nil || coursePage == nil {
		return &content.CoursePage{}, err
	}
	return coursePage, nil
}

// TutorialPage Cache Handlers

func (c *Cache) SetTutorialPageToCache(ctx context.Context, tutorialPage *content.TutorialPage) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(tutorialPage)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, tutorialPage.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedTutorialPage(ctx context.Context, ID string) (*content.TutorialPage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.TutorialPage{}, err
	}
	var tutorialPage *content.TutorialPage

	err = bson.Unmarshal(result, &tutorialPage)

	if err != nil || tutorialPage == nil {
		return &content.TutorialPage{}, err
	}
	return tutorialPage, nil
}

// Course Cache Handlers

func (c *Cache) SetCourseToCache(ctx context.Context, course *content.Course) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(course)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, course.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedCourse(ctx context.Context, ID string) (*content.Course, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.Course{}, err
	}
	var course *content.Course

	err = bson.Unmarshal(result, &course)

	if err != nil || course == nil {
		return &content.Course{}, err
	}

	return course, nil
}

// CourseSeries Cache Handlers

func (c *Cache) SetCourseSeriesToCache(ctx context.Context, courseSeries *content.Course) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(courseSeries)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, courseSeries.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedCourseSeries(ctx context.Context, ID string) (*content.Course, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.Course{}, err
	}
	var courseSeries *content.Course

	err = bson.Unmarshal(result, &courseSeries)

	if err != nil || courseSeries == nil {
		return &content.Course{}, err
	}
	return courseSeries, nil
}

// Video Cache Handlers

func (c *Cache) SetVideoToCache(ctx context.Context, video *content.Video) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(video)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, video.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedVideo(ctx context.Context, ID string) (*content.Video, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.Video{}, err
	}
	var video *content.Video

	err = bson.Unmarshal(result, &video)

	if err != nil || video == nil {
		return &content.Video{}, err
	}

	return video, nil
}

// VideoSeries Cache Handlers

func (c *Cache) SetVideoSeriesToCache(ctx context.Context, videoSeries *content.VideoSeries) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(videoSeries)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, videoSeries.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedVideoSeries(ctx context.Context, ID string) (*content.VideoSeries, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.VideoSeries{}, err
	}
	var videoSeries *content.VideoSeries

	err = bson.Unmarshal(result, &videoSeries)

	if err != nil || videoSeries == nil {
		return &content.VideoSeries{}, err
	}

	return videoSeries, nil
}

// Tutorial Cache Handlers

func (c *Cache) SetTutorialToCache(ctx context.Context, tutorial *content.Tutorial) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(tutorial)
	if err != nil {
		return err
	}

	err = c.Client.Set(ctox, tutorial.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedTutorial(ctx context.Context, ID string) (*content.Tutorial, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.Tutorial{}, err
	}

	var tutorial *content.Tutorial
	err = bson.Unmarshal(result, &tutorial)
	if err != nil || tutorial == nil {

		return &content.Tutorial{}, err
	}
	return tutorial, nil
}

// learningpath Cache Handlers

func (c *Cache) SetLearningPathToCache(ctx context.Context, learningPath *content.LearningPath) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(learningPath)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, learningPath.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedLearningPath(ctx context.Context, ID string) (*content.LearningPath, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res := c.Client.Get(ctox, ID)
	result, err := res.Bytes()
	if err != nil {
		return &content.LearningPath{}, err
	}

	var learningPath *content.LearningPath

	err = bson.Unmarshal(result, &learningPath)

	if err != nil || learningPath == nil {
		return &content.LearningPath{}, err
	}

	return learningPath, nil
}

// podcast Cache Handlers

func (c *Cache) SetPodcastToCache(ctx context.Context, podcast *content.Podcast) error {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := bson.Marshal(podcast)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctox, podcast.Id, res, CacheOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetCachedPodcast(ctx context.Context, ID string) (*content.Podcast, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res := c.Client.Get(ctox, ID)

	result, err := res.Bytes()

	if err != nil {
		return &content.Podcast{}, err
	}
	var podcast *content.Podcast
	err = bson.Unmarshal(result, &podcast)

	if err != nil || podcast == nil {
		return &content.Podcast{}, err
	}

	return podcast, nil
}
