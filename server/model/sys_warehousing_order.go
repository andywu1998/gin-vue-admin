// 自动生成模板WarehousingOrder
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type WarehousingOrder struct {
      global.GVA_MODEL
      NumbersOfCar  string `json:"numbersOfCar" form:"numbersOfCar" gorm:"column:numbers_of_car;comment:车牌号;type:varchar(20);size:20;"`
      TypeOfGoods  float64 `json:"typeOfGoods" form:"typeOfGoods" gorm:"column:type_of_goods;comment:货物类型;type:double;"`
      CompanyName  string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名字;type:varchar(50);size:50;"`
      DriverName  string `json:"driverName" form:"driverName" gorm:"column:driver_name;comment:司机名字;type:varchar(20);size:20;"`
      Driver_phone  string `json:"driver_phone" form:"driver_phone" gorm:"column:driver_phone;comment:;type:varchar(20);size:20;"`
}


