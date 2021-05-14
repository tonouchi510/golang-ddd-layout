package circle_application_service

import (
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

type ICircleApplicationService interface {
	Create(command CircleCreateCommand) error
}

// 実装側はprivateに
type circleApplicationService struct {
	factory    circles.ICircleFactory
	service    circles.CircleService
	circleRepo circles.ICircleRepository
	userRepo   users.IUserRepository
}

func NewCircleApplicationService(
	factory circles.ICircleFactory,
	service circles.CircleService,
	circleRepo circles.ICircleRepository,
	userRepo users.IUserRepository,
) (ICircleApplicationService, error) {
	s := circleApplicationService{
		factory:    factory,
		service:    service,
		circleRepo: circleRepo,
		userRepo:   userRepo,
	}
	return s, nil
}

func (s circleApplicationService) Create(command CircleCreateCommand) error {
	ownerId, err := users.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	name, err := circles.NewName(command.Name)
	if err != nil {
		return err
	}
	circle, err := s.factory.Create(*name, *ownerId) //ownerユーザの存在チェックもこの中で
	if err != nil {
		return err
	}
	exist, err := s.service.Exists(*circle)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("サークル名'%s'はすでに存在しています。", command.Name)
	}
	err = s.circleRepo.Save(*circle)
	return err
}

func (s circleApplicationService) Join(command CircleJoinCommand) error {
	memberId, err := users.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	//aaa
	id, err := circles.NewID(command.CircleId)
	if err != nil {
		return err
	}
	circle, err := s.circleRepo.Find(*id)
	if err != nil {
		return err
	}
	err = circle.Join(*memberId)
	if err != nil {
		return err
	}
	err = s.circleRepo.Save(*circle)
	return err
}
