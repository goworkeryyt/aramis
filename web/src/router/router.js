import Layout from '@/components/Layout'

const routerArray = [
  {
    path: '/',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/exception/404'),
    hidden: true
  },
  {
    path: '/405',
    component: () => import('@/views/exception/405'),
    hidden: true
  },
  {
    path: '/500',
    component: () => import('@/views/exception/500'),
    hidden: true
  },
  {
    path: '/home',
    component: Layout,
    redirect: 'dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'dashboard',
        meta: { title: '数据', icon: 'dashboard', noTag: false }
      }
    ]
  },
  {
    path: '/superAdmin',
    component: Layout,
    redirect: 'superAdmin',
    name: '系统管理',
    children: [
      {
        path: 'authority',
        component: () => import('@/views/superAdmin/authority/authority'),
        name: 'authority',
        meta: { title: '角色管理', noTag: false }
      },
      {
        path: 'menu',
        component: () => import('@/views/superAdmin/menu/menu'),
        name: 'menu',
        meta: { title: '菜单管理', noTag: false }
      },
      {
        path: 'api',
        component: () => import('@/views/superAdmin/api/api'),
        name: 'api',
        meta: { title: 'api管理', noTag: false }
      },
      {
        path: 'user',
        component: () => import('@/views/superAdmin/user/user'),
        name: 'user',
        meta: { title: '用户管理', noTag: false }
      },
      {
        path: 'sysOperationRecord',
        component: () =>
          import('@/views/superAdmin/operation/sysOperationRecord'),
        name: 'sysOperationRecord',
        meta: { title: '操作历史', noTag: false }
      },
      {
        path: 'accessRecord',
        component: () => import('@/views/superAdmin/operation/accessRecord'),
        name: 'accessRecord',
        meta: { title: '操作历史', noTag: false }
      }
    ]
  },
  {
    path: '/corp',
    component: Layout,
    redirect: 'corp',
    name: '运营单位管理',
    children: [
      {
        path: 'corpList',
        component: () => import('@/views/corp/corpList'),
        name: 'corpList',
        meta: { title: '运营单位信息', noTag: false }
      }
    ]
  },
  {
    path: '/merchant',
    component: Layout,
    redirect: 'merchant',
    name: '商户管理',
    children: [
      {
        path: 'merchantList',
        component: () => import('@/views/merchant/merchantList'),
        name: 'merchantList',
        meta: { title: '商户信息', noTag: false }
      }
    ]
  },
  {
    path: '/record',
    component: Layout,
    redirect: 'record',
    name: '交易记录',
    children: [
      {
        path: 'cardRcdList',
        component: () => import('@/views/record/cardRcdList'),
        name: 'cardRcdList',
        meta: { title: '卡交易记录', noTag: false }
      },
      {
        path: 'filePackRcdList',
        component: () => import('@/views/record/filePackRcdList'),
        name: 'filePackRcdList',
        meta: { title: '组包记录', noTag: false }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

export default routerArray
