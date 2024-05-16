package internal

import (
	"context"
	"errors"
	"fmt"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

func (c *ContentService) CreateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("tutorial page cannot be nil")
	}
	result, err := c.Db.CreateTutorialPage(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.SetTutorialPageToCache(ctox, in)
	return result, nil

}

func (c *ContentService) UpdateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("tutorial page cannot be nil")
	}
	result, err := c.Db.UpdateTutorialPage(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.SetTutorialPageToCache(ctox, in)
	return result, nil
}

func (c *ContentService) DeleteTutorialPage(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return nil, errors.New("tutorial page cannot be nil")
	}
	result, err := c.Db.DeleteTutorialPage(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.DeleteFromCache(ctox, in.GetID())
	return result, nil


}

func (c *ContentService) GetTutorialPage(ctx context.Context, in *content.ReqID) (*content.TutorialPage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var tutorialPage *content.TutorialPage

	tutorialPage, err := c.Cache.GetCachedTutorialPage(ctox, in.ID)
	if err == nil && tutorialPage != nil {
		fmt.Println("got from cache")
		return tutorialPage, nil
	}

	tutorialPage, err = c.Db.GetTutorialPage(ctox, in)
	if err == nil && tutorialPage != nil {
		c.Cache.SetTutorialPageToCache(ctox, tutorialPage)
		return tutorialPage, nil
	}
	return tutorialPage, err
}
