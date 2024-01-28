package orderManager

import (
	"log"
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
func (goldGoodsService *GoldGoodsService) CreateGoldGoods(goldGoods *orderManager.GoldGoods, userId int) (err error) {
	goldGoods.CreateId = &userId
	currentTime := time.Now()
	goldGoods.CreateTime = &currentTime
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
			FileType:   new(int),
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
func (goldGoodsService *GoldGoodsService) DeleteGoldGoodsByIds(ids []string, userId int) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldGoods{}, "id in ?", ids).Error
	return err
}

// UpdateGoldGoods 更新goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) UpdateGoldGoods(addGoldGoodsAndFiles orderManagerReq.AddGoldGoodsRequest, userId int) (err error) {
	log.Println("addGoldGoodsAndFiles", addGoldGoodsAndFiles.GoldGoodsFileList)
	currentTime := time.Now()
	goldGoods := &orderManager.GoldGoods{}
	global.GVA_DB.First(&goldGoods, addGoldGoodsAndFiles.ID)
	goldGoods.GoodsTypeId = addGoldGoodsAndFiles.GoodsTypeId
	goldGoods.GoodsName = addGoldGoodsAndFiles.GoodsName
	goldGoods.GoodsPrice = addGoldGoodsAndFiles.GoodsPrice
	goldGoods.UpdateId = &userId
	goldGoods.UpdateTime = &currentTime
	err = global.GVA_DB.Save(goldGoods).Error
	if err != nil {
		return err
	}
	// 获取数据库里该商品所有的工单
	allGoldGoodsFileList := []orderManager.GoldGoodsFile{}
	global.GVA_DB.Find(&allGoldGoodsFileList, "goods_id = ?", goldGoods.ID)
	// 获取数据库中所有工单的ID
	existIdList := []uint{}
	for _, element := range allGoldGoodsFileList {
		existIdList = append(existIdList, element.ID)
	}
	for _, v := range addGoldGoodsAndFiles.GoldGoodsFileList {
		if v.ID != nil {
			// 删除存在的ID
			// 删除切片中的元素
			indexToRemove := -1
			for i, value := range existIdList {
				if value == *v.ID {
					indexToRemove = i
					break
				}
			}
			existIdList = append(existIdList[:indexToRemove], existIdList[indexToRemove+1:]...)
			goldGoodsFile := &orderManager.GoldGoodsFile{}
			global.GVA_DB.First(&goldGoodsFile, v.ID)
			goldGoodsFile.GoodsId = &goldGoods.ID
			goldGoodsFile.FileName = v.FileName
			goldGoodsFile.FileType = new(int)
			goldGoodsFile.FilePath = v.FilePath
			goldGoodsFile.UpdateId = &userId
			goldGoodsFile.UpdateTime = &currentTime
			err = global.GVA_DB.Save(&goldGoodsFile).Error
			if err != nil {
				return err
			}
		} else {
			goldGoodsFile := &orderManager.GoldGoodsFile{
				GoodsId:    &goldGoods.ID,
				FileName:   v.FileName,
				FileType:   new(int),
				FilePath:   v.FilePath,
				UpdateId:   &userId,
				UpdateTime: &currentTime,
			}
			err = global.GVA_DB.Save(&goldGoodsFile).Error
			if err != nil {
				return err
			}
		}
	}
	// 删除已经存在的图片
	for _, element := range existIdList {
		err = global.GVA_DB.Delete(&orderManager.GoldGoodsFile{}, "id = ?", element).Error
		if err != nil {
			return err
		}
	}
	return err
}

// GetGoldGoods 根据id获取goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) GetGoldGoods(id string) (goldGoodsWithFiles orderManagerResp.GoldGoodsResp, err error) {
	goldGoods := orderManager.GoldGoods{}
	err = global.GVA_DB.Where("id = ?", id).First(&goldGoods).Error
	if goldGoods.CreateId == nil {
		goldGoods.CreateId = new(int)
	}
	if goldGoods.UpdateId == nil {
		goldGoods.UpdateId = new(int)
	}
	goldGoodsFileList := []orderManager.GoldGoodsFile{}
	global.GVA_DB.Where("goods_id = ?", id).Find(&goldGoodsFileList)
	log.Printf("%+v\n", goldGoodsFileList)
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
		if item.CreateId == nil {
			item.CreateId = new(int)
		}
		if item.UpdateId == nil {
			item.UpdateId = new(int)
		}
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
	log.Printf("%+v\n", goldGoodsWithFiles)
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

	err = db.Order("id desc").Find(&goldGoodss).Error
	return goldGoodss, total, err
}
