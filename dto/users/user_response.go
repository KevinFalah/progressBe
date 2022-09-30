package usersdto

import (
	postdto "waysgallery/dto/post"
)

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserDetailResponse struct {
	ID       int                    `json:"id"`
	FullName string                 `json:"fullName" form:"name" validate:"required"`
	Email    string                 `json:"email" form:"email" validate:"required"`
	Password string                 `json:"password" form:"password" validate:"required"`
	Posts    []postdto.PostResponse `json:"posts"`
}
