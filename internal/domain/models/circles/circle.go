package circles

import (
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/shared"
)

type Circle struct {
	Id        CircleId // idの公開はアリ
	name      CircleName
	ownerId   shared.UserId
	memberIds []shared.UserId
}

func NewCircle(id CircleId, name CircleName, ownerId shared.UserId, memberIds []shared.UserId) (*Circle, error) {
	circle := Circle{
		Id:        id,
		name:      name,
		ownerId:   ownerId,
		memberIds: members,
	}
	return &circle, nil
}

func (c *Circle) Join(memberId shared.UserId) error {
	if c.IsFull() {
		return fmt.Errorf("CircleFullError: %s", c.Id)
	}
	c.memberIds = append(c.memberIds, memberId)
	return nil
}

func (c Circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c Circle) CountMembers() int {
	return len(c.memberIds) + 1 // ownerユーザ分を足している
}

func (c Circle) Notify(note ICircleNotification) error {
	note.SetId(c.Id)
	note.SetName(c.name)
	note.SetOwnerId(c.ownerId)
	note.SetMemberIds(c.memberIds)
	return nil
}
