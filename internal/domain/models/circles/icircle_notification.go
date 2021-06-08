//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/golang-ddd-layout/internal/domain/models/$GOPACKAGE
package circles

import "github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"

type ICircleNotification interface {
	SetId(id CircleId)
	SetName(name CircleName)
	SetOwnerId(ownerId users.UserId)
	SetMembers(members []users.UserId)
}
