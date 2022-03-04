/*
 * @Description: 用户管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 添加用户
 */
function createUser(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/createUser', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 更新用户信息
 */
function updateUser(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/updateUser', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 更新用户角色
 */
function updateUserRoles(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/updateUserRoles', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 更新用户状态
 */
function updateUserStatus(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/updateUserStatus', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 删除用户
 */
function deleteUser(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/deleteUser', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 重置密码
 */
function updateUserPassword(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/updateUserPassword', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 分页获取用户信息
 */
function getUserPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'user/getUserPage', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 获取当前用户信息
 */
function getSelfUserInfo(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'user/getSelfUserInfo', params)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

export default {
  getUserPage,
  updateUserRoles,
  createUser,
  updateUserStatus,
  deleteUser,
  updateUser,
  updateUserPassword,
  getSelfUserInfo
}
