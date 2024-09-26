package tag_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// 广告上传请求参数
type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标签名"  structs:"title"` //标签名

}

// TagCreateView 添加标签
// @Tags 标签管理
// @Summary 添加标签
// @Description 添加标签
// @Param data body TagRequest    true  "表示多个参数"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {

	var req TagRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//重复判断，是否已添加
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: req.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)

}
