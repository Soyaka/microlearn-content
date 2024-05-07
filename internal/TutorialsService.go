package internal

import (
	"context"

	content "github.com/Soyaka/content/api/protogen/golang"
)

func (c *ContentService) GetTutorial(ctx context.Context, in *content.ReqID) (*content.Tutorial, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var tutorial *content.Tutorial

	tutorial, err := c.Cache.GetCachedTutorial(ctox, in.ID)

	if err == nil && tutorial != nil {
		return tutorial, nil
	}

	tutorial, err = c.Db.GetTutorial(ctox, in)

	if err == nil && tutorial != nil {
		go c.Cache.SetTutorialToCache(ctox, tutorial)
		return tutorial, nil
	}
	return tutorial, err
}

func (c *ContentService) CreateTutorial(ctx context.Context, in *content.Tutorial) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.CreateTutorial(ctox, in)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}

	go c.Cache.SetTutorialToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) UpdateTutorial(ctx context.Context, in *content.Tutorial) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.UpdateTutorial(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go c.Cache.SetTutorialToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) DeleteTutorial(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.DeleteTutorial(ctox, in)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}

	go c.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil
}
