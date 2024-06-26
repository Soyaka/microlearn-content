package internal

import (
	"context"
	"time"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

var (
	TimeOut = 300 * time.Millisecond
)

func (c *ContentService) CreateCourseSeries(ctx context.Context, in *content.Course) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	id, err := c.Db.CreateCourseSeries(ctox, in)
	if err != nil {
		return nil, err
	}
	c.Cache.SetCourseSeriesToCache(ctox, in)
	return id, nil
}

func (c *ContentService) UpdateCourseSeries(ctx context.Context, in *content.Course) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	id, err := c.Db.UpdateCourseSeries(ctox, in)
	if err != nil {
		return nil, err
	}
	c.Cache.SetCourseSeriesToCache(ctox, in)
	return id, nil
}

func (c *ContentService) DeleteCourseSeries(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	id, err := c.Db.DeleteCourseSeries(ctox, in)
	if err != nil {
		return nil, err
	}
	c.Cache.DeleteFromCache(ctox, in.GetID())
	return id, nil
}

func (c *ContentService) GetCourseSeries(ctx context.Context, in *content.ReqID) (*content.Course, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var course *content.Course

	course, err := c.Cache.GetCachedCourseSeries(ctox, in.ID)
	if err == nil && course != nil {
		return course, nil
	}
	course, err = c.Db.GetCourseSeries(ctox, in)
	if err != nil {
		return &content.Course{}, err
	}
	if course != nil {
		c.Cache.SetCourseSeriesToCache(ctox, course) //FIXME: this may be a bug since it may defer than database
	}
	return course, nil
}
