package v1

import (
	"FuckingVersion1/global"
	Msg "FuckingVersion1/global/msg"
	"FuckingVersion1/middleware"
	"FuckingVersion1/model"
	"FuckingVersion1/model/request"
	"FuckingVersion1/model/response"
	"FuckingVersion1/service"
	"FuckingVersion1/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)


// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := &model.AdminUser{
		Username: R.Username,
		NickName: R.NickName,
		Password: R.Password,
		HeaderImg: R.HeaderImg,
		AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		Msg.FailWithDetailed(Msg.ERROR, response.AdminUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		Msg.OkDetailed(response.AdminUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body model.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	U := &model.AdminUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		Msg.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
	} else {
		tokenNext(c, *user)
	}
}


// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ChangePasswordStutrc true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var params request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&params)
	U := &model.AdminUser{Username: params.Username, Password: params.Password}
	if err, _ := service.ChangePassword(U, params.NewPassword); err != nil {
		Msg.FailWithMessage("修改失败，请检查用户名密码", c)
	} else {
		Msg.OkWithMessage("修改成功", c)
	}
}

const (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	USER_HEADER_BUCKET   string = "qm-plus-img"
	ImageSavePath string = "images/"
)


// @Tags AdminUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	//获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*request.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	//便于找到用户 以后从jwt中取
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		//文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			Msg.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := service.UploadHeaderImg(uuid, filePath)
			if err != nil {
				Msg.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				Msg.OkWithData(response.AdminUserResponse{User: *user}, c)
			}
		}
	}
}

// @Tags SysUser
// @Summary 用户上传头像
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/UploadImage [post]
func UploadImage(c *gin.Context) {
	data := make(map[string]string)
	if ok, _ := utils.PathExists(ImageSavePath); !ok {
		// directory not exist
		fmt.Println("create image directory")
		_ = os.Mkdir(ImageSavePath, os.ModePerm)
	}
	_, image, err := c.Request.FormFile("headerImg")
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("读取失败，%v", err), c)
	}
	if image == nil {
		Msg.FailWithMessage(fmt.Sprintf("未找到文件，%v", err), c)
	} else {
		imageName := func (name string) string {
			ext := path.Ext(name)
			fileName := strings.TrimSuffix(name, ext)
			return fileName + ext
		}(image.Filename)
		src := ImageSavePath + imageName
		if err := c.SaveUploadedFile(image, src); err != nil {
			Msg.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
		} else {
			data["image_save_url"] = ImageSavePath + imageName
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

//登录以后签发jwt
func tokenNext(c *gin.Context, user model.AdminUser) {
	j := &middleware.JWT{
		[]byte(global.MyConfig.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "qmPlus",                              //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		Msg.FailWithMessage("获取token失败", c)
	} else {
		if global.MyConfig.System.UseMultipoint {
			var loginJwt model.JwtBlacklist
			loginJwt.Jwt = token
			err, jwtStr := service.GetRedisJWT(user.Username)
			if err == redis.Nil {
				err2 := service.SetRedisJWT(loginJwt, user.Username)
				if err2 != nil {
					Msg.FailWithMessage("设置登录状态失败", c)
				} else {
					Msg.OkWithData(response.LoginResponse{
						User:      user,
						Token:     token,
						ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
					}, c)
				}
			} else if err != nil {
				Msg.FailWithMessage(fmt.Sprintf("%v", err), c)
			} else {
				var blackJWT model.JwtBlacklist
				blackJWT.Jwt = jwtStr
				err3 := service.JsonInBlacklist(blackJWT)
				if err3 != nil {
					Msg.FailWithMessage("jwt作废失败", c)
				} else {
					err2 := service.SetRedisJWT(loginJwt, user.Username)
					if err2 != nil {
						Msg.FailWithMessage("设置登录状态失败", c)
					} else {
						Msg.OkWithData(response.LoginResponse{
							User:      user,
							Token:     token,
							ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
						}, c)
					}
				}
			}
		} else {
			Msg.OkWithData(response.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			}, c)
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		Msg.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	err := service.SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		Msg.FailWithMessage(fmt.Sprintf("修改失败，%v", err), c)
	} else {
		Msg.OkWithMessage("修改成功", c)
	}
}
