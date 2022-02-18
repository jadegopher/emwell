package handlers

import (
	"context"

	"emwell/internal/api/.protobuf/proto"
)

func (h *Handlers) GetEmotionalStatistics(
	ctx context.Context,
	req *proto.GetEmotionalStatisticsRequest,
) (*proto.GetEmotionalStatisticsResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, err
	}

	chart, err := h.linkService.GetByPassword(ctx, req.Password)
	if err != nil {
		return nil, err
	}

	return &proto.GetEmotionalStatisticsResponse{Chart: string(chart)}, nil
}

func validate(req *proto.GetEmotionalStatisticsRequest) error {
	if req == nil {
		return ErrNilRequest
	}

	return nil
}
