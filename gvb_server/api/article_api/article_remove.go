package article_api

import (
	"context"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type ArticleRemoveRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleRemoveView 批量删除文章
// @Tags 文章管理
// @Summary 批量删除文章
// @Description 批量删除文章
// @Param data body ArticleRemoveRequest    true  "文章id列表"
// @Router /api/articles [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleRemoveView(c *gin.Context) {

	// 获取要删除文章的id列表
	var cr ArticleRemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//运用bulk进行批量删除
	// 创建一个Bulk服务
	bulkRequest := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")

	// 遍历ID列表，为每个ID添加一个删除操作
	// 如果文章删除了，用户收藏了这个文章怎么办
	// 顺带把这个文章关联的收藏的数据也删除了


	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkRequest.Add(req)
		go es_service.DeleteFullTextByArticleID(id)
	}

	// 执行Bulk请求
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("删除文章失败", c)
		return
	}
	
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(bulkResponse.Succeeded())), c)
}
