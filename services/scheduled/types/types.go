package types

import (
	"context"
	proto_scheduled "krown/services/genproto/scheduled"
)

type ScheduledService interface {
	GetScheduleds(context.Context, *proto_scheduled.GetScheduledsReq) (*proto_scheduled.GetScheduledsRes, error)
	GetScheduled(context.Context, *proto_scheduled.GetScheduledReq) (*proto_scheduled.GetScheduledRes, error)
	CreateScheduled(context.Context, *proto_scheduled.CreateScheduledReq) (*proto_scheduled.CreateScheduledRes, error)
}