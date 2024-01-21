package users

type (
	CreateUserReq struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserRes struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
