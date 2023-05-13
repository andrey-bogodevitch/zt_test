package main

import (
	"log"

	"zt_test/api"
	"zt_test/entity"
	"zt_test/service"
	"zt_test/storage"

	_ "github.com/lib/pq"
)

func main() {
	config, err := entity.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.PostgresRun(config)
	if err != nil {
		log.Fatal(err)
	}
	cache := storage.NewRedisClient(config.RedisAddr)
	userStorage := storage.New(db, cache)
	userService := service.NewService(userStorage)
	userHandler := api.NewHandler(userService)
	server := api.NewServer(config.ServerPort, userHandler)
	err = server.Run()
	if err != nil {
		log.Fatal(" server run: ", err)
	}

}
