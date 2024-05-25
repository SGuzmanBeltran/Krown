package handlers

import (
	"context"
	proto_tournament "krown/services/genproto/tournament"
	"krown/services/tournament/types"

	"google.golang.org/grpc"
)

type TournamentGrpcHandler struct {
	tournamentService types.TournamentService
	proto_tournament.UnimplementedTournamentServiceServer
}

func NewGrpcTournamentService(grpc *grpc.Server, tournamentService types.TournamentService) {
	gRpcHadler := &TournamentGrpcHandler{tournamentService: tournamentService}
	proto_tournament.RegisterTournamentServiceServer(grpc, gRpcHadler)
}

func (h *TournamentGrpcHandler) GetTournaments(ctx context.Context, req *proto_tournament.GetTournamentsReq) (*proto_tournament.GetTournamentsRes, error) {
	tournaments, err := h.tournamentService.GetTournaments(ctx, req)
	if (err != nil) {
		return nil, err
	}
	return tournaments, nil
}

func (h *TournamentGrpcHandler) GetTournament(ctx context.Context, req *proto_tournament.GetTournamentReq) (*proto_tournament.GetTournamentRes, error) {
	tournament, err := h.tournamentService.GetTournament(ctx, req)
	if (err != nil) {
		return nil, err
	}
	return tournament, nil
}

func (h *TournamentGrpcHandler) CreateTournaments(ctx context.Context, req *proto_tournament.CreateTournamentReq) (*proto_tournament.CreateTournamentRes, error) {
	createTournaments := types.ConvertProtoCreateTournamentsToCreateTournaments(req.CreateTournaments)
	tournaments, err := h.tournamentService.CreateTournaments(ctx, createTournaments)
	if (err != nil) {
		return nil, err
	}
	return tournaments, nil
}