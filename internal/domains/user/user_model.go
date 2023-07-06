package user

type Register struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
