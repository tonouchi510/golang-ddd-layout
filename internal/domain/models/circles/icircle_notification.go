package circles

import "github.com/tonouchi510/golang-ddd-layout/internal/domain/shared"

type ICircleNotification interface {
	SetId(id CircleId)
	SetName(name CircleName)
	SetOwnerId(ownerId shared.UserId)
	SetMemberIds(members []shared.UserId)
}
