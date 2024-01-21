package users

import "github.com/gocql/gocql"

type (
	UserModel struct {
		Id    gocql.UUID `db:"id"`
		Name  string     `db:"name"`
		Email string     `db:"email"`
	}
)
