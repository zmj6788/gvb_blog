package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// 图片更新请求
type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择有效文件id"`
	Name string `json:"name" binding:"required" msg:"请输入修改后的文件名称"`
}

// ImageUpdateView 更新图片
// @Tags 图片管理
// @Summary 更新图片
// @Param token header string  true  "token"
// @Description 更新图片
// @Param data body ImageUpdateRequest    true  "图片的一些参数"
// @Router /api/images [put]
// @Produce json
// @Success 200 {object} res.Response{}
// 图片更新接口函数
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var req ImageUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		//响应提示信息
		res.FailWithError(err, &req, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, req.ID).Error
	if err != nil {
		res.FailWithMessage("图片不存在", c)
		return
	}
	// global.DB.Model(&imageModel).Updates(models.BannerModel{
	// 	Name: req.Name,
	// })
	//修改表数据的两种方法
	err = global.DB.Model(&imageModel).Update("name", req.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
}
