package initialize

import (
	"FuckingVersion1/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.MyConfig.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.MyLog.Error("DEFAULT DB DATABASE START FAIL,err:", err)
	} else {
		global.MyDB = db
		global.MyDB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.MyDB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.MyDB.LogMode(admin.LogMode)
	}
}

