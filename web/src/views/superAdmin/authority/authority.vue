
/** 角色管理 */
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>角色管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <div v-if="userType == 'ADMI'" class="header-item">
        <div class="w-50p">商户：</div>
        <el-select
          v-model="searchMerchantNo"
          placeholder="请选择商户"
          clearable
          class="header-item-inner"
        >
          <el-option
            v-for="item in merchantList"
            :key="item.id"
            :label="item.merchantName"
            :value="item.merchantNo"
          />
        </el-select>
      </div>
      <div class="header-item">
        <el-button class="search" @click="findAllRoles">查询</el-button>
      </div>
      <div class="header-item">
        <el-button class="add" @click="addAuthority()">新增角色</el-button>
      </div>
    </div>
    <el-main>
      <el-table
        ref="table"
        v-loading="listLoading"
        border
        :data="tableData"
        height="100%"
        tooltip-effect="dark"
        row-key="id"
        default-expand-all
        class="table"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
      >
        <el-table-column width="80" label="序号" type="index" />
        <el-table-column label="角色名称" min-width="120" prop="roleName" />
        <el-table-column
          label="角色描述"
          min-width="180"
          show-overflow-tooltip
          prop="remark"
        />
        <el-table-column
          v-if="userType == 'ADMI'"
          label="商户名称"
          prop="merchantName"
          show-overflow-tooltip
          min-width="120"
        />
        <el-table-column
          min-width="155px"
          sortable
          prop="createTime"
          show-overflow-tooltip
          label="添加时间"
        />
        <el-table-column
          label="角色描述"
          min-width="180"
          show-overflow-tooltip
          prop="remark"
        />
        <el-table-column
          v-if="userType == 'ADMI'"
          label="商户名称"
          prop="merchantName"
          show-overflow-tooltip
          min-width="120"
        />
        <el-table-column
          label-class-name="operation"
          fixed="right"
          label="操作"
          width="300"
        >
          <template slot-scope="scope">
            <el-button
              class="edit-btn"
              :disabled="getDisable(scope.row)"
              @click="opdendrawer(scope.row)"
            >设置权限</el-button>
            <el-button
              class="edit-btn"
              :disabled="getDisable(scope.row)"
              @click="editAuthority(scope.row)"
            >编辑</el-button>
            <el-button
              class="edit-btn"
              :disabled="getDisable(scope.row)"
              @click="deleteAuth(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 角色弹窗 -->
      <el-dialog
        :visible.sync="dialogFormVisible"
        :title="dialogTitle"
        width="600px"
      >
        <el-form
          ref="authorityForm"
          :model="form"
          :rules="rules"
          label-width="150px"
        >
          <el-form-item label="角色姓名：" prop="roleName">
            <el-input
              v-model="form.roleName"
              placeholder="请输入角色姓名"
              class="input-width"
              clearable
            />
          </el-form-item>
          <el-form-item label="角色描述：" prop="remark">
            <el-input
              v-model="form.remark"
              placeholder="请输入角色描述"
              class="input-width"
              clearable
            />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button class="cancel-btn" @click="closeDialog">取 消</el-button>
          <el-button class="submit" @click="enterDialog">确 定</el-button>
        </span>
      </el-dialog>
      <el-drawer
        v-if="drawer"
        :show-close="false"
        :with-header="false"
        :visible.sync="drawer"
        :before-close="closedrawer"
        size="60%"
      >
        <el-tabs :before-leave="autoEnter" class="role-box">
          <el-tab-pane class="tab" label="角色菜单">
            <Menus ref="menus" :role-id="roleId" @closedrawer="closedrawer" />
          </el-tab-pane>
          <el-tab-pane class="tab" label="角色api">
            <Apis ref="apis" :role-id="roleId" @closedrawer="closedrawer" />
          </el-tab-pane>
        </el-tabs>
      </el-drawer>
    </el-main>
  </el-container>
