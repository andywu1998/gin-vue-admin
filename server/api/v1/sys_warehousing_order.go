package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags WarehousingOrder
// @Summary 创建WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "创建WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wo/createWarehousingOrder [post]
func CreateWarehousingOrder(c *gin.Context) {
	var wo model.WarehousingOrder
	_ = c.ShouldBindJSON(&wo)
	if err := service.CreateWarehousingOrder(wo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags WarehousingOrder
// @Summary 删除WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "删除WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wo/deleteWarehousingOrder [delete]
func DeleteWarehousingOrder(c *gin.Context) {
	var wo model.WarehousingOrder
	_ = c.ShouldBindJSON(&wo)
	if err := service.DeleteWarehousingOrder(wo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags WarehousingOrder
// @Summary 批量删除WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wo/deleteWarehousingOrderByIds [delete]
func DeleteWarehousingOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteWarehousingOrderByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags WarehousingOrder
// @Summary 更新WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "更新WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wo/updateWarehousingOrder [put]
func UpdateWarehousingOrder(c *gin.Context) {
	var wo model.WarehousingOrder
	_ = c.ShouldBindJSON(&wo)
	if err := service.UpdateWarehousingOrder(wo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags WarehousingOrder
// @Summary 用id查询WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "用id查询WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wo/findWarehousingOrder [get]
func FindWarehousingOrder(c *gin.Context) {
	var wo model.WarehousingOrder
	_ = c.ShouldBindQuery(&wo)
	if err, rewo := service.GetWarehousingOrder(wo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewo": rewo}, c)
	}
}

// @Tags WarehousingOrder
// @Summary 分页获取WarehousingOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WarehousingOrderSearch true "分页获取WarehousingOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wo/getWarehousingOrderList [get]
func GetWarehousingOrderList(c *gin.Context) {
	var pageInfo request.WarehousingOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetWarehousingOrderInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
