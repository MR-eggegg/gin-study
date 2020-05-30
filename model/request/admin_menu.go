package request

import "FuckingVersion1/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.AdminBaseMenu
	AuthorityId string
}

// Get role by id structure
type AuthorityIdInfo struct {
	AuthorityId string
}