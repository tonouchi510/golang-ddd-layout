package circles

type CircleGetSummariesCommand struct {
	Page int
	Size int
}

func NewCircleGetSummariesCommand(page int, size int) (CircleGetSummariesCommand, error) {
	c := CircleGetSummariesCommand{
		Page: page,
		Size: size,
	}
	return c, nil
}
