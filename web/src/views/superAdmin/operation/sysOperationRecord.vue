/** 操作历史 */
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>操作历史</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <el-form :inline="true" :model="searchInfo">
        <!-- <el-form-item class="header-item" label="请求方法：">
          <el-input
            clearable
            class="header-item-inner"
            v-model="searchInfo.method"
            placeholder="请输入请求方法"
          />
        </el-form-item> -->
        <el-form-item class="header-item" label="请求路径：">
          <el-input
            v-model="searchInfo.path"
            clearable
            class="header-item-inner"
            placeholder="请输入请求路径"
          />
        </el-form-item>
        <el-form-item class="header-item" label="结果状态码：">
          <el-input
            v-model="searchInfo.status"
            clearable
            class="header-item-inner"
            placeholder="请输入结果状态码"
          />
        </el-form-item>
        <el-form-item class="header-item">
          <el-button class="search" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item class="header-item">
          <el-button class="cancel-btn" @click="refresh">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-main>
      <el-table
        ref="table"
        v-loading="listLoading"
        border
        :data="tableData"
        height="100%"
        tooltip-effect="dark"
        default-expand-all
      >
        <el-table-column width="80" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作人" min-width="140">
          <template #default="scope">
            <div>{{ scope.row.username }}</div>
          </template>
        </el-table-column>
        <el-table-column label="日期" prop="createTime" min-width="180" />
        <el-table-column label="状态码" prop="status" min-width="120">
          <template #default="scope">
            <div>
              <el-tag type="success">{{ scope.row.status }}</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="请求IP" prop="ip" min-width="120" />
        <el-table-column label="请求方法" prop="method" min-width="120" />
        <el-table-column label="请求路径" prop="path" min-width="240" />
        <el-table-column label="请求" prop="path" min-width="80">
          <template #default="scope">
            <div>
              <el-popover
                v-if="scope.row.body"
                placement="left-start"
                trigger="click"
              >
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.body) }}</pre>
                </div>
                <template slot="reference">
                  <i class="el-icon-view" />
                </template>
              </el-popover>

              <span v-else>无</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="响应" prop="path" min-width="80">
          <template #default="scope">
            <div>
              <el-popover
                v-if="scope.row.resp"
                placement="left-start"
                trigger="click"
              >
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.resp) }}</pre>
                </div>
                <template slot="reference">
                  <i class="el-icon-view" />
                </template>
              </el-popover>
              <span v-else>无</span>
            </div>
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
  </el-container>
</template>
<script>
import infoList from '@/mixins/infoList'
export default {
  mixins: [infoList],
  data() {
    return {
      listApi: this.Api.getOperateRecordPage,
      // 表格数据
      listLoading: false,
      tableData: []
    }
  },
  computed: {
    columnIndex(index) {
      return function(index) {
        return (this.currentPage - 1) * this.pageSize + (index + 1)
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 查询
    onSubmit() {
      this.currentPage = 1
      this.pageCount = 0
      this.pageSize = 10
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    // 选择每页显示的条数
    handleSizeChange(value) {
      this.pageSize = value
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    // 获取当前点击页
    handleCurrentChange(value) {
      this.currentPage = value
      const searchData = this.getSearchInfo()
      this.getTableData(searchData)
    },

    getSearchInfo() {
      return {
        // method: this.searchInfo.method ? `lk:${this.searchInfo.method}` : "",
        path: this.searchInfo.path ? `lk:${this.searchInfo.path}` : '',
        status: this.searchInfo.status ? `eq:${this.searchInfo.status}` : ''
      }
    },

    // 刷新
    refresh() {
      this.searchInfo = {}
      this.getTableData()
    }
  }
}
</script>
<style scoped>
.popover-box {
  background: #112435;
  color: #f08047;
  height: 500px;
  width: 360px;
  overflow: auto;
  padding: 10px;
}
/* .popover-box::-webkit-scrollbar {
  display: none;
} */
</style>
