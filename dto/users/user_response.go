package usersdto

import (
	postdto "waysgallery/dto/post"
)

type UserResponse struct {
	ID       int                    `json:"id"`
	FullName string                 `json:"fullName" form:"name" validate:"required"`
	Email    string                 `json:"email" form:"email" validate:"required"`
	Password string                 `json:"password" form:"password" validate:"required"`
	Greeting string                 `json:"greeting" form:"greeting"`
	Avatar   string                 `json:"avatar" form:"avatar"`
	Posts    []postdto.PostResponse `json:"posts"`
}

type UserDetailResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Greeting string `json:"greeting" form:"greeting"`
	Avatar   string `json:"avatar" form:"avatar"`

	Posts []postdto.PostResponse `json:"posts"`
}
