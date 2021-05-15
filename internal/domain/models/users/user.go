package users

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id   UserId
	name UserName
}

func NewUser(id UserId, name UserName) (*User, error) {
	newUser := User{
		id:   id,
		name: name,
	}
	return &newUser, nil
}

func NewUserByName(name UserName) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("UserIdの生成に失敗しました: %s", err)
	}
	user := User{
		id:   UserId(id.String()),
		name: name,
	}
	return &user, nil
}

func (u *User) ChangeName(name UserName) error {
	u.name = name
	return nil
}

func (u User) Equals(other User) (bool, error) {
	return (u.id == other.id), nil
}

func (u User) Notify(note IUserNotification) error {
	note.SetId(u.id)
	note.SetName(u.name)
	return nil
}
