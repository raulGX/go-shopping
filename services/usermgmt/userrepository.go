package usermgmt

import (
	sql "database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db}
}

func (r *UserPostgresRepository) AddUser(req UserCreateRequest) error {
	user := NewUserFromRequest(&req)
	var nextId int
	r.db.QueryRow("Select nextval(pg_get_serial_sequence('users', 'id'))").Scan(&nextId)
	user.ID = nextId
	_, err := r.db.Exec(`INSERT INTO USERS(id, username, password) VALUES ($1, $2, $3)`, user.ID, user.Username, user.Password)
	return err
}

func (r *UserPostgresRepository) GetUserByUsername(username string) (UserModel, error) {
	const qry = `
SELECT id, username FROM users WHERE username = $1
	`
	row := r.db.QueryRow(qry, username)
	var user = UserModel{}
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		print(err.Error())
		return user, new(UserNotFoundError)
	}
	return user, nil
}

var uuid = 0

func NewUserFromRequest(req *UserCreateRequest) *UserModel {
	newUser := &UserModel{
		Username: req.Username,
		Password: hashPassword(req.Password),
		ID:       uuid,
	}
	uuid++
	return newUser
}

func hashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
