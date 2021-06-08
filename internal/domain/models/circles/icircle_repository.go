//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/golang-ddd-layout/internal/domain/models/$GOPACKAGE
package circles

import (
	"database/sql"
)

type ICircleRepository interface {
	Save(circle Circle, tx *sql.Tx) error
	Find(id CircleId) (*Circle, error)
	FindByName(name CircleName) (*Circle, error)
}
