package users

import (
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

// この中でdbコネクションとか保持？
type userRepository struct{}

func NewUserRepository() (users.IUserRepository, error) {
	return &userRepository{}, nil
}

func (ur userRepository) Find(id users.UserId) (*users.User, error) {
	// ユーザが存在しなかったケースをエラーとして返すか、user = nilで表現するか...
	// 一旦エラーで表現する想定で依存箇所書いてる
	//
	user, err := models.
	return nil, nil
}

func (ur userRepository) FindByName(name users.UserName) (*users.User, error) {
	return nil, nil
}

func (ur userRepository) Save(user users.User) error {
	userDataModelBuilder := &UserDataModelBuilder{}
	user.Notify(userDataModelBuilder)
	userDataModel := userDataModelBuilder.Build()
	fmt.Println(userDataModel.id)
	return nil
}

func (ur userRepository) Delete(user users.User) error {
	return nil
}

type UserDataModel struct {
	id   string
	name string
}

func ToModel(from UserDataModel) (*users.User, error) {
	userId, err := users.NewUserId(from.id)
	if err != nil {
		return nil, err
	}
	userName, err := users.NewUserName(from.name)
	if err != nil {
		return nil, err
	}
	user, err := users.NewUser(userId, userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UserDataModelBuilder struct {
	id   users.UserId
	name users.UserName
}

func NewUserDataModelBuilder() (users.IUserNotification, error) {
	return &UserDataModelBuilder{}, nil
}

func (b *UserDataModelBuilder) Id(id users.UserId) {
	b.id = id
}

func (b *UserDataModelBuilder) Name(name users.UserName) {
	b.name = name
}

func (b UserDataModelBuilder) Build() UserDataModel {
	return UserDataModel{
		id:   string(b.id),
		name: string(b.name),
	}
}
