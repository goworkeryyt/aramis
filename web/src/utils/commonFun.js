/*
 * @Description: 公共方法
 * @LastEditTime: 2021-11-22 09:54:42
 */

// 深拷贝1
function deepcopy(obj) {
  var out = []
  var i = 0
  var len = obj.length
  for (; i < len; i++) {
    if (obj[i] instanceof Array) {
      out[i] = deepcopy(obj[i])
    } else out[i] = obj[i]
  }
  return out
}

// 深拷贝2
function deepcopy2(target) {
  // 定义一个变量
  let result
  // 如果当前需要深拷贝的是一个对象的话
  if (typeof target === 'object') {
    // 如果是一个数组的话
    if (Array.isArray(target)) {
      result = [] // 将result赋值为一个数组，并且执行遍历
      for (const i in target) {
        // 递归克隆数组中的每一项
        result.push(deepcopy2(target[i]))
      }
      // 判断如果当前的值是null的话；直接赋值为null
    } else if (target === null) {
      result = null
      // 判断如果当前的值是一个RegExp对象的话，直接赋值
    } else if (target.constructor === RegExp) {
      result = target
    } else {
      // 否则是普通对象，直接for in循环，递归赋值对象的所有值
      result = {}
      for (const i in target) {
        result[i] = deepcopy2(target[i])
      }
    }
    // 如果不是对象的话，就是基本数据类型，那么直接赋值
  } else {
    result = target
  }
  // 返回最终结果
  return result
}

// 获取该数据的数据类型
const dataType = (data) => {
  return Object.prototype.toString.call(data).replace(/\[|]/g, '').split(' ')[1]
}

// 获取某个元素在数据中的索引
function getIndex(data, prop, value) {
  const index = data.findIndex((val) => val[prop] === value)
  return index
}

// 求取两个数组的差集（arr1中与arr2的差集）
function subSet(arr1, arr2) {
  var set1 = new Set(arr1)
  var set2 = new Set(arr2)

  var subset = []
  for (const item of set1) {
    if (!set2.has(item)) {
      subset.push(item)
    }
  }
  return subset
}

// 获取一个随机数
const genID = (length = 18) => {
  return Number(Math.random().toString().substr(3, length) + Date.now()).toString(36)
}

// 判断两个数组中的元素是否相同（只比较内容，不在乎顺序）
// 方式一：
function arrayCompare(array1, array2) {
  // 如果长度不一样，则直接不相等
  if (array1.length != array2.length) {
    return false
  }
  // 如果长度一样，则for循环比较
  for (let i = 0; i < array1.length; i++) {
    // array1的元素是否都在array2中存在，有一个不存在就不相等
    if (!array2.includes(array1[i])) {
      return false
    }
  }
  // 如果执行到这里，说明全部匹配
  return true
}
// 方式二：
function arrayCompare1(arr1, arr2) {
  const result = arr1.length === arr2.length && arr1.every(a => arr2.some(b => a === b)) && arr2.every(_b => arr1.some(_a => _a === _b))
  return result
}

export default {
  deepcopy,
  deepcopy2,
  dataType,
  getIndex,
  subSet,
  genID,
  arrayCompare,
  arrayCompare1
}
