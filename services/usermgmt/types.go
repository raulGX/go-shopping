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
