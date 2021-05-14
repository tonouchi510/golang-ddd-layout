package circles

import (
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type circleFactory struct{}

func NewCircleFactory() circles.ICircleFactory {
	return &circleFactory{}
}

func (cr circleFactory) Create(name circles.CircleName, owner users.UserId) (*circles.Circle, error) {
	return nil, nil
}
