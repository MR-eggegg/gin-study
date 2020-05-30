package initialize

import (
	"FuckingVersion1/global"
	"FuckingVersion1/model"
)

//注册数据库表专用
func DBTables() {
	db := global.MyDB
	db.AutoMigrate(
		model.AdminUser{},
		model.AdminAuthority{},
		model.JwtBlacklist{},
		model.AdminBaseMenu{},
	)
	global.MyLog.Debug("register table success")
}
