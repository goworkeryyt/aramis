<!--
 * @Description: 介绍
 * @Version:
 * @Author:
 * @LastEditTime: 2021-12-03 12:26:27
 * @LastEditors:
 * @Date: 2021-11-26 10:52:05
-->
<!--
 * @Description: 分页组件
-->
<template>
  <div class="pagination">
    <el-pagination
      background
      :current-page.sync="pager.currentPage"
      :page-sizes="pager.pageSizes"
      :page-size="pager.pageSize"
      :pager-count="5"
      layout="total, sizes, prev, pager, next, jumper"
      :total="pager.total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>
<script>
export default {
  name: 'Pagination',
  props: {
    pager: {
      type: Object,
      default() {
        return {
          currentPage: 1,
          pageSize: 10,
          total: null,
          pageSizes: [5, 10, 50, 100]
        }
      }
    }
  },
  data() {
    return {}
  },
  mounted() {},
  methods: {
    // 点击“每页条数”时触发
    handleSizeChange(val) {
      if (this.pager.total == null) {
        this.pager.total == null
        this.pager.pageSize = val
        return
      }
      this.pager.pageSize = val
      this.pager.currentPage = 1
      this.$emit('change', this.pager)
    },
    // 点击“当前页码”时触发
    handleCurrentChange(val) {
      if (this.pager.total == null) {
        this.pager.total == null
        this.pager.currentPage = 1
        return
      }
      this.pager.currentPage = val
      this.$emit('change', this.pager)
    }
  }
}
</script>
<style scoped>
    .pagination {
        margin-top:10px;
        text-align: right;
    }
</style>
