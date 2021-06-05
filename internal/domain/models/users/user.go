package users

type User struct {
	Id   UserId
	name UserName
}

func NewUser(id UserId, name UserName) (*User, error) {
	newUser := User{
		Id:   id,
		name: name,
	}
	return &newUser, nil
}

func (u *User) ChangeName(name UserName) error {
	u.name = name
	return nil
}

func (u User) Equals(other User) (bool, error) {
	return (u.Id == other.Id), nil
}

func (u User) Notify(note IUserNotification) error {
	note.SetId(u.Id)
	note.SetName(u.name)
	return nil
}
