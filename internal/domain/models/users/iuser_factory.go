//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/golang-ddd-layout/internal/domain/models/$GOPACKAGE
package users

type IUserFactory interface {
	Create(name UserName) (*User, error)
}
