/*
 * @Description: 所有api接口集合
 */

import Login from './Login.js'
import Menu from './Menu.js'
import User from './User.js'
import Authority from './Authority.js'
import Api from './Api.js'
import Operation from './Operation.js'
import Merchant from './Merchant.js'

export default {
  ...Login,
  ...Menu,
  ...User,
  ...Authority,
  ...Api,
  ...Operation,
  ...Merchant
}
