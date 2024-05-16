package internal

import (
	"context"
	"fmt"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

// TODO: add more validation
func (c *ContentService) GetCoursePage(ctx context.Context, in *content.ReqID) (*content.CoursePage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var coursePage *content.CoursePage

	coursePage, err := c.Cache.GetCachedCoursePage(ctox, in.ID)
	if err == nil && coursePage != nil {
		fmt.Println(" got from cache", coursePage)
		return coursePage, nil
	}
	coursePage, err = c.Db.GetCoursePage(ctox, in)
	if err != nil {
		return &content.CoursePage{}, err
	}
	c.Cache.SetCoursePageToCache(ctox, coursePage)
	return coursePage, nil
}
func (c *ContentService) CreateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ReqID, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result, err := c.Db.CreateCoursePage(ctox, in)

	if err != nil || result == nil {
		return nil, err
	}

	c.Cache.SetCoursePageToCache(ctox, in)

	return result, nil
}

func (c *ContentService) UpdateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result, err := c.Db.UpdateCoursePage(ctox, in)

	if err != nil || result == nil {
		return nil, err
	}

	c.Cache.SetCoursePageToCache(ctox, in)

	return result, nil

}

func (c *ContentService) DeleteCoursePage(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	result, err := c.Db.DeleteCoursePage(ctox, in)

	if err != nil || result == nil {
		return nil, err
	}

	c.Cache.DeleteFromCache(ctox, in.GetID())

	return result, nil
}
