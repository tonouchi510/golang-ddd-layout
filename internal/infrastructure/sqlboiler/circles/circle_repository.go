package circles

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/circles"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
	"github.com/tonouchi510/golang-ddd-layout/internal/infrastructure/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type circleRepository struct {
	ctx context.Context
}

func NewCircleRepository(ctx context.Context) (circles.ICircleRepository, error) {
	return &circleRepository{ctx: ctx}, nil
}

func (r circleRepository) Save(circle circles.Circle, tx *sql.Tx) error {
	circleDataModelBuilder := &CircleDataModelBuilder{}
	circle.Notify(circleDataModelBuilder)
	circleData, membersData := circleDataModelBuilder.Build()
	fmt.Println(circleData.ID)

	if tx == nil {
		return fmt.Errorf("")
	} else {
		err := circleData.Upsert(r.ctx, tx, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
		for _, circleMember := range membersData {
			err = circleMember.Upsert(r.ctx, tx, boil.Infer(), boil.Infer())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r circleRepository) Find(id circles.CircleId) (*circles.Circle, error) {
	circleData, err := models.FindCircleG(r.ctx, string(id))
	if err != nil {
		return nil, err
	}
	membersData, err := models.CircleMembers(
		models.CircleMemberWhere.CircleID.EQ(circleData.ID),
	).AllG(r.ctx)
	if err != nil {
		return nil, err
	}
	circle, err := ToModel(*circleData, membersData)
	if err != nil {
		return nil, err
	}
	return circle, nil
}

func (r circleRepository) FindByName(name circles.CircleName) (*circles.Circle, error) {
	circleData, err := models.Circles(models.CircleWhere.Name.EQ(string(name))).OneG(r.ctx)
	if err != nil {
		return nil, err
	}
	membersData, err := models.CircleMembers(
		models.CircleMemberWhere.CircleID.EQ(circleData.ID),
	).AllG(r.ctx)
	if err != nil {
		return nil, err
	}
	circle, err := ToModel(*circleData, membersData)
	if err != nil {
		return nil, err
	}
	return circle, nil
}

func ToModel(circleData models.Circle, membersData models.CircleMemberSlice) (*circles.Circle, error) {
	id, err := circles.NewID(circleData.ID)
	if err != nil {
		return nil, err
	}
	name, err := circles.NewName(circleData.Name)
	if err != nil {
		return nil, err
	}
	ownerId, err := users.NewUserId(circleData.OwnerID)
	if err != nil {
		return nil, err
	}
	members := []users.UserId{}
	for _, circleMember := range membersData {
		memberId, err := users.NewUserId(circleMember.MemberID)
		if err != nil {
			return nil, err
		}
		members = append(members, memberId)
	}
	circle, err := circles.NewCircle(id, name, ownerId, members)
	if err != nil {
		return nil, err
	}
	return circle, nil
}

type CircleDataModelBuilder struct {
	id      circles.CircleId
	name    circles.CircleName
	ownerId users.UserId
	members []users.UserId
}

func NewUserDataModelBuilder() (circles.ICircleNotification, error) {
	return &CircleDataModelBuilder{}, nil
}

func (b *CircleDataModelBuilder) SetId(id circles.CircleId) {
	b.id = id
}

func (b *CircleDataModelBuilder) SetName(name circles.CircleName) {
	b.name = name
}

func (b *CircleDataModelBuilder) SetOwnerId(ownerId users.UserId) {
	b.ownerId = ownerId
}

func (b *CircleDataModelBuilder) SetMembers(members []users.UserId) {
	b.members = members
}

func (b CircleDataModelBuilder) Build() (models.Circle, models.CircleMemberSlice) {
	circle := models.Circle{
		ID:      string(b.id),
		Name:    string(b.name),
		OwnerID: string(b.ownerId),
	}

	circleMembers := models.CircleMemberSlice{}
	for _, userId := range b.members {
		m := models.CircleMember{
			CircleID: string(b.id),
			MemberID: string(userId),
		}
		circleMembers = append(circleMembers, &m)
	}
	return circle, circleMembers
}
