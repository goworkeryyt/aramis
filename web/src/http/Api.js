/*
 * @Description: Api管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 查询API(分页)
 */

function findApiPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'data/findApiPage', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 新建API
 */

function createApi(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'data/createApi', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 删除API(单个)
 */

function deleteApi(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'data/deleteApi', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 更新API
 */

function updateApi(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'data/updateApi', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 查询API(树)
 */

function findApiTree(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'data/findApiTree', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 查询API(角色)
 */

function findRoleApi(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'data/findRoleApi', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 角色API绑定
 */

function roleApiBind(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'data/roleApiBind', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 查询API(所有)
 */

function findApis(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'data/findApis', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
export default {
  findApiPage,
  createApi,
  deleteApi,
  updateApi,
  findApiTree,
  findRoleApi,
  roleApiBind,
  findApis
}
