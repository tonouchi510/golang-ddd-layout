package circles

import (
	"context"
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
	"github.com/volatiletech/sqlboiler/boil"
)

type ICircleApplicationService interface {
	Create(command CircleCreateCommand) error
}

// 実装側はprivateに
type circleApplicationService struct {
	ctx        context.Context
	factory    circles.ICircleFactory
	service    circles.CircleService
	circleRepo circles.ICircleRepository
	userRepo   users.IUserRepository
}

func NewCircleApplicationService(
	ctx context.Context,
	factory circles.ICircleFactory,
	service circles.CircleService,
	circleRepo circles.ICircleRepository,
	userRepo users.IUserRepository,
) (ICircleApplicationService, error) {
	s := circleApplicationService{
		ctx:        ctx,
		factory:    factory,
		service:    service,
		circleRepo: circleRepo,
		userRepo:   userRepo,
	}
	return s, nil
}

func (s circleApplicationService) Create(command CircleCreateCommand) error {
	// 将来的にはトランザクションスコープとか使えるようににしたい
	tx, err := boil.BeginTx(s.ctx, nil)
	if err != nil {
		return err
	}

	ownerId, err := users.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	name, err := circles.NewName(command.Name)
	if err != nil {
		return err
	}
	circle, err := s.factory.Create(name, ownerId)
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

	err = s.circleRepo.Save(*circle, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (s circleApplicationService) Join(command CircleJoinCommand) error {
	tx, err := boil.BeginTx(s.ctx, nil)
	if err != nil {
		return err
	}

	memberId, err := users.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	id, err := circles.NewID(command.CircleId)
	if err != nil {
		return err
	}
	circle, err := s.circleRepo.Find(id)
	if err != nil {
		return err
	}
	err = circle.Join(memberId)
	if err != nil {
		return err
	}
	err = s.circleRepo.Save(*circle, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
