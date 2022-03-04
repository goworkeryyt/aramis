<!--
 * @Description: 卡交易记录
-->
<template>
  <el-container>
    <div class="crumbs-header row-between">
      <el-breadcrumb class="crumbs" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item>卡交易记录</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="header">
      <div class="header-item">
        <div class="w-60p">设备ID：</div>
        <el-input v-model="deviceId" clearable="clearable" class="header-item-inner" placeholder="请输入设备ID" />
      </div>
      <div class="header-item">
        <div class="w-50p">卡号：</div>
        <el-input v-model="cardNo" clearable="clearable" class="header-item-inner" placeholder="请输入卡号" />
      </div>
      <div class="header-item">
        <div class="w-95p">PSAM终端号：</div>
        <el-input v-model="termId" clearable="clearable" class="header-item-inner" placeholder="请输入PSAM终端号" />
      </div>
      <div class="header-item">
        <div class="w-75p">交易日期：</div>
        <el-date-picker v-model="txnDate" type="date" value-format="yyyy-MM-dd" class="header-item-inner" placeholder="请选择交易日期" />
      </div>
      <div class="header-item">
        <div class="w-75p">数据库表：</div>
        <el-select v-model="tableName" clearable class="header-item-inner" placeholder="请选择数据库表">
          <el-option v-for="(item, index) in tableNameList" :key="index" :label="item.tableName" :value="item.tableName" />
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
            <div v-if="col.prop == 'businessType'">
              <span v-html="formatBusinessType(scope.row)" />
            </div>
            <div v-else-if="col.prop == 'tradeFlag'">
              <span v-html="formatTradeFlag(scope.row)" />
            </div>
            <div v-else-if="col.prop == 'currIndustry'">
              <span v-html="formatCurrIndustry(scope.row)" />
            </div>
            <div v-else-if="col.prop == 'origAmt'">
              <span v-html="Formatter.formatPrice(scope.row, 'origAmt')" />
            </div>
            <div v-else-if="col.prop == 'realAmt'">
              <span v-html="Formatter.formatPrice(scope.row, 'realAmt')" />
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
            <div>-</div>
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
      deviceId: '',
      cardNo: '',
      termId: '',
      txnDate: '',
      tableName: '',
      tableNameList: [],

      listLoading: false,
      tableData: [],
      currentPage: 1,
      pageCount: 0,
      pageSize: 10,
      // 选择展示的字段数组
      allColumns: [
        { label: '设备ID', prop: 'deviceId', minWidth: '120', show: true },
        { label: '商户类型编码', prop: 'mchtType', minWidth: '120', show: true },
        { label: '业务类型', prop: 'businessType', minWidth: '120', show: true },
        { label: '交易类型', prop: 'tradeFlag', minWidth: '120', show: true },
        { label: '卡号', prop: 'cardNo', minWidth: '120', show: true },
        { label: '卡计数器', prop: 'cardSeq', minWidth: '120', show: true },
        { label: '行业类型', prop: 'currIndustry', minWidth: '120', show: true },
        { label: '站点编号', prop: 'stationNo', minWidth: '120', show: true },
        { label: 'PSAM终端号', prop: 'termId', minWidth: '120', show: true },
        { label: 'PSAM交易流水号', prop: 'psamSeq', minWidth: '120', show: true },
        { label: 'PSAM版本', prop: 'psamVer', minWidth: '120', show: true },
        { label: '本地流水号', prop: 'localTxnSeq', minWidth: '120', show: true },
        { label: '交易日期', prop: 'txnDate', minWidth: '120', show: true },
        { label: '交易时间', prop: 'txnTime', minWidth: '120', show: true },
        { label: '交易状态', prop: 'tradeStatus', minWidth: '120', show: true },
        { label: '交易代码', prop: 'tradeCode', minWidth: '120', show: true },
        { label: '交易TAC', prop: 'txnTac', minWidth: '120', show: true },
        { label: '交易MAC', prop: 'txnMac', minWidth: '120', show: true },
        { label: '应收金额', prop: 'origAmt', minWidth: '120', show: true },
        { label: '实际扣费金额', prop: 'realAmt', minWidth: '120', show: true },
        { label: '交易前卡余额', prop: 'befBalance', minWidth: '120', show: true },
        { label: '交易后卡余额', prop: 'aftBalance', minWidth: '120', show: true },
        { label: '算法标识', prop: 'algFlag', minWidth: '120', show: true },
        { label: '密钥算法版本', prop: 'algVer', minWidth: '120', show: true },
        { label: '密钥算法索引', prop: 'algIndex', minWidth: '120', show: true },
        { label: '发卡机构代码', prop: 'issueCode', minWidth: '120', show: true },
        { label: '城市代码', prop: 'cityCode', minWidth: '120', show: true },
        { label: '创建时间', prop: 'createTime', minWidth: '120', show: true }
      ],
      columns: [],

      messageContent: '',
      messageVisible: false
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
    this.getPageList() // 获取卡交易记录列表
    this.getCardRcdTables() // 获取卡交易记录的数据库表名列表
  },
  methods: {
    // 格式化数据：业务类型：正常纪录-0；灰记录-1；
    formatBusinessType(row) {
      return row.businessType == '0' ? '正常纪录' : row.businessType == '1' ? '灰记录' : ''
    },
    // 格式化数据：交易类型：电子钱包脱机消费-06；复合消费-09；
    formatTradeFlag(row) {
      return row.tradeFlag == '06' ? '电子钱包脱机消费' : row.tradeFlag == '09' ? '复合消费' : ''
    },
    // 格式化数据：行业类型：公交-0001；地铁-0002；
    formatCurrIndustry(row) {
      return row.currIndustry == '0001' ? '公交' : row.currIndustry == '0002' ? '地铁' : ''
    },

    // 获取 卡交易记录的数据库表名 列表
    async getCardRcdTables() {
      const res = await this.Api.getCardRcdTables()
      if (res.data.code === '0') {
        this.tableNameList = res.data.content
      } else {
        this.messageVisible = true
        this.messageContent = res.data.message
      }
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
      this.deviceId = ''
      this.cardNo = ''
      this.termId = ''
      this.txnDate = ''
      this.tableName = ''
      this.getPageList()
      this.popoverVisible = false
    },
    // 获取卡交易记录列表
    async getPageList(type) {
      this.listLoading = true
      let current
      if (type) {
        current = 1 // 点击查询按钮
        this.currentPage = 1
      } else {
        current = this.currentPage
      }

      const deviceId = encodeURI(this.deviceId)
      const cardNo = encodeURI(this.cardNo)
      const termId = encodeURI(this.termId)
      const txnDate = encodeURI(this.txnDate)
      const txnDate1 = txnDate == '' || txnDate == 'null' ? '' : txnDate.replaceAll('-', '')
      const tableName = encodeURI(this.tableName)
      const rowCount = this.pageSize

      const res = await this.Api.getCardRcdPage({
        deviceId: deviceId == '' ? deviceId : `lk:${deviceId}`,
        cardNo: cardNo == '' ? cardNo : `lk:${cardNo}`,
        termId: termId == '' ? termId : `lk:${termId}`,
        txnDate: txnDate1 == '' ? txnDate1 : `eq:${txnDate1}`,
        tableName: tableName,
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
    // 关闭弹框
    closeMessageBox() {
      this.messageVisible = false
    }
  }
}
</script>
