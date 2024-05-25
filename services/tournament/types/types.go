package types

import (
	"context"
	"krown/db"
	proto_tournament "krown/services/genproto/tournament"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type TournamentService interface {
	GetTournaments(context.Context, *proto_tournament.GetTournamentsReq) (*proto_tournament.GetTournamentsRes, error)
	GetTournament(context.Context, *proto_tournament.GetTournamentReq) (*proto_tournament.GetTournamentRes, error)
	CreateTournaments(c context.Context, tournaments []*CreateTournament) (*proto_tournament.CreateTournamentRes, error)
}

type Tournament struct {
	ID        int64
    Name      string
    EntryFee  int64
    StartTime pgtype.Timestamp
}

type CreateTournament struct {
	Name      string
    EntryFee  int64
    StartTime pgtype.Timestamp
}

func ConvertProtoCreateTournamentsToCreateTournaments(protoTournament []*proto_tournament.CreateTournamentDTO) []*CreateTournament {
	var createTournaments []*CreateTournament
	for _, dbt := range protoTournament {
		pt := ConvertProtoCreateTournamentToCreateTournament(dbt)
		createTournaments = append(createTournaments, pt)
	}
	return createTournaments
}

func ConvertProtoCreateTournamentToCreateTournament(protoTournament *proto_tournament.CreateTournamentDTO) (*CreateTournament) {
	st := time.Unix(protoTournament.StartTime, 0)
	stTimestamp := &pgtype.Timestamp{
		Time: st.UTC(),
		Valid: true,
	}
	return &CreateTournament{
		Name: protoTournament.Name,
		EntryFee: protoTournament.EntryFee,
		StartTime: *stTimestamp,
	}
}

func ConvertDBTournamentsToProtoTournaments(dbtournaments []*db.Tournament) ([]*proto_tournament.Tournament, error) {
	var protoTournaments []*proto_tournament.Tournament
	for _, dbt := range dbtournaments {
		pt, err := ConvertDBTournamentToProto(dbt)
		if err != nil {
			return nil, err
		}
		protoTournaments = append(protoTournaments, pt)
	}

	return protoTournaments, nil
}

func ConvertDBTournamentToProto(dbtournament *db.Tournament) (*proto_tournament.Tournament, error) {
	var protoTournaments *proto_tournament.Tournament
	startTime := dbtournament.StartTime.Time.Unix()
	protoTournaments = &proto_tournament.Tournament{
		Id:        dbtournament.ID,
		Name:      dbtournament.Name,
		EntryFee:  int64(dbtournament.EntryFee),
		StartTime: startTime,
	}
	return protoTournaments, nil
}
