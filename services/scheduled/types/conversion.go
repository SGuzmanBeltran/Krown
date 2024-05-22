package types

import (
	"krown/db"
	proto_scheduled "krown/services/genproto/scheduled"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertDBScheduledToProtos(dbScheduled []db.ScheduledTournament) ([]*proto_scheduled.Scheduled, error) {
	var protoScheduleds []*proto_scheduled.Scheduled

	for _, dbs := range dbScheduled {
		// Convert pgtype.Timestamp to int64
		pt, err := ConvertDBScheduledToProto(dbs)
		if err != nil {
			return nil, err
		}
		protoScheduleds = append(protoScheduleds, pt)
	}

	return protoScheduleds, nil
}

func ConvertDBScheduledToProto(dbScheduled db.ScheduledTournament) (*proto_scheduled.Scheduled, error) {
	var protoScheduled *proto_scheduled.Scheduled

	// Convert pgtype.Timestamp to int64
	startTime := dbScheduled.StartTime.Time.Unix()
	recurrenceStt := dbScheduled.RecurrenceStartTimestamp.Time.Unix()
	recurrenceET := dbScheduled.RecurrenceEndTimestamp.Time.Unix()
	protoScheduled = &proto_scheduled.Scheduled{
		Id:                       dbScheduled.ID,
		Name:                     dbScheduled.Name,
		EntryFee:                 int64(dbScheduled.EntryFee),
		StartTime:                startTime, // Ensure startTime is in milliseconds
		RecurrencePattern:        dbScheduled.RecurrencePattern,
		RecurrenceStartTimestamp: recurrenceStt,
		RecurrenceEndTimestamp:   recurrenceET,
		MustRenew:                dbScheduled.MustRenew.Bool,
	}

	return protoScheduled, nil
}

func ConvertProtoToParams(protoScheduled *proto_scheduled.CreateScheduled) (*db.CreateScheduledTournamentParams) {
	st := time.Unix(protoScheduled.StartTime, 0)
	rstt := time.Unix(protoScheduled.RecurrenceStartTimestamp, 0)
	ret := time.Unix(protoScheduled.RecurrenceEndTimestamp, 0)

	startTime := pgtype.Timestamp{
		Time: st,
		Valid: true,
	}

	recurrenceStt := pgtype.Timestamp{
		Time: rstt,
		Valid: true,
	}

	recurrenceET := pgtype.Timestamp{
		Time: ret,
		Valid: true,
	}

	scheduledParams := &db.CreateScheduledTournamentParams{
		Name:                     protoScheduled.Name,
		EntryFee:                 protoScheduled.EntryFee,
		StartTime:                startTime,
		RecurrencePattern:        protoScheduled.RecurrencePattern,
		RecurrenceStartTimestamp: recurrenceStt,
		RecurrenceEndTimestamp:   recurrenceET,
		MustRenew: pgtype.Bool{
			Bool:  protoScheduled.MustRenew,
			Valid: true,
		},
	}

	return scheduledParams
}
