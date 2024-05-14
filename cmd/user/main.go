package main

import (
	"context"
	"krown/cmd/user/transport"
	"krown/config"
	"krown/db"
	"krown/services/user"
	handler "krown/services/user/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.Envs.DBUrl)
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	defer conn.Close(ctx)

	userQueries := db.New(conn)
	userStore := user.NewStore(userQueries)
	userService := user.NewUserService(userStore)
	userHandler := handler.NewHandler(userService)

	grpcServer := transport.NewGRPCServer(":9001")
	go grpcServer.Run(userStore)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	userHandler.RegisterRoutes(v1)

	log.Println("User service listening on 3000")
	err = app.Listen(":3000")
	if(err != nil){
		log.Fatal("Could init User service", err)
	}
}
