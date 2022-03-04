<!--
 * @Description: 组包记录
-->
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>组包记录</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <div class="header-item">
        <div class="w-85p">组包文件名：</div>
        <el-input v-model="filename" clearable="clearable" class="header-item-inner" placeholder="请输入组包文件名" />
      </div>
      <div class="header-item">
        <div class="w-75p">打包日期：</div>
        <el-date-picker v-model="packDate" type="date" value-format="yyyy-MM-dd" class="header-item-inner" placeholder="请选择打包日期" />
      </div>
      <div class="header-item">
        <div class="w-75p">发送状态：</div>
        <el-select v-model="sendStatus" clearable class="header-item-inner" placeholder="请选择发送状态">
          <el-option v-for="(item, index) in sendStatusList" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </div>
      <div class="header-item">
        <el-button class="search" @click="search">查询</el-button>
      </div>
      <div class="header-item">
        <el-button class="default-btn" type="text" @click="refresh">刷新</el-button>
      </div>
    </div>
    <el-main>
      <el-table ref="multipleTable" v-loading="listLoading" :data="tableData" border height="100%">
        <template slot="empty">
          <div class="tableEmptyData">暂无数据</div>
        </template>
        <el-table-column width="80" label="序号">
          <template slot-scope="scope">
            <span>{{ columnIndex(scope.$index) }}</span>
          </template>
        </el-table-column>
        <el-table-column v-for="(col, index) in showColumns" :key="index" :label="col.label" :prop="col.prop" :min-width="col.minWidth" show-overflow-tooltip>
          <template slot-scope="scope">
            <div v-if="col.prop == 'totalAmt'">
              <span v-html="Formatter.formatPrice(scope.row, 'totalAmt')" />
            </div>
            <div v-else-if="col.prop == 'sendStatus'">
              <span v-html="formatSendStatus(scope.row)" />
            </div>
            <div v-else class="gobyndsingle">
              {{ scope.row[scope.column.property] }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="70" fixed="right">
          <template slot="header" slot-scope="scope">
            <el-popover id="tablePopover" ref="tablePopover" v-model="popoverVisible" width="200" popper-class="tablePopover" trigger="manual" @show="showPopover">
              <span slot="reference" @click="openPopover">操作<i class="el-icon-arrow-down" style="font-weight: bold;" /></span>
              <div class="popoverCheckBoxArea">
                <el-checkbox v-model="checkAll" :indeterminate="isIndeterminate" @change="handleCheckAllChange">全选</el-checkbox>
                <div style="margin: 10px 0;" />
                <el-checkbox-group v-model="checkedColumns" @change="handleCheckedColumnsChange">
                  <el-checkbox v-for="column in columns" :key="column.prop" :label="column.prop">{{ column.label }}</el-checkbox>
                </el-checkbox-group>
              </div>
              <div class="popoverCheckBoxButton">
                <el-button class="cancel-btn" @click="canclePopover">取消</el-button>
                <el-button class="submit" @click="confirmPopover">确定</el-button>
              </div>
            </el-popover>
          </template>
          <template slot-scope="scope">
            <el-button class="edit-btn" @click="downloadInfo(scope.row)">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
    <div class="block">
      <el-pagination background :current-page="currentPage" :page-size="pageSize" :page-sizes="[10, 20, 50, 100]" layout="total, sizes, prev, pager, next, jumper" :total="pageCount" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>

    <message-box :message-content="messageContent" :message-visible="messageVisible" @closeMessageBox="closeMessageBox" />
  </el-container>
</template>
<script>
import popoverColumn from '../../mixins/popoverColumn.js'
export default {
  mixins: [popoverColumn],
  data() {
    return {
      filename: '',
      packDate: '',
      sendStatus: '',

      listLoading: false,
      tableData: [],
      currentPage: 1,
      pageCount: 0,
      pageSize: 10,
      // 选择展示的字段数组
      allColumns: [
        { label: '组包文件名', prop: 'filename', minWidth: '120', show: true },
        { label: '组包文件路径', prop: 'filePath', minWidth: '150', show: true },
        { label: '打包日期', prop: 'packDate', minWidth: '120', show: true },
        { label: '记录条数', prop: 'recordCount', minWidth: '120', show: true },
        { label: '交易金额', prop: 'totalAmt', minWidth: '120', show: true },
        { label: '数据批次号', prop: 'batchNo', minWidth: '120', show: true },
        { label: '发送状态', prop: 'sendStatus', minWidth: '120', show: true },
        { label: '发送时间', prop: 'sendTime', minWidth: '120', show: true },
        { label: '备注', prop: 'remark', minWidth: '120', show: true }
      ],
      columns: [],

      // 发送状态： 空字符-未发送；0-发送成功；1-发送失败；
      sendStatusList: [
        { label: '发送成功', value: '0' },
        { label: '发送失败', value: '1' }
      ],

      messageContent: '',
      messageVisible: false,

      userType: ''
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
    this.columns = JSON.parse(JSON.stringify(this.allColumns))
  },
  mounted() {
    this.getPageList() // 获取组包记录列表
  },
  methods: {
    // 格式化数据：发送状态：0-发送成功；1-发送失败；空字符-未发送；
    formatSendStatus(row) {
      return row.sendStatus == '0' ? '发送成功' : row.sendStatus == '1' ? '发送失败' : row.sendStatus == '' ? '未发送' : ''
    },

    // 选择每页显示的条数
    handleSizeChange(value) {
      this.pageSize = value
      this.getPageList()
    },
    // 获取当前点击页
    handleCurrentChange(value) {
      this.currentPage = value
      this.getPageList()
    },
    // 点击“查询”按钮
    search() {
      this.getPageList('search')
      this.popoverVisible = false
    },
    // 点击“刷新”按钮
    refresh() {
      this.filename = ''
      this.packDate = ''
      this.sendStatus = ''
      this.getPageList()
      this.popoverVisible = false
    },
    // 获取组包记录列表
    async getPageList(type) {
      this.listLoading = true
      let current
      if (type) {
        current = 1 // 点击查询按钮
        this.currentPage = 1
      } else {
        current = this.currentPage
      }

      const filename = encodeURI(this.filename)
      const packDate = encodeURI(this.packDate)
      const packDate1 = packDate == '' || packDate == 'null' ? '' : packDate
      const sendStatus = encodeURI(this.sendStatus)
      const rowCount = this.pageSize

      const res = await this.Api.getFilePackPage({
        filename: filename == '' ? filename : `lk:${filename}`,
        packDate: packDate1 == '' ? packDate1 : `eq:${packDate1}`,
        sendStatus: sendStatus == '' ? sendStatus : `eq:${sendStatus}`,
        orderStr: `createTime:pd:`,
        rowCount: rowCount,
        current: current
      })
      if (res.data.code === '0') {
        this.tableData = res.data.content.rows
        this.pageCount = res.data.content.total
        this.listLoading = false
      } else {
        this.listLoading = false
        this.messageVisible = true
        this.messageContent = res.data.message
      }
    },
    // 点击“下载”按钮
    async downloadInfo(row) {
      await this.Api.downloadPackFile({ id: row.id })
    },
    // 关闭弹框
    closeMessageBox() {
      this.messageVisible = false
    }
  }
}
</script>
