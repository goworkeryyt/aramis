/*
 * @Description: 数组的方法
 * @LastEditTime: 2021-10-09 14:07:47
 */

/**
 * @description: 单一数组去重，返回一个真正的数组
 */
function unique(arr) {
  return Array.from(new Set(arr))
}

/**
 * @description: 求取两个数组的差集（数组arr1中与arr2的差集），返回arr1中剩下的元素组成的数组
 */
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

/**
 * @description: 求取两个数组的交集，返回两个数组相交元素组成的数组
 */
function intSet(arr1, arr2) {
  var set1 = new Set(arr1)
  var set2 = new Set(arr2)

  var intset = []
  for (const item of set1) {
    if (set2.has(item)) {
      intset.push(item)
    }
  }
  return intset
}

/**
 * @description: 获取某个元素在数组中的索引，返回索引值
 */
function getIndex(data, prop, value) {
  const index = data.findIndex((val) => val[prop] === value)
  return index
}

export default {
  unique,
  subSet,
  intSet,
  getIndex
}
