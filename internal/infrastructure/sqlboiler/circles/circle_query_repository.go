package circles

import (
	"context"

	"github.com/tonouchi510/golang-ddd-layout/internal/query/circles"
)

type circleQueryRepository struct {
	ctx context.Context
}

func NewCircleQueryRepository(ctx context.Context) (circles.ICircleQueryRepository, error) {
	return &circleQueryRepository{ctx: ctx}, nil
}

func (r circleQueryRepository) ListCircleSummaries(page int, size int) (circles.CircleGetSummariesResult, error) {
	return nil, nil
}
