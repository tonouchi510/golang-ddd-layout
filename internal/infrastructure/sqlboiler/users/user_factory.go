package users

import (
	"context"

	"github.com/google/uuid"
	models "github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type userFactory struct {
	ctx context.Context
}

func NewUserFactory(ctx context.Context) models.IUserFactory {
	return &userFactory{ctx: ctx}
}

func (r userFactory) Create(name models.UserName) (*models.User, error) {
	//ownerユーザの存在チェックもこの中でやってもアリかも
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	userId, err := models.NewUserId(id.String())
	if err != nil {
		return nil, err
	}
	user, err := models.NewUser(userId, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
