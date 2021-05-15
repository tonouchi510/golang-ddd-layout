package users

type IUserNotification interface {
	SetId(id UserId)
	SetName(name UserName)
}
