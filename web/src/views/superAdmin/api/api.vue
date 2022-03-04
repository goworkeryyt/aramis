/** api管理 */
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>api管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item class="header-item" label="路径：">
          <el-input
            v-model="searchInfo.path"
            clearable
            class="header-item-inner"
            placeholder="请输入路径"
          />
        </el-form-item>
        <el-form-item class="header-item" label="描述：">
          <el-input
            v-model="searchInfo.description"
            clearable
            class="header-item-inner"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item class="header-item" label="API组：">
          <el-input
            v-model="searchInfo.apigroup"
            clearable
            class="header-item-inner"
            placeholder="请输入api组"
          />
        </el-form-item>
        <el-form-item class="header-item" label="请求：">
          <el-select
            v-model="searchInfo.method"
            class="header-item-inner"
            clearable
            placeholder="请选择请求方法"
          >
            <el-option
              v-for="item in Constant.methodOptions"
              :key="item.value"
              :label="item.value"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item class="header-item">
          <el-button class="search" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item class="header-item">
          <el-button class="cancel-btn" @click="refresh">刷新</el-button>
        </el-form-item>
        <el-form-item class="header-item">
          <el-button class="add" @click="openDialog('addApi')">新增</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-main>
      <el-table
        ref="table"
        v-loading="listLoading"
        border
        :data="tableData"
        height="100%"
        tooltip-effect="dark"
        default-expand-all
      >
        <!-- <el-table-column label="id" min-width="60" prop="ID" /> -->
        <el-table-column width="80" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="API路径" min-width="150" prop="path" />
        <el-table-column label="API分组" min-width="150" prop="apiGroup" />
        <el-table-column label="API简介" min-width="150" prop="description" />
        <el-table-column label="请求" min-width="150" prop="method" />
        <el-table-column fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button
              class="edit-btn"
              @click="deleteApi(scope.row.id)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination
          background
          :current-page="currentPage"
          :page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="pageCount"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-main>

    <!-- 新增编辑抽屉 -->
    <el-drawer
      :visible.sync="dialogFormVisible"
      :with-header="false"
      :show-close="false"
      size="625px"
      @close="closeDialog"
    >
      <div class="crumbs-header row-between header-sticky h-70">
        <div class="title">
          {{ dialogTitle }}
        </div>
        <div>
          <el-button class="cancel-btn" @click="closeDialog">取 消</el-button>
          <el-button class="submit" @click="enterDialog">确 定</el-button>
          <i class="el-icon-close ml-10" @click="closeDialog" />
        </div>
      </div>

      <el-form ref="apiForm" :model="form" :rules="rules" label-width="150px">
        <div class="theme row-flex">
          <i />
          <span>基本信息</span>
        </div>

        <div class="form_box">
          <el-form-item label="路径" prop="path">
            <el-input
              v-model="form.path"
              class="input-width"
              placeholder="请输入路径"
              clearable
              autocomplete="off"
            />
          </el-form-item>
          <el-form-item label="请求" prop="method">
            <el-select
              v-model="form.method"
              placeholder="请选择请求"
              class="input-width"
              clearable
            >
              <el-option
                v-for="item in Constant.methodOptions"
                :key="item.value"
                :label="`${item.label}(${item.value})`"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="api分组" prop="apiGroup">
            <el-input
              v-model="form.apiGroup"
              class="input-width"
              placeholder="请输入api分组"
              clearable
              autocomplete="off"
            />
          </el-form-item>
          <el-form-item label="api简介" prop="description">
            <el-input
              v-model="form.description"
              class="input-width"
              placeholder="请输入api简介"
              clearable
              autocomplete="off"
            />
          </el-form-item>
        </div>
      </el-form>
    </el-drawer>
  </el-container>
</template>
<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
import infoList from '@/mixins/infoList'
export default {
  mixins: [infoList],
  data() {
    return {
      listApi: this.Api.findApiPage,
      // 表格数据
      listLoading: false,
      tableData: [],

      // 弹窗
      type: '',
      dialogFormVisible: false,
      dialogTitle: '新增Api',
      form: {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      },
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        apiGroup: [
          { required: true, message: '请输入组名称', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入api介绍', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    columnIndex(index) {
      return function(index) {
        return (this.currentPage - 1) * this.pageSize + (index + 1)
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 查询
    onSubmit() {
      this.currentPage = 1
      this.pageCount = 0
      this.pageSize = 10
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    // 选择每页显示的条数
    handleSizeChange(value) {
      this.pageSize = value
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    // 获取当前点击页
    handleCurrentChange(value) {
      this.currentPage = value
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    getSearchInfo() {
      return {
        path: this.searchInfo.path ? `lk:${this.searchInfo.path}` : '',
        description: this.searchInfo.description
          ? `lk:${this.searchInfo.description}`
          : '',
        apigroup: this.searchInfo.apigroup
          ? `lk:${this.searchInfo.apigroup}`
          : '',
        method: this.searchInfo.method ? `lk:${this.searchInfo.method}` : ''
      }
    },

    // 重置
    refresh() {
      this.searchInfo = {}
      this.getTableData()
    },
    // 新增
    openDialog(type) {
      this.dialogFormVisible = true
      this.type = type
      this.dialogTitle = '新增Api'
    },
    // 删除
    deleteApi(id) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          const params = {
            id
          }
          this.delApi(params)
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    // 删除api请求
    async delApi(params) {
      const res = await this.Api.deleteApi(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (this.tableData.length === 1 && this.page > 1) {
        this.page--
      }
      if (code === '0') {
        const searchData = this.getSearchInfo()
        this.getTableData(searchData)
      }
    },
    // 重置表单
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form = {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      }
    },
    // 关闭弹窗
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    // 确定
    enterDialog() {
      this.$refs.apiForm.validate((valid) => {
        if (valid) {
          this.createApi(this.form)
        }
      })
    },
    // 新增Api请求
    async createApi(params) {
      const res = await this.Api.createApi(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        const searchData = this.getSearchInfo()
        this.getTableData(searchData)
        this.closeDialog()
      }
    }
  }
}
</script>
<style scoped>
</style>

