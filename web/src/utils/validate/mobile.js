/*
 * @Description: 手机号码、电话号码的校验方法
 * @LastEditTime: 2022-02-08 10:04:47
 */

export default {
  isMobile,
  checkPhone,
  checkMPhone
}

/**
 * @description: 判断是不是一个正确的手机号或电话号码
 */
function isMobile(value) {
  value = value.replace(/[^-|\d]/g, '')
  return /^((\+86)|(86))?(1)\d{10}$/.test(value) || /^0[0-9-]{10,13}$/.test(value)
}

/**
 * @description: 校验手机号码
 */
function checkMPhone(rule, value, callback) {
  if (!value) {
    return callback(new Error('手机号不能为空'))
  } else {
    const reg = /^1[3|4|5|7|8][0-9]\d{8}$/
    if (reg.test(value)) {
      callback()
    } else {
      return callback(new Error('请输入正确的手机号'))
    }
  }
}

/**
 * @description: 校验电话号码
 */
function checkPhone(rule, value, callback) {
  if (!value) {
    return callback(new Error('电话号码不能为空'))
  } else {
    const regPhone = /^0\d{2,3}-?\d{7,8}$/
    if (regPhone.test(value)) {
      callback()
    } else {
      return callback(new Error('请输入正确的电话号码'))
    }
  }
}
