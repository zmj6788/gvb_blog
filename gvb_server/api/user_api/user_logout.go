package user_api

import (
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
)

// UserLogoutView 退出登录
// @Tags 用户管理
// @Summary 退出登录
// @Description 退出登录
// @Param token header string true "用户token"
// @Router /api/logout [post]
// @Success 200 {object} res.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.Services.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)

}
