import service from '@/utils/request'

// @Tags GoldGoodsFile
// @Summary 创建goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoodsFile true "创建goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldGoodsFile/createGoldGoodsFile [post]
export const createGoldGoodsFile = (data) => {
  return service({
    url: '/goldGoodsFile/createGoldGoodsFile',
    method: 'post',
    data
  })
}

// @Tags GoldGoodsFile
// @Summary 删除goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoodsFile true "删除goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoodsFile/deleteGoldGoodsFile [delete]
export const deleteGoldGoodsFile = (params) => {
  return service({
    url: '/goldGoodsFile/deleteGoldGoodsFile',
    method: 'delete',
    params
  })
}

// @Tags GoldGoodsFile
// @Summary 批量删除goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoodsFile/deleteGoldGoodsFile [delete]
export const deleteGoldGoodsFileByIds = (params) => {
  return service({
    url: '/goldGoodsFile/deleteGoldGoodsFileByIds',
    method: 'delete',
    params
  })
}

// @Tags GoldGoodsFile
// @Summary 更新goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldGoodsFile true "更新goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldGoodsFile/updateGoldGoodsFile [put]
export const updateGoldGoodsFile = (data) => {
  return service({
    url: '/goldGoodsFile/updateGoldGoodsFile',
    method: 'put',
    data
  })
}

// @Tags GoldGoodsFile
// @Summary 用id查询goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoldGoodsFile true "用id查询goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldGoodsFile/findGoldGoodsFile [get]
export const findGoldGoodsFile = (params) => {
  return service({
    url: '/goldGoodsFile/findGoldGoodsFile',
    method: 'get',
    params
  })
}

// @Tags GoldGoodsFile
// @Summary 分页获取goldGoodsFile表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goldGoodsFile表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldGoodsFile/getGoldGoodsFileList [get]
export const getGoldGoodsFileList = (params) => {
  return service({
    url: '/goldGoodsFile/getGoldGoodsFileList',
    method: 'get',
    params
  })
}
