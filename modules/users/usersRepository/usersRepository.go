package usersrepository

import (
	"context"
	"errors"
	"fmt"
	"tgrzimiar/go-scylla/modules/users"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type (
	UsersRepositoryService interface {
		CreateUser(pctx context.Context, user *users.UserModel) (*gocql.UUID, error)
		GetAllUsers(pctx context.Context) (*[]users.UserModel, error)
	}

	authRepository struct {
		db     *gocql.Session
		logger *zap.Logger
	}
)

func NewUserRepository(db *gocql.Session, logger *zap.Logger) UsersRepositoryService {
	return &authRepository{
		db:     db,
		logger: logger,
	}
}

func (r *authRepository) CreateUser(pctx context.Context, user *users.UserModel) (*gocql.UUID, error) {
	query := `	
	INSERT INTO userdata (
		id,
		name,
		email
	) VALUES (?, ?, ?)
	`

	if err := r.db.Query(query, user.Id, user.Email, user.Name).Exec(); err != nil {
		fmt.Println("insert user failed", err)
		r.logger.Error("insert users.userdata", zap.Error(err))
		return nil, errors.New("insert users.userdata failed")
	}

	return &user.Id, nil
}

func (r *authRepository) GetAllUsers(pctx context.Context) (*[]users.UserModel, error) {
	query := `
    SELECT
		id,
		name,
		email
	FROM userdata
	`
	q := r.db.Query(query)
	usersList := make([]users.UserModel, 0)

	iter := q.Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			r.logger.Warn("Error closing iterator", zap.Error(err))
		}
	}()

	var user users.UserModel
	for iter.Scan(&user.Id, &user.Name, &user.Email) {
		usersList = append(usersList, user)
	}

	if err := iter.Scanner().Err(); err != nil {
		r.logger.Warn("Error scanning iterator", zap.Error(err))
		return nil, err
	}

	return &usersList, nil
}
