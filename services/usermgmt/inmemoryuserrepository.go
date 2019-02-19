package usermgmt

import "sync"

type InMemoryUserRepository struct {
	userMap map[string]UserModel
	sync.Mutex
}

func (r *InMemoryUserRepository) AddUser(u *UserCreateRequest) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.userMap[u.Username]; !ok {
		return new(UserAlreadyExistsError)
	}

	r.userMap[u.Username] = *NewUserFromRequest(u)

	return nil
}

func (r *InMemoryUserRepository) GetUserByUsername(username string) (UserModel, error) {
	r.Lock()
	defer r.Unlock()

	user, ok := r.userMap[username]
	if !ok {
		return UserModel{}, new(UserNotFoundError)
	}
	return user, nil
}
