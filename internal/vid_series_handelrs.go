package internal

import (
	"context"
	"errors"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

func (c *ContentService) GetVideoSeries(ctx context.Context, in *content.ReqID) (*content.VideoSeries, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	var videoSeries *content.VideoSeries

	videoSeries, err := c.Cache.GetCachedVideoSeries(ctox, in.ID)

	if err == nil && videoSeries != nil {
		return videoSeries, nil
	}

	videoSeries, err = c.Db.GetVideoSeries(ctox, in)

	if err == nil && videoSeries != nil {
		c.Cache.SetVideoSeriesToCache(ctox, videoSeries)
		return videoSeries, nil
	}
	return videoSeries, err
}

func (c *ContentService) CreateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return &content.ReqID{ID: ""}, errors.New("video series cannot be nil")
	}
	result, err := c.Db.CreateVideoSeries(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.SetVideoSeriesToCache(ctox, in)
	return result, nil
}
func (c *ContentService) UpdateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return &content.ReqID{ID: ""}, errors.New("video series cannot be nil")
	}
	result, err := c.Db.UpdateVideoSeries(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.SetVideoSeriesToCache(ctox, in)
	return result, nil
}

func (c *ContentService) DeleteVideoSeries(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	if in == nil {
		return &content.ReqID{ID: ""}, errors.New("video series cannot be nil")
	}
	result, err := c.Db.DeleteVideoSeries(ctox, in)
	if err != nil || result == nil {
		return &content.ReqID{ID: ""}, err
	}
	c.Cache.DeleteFromCache(ctox, in.GetID())
	return result, nil
}
