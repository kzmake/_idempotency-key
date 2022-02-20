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
	now := time.Now()

	time.Sleep(3 * time.Second)

	return &pb.NowResponse{Now: now.String()}, nil
}
