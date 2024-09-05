package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"github.com/gin-gonic/gin"
)

// 广告上传请求参数
type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"  structs:"title"`        //广告标题
	Href   string `json:"href" binding:"required,url" msg:"请输入跳转链接"  structs:"href"`   //跳转链接
	Images string `json:"images" binding:"required,url" msg:"请输入图片地址"  structs:"images"` //广告图片
	IsShow bool   `json:"is_show" structs:"is_show"`    //是否显示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {

	var req AdvertRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//重复判断，是否已添加
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  req.Title,
		Href:   req.Href,
		Images: req.Images,
		IsShow: req.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)

}
