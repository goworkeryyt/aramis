<template>
  <div>
    <div class="clearflex">
      <el-button class="fl-right submit" :loading="btnLoading" @click="authApiEnter">确 定</el-button>
    </div>
    <el-tree
      ref="apiTree"
      class="api-tree mt-20"
      :data="apiTreeData"
      show-checkbox
      node-key="id"
      :default-checked-keys="apiTreeIds"
      :props="apiDefaultProps"
    />
  </div>
</template>

<script>
export default {
  name: 'Apis',
  props: {
    roleId: {
      type: String
    }
  },
  data() {
    return {
      apiTreeData: [],
      apiTreeIds: [],
      apiIds: [],
      apiDefaultProps: {
        children: 'apiInfos',
        label: 'description'
      },
      btnLoading: false
    }
  },
  created() {
    this.findApiTree()
    this.findRoleApi(this.roleId)
  },
  methods: {
    // 请求API树
    async findApiTree() {
      const res = await this.Api.findApiTree()
      this.apiTreeData = res.data.content
    },
    // 请求选择的id
    async findRoleApi(id) {
      const res = await this.Api.findRoleApi({ id })
      this.apiTreeIds = []
      res.data.content &&
        res.data.content.forEach((item) => {
          this.apiTreeIds.push(item.id)
        })
    },
    // 确定角色绑定
    authApiEnter() {
      this.btnLoading = true
      this.apiIds = []
      const checkArr = this.$refs.apiTree.getCheckedNodes(false, true)
      checkArr.forEach((item) => {
        this.apiIds.push(item.id)
      })
      const params = {
        roleID: this.roleId,
        apiIds: this.apiIds
      }
      this.roleApiBind(params)
    },
    // 角色API绑定
    async roleApiBind(params) {
      const res = await this.Api.roleApiBind(params)
      this.btnLoading = false
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.$emit('closedrawer')
      }
    }
  }
}
</script>
<style lang="scss">
  .api-tree {
    width: 100%;
    float: left;
    min-height: 100px;
    border: 1px solid #dcdee3;
    border-radius: 6px;

    .el-tree-node__content {
      border-bottom: 0.5px dashed #dcdee3;
      border-top: 0.5px dashed #dcdee3;
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

    .el-tree-node .el-tree-node__children,
    .el-tree-node .el-tree-node__children  .el-tree-node {
      float: left;
      text-align: center;
    }

    .el-tree-node .el-tree-node__children .el-tree-node .el-tree-node__content {
      border: none;
      width: 260px;
      text-align: left;
    }

    .el-tree-node .el-tree-node__children .el-tree-node .el-tree-node__children {
      width: 100%;
      background: #FFFFFF;
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
