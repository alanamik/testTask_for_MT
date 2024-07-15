package main

import (
	"context"
	"flag"
	"mt/internal/config"
	"mt/internal/encryption"
	server "mt/internal/server/http"
	"mt/internal/storage/redis"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/configs/dev.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	conf, err := config.New(configFile)
	if err != nil {
		err = errors.Wrap(err, "[config.New()]")
		panic(err)
	}
	redisClient := redis.NewClient(conf)
	encr := encryption.New()
	server := server.New(encr, redisClient, conf)

	go func() {
		if err := server.Start(); err != nil {
			os.Exit(1)
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)
}
