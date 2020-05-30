package response

import "FuckingVersion1/model"

type AdminUserResponse struct {
	User model.AdminUser `json:"user"`
}

type LoginResponse struct {
	User      model.AdminUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
