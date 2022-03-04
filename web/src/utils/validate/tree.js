/*
 * @Description: 树形结构与平级列表的转换方法
 * @LastEditTime: 2021-10-08 16:30:09
 */

/**
 * @description: 平级结构转化为树形结构
 * @param {Array} data 数据源
 * @param {String} superiorId 父级id
 * @param {String} permissionId 子级id
 * @return {Array} 返回树形结构
 */
function listToTree(data, superiorId, permissionId) {
  let arr = JSON.parse(JSON.stringify(data))
  const listChildren = (obj, filter) => {
    [arr, obj.children] = arr.reduce(
      (res, val) => {
        if (filter(val)) res[1].push(val)
        else res[0].push(val)
        return res
      },
      [[], []]
    )
    obj.children.forEach(val => {
      if (arr.length) { listChildren(val, obj => obj[superiorId] === val[permissionId]) }
    })
  }
  const tree = {}
  listChildren(
    tree,
    val => arr.findIndex(i => i[permissionId] === val[superiorId]) === -1
  )
  return tree.children
}

/**
 * @description: 树形结构转化为平级结构
 * @param {Array} 需要转化的树形结构
 * @return {Array} 返回平级的结构
 */
function treeToList(list) {
  const output = []
  const listData = (list) => {
    list.forEach(function(item) {
      if (item.children && item.children.length > 0) {
        listData(item.children)
        delete item.children
      }
      output.push(item)
    })
  }
  listData(list)
  return output
}

export default {
  listToTree,
  treeToList
}

