<!--
 * @Description: 商户管理
-->
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>商户管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="header">
      <div class="header-item">
        <div class="w-75p">商户名称：</div>
        <el-input v-model="merchantName" class="header-item-inner" clearable placeholder="请输入商户名称" />
      </div>
      <div v-if="userType == 'ADMI' || userType == 'COMM'" class="header-item">
        <div class="w-60p">父商户：</div>
        <el-select v-model="parentId" placeholder="请选择父商户" clearable class="header-item-inner">
          <el-option v-for="item in parentMerchantList" :key="item.id" :label="item.merchantName" :value="item.id" />
        </el-select>
      </div>
      <div class="header-item">
        <kt-button class="search" perms="merchant:merchantList:search" label="查询" @click="getPageList('search')" />
      </div>
      <div class="header-item">
        <el-button class="default-btn" type="text" @click="refresh">刷新</el-button>
      </div>
      <div class="header-item">
        <kt-button class="add" perms="merchant:merchantList:add" label="添加" @click="addInfo" />
      </div>
    </div>
    <el-main>
      <el-table ref="multipleTable" v-loading="listLoading" :data="tableData" border height="100%" :row-style="{ height: '45px' }">
        <el-table-column width="80" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="商户编号" prop="merchantNo" min-width="120" show-overflow-tooltip />
        <el-table-column label="商户名称" prop="merchantName" min-width="120" show-overflow-tooltip />
        <el-table-column v-if="userType == 'ADMI'" label="商户等级" prop="level" min-width="70" :formatter="formatLevel" />
        <el-table-column label="地址" prop="address" min-width="140" show-overflow-tooltip />
        <!-- <el-table-column label="区域代码" prop="regionCode" min-width="120" />
                <el-table-column label="区域名称" prop="regionName" min-width="120" /> -->
        <el-table-column label="状态" prop="status" min-width="80">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.status == '01'" type="danger">停用</el-tag>
            <el-tag v-else-if="scope.row.status == '00'">启用</el-tag>
            <span v-else />
          </template>
        </el-table-column>
        <el-table-column v-if="userType == 'ADMI' || userType == 'COMM'" label="父商户" prop="parentName" min-width="120" show-overflow-tooltip />
        <el-table-column label="收单机构代码" prop="instId" min-width="120" show-overflow-tooltip />
        <el-table-column label="本商户密钥" prop="selfRsaPriKey" min-width="130" show-overflow-tooltip />
        <el-table-column fixed="right" label="操作" width="130">
          <template slot-scope="scope">
            <kt-button class="edit-btn" perms="merchant:merchantList:edit" size="mini" label="编辑" @click="updateInfo(scope.row)" />
            <kt-button class="edit-btn" perms="merchant:merchantList:delete" size="mini" label="删除" @click="deleteInfo(scope.row)" />
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
      <merchant-add ref="addUpdateForm" :count="count" :action-type="actionType" :row-info="rowInfo" @resetForm="resetForm" @submitForm="submitForm" />
    </el-drawer>

    <operation-reminder :operation-content="operationContent" :operation-visible="dialogVisible" @operationResult="getInfo" />
  </el-container>
</template>
<script>
import MerchantAdd from './merchantAdd.vue'
export default {
  components: {
    MerchantAdd
  },
  data() {
    return {
      userType: JSON.parse(sessionStorage.getItem('user')).userType,
      merchantName: '',
      parentId: '',

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
      parentMerchantList: [],

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
    this.getAllMerchants() // 获取所有已添加的商户列表
    this.getPageList()
  },
  methods: {
    // 获取所有商户列表（不分页）
    async getAllMerchants() {
      const res = await this.Api.getAllMerchants()
      const { code, message, content } = res.data

      if (code === '0') {
        const list = []
        res.data.content.forEach(item => {
          if (item.level == '1') {
            list.push(item)
          }
        })

        this.parentMerchantList = list
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    formatLevel(row) {
      return row.level == '1' ? '1 级' : '2 级'
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
      this.merchantName = ''
      this.parentId = ''
      this.getPageList()
    },
    // 获取商户列表
    async getPageList(type) {
      this.listLoading = true
      let current
      if (type) {
        current = 1 // 点击查询按钮
        this.currentPage = 1
      } else {
        current = this.currentPage
      }

      const merchantName = this.merchantName
      const parentId = this.parentId
      const rowCount = this.pageSize

      const res = await this.Api.getMerchantPage({
        merchantName: merchantName ? `lk:${merchantName}` : '',
        parentId: parentId ? `eq:${parentId}` : '',
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
      const data = { ...formValue }

      if (this.actionType == 'add') {
        const res = await this.Api.createMerchant(data)
        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })
        if (code === '0') {
          this.dialogFormVisible = false
          this.getPageList()
          this.getAllMerchants()
        }
      } else {
        data.id = this.selectedId
        const res = await this.Api.updateMerchant(data)
        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })

        if (code === '0') {
          this.dialogFormVisible = false
          this.getPageList()
          this.getAllMerchants()
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
      this.operationContent = '您确定要删除这条商户信息吗？'
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
      const res = await this.Api.deleteMerchant({ id: this.selectedId })
      const { code, message, content } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })

      if (code === '0') {
        this.getPageList()
        this.getAllMerchants()
      }
    }
  }
}
</script>
<style  scoped>

</style>
