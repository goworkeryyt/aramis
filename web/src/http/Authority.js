/*
 * @Description: 角色管理api接口
 */
import { authURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 查询所有角色
 */

function findAllRoles(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'role/findAllRoles', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 角色分页查询
 */

function findRolePage(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'role/findRolePage', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 新建角色
 * role/createRole
 */

function createRole(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'role/createRole', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 删除角色
 * role/deleteRole
 */

function deleteRole(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'role/deleteRole', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 更新角色
 * role/updateRole
 */

function updateRole(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'role/updateRole', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 获取角色菜单ID
 */

function getRoleMenu(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'role/getRoleMenu', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 角色菜单绑定
 */

function roleMenuBind(params) {
  return new Promise((resolve, reject) => {
    postRequest(authURL + 'role/roleMenuBind', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}
/**
 * 获取用户菜单权限树
 */

function getSelfMenuTree(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'role/getSelfMenuTree', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

/**
 * 获取用户按钮权限
 */

function getSelfButtons(params) {
  return new Promise((resolve, reject) => {
    getRequest(authURL + 'role/getSelfButtons', params)
      .then(res => {
        resolve(res)
      })
      .catch(err => {
        reject(err)
      })
  })
}

export default {
  findAllRoles,
  findRolePage,
  createRole,
  deleteRole,
  updateRole,
  getRoleMenu,
  roleMenuBind,
  getSelfMenuTree,
  getSelfButtons
}
