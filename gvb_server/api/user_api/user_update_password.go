package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/untils/jwts"
	"gvb_server/untils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}

// UserUpdateRoleView 用户密码变更
// @Tags 用户管理
// @Summary 用户密码变更
// @Description 用户密码变更
// @Param token header string  true  "token"
// @Param data body UpdatePasswordRequest    true  "用户的一些参数"
// @Router /api/user_pwd [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("旧密码错误", c)
		return
	}
	//新密码是否符合规则
	if cr.Pwd == "" {
		res.FailWithMessage("新密码不能为空", c)
		return
	}
	if cr.Pwd == cr.OldPwd {
		res.FailWithMessage("新旧密码不能相同", c)
		return
	}
	
	//加密新密码
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
}
