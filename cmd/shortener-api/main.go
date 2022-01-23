package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/golinuxcloudnative/go-url-shortener/config"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/golinuxcloudnative/go-url-shortener/internal/api/server"
	repositoryRedis "github.com/golinuxcloudnative/go-url-shortener/repository/redis"
	"github.com/labstack/echo/v4"
)

var Version, Build string

func main() {
	file := flag.String("file", ".env", "env file")
	flag.Parse()

	ctx := context.Background()

	cfg, err := config.NewConfig(*file)
	if err != nil {
		log.Println(err)
	}

	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	opt := &redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	rdb := redis.NewClient(opt)
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("exiting... could not connect database: %v", err)
	}
	log.Printf("database connection successful: %v", ping)
	healthz := domain.Healthz{
		Version:  Version,
		Status:   domain.Healthy,
		Database: ping,
	}

	healthzRepo := repositoryRedis.NewRepositoryHealthzRedis(rdb)
	urlRepo := repositoryRedis.NewRepositoryURLRedis(rdb)

	echo := echo.New()

	s := server.NewServer(echo, cfg, healthz, healthzRepo, urlRepo)
	s.Run()
}
