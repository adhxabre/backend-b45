package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"require"`
	Email    string `json:"email" form:"email" validate:"require"`
	Password string `json:"password" form:"password" validate:"require"`
}
