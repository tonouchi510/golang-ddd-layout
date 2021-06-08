//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/golang-ddd-layout/internal/domain/models/$GOPACKAGE
package circles

import (
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type ICircleFactory interface {
	Create(name CircleName, ownerId users.UserId) (*Circle, error)
}
