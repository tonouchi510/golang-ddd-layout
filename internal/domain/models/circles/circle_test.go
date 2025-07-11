package circles

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/shared"
)

type CircleTestSuite struct {
	suite.Suite
	ownerId shared.UserId
	name    CircleName
	userId1 shared.UserId
	userId2 shared.UserId
}

func createCircleWith29Members(ownerId shared.UserId, name CircleName) *Circle {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	circleId := CircleId(id.String())
	members := []shared.UserId{}
	for i := 0; i < 28; i++ {
		id, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		userId := shared.UserId(id.String())
		members = append(members, userId)
	}
	circle, err := NewCircle(circleId, name, ownerId, members)
	if err != nil {
		return nil
	}
	return circle
}

func TestCircle(t *testing.T) {
	suite.Run(t, new(CircleTestSuite))
}

func (s *CircleTestSuite) SetupSuite() {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	s.ownerId, err = shared.NewUserId(id.String())
	if err != nil {
		panic(err)
	}
	s.name = CircleName("hoge")

	id, err = uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	s.userId1 = shared.UserId(id.String())
	id, err = uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	s.userId2 = shared.UserId(id.String())
}

func (s *CircleTestSuite) TearDownSuite() {
}

func (s *CircleTestSuite) TestCircleModel() {
	t := s.T()
	t.Run("Join/サークルにメンバーが追加できる", func(t *testing.T) {
		circle := createCircleWith29Members(s.ownerId, s.name)
		err := circle.Join(s.userId1)
		require.Nil(t, err)
		assert.Equal(t, 30, circle.CountMembers())
	})

	t.Run("Join/サークルメンバーが上限の時、メンバー追加に失敗する", func(t *testing.T) {
		circle := createCircleWith29Members(s.ownerId, s.name)
		err := circle.Join(s.userId1)
		require.Nil(t, err)
		err = circle.Join(s.userId2)
		require.NotNil(t, err)
	})
}
