package images_api

import (
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_service"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	// 规定图片上传白名单
	WhiteImageList = []string{
		"pjp",
		"svgz",
		"jpg",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"jfif",
		"webp",
		"png",
		"bmp",
		"pjpeg",
		"avif",
	}
)

// 规定文件上传的响应格式，便于客户端解析上传结果信息
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 提示信息
}

// ImageUploodView上传到个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//获取单个上传的文件
	//FileHeader, err := c.FormFile("image")

	//获取多个上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}
	//判断图片文件存储路径是否存在
	//不存在创建
	basepath := global.Config.Upload.Path
	_, err = os.ReadDir(basepath)
	if err != nil {
		err := os.MkdirAll(basepath, fs.ModePerm)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			logrus.Error(err)
			return
		}
	}
	//遍历文件，做上传文件操作
	var resList []image_service.FileUploadResponse
	for _, file := range fileList {

		// 上传文件
		serviceRes := service.Services.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功的
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			//上传成功但是七牛云存储没有开启，说明是本地存储成功，保存一下
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}

	res.Ok(resList, "响应成功", c)
	
}
