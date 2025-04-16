package circles

import (
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/shared"
)

type ICircleFactory interface {
	Create(name CircleName, ownerId shared.UserId) (*Circle, error)
}
