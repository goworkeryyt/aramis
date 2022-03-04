const bsitActiveX = document.getElementById('BsitActiveX')
const readerMap = { isSuccess: false, callBackMessage: '', uid: '' }
const cardMap = { isSuccess: false, callBackMessage: '', cardNo: '' }

/**
 * 获取读卡器信息
 */
function getReaderInfomation() {
  return new Promise((resolve, reject) => {
    let isOpen = -999999

    try {
      isOpen = JSON.parse(bsitActiveX.Open()).Code
    } catch (err) {
      readerMap.callBackMessage = '请检查浏览器是否为IE11,并安装了activeX插件'

      resolve(readerMap)
      return
    }

    if (isOpen < 0) {
      readerMap.callBackMessage = '请先连接读卡器'

      resolve(readerMap)
      return
    }

    const atr = JSON.parse(bsitActiveX.Reset()).Code

    if (atr === '-1') {
      readerMap.callBackMessage = '读卡器复位失败'

      resolve(readerMap)
      return
    }

    const uidMap = JSON.parse(bsitActiveX.GetCardReaderID(2))

    if (uidMap.Code === '9000') {
      readerMap.isSuccess = true
      readerMap.uid = uidMap.Content

      resolve(readerMap)
    } else {
      readerMap.callBackMessage = '获取读卡器ID失败'

      resolve(readerMap)
      return
    }
  })
}

/**
 * 获取卡信息
 */
function getCardInfomation() {
  return new Promise((resolve, reject) => {
    let isOpen = -999999

    try {
      isOpen = JSON.parse(bsitActiveX.Open()).Code
    } catch (err) {
      cardMap.callBackMessage = '请检查浏览器是否为IE11,并安装了activeX插件'

      resolve(cardMap)
      return
    }

    if (isOpen < 0) {
      cardMap.callBackMessage = '请先连接读卡器'

      resolve(cardMap)
      return
    }

    const atr = JSON.parse(bsitActiveX.Reset()).Code

    if (atr === '-1') {
      cardMap.callBackMessage = '读卡器复位失败'

      resolve(cardMap)
      return
    }

    bsitActiveX.SendApdu('00A4000002DF01') // 苏州食堂
    const cardNoMap = JSON.parse(bsitActiveX.SendApdu('00B0950008')) // 苏州食堂

    console.log(cardNoMap)

    if (cardNoMap.Code === '36864') {
      cardMap.isSuccess = true
      cardMap.cardNo = cardNoMap.Content.substr(0, 16)// 苏州食堂

      resolve(cardMap)
    } else if (cardNoMap.Code === '28416') {
      cardMap.callBackMessage = '请先将卡片贴在读卡器的读卡区'

      resolve(cardMap)
      return
    } else {
      cardMap.callBackMessage = '获取卡号失败'

      resolve(cardMap)
      return
    }
  })
}

export { getReaderInfomation, getCardInfomation }
