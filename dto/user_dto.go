package dto

import "errors"

const (
	// SUCCESS
	MESSAGE_SUCCESS_REGISTER_USER     = "success add user"
	MESSAGE_SUCCESS_LOGIN             = "success login user"
	MESSAGE_SUCCESS_FETCH_USERS       = "success to fetch users"
	MESSAGE_SUCCESS_VERIFY_EMAIL_USER = "success to verify user email verification"
	MESSAGE_SUCCESS_FIND_USER         = "success find user"
	MESSAGE_SUCCESS_UPDATE_USER       = "success to update user"
	MESSAGE_SUCCESS_DELETE_USER       = "success to delete user"

	// FAILED
	MESSAGE_FAILED_REGISTER_USER     = "failed add user"
	MESSAGE_FAILED_LOGIN             = "failed login user"
	MESSAGE_FAILED_FETCH_USERS       = "failed to fetch users"
	MESSAGE_FAILED_VERIFY_EMAIL_USER = "failed to verify user email verification"
	MESSAGE_FAILED_FIND_USER         = "failed find user"
	MESSAGE_FAILED_UPDATE_USER       = "failed to update user"
	MESSAGE_FAILED_DELETE_USER       = "failed to delete user"
)

var (
	ErrHashPass           = errors.New("failed to hash password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreateUser         = errors.New("failed create user to db")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("invalid crecentials")
	ErrFailedCreateToken  = errors.New("failed to create token")
)

type (
	RegisterUserRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	RegisterUserResponse struct {
		Username string `json:"username" form:"username"`
		Email    string `json:"email" form:"email"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token" form:"token"`
		Role  string `json:"role" form:"role"`
	}
)
