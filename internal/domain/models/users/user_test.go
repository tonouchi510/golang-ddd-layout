package users

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	id   UserId
	name UserName
	user User
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetUpTest() {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	s.id = UserId(id.String())
	s.name = UserName("hoge")
	s.user = User{
		Id:   s.id,
		name: s.name,
	}
}

func (s *UserTestSuite) TearDownTest() {

}

func (s *UserTestSuite) TestNewUserByName() {
	t := s.T()
	t.Run("success NewUserByName()", func(t *testing.T) {
		user, err := NewUserByName(s.name)
		assert.Nil(t, err)
		assert.Equal(t, 16, len(user.Id))
		assert.Equal(t, s.name, user.name)
	})
}

func (s *UserTestSuite) TestEquals() {
	t := s.T()
	t.Run("識別子が違うユーザ同士を比較したらfalseになる", func(t *testing.T) {
		sameNameUser, err := NewUserByName(s.name)
		if err != nil {
			panic(err)
		}
		result, err := s.user.Equals(*sameNameUser)
		assert.Nil(t, err)
		assert.Negative(t, result)
	})

	t.Run("識別子が同じユーザ同士を比較したらtrueになる", func(t *testing.T) {
		sameIdUser, err := NewUser(s.id, s.name)
		if err != nil {
			panic(err)
		}
		result, err := s.user.Equals(*sameIdUser)
		assert.Nil(t, err)
		assert.Positive(t, result)
	})
}
