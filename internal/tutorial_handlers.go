package internal

import (
	"context"
	"fmt"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

func (c *ContentService) GetTutorial(ctx context.Context, in *content.ReqID) (*content.Tutorial, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var tutorial *content.Tutorial

	tutorial, err := c.Cache.GetCachedTutorial(ctox, in.ID)

	if err == nil && tutorial != nil {
		fmt.Println("got from cache")

		return tutorial, nil
	}

	tutorial, err = c.Db.GetTutorial(ctox, in)

	if err == nil && tutorial != nil {
		c.Cache.SetTutorialToCache(ctox, tutorial)
		return tutorial, nil
	}
	return tutorial, err
}

func (c *ContentService) CreateTutorial(ctx context.Context, in *content.Tutorial) (*content.ReqID, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	id, err := c.Db.CreateTutorial(ctox, in)
	if err != nil || id == nil {
		return nil, err
	}
	c.Cache.SetTutorialToCache(ctox, in)
	return id, nil
}

func (c *ContentService) UpdateTutorial(ctx context.Context, in *content.Tutorial) (*content.ReqID, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	id, err := c.Db.UpdateTutorial(ctox, in)
	if err != nil || id == nil {
		return nil, err
	}
	c.Cache.SetTutorialToCache(ctox, in)
	return id, nil
}

func (c *ContentService) DeleteTutorial(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res, err := c.Db.DeleteTutorial(ctox, in)
	if err != nil || res == nil {

		return nil, err
	}

	 c.Cache.DeleteFromCache(ctox, in.ID)
	return res, nil
}
