package article_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleCollectRequest struct {
	ID string `json:"id" `
}

// ArticleCollectCreateView文章收藏或者取消收藏接口
// @Tags 文章管理
// @Summary 文章收藏或取消
// @Description 文章收藏或取消
// @Param token header string true "用户token"
// @Param data body ArticleCollectRequest    true  "表示多个参数"
// @Router /api/articles/collect [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollectCreateView(c *gin.Context) {
	// 获取当前登录用户ID
	//通过中间件拿到解析token后的用户信息
	_claims, _ := c.Get("claims")
	//将获取的值转换为 jwts.CustomClaims 类型，并存储在变量 claims 中
	claims := _claims.(*jwts.CustomClaims)

	// fmt.Println(claims.UserID)
	// 获取文章ID
	var cr ArticleCollectRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("参数错误", c)
		return
	}
	// 查询文章是否存在
	model, err := es_service.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	// fmt.Println(cr.ID)
	// 这一张文章收藏记录是否存在，存在取消收藏，不存在收藏
	var num = -1
	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		// 没有查找到，创建
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})

		// 文章收藏数加一
		num = 1
	}
	// 查找到，删除
	// 文案收藏数减一
	global.DB.Delete(&coll)

	// 更新文章数
	// 更新文章收藏数
	err = es_service.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", c)
	} else {
		res.OkWithMessage("取消收藏成功", c)
	}
}
