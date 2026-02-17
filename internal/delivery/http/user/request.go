package user

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ChangeRoleRequest struct {
	UserID int `json:"user_id" binding:"required"`
	RoleID int `json:"role_id" binding:"required"`
}

type FindByIdRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

type FindByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
}

type GetUserInfoRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=20"`
	Email	 string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"required,min=8"`
}
