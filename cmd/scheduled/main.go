package main

import (
	"context"
	"krown/cmd/scheduled/transport"
	"krown/config"
	"krown/db"
	"krown/services/scheduled"
	"log"

	"github.com/jackc/pgx/v5"
)


func main() {
    ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.Envs.DBUrl)
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	defer conn.Close(ctx)

	scheduledQueries := db.New(conn)
	scheduledStore := scheduled.NewScheduledStore(scheduledQueries)

	scheduledGRpcServer := transport.NewScheduledGRPCServer(":9003")
	scheduledGRpcServer.Run(scheduledStore)
}
