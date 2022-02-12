import service from '@/utils/request'

// @Tags WarehousingOrder
// @Summary 创建WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "创建WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wo/createWarehousingOrder [post]
export const createWarehousingOrder = (data) => {
     return service({
         url: "/wo/createWarehousingOrder",
         method: 'post',
         data
     })
 }


// @Tags WarehousingOrder
// @Summary 删除WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "删除WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wo/deleteWarehousingOrder [delete]
 export const deleteWarehousingOrder = (data) => {
     return service({
         url: "/wo/deleteWarehousingOrder",
         method: 'delete',
         data
     })
 }

// @Tags WarehousingOrder
// @Summary 删除WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wo/deleteWarehousingOrder [delete]
 export const deleteWarehousingOrderByIds = (data) => {
     return service({
         url: "/wo/deleteWarehousingOrderByIds",
         method: 'delete',
         data
     })
 }

// @Tags WarehousingOrder
// @Summary 更新WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "更新WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wo/updateWarehousingOrder [put]
 export const updateWarehousingOrder = (data) => {
     return service({
         url: "/wo/updateWarehousingOrder",
         method: 'put',
         data
     })
 }


// @Tags WarehousingOrder
// @Summary 用id查询WarehousingOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehousingOrder true "用id查询WarehousingOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wo/findWarehousingOrder [get]
 export const findWarehousingOrder = (params) => {
     return service({
         url: "/wo/findWarehousingOrder",
         method: 'get',
         params
     })
 }


// @Tags WarehousingOrder
// @Summary 分页获取WarehousingOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取WarehousingOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wo/getWarehousingOrderList [get]
 export const getWarehousingOrderList = (params) => {
     return service({
         url: "/wo/getWarehousingOrderList",
         method: 'get',
         params
     })
 }