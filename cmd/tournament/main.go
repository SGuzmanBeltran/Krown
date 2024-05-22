package main

import (
	"context"
	"krown/cmd/tournament/transport"
	"krown/config"
	"krown/db"
	"krown/services/tournament"
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

	tournamentQueries := db.New(conn)
	tournamentStore := tournament.NewTournamentStore(tournamentQueries)

	tournamentGRpcServer := transport.NewTournamentGRPCServer(":9002")
	tournamentGRpcServer.Run(tournamentStore)
}
