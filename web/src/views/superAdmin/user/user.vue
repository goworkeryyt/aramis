<!--
 * @Description: 用户管理
-->
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>用户管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="header">
      <div class="header-item">
        <div class="w-65p">用户名：</div>
        <el-input
          v-model="searchInfo.userName"
          clearable
          class="header-item-inner"
          placeholder="请输入用户名"
        />
      </div>
      <div class="header-item">
        <div class="w-75p">用户类型：</div>
        <el-select
          v-model="searchInfo.userType"
          class="header-item-inner"
          clearable
          placeholder="选择用户类型"
        >
          <el-option
            v-for="(item, index) in searchUserTypes"
            :key="index"
            :label="item"
            :value="index"
          />
        </el-select>
      </div>
      <div class="header-item">
        <div class="w-75p">用户状态：</div>
        <el-select
          v-model="searchInfo.status"
          clearable
          class="header-item-inner"
          placeholder="用户状态"
        >
          <el-option
            v-for="(item, index) in Constant.userStatus"
            :key="index"
            :label="item"
            :value="index"
          />
        </el-select>
      </div>
      <div class="header-item">
        <el-button class="search" @click="onSubmit">查询</el-button>
      </div>
      <div class="header-item">
        <el-button class="default-btn" @click="refresh">刷新</el-button>
      </div>
      <div class="header-item">
        <el-button class="add" @click="addUser">新增用户</el-button>
      </div>
    </div>

    <el-main>
      <el-table
        ref="multipleTable"
        v-loading="listLoading"
        :data="tableData"
        border
        height="100%"
      >
        <el-table-column width="50" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="头像"
          width="70"
          prop="headerImg"
          align="center"
        >
          <template #default="scope">
            <img
              v-if="scope.row.headerImg"
              :src="getHeaderImg(scope.row.headerImg)"
              style="height: 42px; border-radius: 21px"
            >
          </template>
        </el-table-column>
        <el-table-column
          label="用户名"
          min-width="100"
          show-overflow-tooltip
          prop="username"
        />
        <el-table-column
          min-width="100"
          prop="phone"
          show-overflow-tooltip
          label="手机号"
        />
        <el-table-column label="昵称" min-width="150" prop="nickName">
          <!-- 行内编辑昵称功能 -->
          <!-- <template #default="scope"> -->
          <!-- <p v-if="!scope.row.editFlag" class="nickName">{{ scope.row.nickName }} <i class="el-icon-edit pointer" style="color:#1292FF" @click="openEidt(scope.row)" /></p>
            <p v-if="scope.row.editFlag" class="nickName">
              <el-input v-model="scope.row.nickName" />
              <i class="el-icon-check pointer" size="14px" style="color:#1292FF;font-weight:bold" @click="enterEdit(scope.row)" />
              <i class="el-icon-close pointer" size="14px" style="color:#f23c3c;font-weight:bold" @click="closeEdit(scope.row)" />
            </p> -->
          <!-- </template> -->
        </el-table-column>
        <el-table-column
          prop="merchantName"
          label="商户名称"
          show-overflow-tooltip
          min-width="120"
        />
        <el-table-column
          prop="userType"
          label="用户类型"
          :formatter="Formatter.formatUserType"
          min-width="90"
          show-overflow-tooltip
        />
        <el-table-column label="用户角色" min-width="300" prop="roleIds">
          <template #default="scope">
            <el-tag
              v-for="item in scope.row.roles"
              :key="item.id"
              style="margin-right: 10px"
            >{{ item.roleName }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          width="80px"
          prop="status"
          show-overflow-tooltip
          label="状态"
          :formatter="Formatter.userStatus"
        >
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.status"
              active-value="00"
              inactive-value="99"
              @change="updateUserStatus(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column
          width="100px"
          sortable
          prop="loginCount"
          label="登录次数"
        />
        <el-table-column
          min-width="155px"
          sortable
          prop="loginTime"
          label="最后一次登录时间"
        />
        <el-table-column
          min-width="155px"
          sortable
          prop="createTime"
          show-overflow-tooltip
          label="添加时间"
        />
        <el-table-column fixed="right" label="操作" width="200">
          <template slot-scope="scope">
            <el-button
              class="edit-btn"
              :disabled="userType == 'ADMI' && scope.row.creatorType != 'ADMI'"
              @click="editUserInfo(scope.row)"
            >编辑</el-button>
            <el-button
              class="edit-btn"
              :disabled="userType == 'ADMI' && scope.row.creatorType != 'ADMI'"
              @click="deleteUser(scope.row)"
            >删除</el-button>
            <el-button
              class="edit-btn"
              @click="updateUserPassword(scope.row)"
            >重置密码</el-button>
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

    <!-- 添加、编辑抽屉-->
    <el-drawer
      :visible.sync="dialogFormVisible"
      :with-header="false"
      :show-close="false"
      size="625px"
      @close="resetForm"
    >
      <add-user
        ref="addUpdateForm"
        :count="count"
        :user-roles="userRoles"
        :user-types="userTypes"
        :action-type="actionType"
        :row-info="rowInfo"
        @resetForm="resetForm"
        @submitForm="submitForm"
      />
    </el-drawer>

    <operation-reminder
      :operation-content="operationContent"
      :operation-visible="dialogVisible"
      @operationResult="getInfo"
    />
  </el-container>
</template>
<script>
import AddUser from './addUser.vue'
import infoList from '@/mixins/infoList'
export default {
  components: {
    AddUser
  },
  mixins: [infoList],
  data() {
    return {
      listLoading: false,
      backNickName: '',
      userType: '',
      searchUserTypes: [],

      tableData: [],
      currentPage: 1,
      pageCount: 0,
      pageSize: 10,
      orderStr: `createTime:pd:`,
      searchInfo: {},

      actionType: 'add',
      dialogFormVisible: false,
      rowInfo: {},
      count: 0,
      userTypes: [],
      userRoles: [],

      selectedId: '',
      dialogVisible: false,
      operationContent: '您确定要删除这条用户信息吗？'
    }
  },
  computed: {
    columnIndex(index) {
      return function(index) {
        return (this.currentPage - 1) * this.pageSize + (index + 1)
      }
    },
    getHeaderImg(headerImg) {
      return function(headerImg) {
        return headerImg.slice(0, 4) == 'http'
          ? headerImg
          : this.Constant.headerImgMap[headerImg]
      }
    }
  },
  watch: {
    tableData() {}
  },
  async created() {
    const data = {}
    if ((this.userType = 'ADMI')) {
      data.isCreate = '1'
    }

    const res = await this.Api.findAllRoles(data)
    this.userRoles = res.data.content
    await this.getUserPage()

    const user = JSON.parse(sessionStorage.getItem('user'))
    this.userType = user.userType

    if (this.userType === 'ADMI') {
      this.userTypes = this.Constant.adminUserTypes
      this.searchUserTypes = this.Constant.allUserTypes
    } else if (this.userType === 'MERT') {
      this.userTypes = this.Constant.mertUserTypes
      this.searchUserTypes = this.Constant.mertUserTypes
    }
  },

  mounted() {},
  methods: {
    // 查询
    onSubmit() {
      this.currentPage = 1
      this.pageCount = 0
      this.pageSize = 10
      this.getUserPage()
    },

    // 选择每页显示的条数
    handleSizeChange(value) {
      this.pageSize = value
      this.getUserPage()
    },

    // 获取当前点击页
    handleCurrentChange(value) {
      this.currentPage = value
      this.getUserPage()
    },

    // 刷新
    refresh() {
      this.searchInfo = {}
      this.getUserPage()
    },

    // 获取用户列表
    async getUserPage(type) {
      this.listLoading = true

      const data = {
        orderStr: `createTime:pd:`,
        rowCount: this.pageSize,
        current: this.currentPage,
        userName: this.searchInfo.userName
          ? `lk:${this.searchInfo.userName}`
          : '',
        userType: this.searchInfo.userType
          ? `eq:${this.searchInfo.userType}`
          : '',
        status: this.searchInfo.status ? `eq:${this.searchInfo.status}` : ''
      }

      const res = await this.Api.getUserPage(data)
      const { content, code, message } = res.data
      this.listLoading = false

      if (code === '0') {
        content.rows.forEach((item) => {
          item.editFlag = false
        })

        this.tableData = res.data.content.rows
        this.setRoleIds()
        this.pageCount = res.data.content.total
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    // 编辑提示
    openEidt(row) {
      if (this.tableData.some((item) => item.editFlag)) {
        this.$message({
          type: 'warning',
          center: true,
          message: '当前存在正在编辑的用户'
        })
        return
      }

      this.backNickName = row.nickName
      row.editFlag = true
    },

    // 修改昵称
    async enterEdit(row) {
      const data = { ...row }
      delete data.roles

      const res = await this.Api.updateUser(data)

      const { code, message, content } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })

      if (code === '0') {
        this.backNickName = ''
        row.editFlag = false
      }
    },

    // 关闭编辑昵称
    closeEdit(row) {
      row.nickName = this.backNickName
      this.backNickName = ''
      row.editFlag = false
    },

    // 用户角色展示
    setRoleIds() {
      this.tableData &&
        this.tableData.forEach((user) => {
          const roleIds =
            user.roles &&
            user.roles.map((i) => {
              return i.id
            })

          user.roleIds = roleIds
        })
    },

    // 点击“添加”按钮
    addUser() {
      this.dialogFormVisible = true
      this.actionType = 'add'
      this.rowInfo = {}
      this.count = this.count + 1
    },

    // 编辑用户信息
    editUserInfo(row) {
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
      this.dialogFormVisible = false
      this.$nextTick(() => {
        this.$refs.addUpdateForm.resetFields()
      })
    },

    // 点击表单的“确定”按钮 添加用户信息
    async submitForm(dialogForm) {
      if (this.actionType == 'add') {
        const res = await this.Api.createUser(dialogForm)
        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })

        if (code === '0') {
          this.dialogFormVisible = false
          this.getUserPage()
        }
      } else {
        dialogForm.id = this.selectedId
        const res = await this.Api.updateUser(dialogForm)
        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })

        if (code === '0') {
          this.dialogFormVisible = false
          this.getUserPage()
        }
      }
    },

    // 启用/禁用
    async updateUserStatus(row) {
      const res = await this.Api.updateUserStatus({
        id: row.id,
        status: row.status
      })
      const { code, message, content } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })

      if (code != '0') {
        this.getUserPage()
      }
    },

    // 点击“删除”按钮
    deleteUser(row) {
      this.dialogVisible = true
      this.selectedId = row.id
    },

    // 点击“确认删除”按钮
    getInfo(value) {
      this.dialogVisible = false
      if (value == 'yes') {
        this.deleteInfoConfirm()
      }
    },

    // 删除用户
    async deleteInfoConfirm() {
      const res = await this.Api.deleteUser({ id: this.selectedId })

      const { code, message, content } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })

      if (code == '0') {
        this.getUserPage()
      }
    },

    // 重置密码
    updateUserPassword(row) {
      this.$confirm('是否将此用户密码重置为bsit123456?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async() => {
        const res = await this.Api.updateUserPassword({
          id: row.id
        })

        const { code, message, content } = res.data
        this.$message({
          type: code !== '0' ? 'warning' : 'success',
          center: true,
          message
        })
      })
    }
  }
}
</script>
<!-- <style lang="scss" scoped>
.del-active >>>.el-switch__label.is-active {
    color: #FF1300 !important;
}
</style> -->
