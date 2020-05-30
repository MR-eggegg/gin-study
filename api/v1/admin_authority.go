package v1

import (
	Msg "FuckingVersion1/global/msg"
	"fmt"
	"FuckingVersion1/model"
	"FuckingVersion1/model/request"
	resp "FuckingVersion1/model/response"
	"FuckingVersion1/service"
	"github.com/gin-gonic/gin"
)

// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "创建角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var auth model.AdminAuthority
	_ = c.ShouldBindJSON(&auth)
	err, authBack := service.CreateAuthority(auth)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		Msg.OkWithData(resp.AdminAuthorityResponse{Authority: authBack}, c)
	}
}

// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var a model.AdminAuthority
	_ = c.ShouldBindJSON(&a)
	//删除角色之前需要判断是否有用户正在使用此角色
	err := service.DeleteAuthority(&a)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		Msg.OkWithMessage("删除成功", c)
	}
}

// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetAuthorityInfoList(pageInfo)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		Msg.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func SetDataAuthority(c *gin.Context) {
	var auth model.AdminAuthority
	_ = c.ShouldBindJSON(&auth)
	err := service.SetDataAuthority(auth)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("设置关联失败，%v", err), c)
	} else {
		Msg.Ok(c)
	}
}
