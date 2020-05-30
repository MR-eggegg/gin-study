package service

import (
	"errors"
	"FuckingVersion1/global"
	"FuckingVersion1/model"
	"FuckingVersion1/model/request"
)

// @title    CreateAuthority
// @description   创建一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority
func CreateAuthority(auth model.AdminAuthority) (err error, authority model.AdminAuthority) {
	err = global.MyDB.Create(&auth).Error
	return err, auth
}

// @title    DeleteAuthority
// @description   删除角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// 删除角色
func DeleteAuthority(auth *model.AdminAuthority) (err error) {
	err = global.MyDB.Where("authority_id = ?", auth.AuthorityId).Find(&model.AdminUser{}).Error
	if err == nil {
		err = errors.New("此角色有用户正在使用禁止删除")
		return
	}
	err = global.MyDB.Where("parent_id = ?", auth.AuthorityId).Find(&model.AdminAuthority{}).Error
	if err == nil {
		err = errors.New("此角色存在子角色不允许删除")
		return
	}
	db := global.MyDB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth)
	if len(auth.AdminBaseMenus) > 0 {
		err = db.Association("SysBaseMenus").Delete(auth.AdminBaseMenus).Error
	} else {
		err = db.Error
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

// @title    GetInfoList
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     info            request.PaveInfo
// @return                    error
// 分页获取数据
func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MyDB
	var authority []model.AdminAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k, _ := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// @title    GetAuthorityInfo
// @description   获取所有角色信息
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @param     authority       model.SysAuthority
func GetAuthorityInfo(auth model.AdminAuthority) (err error, sa model.AdminAuthority) {
	err = global.MyDB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// @title    SetDataAuthority
// @description   设置角色资源权限
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
func SetDataAuthority(auth model.AdminAuthority) error {
	var s model.AdminAuthority
	global.MyDB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.MyDB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId).Error
	return err
}

// @title    SetMenuAuthority
// @description   菜单与角色绑定
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error
func SetMenuAuthority(auth *model.AdminAuthority) error {
	var s model.AdminAuthority
	global.MyDB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.MyDB.Model(&s).Association("SysBaseMenus").Replace(&auth.AdminBaseMenus).Error
	return err
}

// @title    findChildrenAuthority
// @description   查询子角色
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error
func findChildrenAuthority(authority *model.AdminAuthority) (err error) {
	err = global.MyDB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k, _ := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
