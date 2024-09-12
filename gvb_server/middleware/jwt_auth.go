package middleware

import (
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/untils"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
)

//使用中间件后的接口
//可以理解为只有登录的用户才能调用绑定用户登录中间件的接口
// 验证用户登录状态的中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("token不存在", c)
			// Abort() 方法会停止后续中间件或路由处理器的执行，
			// 并立即返回响应给客户端。主要用于错误处理或提前完成请求。
			c.Abort()
			return
		}
		
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}

		//验证token是否在redis注销列表token中
		prefix := "logout_"
		keys := global.Redis.Keys(prefix + "*").Val()
		global.Log.Info(keys)
		if untils.InList(prefix + token, keys) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		//验证成功
		//已经登陆的用户
		c.Set("claims", claims)
	}
}

//只有登录的用户是管理员才能调用绑定管理员登录中间件的接口
//管理员使用的验证登陆状态的中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("token不存在", c)
			c.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}

		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
		//验证成功
		//已经登陆的用户
		c.Set("claims", claims)
	}
}
