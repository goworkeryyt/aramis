/*
 * @Description: 日志管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 操作记录查询(分页)
 */

function getOperateRecordPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'record/getOperateRecordPage', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 访问记录查询(分页)
 */

function getAccessRecordPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'record/getAccessRecordPage', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
export default {
  getOperateRecordPage,
  getAccessRecordPage
}
