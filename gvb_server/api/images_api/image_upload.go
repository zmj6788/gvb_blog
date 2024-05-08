package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/untils"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

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

	var iml []FileUploadResponse
	//图片文件的存储
	for _, file := range fileList {
		//获取文件后缀判读是否在白名单中，即是否允许上传
		filename := file.Filename
		namelist := strings.Split(filename, ".")
		//获取文件后缀并小写处理
		suffix := strings.ToLower(namelist[len(namelist)-1])
		//判断是否在白名单中
		//编写工具类便于后续使用
		legal := untils.InList(suffix, WhiteImageList)
		if !legal {
			iml = append(iml, FileUploadResponse{
				FileName:  filename,
				IsSuccess: false,
				Msg:       "文件不合法，请上传有效图片文件",
			})
			continue
		}

		//根据图片大小判断是否存储
		filesize := float64(file.Size) / float64(1024*1024)
		if filesize > global.Config.QiNiu.Size {
			//图片大于5M，无法存储
			iml = append(iml, FileUploadResponse{
				FileName:  filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大于%dMB，当前文件大小为：%.2fMB，无法存储", int(global.Config.QiNiu.Size), filesize),
			})
			continue

		}
		//获取文件内容加密hash值
		fileObj, err := file.Open()
		if err != nil {
			logrus.Error(err)
		}
		byteData, err := ioutil.ReadAll(fileObj)
		if err != nil {
			logrus.Error(err)
		}
		hashstr := untils.MD5(byteData)
		//判断文件是否存在数据库中
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", hashstr).Error
		if err == nil {
			//查找到该文件
			iml = append(iml, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "文件已存在，无需重复上传",
			})
			continue
		}
		//保存文件到本地目录
		filePath := basepath + filename
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			logrus.Error(err)
			iml = append(iml, FileUploadResponse{
				FileName:  filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			return
		}

		iml = append(iml, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
		//存储文件信息到数据库
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: hashstr,
			Name: filename,
		})

	}
	res.Ok(iml, "响应成功", c)
}
