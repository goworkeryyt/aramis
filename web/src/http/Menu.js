/*
 * @Description: 菜单管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 创建菜单
 */

function createMenu(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'menu/createMenu', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

/**
 * 删除菜单(单个)
 */

function deleteMenu(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'menu/deleteMenu', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 更新菜单
 */

function updateMenu(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'menu/updateMenu', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

/**
 * 查询菜单(详情)
 */

function findMenu(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'menu/findMenu', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

export default {
  createMenu,
  deleteMenu,
  updateMenu,
  findMenu
}
