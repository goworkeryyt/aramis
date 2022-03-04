import { formatTimeToStr } from '@/utils/date'
export default {
  data() {
    return {
      // 分页
      currentPage: 1,
      pageCount: 0,
      pageSize: 10,
      orderStr: `createTime:pd:`,
      searchInfo: {}
    }
  },
  methods: {
    fmtBody(value) {
      try {
        return JSON.parse(value)
      } catch (err) {
        return value
      }
    },
    formatDate(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },

    // @params beforeFunc function 请求发起前执行的函数 默认为空函数
    // @params afterFunc function 请求完成后执行的函数 默认为空函数
    async getTableData(
      searchData,
      beforeFunc = () => {},
      afterFunc = () => {}
    ) {
      this.listLoading = true
      beforeFunc()
      const res = await this.listApi({
        current: this.currentPage,
        rowCount: this.pageSize,
        orderStr: this.orderStr,
        ...searchData
      })
      const { content, code, message } = res.data
      this.listLoading = false
      if (code === '0') {
        this.tableData = content.rows
        this.pageCount = content.total
      }
      afterFunc()
    },
    async getAllTableData(
      searchData,
      beforeFunc = () => {},
      afterFunc = () => {}
    ) {
      this.listLoading = true
      beforeFunc()
      const res = await this.listApi({
        ...searchData
      })
      const { content, code, message } = res.data
      this.listLoading = false
      if (code === '0') {
        this.tableData = content
      }
      afterFunc()
    }
  }
}
