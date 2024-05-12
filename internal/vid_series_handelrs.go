package internal

import (
	"context"

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

func (c *ContentService) CreateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := c.Db.CreateVideoSeries(ctox, in)

	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}

	c.Cache.SetVideoSeriesToCache(ctox, in)

	return ok, nil
}
func (c *ContentService) UpdateVideoSeries(ctx context.Context, in *content.VideoSeries) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	ok, err := c.Db.UpdateVideoSeries(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.SetVideoSeriesToCache(ctox, in)
	return ok, nil
}

func (c *ContentService) DeleteVideoSeries(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok ,err := c.Db.DeleteVideoSeries(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	c.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil
}
