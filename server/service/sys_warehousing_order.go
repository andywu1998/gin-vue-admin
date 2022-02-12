package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateWarehousingOrder
//@description: 创建WarehousingOrder记录
//@param: wo model.WarehousingOrder
//@return: err error

func CreateWarehousingOrder(wo model.WarehousingOrder) (err error) {
	err = global.GVA_DB.Create(&wo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteWarehousingOrder
//@description: 删除WarehousingOrder记录
//@param: wo model.WarehousingOrder
//@return: err error

func DeleteWarehousingOrder(wo model.WarehousingOrder) (err error) {
	err = global.GVA_DB.Delete(&wo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteWarehousingOrderByIds
//@description: 批量删除WarehousingOrder记录
//@param: ids request.IdsReq
//@return: err error

func DeleteWarehousingOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.WarehousingOrder{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateWarehousingOrder
//@description: 更新WarehousingOrder记录
//@param: wo *model.WarehousingOrder
//@return: err error

func UpdateWarehousingOrder(wo model.WarehousingOrder) (err error) {
	err = global.GVA_DB.Save(&wo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetWarehousingOrder
//@description: 根据id获取WarehousingOrder记录
//@param: id uint
//@return: err error, wo model.WarehousingOrder

func GetWarehousingOrder(id uint) (err error, wo model.WarehousingOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&wo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetWarehousingOrderInfoList
//@description: 分页获取WarehousingOrder记录
//@param: info request.WarehousingOrderSearch
//@return: err error, list interface{}, total int64

func GetWarehousingOrderInfoList(info request.WarehousingOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.WarehousingOrder{})
    var wos []model.WarehousingOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.NumbersOfCar != "" {
        db = db.Where("`numbers_of_car` LIKE ?","%"+ info.NumbersOfCar+"%")
    }
    if info.TypeOfGoods != 0 {
        db = db.Where("`type_of_goods` <> ?",info.TypeOfGoods)
    }
    if info.CompanyName != "" {
        db = db.Where("`company_name` LIKE ?","%"+ info.CompanyName+"%")
    }
    if info.DriverName != "" {
        db = db.Where("`driver_name` LIKE ?","%"+ info.DriverName+"%")
    }
    if info.Driver_phone != "" {
        db = db.Where("`driver_phone` LIKE ?","%"+ info.Driver_phone+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&wos).Error
	return err, wos, total
}