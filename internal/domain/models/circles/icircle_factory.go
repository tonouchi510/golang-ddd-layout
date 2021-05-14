package circles

import (
	"github.com/huroshotoku/golang-ddd-layout/internal/domain/models/users"
)

type ICircleFactory interface {
	Create(name CircleName, owner users.UserId) (*Circle, error)
}
