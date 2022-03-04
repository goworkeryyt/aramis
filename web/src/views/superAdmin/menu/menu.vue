
/** 菜单管理 */
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>菜单管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <div class="header-item">
        <el-button class="add" @click="addMenu('0')">新增根菜单</el-button>
      </div>
    </div>
    <el-main>
      <el-table
        ref="table"
        border
        :data="menuList"
        height="100%"
        tooltip-effect="dark"
        default-expand-all
        class="table"
        row-key="id"
      >
        <el-table-column
          align="left"
          label="展示名称"
          min-width="180"
          prop="menuName"
          fixed
        />
        <el-table-column
          align="left"
          label="路由Name"
          show-overflow-tooltip
          min-width="160"
          prop="routerName"
        />
        <el-table-column
          align="left"
          label="路由Path"
          show-overflow-tooltip
          min-width="160"
          prop="url"
        />
        <el-table-column
          align="left"
          label="是否显示"
          min-width="100"
          prop="IsShow"
        >
          <template #default="scope">
            <span>{{ scope.row.IsShow === -1 ? "隐藏" : "显示" }}</span>
          </template>
        </el-table-column>
        <el-table-column min-width="300" prop="funName" label="功能（按钮）">
          <template slot-scope="scope">
            <ul
              v-if="scope.row.buttonList && scope.row.buttonList.length > 0"
              class="function-btn-ul"
            >
              <li v-for="(item, index) in scope.row.buttonList" :key="index">
                {{ item.buttonName }}
              </li>
            </ul>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" min-width="70" prop="sort" />
        <el-table-column
          align="left"
          label="文件路径"
          min-width="300"
          prop="filePath"
        />
        <el-table-column
          align="left"
          label="图标"
          min-width="100"
          prop="menuIcon"
        >
          <template #default="scope">
            <i :class="`el-icon-${scope.row.menuIcon}`" />
            <span>{{ scope.row.menuIcon }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label-class-name="operation"
          fixed="right"
          label="操作"
          width="220"
        >
          <template slot-scope="scope">
            <el-button
              class="edit-btn"
              @click="editMenu(scope.row)"
            >编辑</el-button>
            <el-button
              class="edit-btn"
              @click="delMenu(scope.row.id)"
            >删除</el-button>
            <el-button
              v-if="!scope.row.parentId || scope.row.parentId == '0'"
              class="edit-btn"
              @click="addMenu(scope.row)"
            >添加子菜单</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
    <el-drawer
      size="50%"
      :visible.sync="drawer"
      :with-header="false"
      :show-close="false"
    >
      <AddMenu
        ref="addMenu"
        :title="title"
        :form="form"
        :type="type"
        :is-edit="isEdit"
        @closeDrawer="closeDrawer"
      />
    </el-drawer>
  </el-container>
</template>
<script>
import AddMenu from './addMenu.vue'
export default {
  components: {
    AddMenu
  },
  data() {
    return {
      menuList: [],
      // 弹窗
      drawer: false,
      title: '',
      type: '',
      form: {},
      isEdit: false
    }
  },
  mounted() {
    this.getSelfMenuTree()
  },
  methods: {
    async getSelfMenuTree() {
      const res = await this.Api.getSelfMenuTree()

      if (res.data.code === '0') {
        res.data.content.forEach((item) => {
          item.children &&
            item.children.forEach((val) => {
              val.buttonList = val.children || []
              delete val.children
            })
        })
        this.menuList = res.data.content
        this.$store.dispatch('router/getSelfMenuTree')
      }
    },
    // 新增菜单
    addMenu(val) {
      this.drawer = true
      this.title = '新增菜单'
      this.type = 'add'
      this.isEdit = false
      this.$nextTick(() => {
        const data = val == '0' ? val : { ...val }
        this.$refs.addMenu.initForm()
        this.$refs.addMenu.setOptions(this.menuList, data)
      })
    },
    // 关闭弹窗
    closeDrawer(success) {
      this.drawer = false
      if (success) {
        this.getSelfMenuTree()
      }
    },
    // 编辑菜单
    editMenu(val) {
      this.drawer = true
      this.title = '编辑菜单'
      this.type = 'edit'
      this.isEdit = true
      this.$nextTick(() => {
        this.form = { ...val }
        this.$refs.addMenu.setOptions(this.menuList, this.form)
        this.$refs.addMenu.findMenu(this.form.id)
      })
    },
    // 删除菜单
    delMenu(id) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          const params = {
            id: id
          }
          this.deleteMenu(params)
        })
        .catch(() => {
          this.$message({
            type: 'info',
            center: true,
            message: '已取消删除'
          })
        })
    },
    // 删除请求
    async deleteMenu(params) {
      const res = await this.Api.deleteMenu(params)
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.getSelfMenuTree()
      }
    }
  }
}
</script>
<style lang="scss" scoped>
.table /deep/ .el-table__expand-icon {
  line-height: 10px !important;
  height: 10px !important;
}
.role-box /deep/ .el-tabs__item {
  width: 33.33% !important;
}
.role-box {
  height: 100%;
}

.function-btn-ul {
  padding: 0 10px;
  margin: 0;

  li {
    display: inline-block;
    margin-right: 8%;
    position: relative;

    &::before {
      content: "•";
      position: absolute;
      left: -10px;
      font-size: 16px;
      font-weight: bold;
      color: #82848a;
    }
  }
}
</style>
