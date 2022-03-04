/*
 * @Description: 选择表格展示列
 */
export default {
  data() {
    return {
      popoverVisible: false,
      checkAll: true,
      checkedColumns: [],
      isIndeterminate: false
    }
  },
  computed: {
    // 表格展示的字段列
    showColumns() {
      const result = []
      this.columns.map(item => {
        if (item.show) {
          result.push(item)
        }
      })
      return result
    }
  },
  methods: {
    // -------------------------- Popover相关方法 START -------------------------
    // 点击“操作”打开菜单选择框
    openPopover() {
      this.popoverVisible = true
      // 去掉第二个popover弹框（修改问题：使用fixed固定列后popover会出现两个）
      this.$nextTick(() => {
        document.getElementsByClassName('tablePopover')[1].style.display = 'none'
      })
    },
    // 弹出框打开时触发
    showPopover() {
      // 选中目前已展示的字段值
      this.checkedColumns = []
      this.columns.map(item => {
        if (item.show) {
          this.checkedColumns.push(item.prop)
        }
      })
      // 如果目前展示的是全部字段，则需要勾选上“全选”按钮
      if (this.columns.length == this.checkedColumns.length) {
        this.checkAll = true
        this.isIndeterminate = false
      }
      // 重新渲染表格
      this.$nextTick(() => {
        this.$refs.multipleTable.doLayout()
      })
    },
    // 点击弹出框的“全选”按钮
    handleCheckAllChange(val) {
      const columnsPropList = []
      this.columns.map(item => columnsPropList.push(item.prop))
      this.checkedColumns = val ? columnsPropList : []
      this.isIndeterminate = false
    },
    // 点击弹出框的选择展示菜单的复选框
    handleCheckedColumnsChange(value) {
      const checkedCount = value.length
      this.checkAll = checkedCount === this.columns.length
      this.isIndeterminate = checkedCount > 0 && checkedCount < this.columns.length
    },
    // 点击弹出框的“取消”按钮
    canclePopover() {
      this.popoverVisible = false
      // this.$refs.tablePopover.doClose();
    },
    // 点击弹出框的“确定”按钮
    confirmPopover() {
      if (this.checkedColumns.length == 0) {
        this.$message({
          message: '请选择需要展示的表格字段',
          center: true,
          customClass: 'message-error'
        })
        return
      }
      // 根据选择结果，遍历修改列是否展示的属性值
      this.columns.forEach(item => {
        if (this.checkedColumns.some(el => el == item.prop)) {
          item.show = true
        } else {
          item.show = false
        }
      })
      this.popoverVisible = false
      // this.$refs.tablePopover.doClose();
      // 重新渲染表格
      this.$nextTick(() => {
        this.$refs.multipleTable.doLayout()
      })
    }
    // -------------------------- Popover相关方法 END -------------------------
  }
}
