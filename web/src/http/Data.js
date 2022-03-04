import { tradeURL } from '../commen/url/index.js'
import { getRequest } from '../commen/axios/func.js'

/**
 * 数据管理，统计消费充值订单量
 */
function statisticsOrder() {
  return new Promise((resolve, reject) => {
    getRequest(tradeURL + 'api/statisticsOrder')
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

function statisticsDevice() {
  return new Promise((resolve, reject) => {
    getRequest(tradeURL + 'api/statisticsDevice')
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

export default {
  statisticsOrder,
  statisticsDevice
}
