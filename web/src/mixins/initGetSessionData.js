/*
 * @Description: 初次进入页面时根据角色获取相应的存储信息（商户信息、运营单位信息、线路信息）
 */
export default {
  data() {
    return {
      userType: ''
    }
  },
  created() {
    this.userType = sessionStorage.getItem('userType')
    if (this.userType == 'ADMI') {
      this.merchantNoList = JSON.parse(sessionStorage.getItem('merchantListData'))
      this.getCorpList('', 'tableCropList')

      if (this.columns) {
        this.columns = JSON.parse(JSON.stringify(this.allColumns))
      }
    } else if (this.userType == 'MERT') {
      this.corpIdList = JSON.parse(sessionStorage.getItem('corpListData'))
      this.corpList = JSON.parse(sessionStorage.getItem('corpListData'))
      this.merchantNo = sessionStorage.getItem('merchantNo')
      this.merchantNoList = JSON.parse(sessionStorage.getItem('merchantListData'))

      if (this.columns) {
        this.columns = JSON.parse(JSON.stringify(this.allColumns))
        this.columns.map((item, index) => {
          if (item.prop == 'merchantNo') {
            this.columns.splice(index, 1)
          }
        })
      }
    } else if (this.userType == 'CORP') {
      this.merchantNo = sessionStorage.getItem('merchantNo')
      this.corpId = sessionStorage.getItem('corpNo')
      this.lineNoList = JSON.parse(sessionStorage.getItem('lineListData'))

      if (this.columns) {
        this.columns = JSON.parse(JSON.stringify(this.allColumns))
        this.columns.map((item, index) => {
          if (item.prop == 'merchantNo') {
            this.columns.splice(index, 1)
          }
        })
        this.columns.map((item, index) => {
          if (item.prop == 'corpId') {
            this.columns.splice(index, 1)
          }
        })
      }
    }
  }
}
