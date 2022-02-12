package request

import "gin-vue-admin/model"

type WarehousingOrderSearch struct{
    model.WarehousingOrder
    PageInfo
}