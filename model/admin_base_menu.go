package model

import (
	"github.com/jinzhu/gorm"
)

type AdminBaseMenu struct {
	gorm.Model
	MenuLevel       uint   `json:"-"`
	ParentId        string `json:"parentId"`
	Path            string `json:"path"`
	Name            string `json:"name"`
	Hidden          bool   `json:"hidden"`
	Component       string `json:"component"`
	Sort            int    `json:"sort"`
	Meta            `json:"meta"`
	AdminAuthoritys []AdminAuthority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children        []AdminBaseMenu  `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}