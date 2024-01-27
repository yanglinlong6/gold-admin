package orderManager

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
	orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
	orderManagerResp "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/response"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type GoldGoodsService struct {
}

// CreateGoldGoods 创建goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) CreateGoldGoods(goldGoods *orderManager.GoldGoods) (err error) {
	err = global.GVA_DB.Create(goldGoods).Error
	return err
}

func (goldGoodsService *GoldGoodsService) AddGoldGoodsAndFiles(addGoldGoodsAndFiles *orderManagerReq.AddGoldGoodsRequest, userId int) (err error) {
	currentTime := time.Now()
	goldGoods := &orderManager.GoldGoods{
		GoodsTypeId: addGoldGoodsAndFiles.GoodsTypeId,
		GoodsName:   addGoldGoodsAndFiles.GoodsName,
		GoodsPrice:  addGoldGoodsAndFiles.GoodsPrice,
		CreateId:    &userId,
		CreateTime:  &currentTime,
	}
	err = global.GVA_DB.Create(goldGoods).Error
	if err != nil {
		return err
	}

	for _, v := range addGoldGoodsAndFiles.GoldGoodsFileList {
		goldGoodsFile := &orderManager.GoldGoodsFile{
			GoodsId:    &goldGoods.ID,
			FileName:   v.FileName,
			FileType:   v.FileType,
			FilePath:   v.FilePath,
			CreateId:   &userId,
			CreateTime: &currentTime,
		}

		err = global.GVA_DB.Create(goldGoodsFile).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteGoldGoods 删除goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) DeleteGoldGoods(id string) (err error) {
	err = global.GVA_DB.Delete(&orderManager.GoldGoods{}, "id = ?", id).Error
	return err
}

// DeleteGoldGoodsByIds 批量删除goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) DeleteGoldGoodsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldGoods{}, "id in ?", ids).Error
	return err
}

// UpdateGoldGoods 更新goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) UpdateGoldGoods(goldGoods orderManager.GoldGoods) (err error) {
	err = global.GVA_DB.Save(&goldGoods).Error
	return err
}

// GetGoldGoods 根据id获取goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) GetGoldGoods(id string) (goldGoodsWithFiles orderManagerResp.GoldGoodsResp, err error) {
	goldGoods := orderManager.GoldGoods{}
	err = global.GVA_DB.Where("id = ?", id).First(&goldGoods).Error
	goldGoodsFileList := []orderManager.GoldGoodsFile{}
	global.GVA_DB.Where("goods_id = ?", id).Find(&goldGoodsFileList)
	sysUserList := []systemModel.SysUser{}
	global.GVA_DB.Find(&sysUserList)
	sysUserMap := make(map[uint]systemModel.SysUser)
	for _, element := range sysUserList {
		sysUserMap[element.ID] = element
	}
	goldGoodsWithFiles = orderManagerResp.GoldGoodsResp{
		ID:          goldGoods.ID,
		GoodsTypeId: goldGoods.GoodsTypeId,
		GoodsName:   goldGoods.GoodsName,
		GoodsPrice:  goldGoods.GoodsPrice,
		CreateName:  sysUserMap[uint(*goldGoods.CreateId)].Username,
		CreateTime:  goldGoods.CreateTime,
		UpdateName:  sysUserMap[uint(*goldGoods.UpdateId)].Username,
		UpdateTime:  goldGoods.UpdateTime,
	}
	for _, item := range goldGoodsFileList {
		resp := orderManagerResp.GoldGoodsFileResp{
			ID:         item.ID,
			GoodsId:    item.GoodsId,
			FileName:   item.FileName,
			FilePath:   item.FilePath,
			CreateName: sysUserMap[uint(*item.CreateId)].Username,
			CreateTime: item.CreateTime,
			UpdateName: sysUserMap[uint(*item.UpdateId)].Username,
			UpdateTime: item.UpdateTime,
		}
		goldGoodsWithFiles.GoldGoodsFileList = append(goldGoodsWithFiles.GoldGoodsFileList, resp)
	}
	return goldGoodsWithFiles, err
}

// GetGoldGoodsInfoList 分页获取goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) GetGoldGoodsInfoList(info orderManagerReq.GoldGoodsSearch) (list []orderManager.GoldGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&orderManager.GoldGoods{})
	var goldGoodss []orderManager.GoldGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&goldGoodss).Error
	return goldGoodss, total, err
}
