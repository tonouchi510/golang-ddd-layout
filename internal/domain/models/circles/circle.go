package circles

import (
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type Circle struct {
	id      CircleId // idの公開はアリ
	name    CircleName
	ownerId users.UserId
	members []users.UserId
}

func NewCircle(id CircleId, name CircleName, ownerId users.UserId, members []users.UserId) (*Circle, error) {
	circle := Circle{
		id:      id,
		name:    name,
		ownerId: ownerId,
		members: members,
	}
	return &circle, nil
}

func (c *Circle) Join(memberId users.UserId) error {
	if c.IsFull() {
		return fmt.Errorf("CircleFullError: %s", c.id)
	}
	c.members = append(c.members, memberId)
	return nil
}

func (c Circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c Circle) CountMembers() int {
	return len(c.members) + 1 // ownerユーザ分を足している
}

func (c Circle) Notify(note ICircleNotification) error {
	note.SetId(c.id)
	note.SetName(c.name)
	note.SetOwnerId(c.ownerId)
	note.SetMembers(c.members)
	return nil
}
