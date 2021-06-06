//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/golang-ddd-layout/internal/domain/models/$GOPACKAGE
package users

type IUserRepository interface {
	Find(id UserId) (*User, error)
	FindByName(name UserName) (*User, error)
	Save(user User) error
	Delete(user User) error
}
