package models

import (
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片类型，本地还是七牛云
}

// 钩子函数，删除数据库数据之前，删除图片文件
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	
	if b.ImageType == ctype.Local {
		//本地存储，删除数据库数据之前，删除本地图片文件
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
