package main

import (
	"championForge/config"
	"championForge/db"
	"championForge/services/user"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.Envs.DBUrl)
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	defer conn.Close(ctx)

	userQueries := db.New(conn)
	userStore := user.NewStore(userQueries)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	log.Println("User service listening on 3000")
	err = http.ListenAndServe(":3000", r)
	if(err != nil){
		log.Fatal("Could init User service", err)
	}
}
