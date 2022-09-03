package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api-test/internal/config"
	"rest-api-test/internal/user"
	"rest-api-test/internal/user/cache"
	"rest-api-test/internal/user/db"
	"rest-api-test/pkg/client/mongodb"
	"rest-api-test/pkg/client/redisdb"
	"rest-api-test/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()

	logger.Info("create router")
	router := httprouter.New()

	logger.Info("initializing new redis client")
	rdb, err := redisdb.NewClient(context.Background(), &cfg.Redis)
	if err != nil {
		logger.Fatalln("failed to initialize redis client")
		panic(err)
	}

	logger.Info("initializing new mongo client")
	mongoClient, err := mongodb.NewClient(context.Background(), cfg)
	if err != nil {
		logger.Fatalln("failed to initialize mongo client")
		return
	}

	logger.Info("initializing new user storage")
	storage := db.NewStorage(mongoClient, cfg.Collection, logger)

	logger.Info("initializing new user cache")
	c := cache.NewCache(rdb)

	logger.Info("initializing new user service")
	service := user.NewService(storage, logger, c)

	logger.Info("register new user handler")
	handler := user.NewUserHandler(logger, cfg, service)

	handler.Register(router)
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	logger.Info("start application")

	address := fmt.Sprintf("%s:%s", cfg.BindIp, cfg.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server is listening address %s", address)
	log.Fatalln(srv.Serve(listener))
}
