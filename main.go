package main

import (
	"FuckingVersion1/core"
	"FuckingVersion1/global"
	"FuckingVersion1/initialize"
)

func main () {
	switch global.MyConfig.System.DbType {
	case "mysql":
		initialize.Mysql()
	//case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}

	initialize.DBTables()

	defer global.MyDB.Close()

	core.RunWindowsServer()
}
