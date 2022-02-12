package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWarehousingOrderRouter(Router *gin.RouterGroup) {
	WarehousingOrderRouter := Router.Group("wo").Use(middleware.OperationRecord())
	{
		WarehousingOrderRouter.POST("createWarehousingOrder", v1.CreateWarehousingOrder)   // 新建WarehousingOrder
		WarehousingOrderRouter.DELETE("deleteWarehousingOrder", v1.DeleteWarehousingOrder) // 删除WarehousingOrder
		WarehousingOrderRouter.DELETE("deleteWarehousingOrderByIds", v1.DeleteWarehousingOrderByIds) // 批量删除WarehousingOrder
		WarehousingOrderRouter.PUT("updateWarehousingOrder", v1.UpdateWarehousingOrder)    // 更新WarehousingOrder
		WarehousingOrderRouter.GET("findWarehousingOrder", v1.FindWarehousingOrder)        // 根据ID获取WarehousingOrder
		WarehousingOrderRouter.GET("getWarehousingOrderList", v1.GetWarehousingOrderList)  // 获取WarehousingOrder列表
	}
}
