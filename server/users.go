package server

import (
	usershandler "tgrzimiar/go-scylla/modules/users/usersHandler"
	usersrepository "tgrzimiar/go-scylla/modules/users/usersRepository"
	usersusecase "tgrzimiar/go-scylla/modules/users/usersUsecase"
)

func (s *server) usersService() {
	usersRepo := usersrepository.NewUserRepository(s.db, s.logger)
	usersUsecase := usersusecase.NewUsersUsecase(usersRepo)
	usersHttpHandler := usershandler.NewUsersHttpHandler(s.cfg, usersUsecase)

	users := s.app.Group("/users_v1")
	users.Post("/register", usersHttpHandler.CreateUser)
	users.Get("/get-users", usersHttpHandler.GetAllUsers)
}
