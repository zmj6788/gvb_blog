package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//图片删除接口，接受一个删除请求，请求参数为idlist的json数据
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	//获得删除请求参数信息
	var rq models.RemoveRequest
	err := c.ShouldBindJSON(&rq)
	if err != nil{
		res.FailWithCode(res.ArgumentError, c)
		logrus.Info(err)
		return
	}

	var imagelist []models.BannerModel
	//获取要删除的数据到imagelist中
	count := global.DB.Find(&imagelist, rq.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	//imagelist中既是需要删除的图片数据
	global.DB.Delete(&imagelist)
	res.OkWithMessage(fmt.Sprintf("共删除%d张图片", count), c)
}
