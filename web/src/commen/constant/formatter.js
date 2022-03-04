/**
 ******************** 格式化数据显示：系统管理模块 ********************
 */
// 用户类型
function formatUserType(row) {
  if (row.userType == 'ADMI') {
    return '超级管理员'
  } else if (row.userType == 'MERT') {
    return '商户'
  } else if (row.userType == 'CORP') {
    return '运营单位'
  } else if (row.userType == 'SUPP') {
    return '供应商'
  } else if (row.userType == 'COMM') {
    return '普通用户'
  }
}

// 用户状态
function userStatus(row) {
  return row.status === '0' ? '已禁用' : '已激活'
}

/**
 ******************** 格式化数据显示：运营单位名称和商户名称 ********************
 */
// 格式化数据显示：运营单位名称
function formatCorp(row, column, corpList) {
  if (corpList.length > 0 && row.corpId) {
    return corpList.find((item) => item.corpNo == row.corpId) ? corpList.find((item) => item.corpNo == row.corpId).corpName : ''
  }
}
// 格式化数据显示：商户名称
function formatMerchant(row, column, merchantNoList) {
  if (merchantNoList.length > 0 && row.merchantNo) {
    return merchantNoList.find((item) => item.merchantNo == row.merchantNo) ? merchantNoList.find((item) => item.merchantNo == row.merchantNo).merchantName : ''
  }
}

/**
 ******************** 格式化数据显示：交易订单管理模块 ********************
 */
// 卡组织： 000-本地⾃建；001-住建部；003-交通部；
function formatCardOrg(row) {
  if (row.cardOrg == '000') {
    return '本地⾃建'
  } else if (row.cardOrg == '001') {
    return '住建部'
  } else if (row.cardOrg == '003') {
    return '交通部'
  }
}

// 卡型： 02-M1卡；01-CPU卡；
function formatCardModel(row) {
  if (row.cardModel == '01') {
    return 'CPU卡'
  } else if (row.cardModel == '02') {
    return 'M1卡'
  }
}

// 优惠类型： 00-优惠⾦额；01-优惠百分⽐；
function formatDiscountType(row) {
  if (row.discountType == '00') {
    return '优惠⾦额'
  } else if (row.discountType == '01') {
    return '优惠百分⽐'
  }
}

// 票制类型： 00-单票制；01-递增多票制；02-⼿动按票价卡；03-阶梯票价；
function formatPriceType(row) {
  if (row.priceType == '00') {
    return '单票制'
  } else if (row.priceType == '01') {
    return '递增多票制'
  } else if (row.priceType == '02') {
    return '⼿动票价'
  } else if (row.priceType == '03') {
    return '阶梯票价'
  }
}

// ⾏驶⽅向： 00-下⾏；01-上⾏；FF-未知；
function formatDirection(row) {
  if (row.direction == '00') {
    return '下⾏'
  } else if (row.direction == '01') {
    return '上⾏'
  } else if (row.direction == 'FF') {
    return '未知'
  }
}

// ⻋辆类型： 01-普通巴士；02-双层巴士；03-新能源巴士；04-燃油巴士；
function formatBusType(row) {
  if (row.busType == '01') {
    return '普通巴士'
  } else if (row.busType == '02') {
    return '双层巴士'
  } else if (row.busType == '03') {
    return '新能源巴士'
  } else if (row.busType == '04') {
    return '燃油巴士'
  }
}

// ⾏业类型： 01-公交；02-地铁；
function formatCurrIndustry(row) {
  if (row.currIndustry == '01') {
    return '公交'
  } else if (row.currIndustry == '02') {
    return '地铁'
  }
}

// 交易类型标识： 06-电⼦钱包脱机消费；09-复合消费；
function formatTradeFlag(row) {
  if (row.tradeFlag == '06') {
    return '电⼦钱包脱机消费'
  } else if (row.tradeFlag == '09') {
    return '复合消费'
  }
}

// 算法标识： 03-3DES；02-SM2；04-SM4；
function formatAlgFlag(row) {
  if (row.algFlag == '02') {
    return 'SM2'
  } else if (row.algFlag == '03') {
    return '3DES'
  } else if (row.algFlag == '04') {
    return 'SM4'
  }
}

// 扣费类型： 00-正常消费；01-上⻋/进站；02-下⻋/出站；03-补扣/追缴；
function formatHandlerType(row) {
  if (row.handlerType == '00') {
    return '正常消费'
  } else if (row.handlerType == '01') {
    return '上⻋/进站'
  } else if (row.handlerType == '02') {
    return '下⻋/出站'
  } else if (row.handlerType == '03') {
    return '补扣/追缴'
  }
}

// 订单来源： LOC-本地码；MOT-交通部互联互通码；TEN-腾讯乘车码；ALI-支付宝乘车码；UNI-银联乘车码；
function formatOrderSource(row) {
  if (row.orderSource == 'LOC') {
    return '本地码'
  } else if (row.orderSource == 'MOT') {
    return '交通部互联互通码'
  } else if (row.orderSource == 'TEN') {
    return '腾讯乘车码'
  } else if (row.orderSource == 'ALI') {
    return '支付宝乘车码'
  } else if (row.orderSource == 'UNI') {
    return '银联乘车码'
  }
}

// 格式化数据显示：价格 分 -> 元
function formatPrice(row, prop) {
  return Number(row[`${prop}`]) / 100 + '元'
}

export default {
  formatUserType,
  userStatus,

  formatCorp,
  formatMerchant,

  formatCardOrg,
  formatCardModel,
  formatDiscountType,
  formatPriceType,
  formatDirection,
  formatBusType,
  formatCurrIndustry,
  formatTradeFlag,
  formatAlgFlag,
  formatHandlerType,
  formatOrderSource,

  formatPrice
}
