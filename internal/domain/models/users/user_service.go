package users

type UserService struct {
	userRepository IUserRepository
}

func NewUserSirvice(userRepository IUserRepository) (UserService, error) {
	us := UserService{userRepository: userRepository}
	return us, nil
}

func (us UserService) Exists(user User) (bool, error) {
	_, err := us.userRepository.FindByName(user.name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
