package internal

import (
	"log"
	"os"

	content "github.com/Soyaka/microlearn-content/api/protogen/golang"
	"github.com/Soyaka/microlearn-content/internal/cache"
	"github.com/Soyaka/microlearn-content/internal/database"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	redis_addr := os.Getenv("REDIS_ADDR")

	return &GlobalConfig{
		Db_config:    database.GetConfig(),
		Cache_Config: cache.NewCacheConfig(redis_addr), //TODO: Get from env
	}
}

func NewContentService() *ContentService {
	return &ContentService{
		Db:    database.New(GetGlobalConfig().Db_config),
		Cache: cache.NewCache(GetGlobalConfig().Cache_Config),
	}
}
