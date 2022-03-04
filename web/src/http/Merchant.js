/*
 * @Description: 商户管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 添加商户
 */
function createMerchant(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'merchant/createMerchant', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 修改商户
 */
function updateMerchant(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'merchant/updateMerchant', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 删除单个商户信息
 */
function deleteMerchant(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'merchant/deleteMerchant', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 分页获取所有的商户列表
 */
function getMerchantPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'merchant/getMerchantPage', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 获取所有商户列表（不分页）
 */
function getAllMerchants(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'merchant/getAllMerchants')
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 根据id获取商户信息
 */
function findMerchantInfo(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'merchant/findMerchantInfo?id=' + params.id)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

export default {
  createMerchant,
  updateMerchant,
  deleteMerchant,
  getMerchantPage,
  getAllMerchants,
  findMerchantInfo
}
