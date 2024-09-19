package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/redis_service"
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

	//将注销用户的token放入redis中
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMessage("token不存在", c)
		return
	}

	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token错误", c)
		c.Abort()
		return
	}
	//退出登录使用双重验证
	//验证token是否在redis注销列表token中，如果在，两种情况
	//用户被更改权限，管理员强制重新登陆，or，用户已退出登录，token失效
	//验证mysql中该用户的token是否存在于redis中，仍应该可以退出登录
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = service.Services.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("token处理失败", c)
		return
	}
	token_mysql:= user.Token
	//分析：情况分为，权限更改后，用户退出，权限未更改，用户退出
	if !redis_service.CheckLogout(token_mysql) || token_mysql != token {
		res.FailWithMessage("请手动退出登录", c)
		return
	}

	res.OkWithMessage("退出成功", c)

}
