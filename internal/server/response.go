package server

import (
	"context"
	"errors"

	"github.com/nanchano/hathor/internal/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error core.Error `json:"error"`
}

// renderErrorResponse parses the error type and writes a GRPC error response
func renderErrorResponse(ctx context.Context, err error) error {
	st := codes.Internal
	resp := ErrorResponse{
		Error: core.NewError("Internal Server Error", int(st)),
	}

	var ierr *core.Error
	if !errors.As(err, &ierr) {
		return status.Errorf(codes.Internal, "Unexpected error while creating event: %s", err.Error())
	} else {
		switch ierr.Code() {
		case core.ErrorNotFound:
			st = codes.NotFound
		case core.ErrorInvalidArgument:
			st = codes.InvalidArgument
		case core.ErrorUnknown:
			fallthrough
		default:
			st = codes.Unknown
		}
	}

	return status.Errorf(st, "Error processing the request: %s", resp.Error.Error())
}
