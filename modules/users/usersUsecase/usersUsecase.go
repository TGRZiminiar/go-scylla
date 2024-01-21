package usersusecase

import (
	"context"
	"tgrzimiar/go-scylla/config"
	"tgrzimiar/go-scylla/modules/users"
	usersrepository "tgrzimiar/go-scylla/modules/users/usersRepository"
	"tgrzimiar/go-scylla/pkg/utils"
)

type (
	UsersUsecaseService interface {
		CreateUser(cfg *config.Config, pctx context.Context, req *users.CreateUserReq) (*users.UserRes, error)
		GetAllUsers(cfg *config.Config, pctx context.Context) ([]users.UserRes, error)
	}

	usersUsecase struct {
		usersRepo usersrepository.UsersRepositoryService
	}
)

func NewUsersUsecase(usersRepo usersrepository.UsersRepositoryService) UsersUsecaseService {
	return &usersUsecase{
		usersRepo: usersRepo,
	}
}

func (u *usersUsecase) CreateUser(cfg *config.Config, pctx context.Context, req *users.CreateUserReq) (*users.UserRes, error) {

	userId, err := u.usersRepo.CreateUser(pctx, &users.UserModel{
		Id:    utils.GenrateUUId(),
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &users.UserRes{
		Id:    userId.String(),
		Name:  req.Name,
		Email: req.Email,
	}, nil

}

func (u *usersUsecase) GetAllUsers(cfg *config.Config, pctx context.Context) ([]users.UserRes, error) {

	listUsers, err := u.usersRepo.GetAllUsers(pctx)
	if err != nil {
		return []users.UserRes{}, err
	}

	usersRes := make([]users.UserRes, 0)

	for _, v := range *listUsers {
		usersRes = append(usersRes, users.UserRes{
			Id:    v.Id.String(),
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return usersRes, nil
}
