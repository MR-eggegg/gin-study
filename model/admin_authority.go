package model

import (
	"time"
)

type AdminAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key"`
	AuthorityName   string         `json:"authorityName"`
	ParentId        string         `json:"parentId"`
	DataAuthorityId []AdminAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;association_jointable_foreignkey:data_authority_id"`
	Children        []AdminAuthority `json:"children"`
	AdminBaseMenus  []AdminBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}