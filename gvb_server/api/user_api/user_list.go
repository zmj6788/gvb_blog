package user_api

import (
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/untils/desens"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
)

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param token header string  true  "token"
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	// 获取token
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMessage("未携带token", c)
		return
	}
	// 解析token
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token解析错误", c)
		return
	}
	// 不需要验证权限，在最终返回数据时，根据角色显示不同的数据

	// 查询所有用户

	var page models.PageInfo
	err = c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.UserModel{},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	// 根据用户角色过滤信息
	var users []models.UserModel
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员，过滤数据
			user.UserName = ""
		}
		// 数据脱敏
		//修改或隐藏原始数据，使得第三方可以访问和使用这些数据，
		// 而不会泄露真实的数据信息。
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	if err != nil {
		res.FailWithMessage("用户列表获取失败", c)
		return
	}
	res.OkWithList(users, count, c)
}