</template>
<script>
import Menus from './components/menus.vue'
import Apis from './components/apis.vue'
export default {
  components: {
    Menus,
    Apis
  },
  data() {
    return {
      userType: JSON.parse(sessionStorage.getItem('user')).userType,
      searchMerchantNo: '',
      merchantList: [],
      listLoading: false,
      // 表格数据
      tableData: [],

      // 角色弹窗
      dialogType: 'add',
      dialogFormVisible: false,
      dialogTitle: '',
      id: '',
      form: {
        roleName: '',
        remark: ''
      },
      rules: {
        roleName: [
          { required: true, message: '请输入角色名', trigger: 'blur' },,
          {
            pattern: /^[\u4E00-\u9FA5A-Za-z]{2,10}$/,
            message: '角色名称为2-10位中文或字母组合'
          }
        ]
      },

      // 设置权限
      drawer: false,
      roleId: '',
      merchantNo: ''
    }
  },
  computed: {
    getDisable(row) {
      return function(row) {
        return !!(this.userType == 'ADMI' && row.merchantNo)
      }
    }
  },
  created() {
    this.findAllRoles()
    this.getAllMerchants()
  },
  methods: {
    autoEnter(activeName, oldActiveName) {
      const paneArr = ['menus', 'apis']
      // this.$nextTick(() => {
      //   this.$refs[paneArr[oldActiveName]].getRoleMenu(this.roleId);
      // });
    },

    async findAllRoles() {
      const data = {}

      if (this.userType == 'ADMI') {
        data.merchantNo = this.searchMerchantNo
      }

      const res = await this.Api.findAllRoles(data)
      this.tableData = res.data.content
    },

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

    // 设置权限
    opdendrawer(row) {
      this.drawer = true
      this.roleId = row.id
    },
    // 新增角色
    addAuthority() {
      this.initForm()
      this.dialogTitle = '新增角色'
      this.dialogType = 'add'
      this.dialogFormVisible = true
    },

    // 编辑
    editAuthority(row) {
      this.dialogTitle = '编辑角色'
      this.dialogType = 'edit'
      this.id = row.id
      for (const k in this.form) {
        this.form[k] = row[k]
      }
      this.dialogFormVisible = true
    },
    // 删除角色
    deleteAuth(row) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const params = {
            id: row.id
          }
          this.deleteRole(params)
        })
        .catch(() => {
          this.$message({
            type: 'info',
            center: true,
            message: '已取消删除'
          })
        })
    },
    // 初始化表单
    initForm() {
      if (this.$refs.authorityForm) {
        this.$refs.authorityForm.resetFields()
      }
      this.form = {
        roleName: '',
        remark: ''
      }
    },

    // 弹窗关闭
    closeDialog() {
      this.dialogFormVisible = false
    },
    // 抽屉关闭
    closedrawer() {
      this.drawer = false
    },
    // 弹窗确定
    enterDialog() {
      this.$refs.authorityForm.validate((valid) => {
        if (valid) {
          switch (this.dialogType) {
            case 'add':
              this.createRole(this.form)
              break
            case 'edit':
              const params = {
                ...this.form,
                id: this.id
              }
              this.updateRole(params)
              break
          }
        }
      })
    },
    // 新增角色请求
    async createRole(params) {
      const res = await this.Api.createRole(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.findAllRoles()
        this.dialogFormVisible = false
      }
    },
    // 编辑角色
    async updateRole(params) {
      const res = await this.Api.updateRole(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.findAllRoles()
        this.dialogFormVisible = false
      }
    },
    // 删除角色请求
    async deleteRole(params) {
      const res = await this.Api.deleteRole(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.findAllRoles()
      }
    }
  }
}
</script>
<style  scoped>
.table >>> .el-table__expand-icon {
  line-height: 10px !important;
  height: 10px !important;
}
.role-box >>> .el-tabs__item {
  width: 50% !important;
}

.role-box {
  height: 100% !important;
}

.role-box >>> .el-tabs__header {
  border-top: none !important;
  margin: -1px 0 0 -1px !important;
}

.role-box >>> .el-tabs__item:first-child {
  margin-left: -2px;
}

.role-box >>> .el-tabs__item:last-child {
  margin-right: -2px;
}

.role-box >>> .el-tabs__content {
  padding: 20px;
}
</style>
