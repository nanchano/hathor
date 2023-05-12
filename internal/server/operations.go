package server

import (
	"context"
	"errors"

	"github.com/nanchano/hathor/internal/core"
	pb "github.com/nanchano/hathor/pb"
)

var invalidFieldMask error = errors.New("Field Mask is not valid")

// Ping checks the server is running
func (s server) Ping(context.Context, *pb.EmptyRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

// Create event receives a request and calls the service to create a new event
// given the payload, sending the event itself as a response afterwards.
// If unsuccessful, an error response will be sent.
func (s server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.Event, error) {
	e := core.Event{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Category:    req.GetCategory(),
		Location:    req.GetLocation(),
		Publisher:   req.GetPublisher(),
		Lineup:      req.GetLineup(),
		StartTS:     req.GetStartTs().AsTime(),
		EndTS:       req.GetEndTs().AsTime(),
	}
	event, err := s.service.CreateEvent(ctx, e)
	if err != nil {
		return nil, renderErrorResponse(ctx, err)
	}

	return modelToProto(event), nil
}

// GetEvent receives a request and calls the service for the required
// event given the ID, return the event as a response.
// If unsuccessful, an error response will be sent.
func (s server) GetEvent(ctx context.Context, req *pb.EventIDRequest) (*pb.Event, error) {
	id := req.GetEventId()

	event, err := s.service.GetEvent(ctx, id)
	if err != nil {
		return nil, renderErrorResponse(ctx, err)
	}

	return modelToProto(event), nil
}

// UpdateEvent receives a request and calls the service to
// update an event, sending a response with the updated event afterwards.
// If unsuccessful, an error response will be sent.
func (s server) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.Event, error) {
	e := core.Event{
		Name:        req.GetEvent().GetName(),
		Description: req.GetEvent().GetDescription(),
		Category:    req.GetEvent().GetCategory(),
		Location:    req.GetEvent().GetLocation(),
		Publisher:   req.GetEvent().GetPublisher(),
		Lineup:      req.GetEvent().GetLineup(),
		StartTS:     req.GetEvent().GetStartTs().AsTime(),
		EndTS:       req.GetEvent().GetEndTs().AsTime(),
		ID:          req.GetEvent().GetId(),
	}

	event, err := s.service.UpdateEvent(ctx, e)
	if err != nil {
		return nil, renderErrorResponse(ctx, err)
	}

	return modelToProto(event), nil
}

// DeleteEvent receives a request and calls the service to
// delete an event by ID, sending an empty response afterwards.
// If unsuccessful, an error response will be sent.
func (s server) DeleteEvent(ctx context.Context, req *pb.EventIDRequest) (*pb.EmptyResponse, error) {
	id := req.GetEventId()

	err := s.service.DeleteEvent(ctx, id)
	if err != nil {
		return nil, renderErrorResponse(ctx, err)
	}

	return &pb.EmptyResponse{}, nil
}
