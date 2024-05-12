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
func (c *ContentService) CreateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.CreateCoursePage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.SetCoursePageToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) UpdateCoursePage(ctx context.Context, in *content.CoursePage) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.UpdateCoursePage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.SetCoursePageToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) DeleteCoursePage(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.DeleteCoursePage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil
}
