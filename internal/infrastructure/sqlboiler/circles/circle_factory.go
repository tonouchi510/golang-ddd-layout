package circles

import (
	"context"

	"github.com/google/uuid"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type circleFactory struct {
	ctx context.Context
}

func NewCircleFactory(ctx context.Context) (circles.ICircleFactory, error) {
	return &circleFactory{ctx: ctx}, nil
}

func (r circleFactory) Create(name circles.CircleName, ownerId users.UserId) (*circles.Circle, error) {
	//ownerユーザの存在チェックもこの中でやってもアリかも
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	circleId, err := circles.NewID(id.String())
	if err != nil {
		return nil, err
	}
	members := []users.UserId{}
	circle, err := circles.NewCircle(circleId, name, ownerId, members)
	if err != nil {
		return nil, err
	}
	return circle, nil
}
