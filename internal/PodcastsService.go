package internal

import (
	"context"

	content "github.com/Soyaka/content/api/protogen/golang"
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
		go p.Cache.SetPodcastToCache(ctox, podcast)
	}
	return podcast, nil
}

func (p *ContentService) CreatePodcast(ctx context.Context, in *content.Podcast) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := p.Db.CreatePodcast(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go p.Cache.SetPodcastToCache(ctox, in)
	return ok, nil
}

func (p *ContentService) UpdatePodcast(ctx context.Context, in *content.Podcast) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := p.Db.UpdatePodcast(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}

	go p.Cache.SetPodcastToCache(ctox, in)

	return ok, nil
}

func (p *ContentService) DeletePodcast(ctx context.Context, in *content.ReqID) (*content.ResOK, error) {
	ctox, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	ok, err := p.Db.DeletePodcast(ctox, in)
	if err != nil || !ok.OK {
		return &content.ResOK{OK: false}, err
	}
	go p.Cache.DeleteFromCache(ctox, in.ID)
	return ok, nil
}
