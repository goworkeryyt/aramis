/*
 * @Description: 数字的几种校验方法 + 解决加减乘除的精度丢失问题函数
 * @LastEditTime: 2022-02-08 10:05:06
 */

export default {
  isNumberic,
  isNaN,
  isNumber,
  isNumberP,
  checkPswd,
  add,
  sub,
  mcl,
  division
}

/**
 * @description: 校验是不是一个小数
 */
function isNumberic(val) {
  return /^\d+(\.\d+)?$/.test(val)
}

/**
 * @description: 校验是不是一个小数
 */
function isNumber(rule, value, callback) {
  if (!value && value !== 0) {
    return callback(new Error('请输入内容'))
  } else {
    const reg = /^\d+(\.\d+)?$/
    if (reg.test(value)) {
      callback()
    } else {
      return callback(new Error('请输入正确的数字'))
    }
  }
}

/**
 * @description: 校验是不是1~9之间的整数
 */
function isNumberP(rule, value, callback) {
  if (!value && value !== 0) {
    return callback(new Error('请输入内容'))
  } else {
    const reg = /[^1-9]/
    if (reg.test(value)) {
      callback()
    } else {
      return callback(new Error('请输入正确的数字'))
    }
  }
}

/**
 * @description: 校验是不是一个数值
 */
function isNaN(val) {
  if (Number.isNaN) {
    return Number.isNaN(val)
  }
  return val !== val
}

/**
 * @description: 校验6~16位密码
 */
function checkPswd(rule, value, callback) {
  if (!value && value !== 0) {
    return callback(new Error('请输入内容'))
  } else {
    const reg = /^[\u4e00-\u9fa5]{0,}$/
    if (!reg.test(value) && value.length >= 6 && value.length < 17) {
      callback()
    } else {
      return callback(new Error('只能输入6-16位数字、字母、符号'))
    }
  }
}

/**
 * @description: 加减乘除精度丢失问题的函数
 */
// 加法函数（精度丢失问题）
function add(arg1, arg2) {
  let r1, r2, m
  try {
    r1 = arg1.toString().split('.')[1].length
  } catch (e) {
    r1 = 0
  }
  try {
    r2 = arg2.toString().split('.')[1].length
  } catch (e) {
    r2 = 0
  }
  m = Math.pow(10, Math.max(r1, r2))
  return (arg1 * m + arg2 * m) / m
}

// 减法函数（精度丢失问题）
function sub(arg1, arg2) {
  let r1, r2, m, n
  try {
    r1 = arg1.toString().split('.')[1].length
  } catch (e) {
    r1 = 0
  }
  try {
    r2 = arg2.toString().split('.')[1].length
  } catch (e) {
    r2 = 0
  }
  m = Math.pow(10, Math.max(r1, r2))
  n = (r1 >= r2) ? r1 : r2
  return Number(((arg1 * m - arg2 * m) / m).toFixed(n))
}

// 乘法函数（精度丢失问题）
function mcl(arg1, arg2) {
  var m = 0
  var s1 = arg1.toString()
  var s2 = arg2.toString()
  try {
    m += s1.split('.')[1].length
  } catch (e) {}
  try {
    m += s2.split('.')[1].length
  } catch (e) {}
  return Number(s1.replace('.', '')) * Number(s2.replace('.', '')) / Math.pow(10, m)
}

// 除法函数（精度丢失问题）
function division(num1, num2) {
  let t1, t2, r1, r2
  try {
    t1 = num1.toString().split('.')[1].length
  } catch (e) {
    t1 = 0
  }
  try {
    t2 = num2.toString().split('.')[1].length
  } catch (e) {
    t2 = 0
  }
  r1 = Number(num1.toString().replace('.', ''))
  r2 = Number(num2.toString().replace('.', ''))
  return (r1 / r2) * Math.pow(10, t2 - t1)
}
