/*
 * @Description: 登录模块api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 *
 * @param {*} params
 * @returns
 * 获取验证码
 * user/captcha
 */

function captcha(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'user/captcha', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 *
 * @param {*} params
 * @returns
 * 获取验证码
 * user/login
 */

function login(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/login', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 *
 * @param {*} params
 * @returns
 * 修改密码
 * /user/changePassword
 */

function changePassword(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'user/changePassword', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

export default {
  captcha,
  login,
  changePassword
}
