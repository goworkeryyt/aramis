<!--
 * @Description: 运营单位管理
-->
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>运营单位管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="header">
      <div class="header-item">
        <div class="w-90p">运营单位名称：</div>
        <el-input v-model="corpName" class="header-item-inner" clearable placeholder="请输入运营单位名称" />
      </div>
      <div v-if="userType == 'ADMI' || userType == 'COMM'" class="header-item">
        <div class="w-75p">关联商户：</div>
        <el-select v-model="merchantNo" placeholder="请选择关联商户" clearable class="header-item-inner">
          <el-option v-for="item in merchantList" :key="item.id" :label="item.merchantName" :value="item.merchantNo" />
        </el-select>
      </div>
      <div class="header-item">
        <kt-button class="search" perms="corp:corpList:search" label="查询" @click="getPageList('search')" />
      </div>
      <div class="header-item">
        <el-button class="default-btn" type="text" @click="refresh">刷新</el-button>
      </div>
      <div class="header-item">
        <kt-button class="add" perms="corp:corpList:add" label="添加" @click="addInfo" />
      </div>
    </div>
    <el-main>
      <el-table ref="multipleTable" v-loading="listLoading" :data="tableData" border height="100%" :row-style="{ height: '45px' }">
        <el-table-column width="80" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="运营单位编号" prop="corpNo" min-width="120" show-overflow-tooltip />
        <el-table-column label="运营单位名称" prop="corpName" min-width="120" show-overflow-tooltip />
        <el-table-column v-if="userType == 'ADMI' || userType == 'COMM'" label="关联商户" prop="merchantName" min-width="120" show-overflow-tooltip />
        <el-table-column fixed="right" label="操作" width="130">
          <template slot-scope="scope">
            <kt-button class="edit-btn" perms="corp:corpList:edit" size="mini" label="编辑" @click="updateInfo(scope.row)" />
            <kt-button class="edit-btn" perms="corp:corpList:delete" size="mini" label="删除" @click="deleteInfo(scope.row)" />
          </template>
        </el-table-column>
      </el-table>

      <div class="block">
        <el-pagination background :current-page="currentPage" :page-size="pageSize" :page-sizes="[10, 20, 50, 100]" layout="total, sizes, prev, pager, next, jumper" :total="pageCount" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
      </div>
    </el-main>

    <!-- 添加、编辑抽屉-->
    <el-drawer
      :visible.sync="dialogFormVisible"
      :with-header="false"
      :show-close="false"
      size="625px"
      @close="resetForm"
    >
      <corp-add ref="addUpdateForm" :count="count" :merchant-list="merchantList" :action-type="actionType" :row-info="rowInfo" @resetForm="resetForm" @submitForm="submitForm" />
    </el-drawer>

    <operation-reminder :operation-content="operationContent" :operation-visible="dialogVisible" @operationResult="getInfo" />
  </el-container>
</template>
<script>
import CorpAdd from './corpAdd.vue'
export default {
  components: {
    CorpAdd
  },
  data() {
    return {
      userType: JSON.parse(sessionStorage.getItem('user')).userType,
      corpName: '',
      merchantNo: '',

      listLoading: false,
      tableData: [],
      selectedId: '', // 保存当前点击行信息的id
      currentPage: 1,
      pageCount: 0,
      pageSize: 10,

      actionType: 'add',
      dialogFormVisible: false,
      rowInfo: {},
      count: 0,
      merchantList: [],

      dialogVisible: false,
      operationContent: ''
    }
  },
  computed: {
    columnIndex(index) {
      return function(index) {
        return (this.currentPage - 1) * this.pageSize + (index + 1)
      }
    }
  },
  mounted() {
    this.getAllMerchants()
    this.getPageList()
  },
  methods: {
    // 获取所有商户列表（不分页）
    async getAllMerchants() {
      const res = await this.Api.getAllMerchants()
      const { code, message, content } = res.data

      if (code === '0') {
        this.merchantList = res.data.content
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    // 选择每页显示的条数
    handleSizeChange(value) {
      this.pageSize = value
      this.getPageList()
    },

    // 获取当前点击页
    handleCurrentChange(value) {
      this.currentPage = value
      this.getPageList()
    },

    // 点击“刷新”按钮
    refresh() {
      this.corpName = ''
      this.merchantNo = ''
      this.getPageList()
    },

    // 获取运营单位列表
    async getPageList(type) {
      this.listLoading = true
      let current
      if (type) {
        current = 1 // 点击查询按钮
        this.currentPage = 1
      } else {
        current = this.currentPage
      }

      const corpName = this.corpName
      const merchantNo = this.merchantNo
      const rowCount = this.pageSize

      const res = await this.Api.getCorpInfoPage({
        corpName: corpName ? `lk:${corpName}` : '',
        merchantNo: merchantNo ? `eq:${merchantNo}` : '',
        orderStr: `createTime:pd:`,
        rowCount: rowCount,
        current: current
      })

      const { code, message, content } = res.data
      if (code === '0') {
        this.tableData = content.rows
        this.pageCount = content.total
        this.listLoading = false
      } else {
        this.listLoading = false
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },
    // 点击“添加”按钮
    addInfo() {
      this.dialogFormVisible = true
      this.actionType = 'add'
      this.rowInfo = {}
      this.count = this.count + 1
    },
    // 点击“编辑”按钮
    updateInfo(row) {
      this.dialogFormVisible = true
      this.actionType = 'edit'
      this.selectedId = row.id
      this.rowInfo = row
      this.count = this.count + 1
      this.$nextTick(() => {
        this.$refs.addUpdateForm.clearValidate()
      })
    },
    // 点击表单的“取消”按钮
    resetForm() {
      this.$nextTick(() => {
        this.$refs.addUpdateForm.resetFields()
      })
      this.dialogFormVisible = false
    },
    // 点击表单的“确定”按钮
    async submitForm(formValue) {
      if (this.actionType == 'add') {
        const res = await this.Api.createCorp(formValue)
        const { code, message, content } = res.data

        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })
        if (code === '0') {
          this.dialogFormVisible = false
          this.getPageList()
        }
      } else {
        formValue.id = this.selectedId
        const res = await this.Api.updateCorp(formValue)
        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })

        if (code === '0') {
          this.dialogFormVisible = false
          this.getPageList()
        }
      }
    },
    // 关闭弹框
    closeMessageBox() {
      this.messageVisible = false
    },
    // 点击“删除”按钮
    deleteInfo(row) {
      this.dialogVisible = true
      this.operationContent = '您确定要删除这条运营单位信息吗？'
      this.selectedId = row.id
    },
    // 点击“确认删除”按钮
    getInfo(value) {
      this.dialogVisible = false
      if (value == 'yes') {
        this.deleteInfoConfirm()
      }
    },
    // 确认删除
    async deleteInfoConfirm() {
      const res = await this.Api.deleteCorp({ id: this.selectedId })
      const { code, message, content } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })

      if (code === '0') {
        this.getPageList()
      }
    }
  }
}
</script>
<style  scoped>

</style>
