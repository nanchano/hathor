package server

import (
	"github.com/nanchano/hathor/internal/core"
	pb "github.com/nanchano/hathor/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func modelToProto(e *core.Event) *pb.Event {
	return &pb.Event{
		Id:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		Publisher:   e.Publisher,
		Lineup:      e.Lineup,
		StartTs:     timestamppb.New(e.StartTS),
		EndTs:       timestamppb.New(e.EndTS),
	}
}
