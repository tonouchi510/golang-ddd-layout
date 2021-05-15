package circles

import "github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"

type ICircleNotification interface {
	Id(id CircleId)
	Name(name CircleName)
	OwnerId(ownerId users.UserId)
	Members(members []users.UserId)
}
