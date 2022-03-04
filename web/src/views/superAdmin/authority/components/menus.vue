<template>
  <div>
    <div class="clearflex">
      <el-button
        class="fl-right submit"
        :loading="btnLoading"
        @click="relation"
      >确 定</el-button>
    </div>
    <el-tree
      ref="menuTree"
      class="role-tree mt-20"
      :data="userMenuList"
      show-checkbox
      node-key="id"
      :props="defaultProps"
      :default-checked-keys="menuTreeIds"
      default-expand-all
    />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Menus',
  props: {
    roleId: {
      type: String
    }
  },
  data() {
    return {
      menuTreeIds: [],
      menuIds: [],
      buttonIds: [],
      defaultProps: {
        children: 'children',
        label: 'menuName',
        label: function(data, node) {
          if (data.buttonName) {
            return data.buttonName
          } else {
            return data.menuName
          }
        }
      },
      btnLoading: false
    }
  },
  watch: {},
  computed: {
    ...mapGetters(['userMenuList'])
  },
  created() {
    this.getRoleMenu(this.roleId)
  },
  mounted() {},
  methods: {
    async getRoleMenu(id) {
      const res = await this.Api.getRoleMenu({ id })
      const menus = res.data.content
      const arr = []
      menus.forEach((item) => {
        // 防止直接选中父级造成全选
        if (!menus.some((same) => same.parentId === item.id)) {
          arr.push(item.id)
        }
      })
      this.menuTreeIds = arr
    },
    relation() {
      this.btnLoading = true
      this.menuIds = []
      this.buttonIds = []

      const checkArr = this.$refs.menuTree.getCheckedNodes(false, true)
      checkArr.forEach((item) => {
        if (item.buttonId) {
          this.buttonIds.push(item.buttonId)
        } else {
          this.menuIds.push(item.id)
        }
      })
      const params = {
        roleId: this.roleId,
        menuIds: this.menuIds,
        buttonIds: this.buttonIds
      }
      this.roleMenuBind(params)
    },
    // 角色菜单绑定
    async roleMenuBind(params) {
      const res = await this.Api.roleMenuBind(params).catch((err) => {
        this.btnLoading = false
      })

      this.btnLoading = false
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.$emit('closedrawer')
        this.$store.dispatch('router/getSelfMenuTree')
      }
    }
  }
}
</script>
<style lang="scss">
.role-tree {
  width: 100%;
  float: left;
  min-height: 100px;
  border: 1px solid #dcdee3;
  border-radius: 6px;

  .el-tree-node__content {
    border-bottom: 0.5px dashed #dcdee3;
    padding-top: 18px;
    padding-bottom: 12px;
    height: auto;
    display: inherit !important;
  }

  .el-tree-node {
    overflow: hidden;
  }

  .el-tree-node__content:last-child {
    border-bottom: none !important;
  }

  .el-tree-node__label {
    font-size: 14px;
    color: #333333;
  }

  .el-tree-node .el-tree-node__children .el-tree-node .el-tree-node__children,
  .el-tree-node
    .el-tree-node__children
    .el-tree-node
    .el-tree-node__children
    .el-tree-node {
    float: left;
    text-align: center;
  }

  .el-tree-node
    .el-tree-node__children
    .el-tree-node
    .el-tree-node__children
    .el-tree-node
    .el-tree-node__content {
    border: none;
    width: 200px;
    text-align: left;
  }

  .el-tree-node .el-tree-node__children .el-tree-node .el-tree-node__children {
    width: 100%;
    background: #ffffff;
  }

  .el-checkbox__inner {
    border-width: 1px;
    width: 18px;
    height: 18px;

    &::after {
      height: 9px;
      left: 5px;
      top: 0px;
      width: 4px;
      border: 2px solid #fff;
      border-left: 0;
      border-top: 0;
    }
  }
}
</style>
