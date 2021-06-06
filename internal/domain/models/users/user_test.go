package users

import (
	"testing"

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
