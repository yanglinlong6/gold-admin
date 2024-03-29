<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="商品分类" width="120">
          <template #default="scope">{{
            convertGoodsTypeId(scope.row.goodsTypeId)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="商品名称"
          prop="goodsName"
          width="120"
        />
        <el-table-column
          align="left"
          label="商品价格"
          prop="goodsPrice"
          width="120"
        />
        <el-table-column
          align="left"
          label="商品创建人"
          prop="createName"
          width="120"
        />
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.createTime)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="商品更新人"
          prop="updateName"
          width="120"
        />
        <el-table-column align="left" label="更新时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.updateTime)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="120">
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateGoldGoodsFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          :model="formData"
          label-position="right"
          ref="elFormRef"
          :rules="rule"
          label-width="100px"
        >
          <el-form-item label="商品分类:" prop="goodsTypeId">
            <el-select
              v-model.number="formData.goodsTypeId"
              class="m-2"
              placeholder="请选择商品分类"
              size="small"
              :clearable="true"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="商品名称:" prop="goodsName">
            <el-input
              v-model="formData.goodsName"
              :clearable="true"
              placeholder="请输入商品名称"
            />
          </el-form-item>
          <el-form-item label="商品价格:" prop="goodsPrice">
            <el-input
              v-model.number="formData.goodsPrice"
              :clearable="true"
              placeholder="请输入商品价格"
            />
          </el-form-item>
          <el-form-item label="上传图片">
            <div>
              <el-upload
                v-model:file-list="fileList"
                :action="`${path}/fileUploadAndDownload/upload`"
                list-type="picture-card"
                :on-preview="handlePictureCardPreview"
                :on-remove="handleRemove"
                :on-success="uploadSuccess"
                :on-error="uploadError"
                :show-file-list="true"
              >
                <el-icon><Plus /></el-icon>
              </el-upload>

              <el-dialog v-model="dialogVisible">
                <img w-full :src="dialogImageUrl" alt="Preview Image" />
              </el-dialog>
            </div>
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="商品分类">
            {{ convertGoodsTypeId(formData.goodsTypeId) }}
          </el-descriptions-item>
          <el-descriptions-item label="商品名称">
            {{ formData.goodsName }}
          </el-descriptions-item>
          <el-descriptions-item label="商品价格">
            {{ formData.goodsPrice }}
          </el-descriptions-item>
          <el-descriptions-item label="商品创建人">
            {{ formData.createName }}
          </el-descriptions-item>
          <el-descriptions-item label="创建时间">
            {{ formatDate(formData.createTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="商品更新人">
            {{ formData.updateName }}
          </el-descriptions-item>
          <el-descriptions-item label="更新时间">
            {{ formatDate(formData.updateTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="照片">
            <div>
              <el-upload
                v-model:file-list="fileList"
                :action="`${path}/fileUploadAndDownload/upload`"
                list-type="picture-card"
                :on-preview="handlePictureCardPreview"
                :on-remove="handleRemove"
                :on-success="uploadSuccess"
                :on-error="uploadError"
                :show-file-list="true"
                :disabled="true"
              >
              </el-upload>
              <el-dialog v-model="dialogVisible">
                <img w-full :src="dialogImageUrl" alt="Preview Image" />
              </el-dialog>
            </div>
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
createGoldGoods,
deleteGoldGoods,
deleteGoldGoodsByIds,
findGoldGoods,
getGoldGoodsList,
updateGoldGoods,
} from "@/api/goldGoods";

import { getSysDictionaryDetailList } from "@/api/sysDictionaryDetail";
import { Plus } from "@element-plus/icons-vue";

// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { reactive, ref } from "vue";

defineOptions({
  name: "GoldGoods",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  ID: 0,
  goodsTypeId: 1,
  goodsName: "",
  goodsPrice: 0,
  createName: "",
  createTime: new Date(),
  updateName: "",
  updateTime: new Date(),
  goldGoodsFileList: [],
});

const goldGoodsFile = ref({
  ID: 0,
  goodsId: 0,
  fileName: "",
  fileType: 0,
  filePath: "",
  createName: "",
  createTime: new Date(),
  updateName: "",
  updateTime: new Date(),
});

const options = ref([]);
const path = ref(import.meta.env.VITE_BASE_API);
const fileList = ref([]);

const dialogImageUrl = ref("");
const dialogVisible = ref(false);

// 查询
const getGoodsTypeIdList = async () => {
  const table = await getSysDictionaryDetailList({
    page: 1,
    pageSize: 10000,
    sysDictionaryID: 7,
  });
  if (table.code === 0) {
    console.log(table.data.list);
    if (table.data.list.length > 0) {
      table.data.list.forEach((item) => {
        options.value.push({
          value: item.value,
          label: item.label,
        });
      });
    }
  }
};

getGoodsTypeIdList();

const convertGoodsTypeId = (val) => {
  console.log("val:" + val);
  let value;
  options.value.forEach((item) => {
    console.log(item.value);
    if (item.value == val) {
      console.log(item.label);
      value = item.label;
    }
  });
  return value;
};

// 验证规则
const rule = reactive({});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getGoldGoodsList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteGoldGoodsFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteGoldGoodsByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateGoldGoodsFunc = async (row) => {
  const res = await findGoldGoods({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.regoldGoods;
    console.log(res.data.regoldGoods.goldGoodsFileList);
    if (res.data.regoldGoods.goldGoodsFileList != null) {
      fileList.value = res.data.regoldGoods.goldGoodsFileList;
    }
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteGoldGoodsFunc = async (row) => {
  const res = await deleteGoldGoods({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 查看详情控制标记
const detailShow = ref(false);

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true;
};

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findGoldGoods({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.regoldGoods;
    console.log("formData", res.data.regoldGoods);
    console.log(
      "formData.goldGoodsFileList",
      res.data.regoldGoods.goldGoodsFileList
    );
    if (res.data.regoldGoods.goldGoodsFileList != null) {
      fileList.value = res.data.regoldGoods.goldGoodsFileList;
    }
    console.log("fileList.value", fileList.value);
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  formData.value = {
    goodsTypeId: 1,
    goodsName: "",
    goodsPrice: 0,
    createName: "",
    createTime: new Date(),
    updateName: "",
    updateTime: new Date(),
  };
};

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    goodsTypeId: 1,
    goodsName: "",
    goodsPrice: 0,
    createName: "",
    createTime: new Date(),
    updateName: "",
    updateTime: new Date(),
  };
  fileList.value = [];
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createGoldGoods(formData.value);
        break;
      case "update":
        res = await updateGoldGoods(formData.value);
        break;
      default:
        res = await createGoldGoods(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};

const handleRemove = (uploadFile, uploadFiles) => {
  console.log("handleRemove");
  console.log(uploadFile, uploadFiles);
  uploadFiles.splice(index, uploadFiles.indexOf(uploadFile) - 1);
  formData.value.goldGoodsFileList.splice(
    index,
    formData.value.goldGoodsFileList.indexOf(uploadFile) - 1
  );
  ElMessage({
    message: "删除成功",
    type: "success",
  });
};

const handlePictureCardPreview = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url;
  dialogVisible.value = true;
};

const uploadSuccess = (res) => {
  console.log("uploadSuccess", res.data);
  console.log(fileList.value);
  if (
    formData.value.goldGoodsFileList == null ||
    formData.value.goldGoodsFileList.length == 0
  ) {
    formData.value.goldGoodsFileList = [
      {
        fileName: res.data.file.key,
        filePath: res.data.file.url,
      },
    ];
  } else {
    formData.value.goldGoodsFileList.push({
      fileName: res.data.file.key,
      filePath: res.data.file.url,
    });
  }
  console.log("fileList", fileList.value);
  console.log("formData", formData.value.goldGoodsFileList);
};

const uploadError = () => {
  console.log("baocuo");
  ElMessage({
    type: "error",
    message: "上传失败",
  });
};
</script>

<style></style>
