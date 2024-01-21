import service from '@/utils/request'

// @Tags GoldGoods
// @Summary 创建goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoods true "创建goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldGoods/createGoldGoods [post]
export const createGoldGoods = (data) => {
  return service({
    url: '/goldGoods/createGoldGoods',
    method: 'post',
    data
  })
}

// @Tags GoldGoods
// @Summary 删除goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoods true "删除goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoods/deleteGoldGoods [delete]
export const deleteGoldGoods = (params) => {
  return service({
    url: '/goldGoods/deleteGoldGoods',
    method: 'delete',
    params
  })
}

// @Tags GoldGoods
// @Summary 批量删除goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoods/deleteGoldGoods [delete]
export const deleteGoldGoodsByIds = (params) => {
  return service({
    url: '/goldGoods/deleteGoldGoodsByIds',
    method: 'delete',
    params
  })
}

// @Tags GoldGoods
// @Summary 更新goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoods true "更新goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldGoods/updateGoldGoods [put]
export const updateGoldGoods = (data) => {
  return service({
    url: '/goldGoods/updateGoldGoods',
    method: 'put',
    data
  })
}

// @Tags GoldGoods
// @Summary 用id查询goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoldGoods true "用id查询goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldGoods/findGoldGoods [get]
export const findGoldGoods = (params) => {
  return service({
    url: '/goldGoods/findGoldGoods',
    method: 'get',
    params
  })
}

// @Tags GoldGoods
// @Summary 分页获取goldGoods表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goldGoods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldGoods/getGoldGoodsList [get]
export const getGoldGoodsList = (params) => {
  return service({
    url: '/goldGoods/getGoldGoodsList',
    method: 'get',
    params
  })
}
