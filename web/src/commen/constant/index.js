/**
 * 权限图标
 */
const functionIcons = {
  shuju: require('../../images/functionIcon/shuju-moren.png'),
  rizhi: require('../../images/functionIcon/rizhi-moren.png'),
  shanghu: require('../../images/functionIcon/shanghu-moren.png'),
  kapian: require('../../images/functionIcon/kapian-moren.png'),
  jilu: require('../../images/functionIcon/jilu-moren.png'),
  caozuo: require('../../images/functionIcon/caozuo-moren.png'),
  guize: require('../../images/functionIcon/guize-moren.png'),
  chongzhi: require('../../images/functionIcon/chongzhi-moren.png'),
  dingdan: require('../../images/functionIcon/dingdan-moren.png'),
  zhangdan: require('../../images/functionIcon/zhangdan-moren.png'),
  xitong: require('../../images/functionIcon/xitong-moren.png'),
  shebei: require('../../images/functionIcon/shebei-moren.png'),
  'shuju-active': require('../../images/functionIcon/shuju-xuanzhong.png'),
  'rizhi-active': require('../../images/functionIcon/rizhi-xuanzhong.png'),
  'shanghu-active': require('../../images/functionIcon/shanghu-xuanzhong.png'),
  'kapian-active': require('../../images/functionIcon/kapian-xuanzhong.png'),
  'jilu-active': require('../../images/functionIcon/jilu-xuanzhong.png'),
  'caozuo-active': require('../../images/functionIcon/caozuo-xuanzhong.png'),
  'guize-active': require('../../images/functionIcon/guize-xuanzhong.png'),
  'chongzhi-active': require('../../images/functionIcon/chongzhi-xuanzhong.png'),
  'dingdan-active': require('../../images/functionIcon/dingdan-xuanzhong.png'),
  'zhangdan-active': require('../../images/functionIcon/zhangdan-xuanzhong.png'),
  'xitong-active': require('../../images/functionIcon/xitong-xuanzhong.png'),
  'shebei-active': require('../../images/functionIcon/shebei-xuanzhong.png')
}

// 头像资源
const headerImgMap = {
  boy: require('@/images/icon_user1.png'),
  girl: require('@/images/user-image-nv.png')
}

/**
 * 用户身份
 */
const userIndentity = {
  ADMI: 'A', // 管理员
  MERT: 'M' // 商户
}

/**
 * admin账号查询时用户类型 all
 */
const allUserTypes = {
  ADMI: '超级管理员',
  MERT: '商户',
  CORP: '运营单位',
  SUPP: '供应商',
  COMM: '普通用户'
}

/**
 * admin账号新增时类型
 */
const adminUserTypes = {
  ADMI: '超级管理员',
  MERT: '商户',
  COMM: '普通用户'
}

/**
 * 商户账号用户类型
 */

const mertUserTypes = {
  MERT: '商户',
  CORP: '运营单位',
  SUPP: '供应商'
}

/**
 *  用户状态
 */
const userStatus = {
  '99': '禁用',
  '00': '正常'
}

/**
 *  商户状态  00：启用  01：停用
 */
const merchantStatus = {
  '00': '启用',
  '01': '停用'
}

/**
 * 请求方法
 */

const methodOptions = [
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  }
  // {
  //   value: "PUT",
  //   label: "更新",
  //   type: "warning"
  // },
  // {
  //   value: "DELETE",
  //   label: "删除",
  //   type: "danger"
  // }
]

/**
 * 按钮基础信息
 */

const baseButtons = [
  {
    buttonNo: '10001',
    buttonName: '查询',
    buttonIdentity: 'search'
  },
  {
    buttonNo: '10002',
    buttonName: '新增',
    buttonIdentity: 'add'
  },
  {
    buttonNo: '10003',
    buttonName: '编辑',
    buttonIdentity: 'edit'
  },
  {
    buttonNo: '10004',
    buttonName: '删除',
    buttonIdentity: 'delete'
  }
]

export default {
  functionIcons,
  headerImgMap,
  allUserTypes,
  adminUserTypes,
  mertUserTypes,
  userIndentity,
  userStatus,
  merchantStatus,
  methodOptions,
  baseButtons
}
