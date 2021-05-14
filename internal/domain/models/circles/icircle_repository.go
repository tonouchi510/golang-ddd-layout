package circles

type ICircleRepository interface {
	Save(circle Circle) error
	Find(id CircleId) (*Circle, error)
	FindByName(name CircleName) (*Circle, error)
}
