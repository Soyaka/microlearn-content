package internal

import (
	"context"

	content "github.com/Soyaka/content/api/protogen/golang"
)

func (c *ContentService) CreateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	ok, err := c.Db.CreateTutorialPage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go c.Cache.SetTutorialPageToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) UpdateTutorialPage(ctx context.Context, in *content.TutorialPage) (*content.ResOK, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.UpdateTutorialPage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go c.Cache.SetTutorialPageToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) DeleteTutorialPage(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.DeleteTutorialPage(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go c.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil

}

func (c *ContentService) GetTutorialPage(ctx context.Context, in *content.ReqID) (*content.TutorialPage, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var tutorialPage *content.TutorialPage

	tutorialPage, err := c.Cache.GetCachedTutorialPage(ctox, in.ID)
	if err == nil && tutorialPage != nil {
		return tutorialPage, nil
	}

	tutorialPage, err = c.Db.GetTutorialPage(ctox, in)
	if err == nil && tutorialPage != nil {
		go c.Cache.SetTutorialPageToCache(ctox, tutorialPage)
		return tutorialPage, nil
	}
	return tutorialPage, err
}
