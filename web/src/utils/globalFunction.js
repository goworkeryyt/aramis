/*
 * @Description: 全局公共方法注册
 * @LastEditTime: 2021-10-08 16:25:09
 */
import commonFun from './commonFun.js'
import array from './validate/array.js'
import date from './validate/date.js'
import email from './validate/email.js'
import mobile from './validate/mobile.js'
import number from './validate/number.js'
import tree from './validate/tree.js'

export default {
  ...commonFun,
  ...array,
  ...date,
  ...email,
  ...mobile,
  ...number,
  ...tree
}
