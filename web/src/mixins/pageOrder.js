/*
 * @Description: 列表分页的序号问题计算（结合分页组件一起使用）
 */
export default {
  computed: {
    columnIndex(index) {
      return function(index) {
        return (this.pager.currentPage - 1) * this.pager.pageSize + (index + 1)
      }
    }
  }
}
