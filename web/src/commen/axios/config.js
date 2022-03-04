import axios from 'axios'
import qs from 'qs'
import Vue from 'vue'
import { Message, MessageBox } from 'element-ui'
import 'normalize.css/normalize.css'
import router from '../../router'

const instance = axios.create({
  method: 'GET',
  headers: {
    'Content-Type': 'application/json;charset=UTF-8'
  },
  withCredentials: true,
  responseType: 'json',
  timeout: 100000
})

instance.interceptors.request.use(
  config => {
    const localToken = sessionStorage.getItem('token')
    if (localToken) {
      config.headers.ACCESS_TOKEN = localToken
    }

    if (config.method === 'get') {
      config.params = {
        _t: Date.now(),
        ...config.params
      }
    } else if (config.method === 'post') {
      if (Object.prototype.toString.call(config.data) === '[object FormData]') {
        config.headers['Content-Type'] = 'multipart/form-data'
        config.data = config.data
      } else {
        config.headers['Content-Type'] = 'application/json;charset=UTF-8'
        config.data = config.data
        // config.data = qs.stringify(config.data);
      }
    }

    return config
  },
  error => {
    Promise.reject(error)
  }
)

instance.interceptors.response.use(
  config => {
    return config
  },
  error => {
    if (error.response.status === 401) {
      localStorage.clear()
      sessionStorage.clear()
      router.replace({
        path: '/',
        query: { redirect: router.currentRoute.fullPath || '' }
      })
      if (error.response.data.message) {
        Message({
          message: error.response.data.message,
          center: true,
          duration: '5000',
          type: 'error'
        })
      }

      return
    } else if (error.response.status === 404) {
      // router.replace({
      //   path: "/404",
      //   query: { redirect: router.currentRoute.fullPath }
      // });
      if (error.response.data) {
        MessageBox.alert(error.response.data.message, '404', {
          confirmButtonText: '确定'
        })
        return
      }
      return
    } else if (error.response.status === 405) {
      // router.replace({
      //   path: "/405",
      //   query: { redirect: router.currentRoute.fullPath }
      // });
      if (error.response.data) {
        MessageBox.alert(error.response.data.message, '405', {
          confirmButtonText: '确定'
        })
      }
      return
    } else if (error.response.status === 500) {
      // router.replace({
      //   path: "/500",
      //   query: { redirect: router.currentRoute.fullPath }
      // });
      if (error.response.data) {
        MessageBox.alert(error.response.data.message, '500', {
          confirmButtonText: '确定'
        })
      }
      return
    }

    return Promise.reject(error)
  }
)

export default instance
