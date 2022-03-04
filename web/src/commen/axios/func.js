import axios from './config.js'

export function getRequest(url, params) {
  return axios({
    url: url,
    method: 'get',
    params
  })
}

export function postRequest(url, params) {
  return axios({
    url: url,
    method: 'post',
    data: params
  })
}

export function postRequestJson(url, params) {
  return axios.post(url, params)
}
