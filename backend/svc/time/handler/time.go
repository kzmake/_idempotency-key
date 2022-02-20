package handler

import (
	"context"
	"time"

	pb "github.com/kzmake/_idempotency-key/gen/go/time/v1"
)

type timeH struct{}

var _ pb.TimeServer = new(timeH)

func NewTime() pb.TimeServer { return &timeH{} }

func (h *timeH) Now(_ context.Context, _ *pb.NowRequest) (*pb.NowResponse, error) {
	return &pb.NowResponse{Now: time.Now().String()}, nil
}
