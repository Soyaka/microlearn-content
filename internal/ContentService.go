package internal

import (
	content "github.com/Soyaka/content/api/protogen/golang"
	"github.com/Soyaka/content/internal/cache"
	"github.com/Soyaka/content/internal/database"
)

type ContentService struct {
	Db    *database.Service
	Cache *cache.Cache
	content.UnimplementedLearningPathServiceServer
	content.UnimplementedCoursePageServiceServer
	content.UnimplementedTutorialPageServiceServer
	content.UnimplementedTutorialServiceServer
	content.UnimplementedPodcastServiceServer
	content.UnimplementedVideoSeriesServiceServer
	content.UnimplementedCourseSeriesServiceServer
}

type GlobalConfig struct {
	Db_config    *database.DataConfig
	Cache_Config *cache.CacheConfig
}

func GetGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		Db_config:    database.GetConfig(),
		Cache_Config: cache.NewLocalConfig("localhost:6379"), //TODO: Get from env
	}
}

func NewContentService() ContentService {
	return ContentService{
		Db:    database.New(GetGlobalConfig().Db_config),
		Cache: cache.NewCache(GetGlobalConfig().Cache_Config),
	}
}
