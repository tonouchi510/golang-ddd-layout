package users

import (
	"context"
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
	"github.com/tonouchi510/golang-ddd-layout/internal/infrastructure/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type userRepository struct {
	ctx context.Context
}

func NewUserRepository(ctx context.Context) (users.IUserRepository, error) {
	return &userRepository{ctx: ctx}, nil
}

func (ur userRepository) Find(id users.UserId) (*users.User, error) {
	// ユーザが存在しない場合はエラーが変える
	userData, err := models.FindUserG(ur.ctx, string(id))
	if err != nil {
		return nil, err
	}
	user, err := ToModel(*userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) FindByName(name users.UserName) (*users.User, error) {
	userData, err := models.Users(models.UserWhere.Name.EQ(string(name))).OneG(ur.ctx)
	if err != nil {
		return nil, err
	}
	user, err := ToModel(*userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) Save(user users.User) error {
	// userオブジェクトのパラメータ取得
	// privateな変数を取得するために、通知オブジェクトを使う
	userDataModelBuilder := &UserDataModelBuilder{}
	user.Notify(userDataModelBuilder)
	userData := userDataModelBuilder.Build()
	fmt.Println(userData.ID)

	err := userData.InsertG(ur.ctx, boil.Infer())
	return err
}

func (ur userRepository) Delete(user users.User) error {
	userDataModelBuilder := &UserDataModelBuilder{}
	user.Notify(userDataModelBuilder)
	userData := userDataModelBuilder.Build()
	fmt.Println(userData.ID)

	_, err := userData.DeleteG(ur.ctx, false)
	return err
}

func ToModel(from models.User) (*users.User, error) {
	userId, err := users.NewUserId(from.ID)
	if err != nil {
		return nil, err
	}
	userName, err := users.NewUserName(from.Name)
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

func (b *UserDataModelBuilder) SetId(id users.UserId) {
	b.id = id
}

func (b *UserDataModelBuilder) SetName(name users.UserName) {
	b.name = name
}

func (b UserDataModelBuilder) Build() models.User {
	return models.User{
		ID:   string(b.id),
		Name: string(b.name),
	}
}
