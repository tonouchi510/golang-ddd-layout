package users

type IUserFactory interface {
	Create(name UserName) (*User, error)
}
