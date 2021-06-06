package users

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	name        UserName
	user        *User
	userService UserService
}

func (s *UserTestSuite) createUser(name UserName) *User {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	userId, err := NewUserId(id.String())
	if err != nil {
		panic(err)
	}
	user, err := NewUser(userId, name)
	if err != nil {
		panic(err)
	}
	return user
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupSuite() {
	s.name = UserName("hoge") // testなので直接生成
	s.user = s.createUser(s.name)
}

func (s *UserTestSuite) TearDownSuite() {
}

func (s *UserTestSuite) TestUserModel() {
	t := s.T()
	t.Run("Equals/識別子が違うユーザ同士を比較したらfalseになる", func(t *testing.T) {
		sameNameUser := s.createUser(s.name)
		result, err := s.user.Equals(*sameNameUser)
		require.Nil(t, err)
		assert.False(t, result)
	})

	t.Run("Equals/識別子が同じユーザ同士を比較したらtrueになる", func(t *testing.T) {
		name := UserName("fuga")
		sameIdUser, err := NewUser(s.user.Id, name)
		if err != nil {
			panic(err)
		}
		result, err := s.user.Equals(*sameIdUser)
		require.Nil(t, err)
		assert.True(t, result)
	})

	t.Run("ChangeName/Success", func(t *testing.T) {
		user := s.createUser(s.name)
		newName := UserName("hogehoge")
		err := user.ChangeName(newName)
		require.Nil(t, err)
		assert.Equal(t, newName, user.name)
	})
}

func (s *UserTestSuite) TestUserService() {
	t := s.T()
	t.Run("Exists/同じ名前のユーザが存在する場合、trueを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := NewMockIUserRepository(ctrl)
		repo.EXPECT().FindByName(s.user.name).Return(s.user, nil)
		userService := NewUserService(repo)
		exist, err := userService.Exists(*s.user)
		require.Nil(t, err)
		assert.True(t, exist)
	})

	t.Run("Exists/同じ名前のユーザが存在しない場合、falseを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := NewMockIUserRepository(ctrl)
		repo.EXPECT().FindByName(s.name).Return(nil, fmt.Errorf("sql: no rows in result set"))
		userService := NewUserService(repo)
		exist, err := userService.Exists(*s.user)
		require.Nil(t, err)
		assert.False(t, exist)
	})
}
