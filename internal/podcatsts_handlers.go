package internal

import (
	"context"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
)

func (p *ContentService) GetPodcast(ctx context.Context, in *content.ReqID) (*content.Podcast, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var podcast *content.Podcast
	podcast, err := p.Cache.GetCachedPodcast(ctox, in.ID)
	if err == nil && podcast != nil {
		return podcast, nil
	}
	podcast, err = p.Db.GetPodcast(ctox, in)
	if err != nil {
		return &content.Podcast{}, err
	}
	if podcast != nil {
		p.Cache.SetPodcastToCache(ctox, podcast)
	}
	return podcast, nil
}

func (p *ContentService) CreatePodcast(ctx context.Context, in *content.Podcast) (*content.ReqID, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := p.Db.CreatePodcast(ctox, in)
	if err != nil || res.ID == "" {
		return nil, err
	}

	p.Cache.SetPodcastToCache(ctox, in)

	return res, nil

}

func (p *ContentService) UpdatePodcast(ctx context.Context, in *content.Podcast) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := p.Db.UpdatePodcast(ctox, in)
	if err != nil || res.ID == "" {
		return nil, err
	}
	return res, nil
}

func (p *ContentService) DeletePodcast(ctx context.Context, in *content.ReqID) (*content.ReqID, error) {

	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	res, err := p.Db.DeletePodcast(ctox, in)
	if err != nil || res.ID == "" {
		return nil, err
	}

	p.Cache.DeleteFromCache(ctox, in.GetID())
	return res, nil
}
