package users

import "github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"

type UserData struct {
	Id   string
	Name string
}

func NewUserData(user users.User) (UserData, error) {
	userData := &UserData{}
	user.Notify(userData)

	return *userData, nil
}

func (d *UserData) SetId(id users.UserId) {
	d.Id = string(id)
}

func (b *UserData) SetName(name users.UserName) {
	b.Name = string(name)
}
