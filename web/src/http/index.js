/*
 * @Description: 所有api接口集合
 */

import Data from './Data.js'
import Login from './Login.js'
import Menu from './Menu.js'
import User from './User.js'
import Authority from './Authority.js'
import Api from './Api.js'
import Operation from './Operation.js'
import Merchant from './Merchant.js'
import Corp from './Corp.js'
import Record from './Record.js'

export default {
  ...Data,
  ...Login,
  ...Menu,
  ...User,
  ...Authority,
  ...Api,
  ...Operation,
  ...Merchant,
  ...Corp,
  ...Record
}
