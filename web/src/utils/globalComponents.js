/*
 * @Description: 统一注册公共组件
 * @LastEditTime: 2021-11-22 12:07:25
 */
import Vue from 'vue'
import GlobalSearch from '../components/global-searchBox/index.vue'
import GlobalForm from '../components/global-form/index.vue'
import GlobalCity from '../components/global-city/index.vue'
import GlobalUploadImg from '../components/global-upload-img/index.vue'
import GlobalUploadFile from '../components/global-upload-file/index.vue'
import GlobalEditor from '../components/global-editor/index.vue'
import GlobalPagination from '../components/global-pagination/index.vue'
import GlobalSelectTree from '../components/global-select-tree/index.vue'

import OperationReminder from '../components/OperationReminder/index'
import MessageBox from '../components/MessageBox/index'

import KtButton from '../components/KtButton/index'

const components = {
  GlobalSearch,
  GlobalForm,
  GlobalCity,
  GlobalUploadImg,
  GlobalUploadFile,
  GlobalEditor,
  GlobalPagination,
  GlobalSelectTree,

  OperationReminder,
  MessageBox,

  KtButton
}

const global = {
  ...components
}
function install() {
  if (install.installed) {
    return
  }
  Object.keys(global).forEach(key => {
    Vue.component(key, global[key])
  })
}

export default install
