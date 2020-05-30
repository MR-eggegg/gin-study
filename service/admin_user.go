package service

import (
	"FuckingVersion1/global"
	"FuckingVersion1/model"
	"FuckingVersion1/model/request"
	"FuckingVersion1/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser
func Register(u model.AdminUser) (err error, userInter model.AdminUser) {
	var user model.AdminUser
	//判断用户名是否注册
	notRegister := global.MyDB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	//notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.MyDB.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser
func Login(u *model.AdminUser) (err error, userInter *model.AdminUser) {
	var user model.AdminUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.MyDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	//err = global.MyDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}


// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser
func ChangePassword(u *model.AdminUser, newPassword string) (err error, userInter *model.AdminUser) {
	var user model.AdminUser
	//后期修改jwt+password模式
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.MyDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}


// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser
func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.AdminUser) {
	var user model.AdminUser
	err = global.MyDB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MyDB
	var userList []model.AdminUser
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error
func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.MyDB.Where("uuid = ?", uuid).First(&model.AdminUser{}).Update("authority_id", authorityId).Error
	return err
}