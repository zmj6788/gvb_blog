package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

// MenuCreateView 添加菜单
// @Tags 菜单管理
// @Summary 添加菜单
// @Description 添加菜单
// @Param data body MenuRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuCreateView(c *gin.Context) {
	//存在bug，首次添加菜单图片关联失败后，无法再关联，只能删除该菜单重新添加
	var req MenuRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//重复判断，是否已添加
	var menuList models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", req.Title, req.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("该菜单已存在", c)
		return
	}
  //创建菜单数据入库
	menuModel := models.MenuModel{
		Title:    req.Title,
    Path:      req.Path,
		Slogan:       req.Slogan,
		Abstract:     req.Abstract,
		AbstractTime: req.AbstractTime,
		BannerTime:   req.BannerTime,
		Sort:         req.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(req.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range req.ImageSortList {
		// 这里也得判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
