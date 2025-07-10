package request

import (
	model "go_auth/databases/models"
	pagination "go_auth/lib"
)

type (
	ListUserRequest struct {
		pagination.PaginationClientRequest
		Username  *string `json:"username"        validate:"omitempty"`
		Firstname *string `json:"firstname" validate:"omitempty"`
		Lastname  *int32  `json:"lastname"     validate:"omitempty"`
		model.Users
	}

	CreateUserRequest struct {
		Username  string `json:"username"                      validate:"required"`
		Password  string `json:"password"                validate:"required"`
		Firstname string `json:"firstname"                   validate:"required"`
		Lastname  string `json:"lastname"                   validate:"required"`
	}

	LoginRequest struct {
		Username string `json:"username"                      validate:"required"`
		Password string `json:"password"                validate:"required"`
	}

	LogoutRequest struct {
		id int32 `query:"id" validate:"required"`
	}
)
