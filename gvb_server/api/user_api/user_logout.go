package user_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/untils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)


// UserLogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销，删除用户数据并将注销token放入Redis
// @Param token header string true "用户token"
// @Success 200 {object} res.Response{}
// @Failure 400 {object} res.Response{msg="注销token处理失败"}
// @Failure 500 {object} res.Response{msg="用户数据清除失败"}
// @Router /api/logout [post]
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	//claims.ExpiresAt token的过期时间
	fmt.Println(claims.ExpiresAt)
	//计算距离过期的剩余时间
	exp := claims.ExpiresAt
	now := time.Now()

	diff := exp.Time.Sub(now)

	fmt.Println(diff)
	//将注销用户的token放入redis中
	prefix := "logout_"
	token := c.Request.Header.Get("token")
	err := global.Redis.Set(prefix+token, "", diff).Err()
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("注销token处理失败", c)
		return
	}
	// 清空用户数据
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Delete(&user).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("用户数据清除失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)

}
