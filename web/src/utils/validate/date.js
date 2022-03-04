/*
 * @Description: 时间 日期 的转换方法
 * @LastEditTime: 2022-02-08 10:04:01
 */

export default {
  isDate,
  convertDate,
  formatDate
}

/**
 * @description: 判断该值是不是一个日期
 */
var _number = require('./number')
function isDate(val) {
  return Object.prototype.toString.call(val) === '[object Date]' && !(0, _number.isNaN)(val.getTime())
}

/**
 * @description: 日期时间格式转换
 * @param {*} type ：all   YYYY-MM-DD    YYYY-MM-DD    HH:MM    HH:MM:SS    T
 */
function convertDate(date, type = 'all') {
  if (date === null) {
    return ''
  }
  if (date.indexOf('T') != -1) {
    date = date.replace('T', ' ')
  }
  date = new Date(date)
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()
  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()
  if (type == 'all') {
    return year + '-' + formatTen(month) + '-' + formatTen(day) + ' ' + formatTen(hour) + ':' + formatTen(minute) + ':' + formatTen(second)
  }
  if (type == 'YYYY-MM-DD') {
    return year + '-' + formatTen(month) + '-' + formatTen(day)
  }
  if (type == 'YYYY-MM-DD HH:MM') {
    return year + '-' + formatTen(month) + '-' + formatTen(day) + ' ' + formatTen(hour) + ':' + formatTen(minute)
  }
  if (type == 'HH:MM:SS') {
    return formatTen(hour) + ':' + formatTen(minute) + ':' + formatTen(second)
  }
}

/**
 * @description: 当前时间转换时区时间
 */
function formatDate(date) {
  date = new Date(date)
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()
  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()
  return year + '-' + formatTen(month) + '-' + formatTen(day) + 'T' + formatTen(hour) + ':' + formatTen(minute) + ':' + formatTen(second)
}

/**
 * @description: 数字转化（不够10的前面补0）
 */
function formatTen(num) {
  return num > 9 ? `${num}` : `0${num}`
}
