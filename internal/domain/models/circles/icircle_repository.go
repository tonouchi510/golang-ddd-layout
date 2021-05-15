package circles

import (
	"database/sql"
)

type ICircleRepository interface {
	Save(circle Circle, tx *sql.Tx) error
	Find(id CircleId) (*Circle, error)
	FindByName(name CircleName) (*Circle, error)
}
