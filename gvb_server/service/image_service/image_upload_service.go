package image_service

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/untils"
	"io/ioutil"
	"mime/multipart"
	"strings"

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

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	filename := file.Filename
	basepath := global.Config.Upload.Path
	filePath := basepath + filename
	res.FileName = filePath

	//获取文件后缀判读是否在白名单中，即是否允许上传
	namelist := strings.Split(filename, ".")
	//获取文件后缀并小写处理
	suffix := strings.ToLower(namelist[len(namelist)-1])
	//判断是否在白名单中
	//编写工具类便于后续使用
	legal := untils.InList(suffix, WhiteImageList)
	if !legal {
		res.IsSuccess = false
		res.Msg = "文件格式错误，请上传图片文件"
		return res
	}

	//根据图片大小判断是否存储
	filesize := float64(file.Size) / float64(1024*1024)
	if filesize > global.Config.Upload.Size {
		//图片大于5M，无法存储
		res.IsSuccess = false
		res.Msg = fmt.Sprintf("图片大于%dMB，当前文件大小为：%.2fMB，无法存储", int(global.Config.Upload.Size), filesize)
		return res
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
		res.FileName = bannerModel.Path
		res.IsSuccess = false
		res.Msg = "文件已存在，无需重复上传"
		return res
	}
	//保存文件到本地目录
	//必须放在这里确保,若文件上传七牛云成功，res.Msg更新
	res.IsSuccess = true
	res.Msg = "上传本地成功"
	fileType := ctype.Local
	//保存至本地之前判断是否上传至七牛云
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, filename, "images")
		if err != nil {
			logrus.Error(err)
			res.FileName = filename
			res.IsSuccess = false
			res.Msg = err.Error()
			return res
		}
		res.FileName = filePath
		res.IsSuccess = true
		res.Msg = "上传七牛云成功"
		fileType = ctype.QiNiu
	}

	//存储文件信息到数据库
	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      hashstr,
		Name:      filename,
		ImageType: fileType,
	})
	return res
}
