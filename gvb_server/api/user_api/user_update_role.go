package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更,昵称变更
// @Tags 用户管理
// @Summary 用户权限变更,昵称变更
// @Description 用户权限变更,昵称变更
// @Param token header string  true  "token"
// @Param data body UserRole    true  "用户的一些参数"
// @Router /api/user_role [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户id错误，用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	// //将用户当前token失效，使得用户强制重新登陆更新权限
	// token := user.Token
	// if token != "" {
	// 	//避免注册后从未登录的用户权限修改问题
	// 	claims, err := jwts.ParseToken(token)
	// 	if err != nil {
	// 		res.FailWithMessage("token错误", c)
	// 		c.Abort()
	// 		return
	// 	}
	// 	err = service.Services.UserService.Logout(claims, token)
	// 	if err != nil {
	// 		res.FailWithMessage("token注销失败", c)
	// 		return
	// 	}
	// }

	res.OkWithMessage("修改权限成功", c)
}
