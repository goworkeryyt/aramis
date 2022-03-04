/*
 * @Description: 运营单位管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 添加运营单位
 */
function createCorp(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'corp/createCorp', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 修改运营单位
 */
function updateCorp(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'corp/updateCorp', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 删除单个运营单位信息
 */
function deleteCorp(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'corp/deleteCorp', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 分页获取所有的运营单位列表
 */
function getCorpInfoPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'corp/getCorpInfoPage', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 获取所有运营单位列表（不分页）
 */
function findAllCorps(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'corp/findAllCorps', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

export default {
  createCorp,
  updateCorp,
  deleteCorp,
  getCorpInfoPage,
  findAllCorps
}
