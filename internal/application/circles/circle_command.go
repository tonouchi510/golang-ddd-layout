package circles

type CircleCreateCommand struct {
	UserId string
	Name   string
}

func NewCircleCreateCommand(userId string, name string) (CircleCreateCommand, error) {
	c := CircleCreateCommand{
		UserId: userId,
		Name:   name,
	}
	return c, nil
}

type CircleJoinCommand struct {
	UserId   string
	CircleId string
}

func NewCircleJoinCommand(userId string, circleId string) (CircleJoinCommand, error) {
	c := CircleJoinCommand{
		UserId:   userId,
		CircleId: circleId,
	}
	return c, nil
}

type CircleGetCommand struct {
	CircleId string
}
