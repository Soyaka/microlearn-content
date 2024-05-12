package internal

import (
	"context"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

func (c *ContentService) GetLearningPath(ctx context.Context, in *content.ReqID) (*content.LearningPath, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var learningPath *content.LearningPath

	learningPath, err := c.Cache.GetCachedLearningPath(ctox, in.ID)

	if err == nil && learningPath != nil {
		return learningPath, nil
	}

	learningPath, err = c.Db.GetLearningPath(ctox, in)

	if err != nil {
		return &content.LearningPath{}, err
	}
	if learningPath != nil {
		c.Cache.SetLearningPathToCache(ctox, learningPath) //FIXME: this may be a bug since it may defer than database
	}
	return learningPath, nil
}

func (c *ContentService) CreateLearningPath(ctx context.Context, learningPath *content.LearningPath) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.CreateLearningPath(ctox, learningPath)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.SetLearningPathToCache(ctox, learningPath)
	return ok, nil
}

func (c *ContentService) UpdateLearningPath(ctx context.Context, learningPath *content.LearningPath) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.UpdateLearningPath(ctox, learningPath)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.SetLearningPathToCache(ctox, learningPath)
	return ok, nil
}

func (c *ContentService) DeleteLearningPath(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.DeleteLearningPath(ctox, in)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go c.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil
}
