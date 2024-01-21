package usershandler

import (
	"context"
	"net/http"
	"tgrzimiar/go-scylla/config"
	"tgrzimiar/go-scylla/modules/users"
	usersusecase "tgrzimiar/go-scylla/modules/users/usersUsecase"
	"tgrzimiar/go-scylla/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type (
	UsersHttpHandlerService interface {
		CreateUser(c *fiber.Ctx) error
		GetAllUsers(c *fiber.Ctx) error
	}

	usersHttpHandler struct {
		cfg          *config.Config
		usersUsecase usersusecase.UsersUsecaseService
	}
)

func NewUsersHttpHandler(cfg *config.Config, usersUsecase usersusecase.UsersUsecaseService) UsersHttpHandlerService {
	return &usersHttpHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}

func (h *usersHttpHandler) CreateUser(c *fiber.Ctx) error {

	ctx := context.Background()

	// wrapper := request.NewContextWrapper(c)

	req := new(users.CreateUserReq)
	// if err := wrapper.ParseJson(req); err != nil {
	// 	return response.ErrorRes(c, http.StatusBadRequest, err.Error())
	// }
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// if errs := wrapper.Validate(req); len(errs) > 0 && errs[0].Error {
	// 	errMsgs := make([]string, 0)

	// 	for _, err := range errs {
	// 		errMsgs = append(errMsgs, fmt.Sprintf(
	// 			"[%s]: '%v' | Needs to implement '%s'",
	// 			err.FailedField,
	// 			err.Value,
	// 			err.Tag,
	// 		))
	// 	}
	// 	return response.ErrorRes(c, http.StatusBadRequest, strings.Join(errMsgs, " and "))
	// }

	user, err := h.usersUsecase.CreateUser(h.cfg, ctx, req)
	if err != nil {
		return response.ErrorRes(c, http.StatusBadRequest, err.Error())
	}

	// oneWeek := time.Now().Add(7 * 24 * time.Hour)

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "accessToken",
	// 	Value:    token,
	// 	HTTPOnly: true,
	// 	Secure:   true,
	// 	Expires:  oneWeek,
	// })
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "accessChecker",
	// 	Value:    user.Id.String(),
	// 	HTTPOnly: false,
	// 	Secure:   false,
	// 	Expires:  oneWeek,
	// })

	return response.SuccessRes(c, 201, user)
}

func (h *usersHttpHandler) GetAllUsers(c *fiber.Ctx) error {

	ctx := context.Background()

	users, err := h.usersUsecase.GetAllUsers(h.cfg, ctx)
	if err != nil {
		return response.ErrorRes(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessRes(c, 201, users)
}
