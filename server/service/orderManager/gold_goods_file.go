package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
    orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
)

type GoldGoodsFileService struct {
}

// CreateGoldGoodsFile 创建goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService) CreateGoldGoodsFile(goldGoodsFile *orderManager.GoldGoodsFile) (err error) {
	err = global.GVA_DB.Create(goldGoodsFile).Error
	return err
}

// DeleteGoldGoodsFile 删除goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService)DeleteGoldGoodsFile(id string) (err error) {
	err = global.GVA_DB.Delete(&orderManager.GoldGoodsFile{},"id = ?",id).Error
	return err
}

// DeleteGoldGoodsFileByIds 批量删除goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService)DeleteGoldGoodsFileByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldGoodsFile{},"id in ?",ids).Error
	return err
}

// UpdateGoldGoodsFile 更新goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService)UpdateGoldGoodsFile(goldGoodsFile orderManager.GoldGoodsFile) (err error) {
	err = global.GVA_DB.Save(&goldGoodsFile).Error
	return err
}

// GetGoldGoodsFile 根据id获取goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService)GetGoldGoodsFile(id string) (goldGoodsFile orderManager.GoldGoodsFile, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goldGoodsFile).Error
	return
}

// GetGoldGoodsFileInfoList 分页获取goldGoodsFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsFileService *GoldGoodsFileService)GetGoldGoodsFileInfoList(info orderManagerReq.GoldGoodsFileSearch) (list []orderManager.GoldGoodsFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&orderManager.GoldGoodsFile{})
    var goldGoodsFiles []orderManager.GoldGoodsFile
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&goldGoodsFiles).Error
	return  goldGoodsFiles, total, err
}
