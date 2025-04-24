package circles

type UserData struct {
	UserId   string
	UserName string
}

type CircleData struct {
	Id        string
	Name      string
	OwnerId   string
	OwnerName string
	Members   []UserData
}
