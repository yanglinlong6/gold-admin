import service from '@/utils/request'

// @Tags GoldShop
// @Summary 创建goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldShop true "创建goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldShop/createGoldShop [post]
export const createGoldShop = (data) => {
  return service({
    url: '/goldShop/createGoldShop',
    method: 'post',
    data
  })
}

// @Tags GoldShop
// @Summary 删除goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldShop true "删除goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldShop/deleteGoldShop [delete]
export const deleteGoldShop = (params) => {
  return service({
    url: '/goldShop/deleteGoldShop',
    method: 'delete',
    params
  })
}

// @Tags GoldShop
// @Summary 批量删除goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldShop/deleteGoldShop [delete]
export const deleteGoldShopByIds = (params) => {
  return service({
    url: '/goldShop/deleteGoldShopByIds',
    method: 'delete',
    params
  })
}

// @Tags GoldShop
// @Summary 更新goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldShop true "更新goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldShop/updateGoldShop [put]
export const updateGoldShop = (data) => {
  return service({
    url: '/goldShop/updateGoldShop',
    method: 'put',
    data
  })
}

// @Tags GoldShop
// @Summary 用id查询goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoldShop true "用id查询goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldShop/findGoldShop [get]
export const findGoldShop = (params) => {
  return service({
    url: '/goldShop/findGoldShop',
    method: 'get',
    params
  })
}

// @Tags GoldShop
// @Summary 分页获取goldShop表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goldShop表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldShop/getGoldShopList [get]
export const getGoldShopList = (params) => {
  return service({
    url: '/goldShop/getGoldShopList',
    method: 'get',
    params
  })
}
