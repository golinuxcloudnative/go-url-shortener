package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/golinuxcloudnative/go-url-shortener/config"
	"github.com/golinuxcloudnative/go-url-shortener/internal/api/server"
	repositoryRedis "github.com/golinuxcloudnative/go-url-shortener/repository/url/redis"
	"github.com/golinuxcloudnative/go-url-shortener/usecases/url"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	opt := &redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       0,
	}

	rdb := redis.NewClient(opt)
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("could not connect database: %v", err)
	}
	log.Printf("database connection successful: %v", ping)

	dbRepo := repositoryRedis.NewRepositoryRedis(rdb)
	urlService := url.NewService(dbRepo)
	url, err := urlService.CreateURL("google.com")
	if err != nil {
		log.Fatalf("could create short version: %v", err)
	}
	fmt.Println(url)

	echo := echo.New()

	s := server.NewServer(echo, cfg)
	s.Run()
}
