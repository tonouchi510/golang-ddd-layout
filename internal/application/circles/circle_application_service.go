package circles

import (
	"context"
	"fmt"

	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/shared"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ICircleApplicationService interface {
	Create(command CircleCreateCommand) error
	Join(command CircleJoinCommand) error
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

	ownerId, err := shared.NewUserId(command.UserId)
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

	memberId, err := shared.NewUserId(command.UserId)
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

// domainモデルを介した非効率なQuery例
func (s circleApplicationService) Get(command CircleGetCommand) (*CircleGetResponse, error) {
	id, err := circles.NewID(command.CircleId)
	if err != nil {
		return nil, err
	}
	circle, err := s.circleRepo.Find(id)
	if err != nil {
		return nil, err
	}
	owner, err := s.userRepoitory.Find(circle.OwnerId)
	if err != nil {
		return nil, err
	}

	// ここでN+1が発生する
	MemberList := []UserResponse{}
	for _, memberId := range circle.MemberIds() {
		user, err := s.userRepoitory.Find(memberId)
		if err != nil {
			return nil, err
		}
		userResponse := UserResponse{
			Id:   string(user.Id),
			Name: string(user.Name),
		}
		MemberList = append(MemberList, userResponse)
	}

	response := CircleGetResponse{
		Id:        string(circle.Id),
		Name:      string(circle.Name),
		OwnerId:   string(circle.OwnerId),
		OwnerName: string(owner.Name),
		Members:   MemberList,
	}
	return &response, nil
}
