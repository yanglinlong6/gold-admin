<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品Id:" prop="goodsId">
          <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="购买人Id:" prop="buyId">
          <el-input v-model.number="formData.buyId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人Id:" prop="createId">
          <el-input v-model.number="formData.createId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新人Id:" prop="updateId">
          <el-input v-model.number="formData.updateId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createGoldShop,
  updateGoldShop,
  findGoldShop
} from '@/api/goldShop'

defineOptions({
    name: 'GoldShopForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            goodsId: 0,
            buyId: 0,
            createId: 0,
            createTime: new Date(),
            updateId: 0,
            updateTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGoldShop({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.regoldShop
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createGoldShop(formData.value)
               break
             case 'update':
               res = await updateGoldShop(formData.value)
               break
             default:
               res = await createGoldShop(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
