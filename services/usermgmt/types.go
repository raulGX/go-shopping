package usermgmt

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) IsValid() bool {
	// can return formatted error
	isValid := true
	if u.Username == "" {
		isValid = false
	}
	if len(u.Password) < 6 {
		isValid = false
	}
	return isValid
}

type UserRepository interface {
	AddUser(UserCreateRequest) error
	GetUserByUsername(string) (UserModel, error)
}

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "User not found"
}

type UserAlreadyExistsError struct{}

func (e *UserAlreadyExistsError) Error() string {
	return "User already exists"
}
