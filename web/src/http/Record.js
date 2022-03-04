/*
 * @Description: 介绍
 * @Version:
 * @Author:
 * @LastEditTime: 2022-02-16 17:16:11
 * @LastEditors: Please set LastEditors
 * @Date: 2022-02-16 16:39:43
 */
/*
 * @Description: 苏州小兰卡（卡交易记录、组包记录）
 * @Description: 卡交易记录模块api
 */
import { lanURL } from '../commen/url/index.js'
import { getRequest, postRequest } from '../commen/axios/func.js'

/**
 * 获取所有的卡交易记录 数据库表名 列表（不分页）
 */
function getCardRcdTables(params) {
  return new Promise((resolve, reject) => {
    getRequest(lanURL + 'card/getCardRcdTables')
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 分页获取卡交易记录列表
 */
function getCardRcdPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(lanURL + 'card/getCardRcdPage?deviceId=' + params.deviceId +
            '&cardNo=' + params.cardNo +
            '&termId=' + params.termId +
            '&txnDate=' + params.txnDate +
            '&tableName=' + params.tableName +
            '&orderStr=' + params.orderStr +
            '&current=' + params.current +
            '&rowCount=' + params.rowCount)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 分页获取组包记录列表
 */
function getFilePackPage(params) {
  return new Promise((resolve, reject) => {
    getRequest(lanURL + 'pack/getFilePackPage?filename=' + params.filename +
            '&packDate=' + params.packDate +
            '&sendStatus=' + params.sendStatus +
            '&orderStr=' + params.orderStr +
            '&current=' + params.current +
            '&rowCount=' + params.rowCount)
      .then((res) => {
        resolve(res)
      }).catch((err) => {
        reject(err)
      })
  })
}

/**
 * 下载组包记录文件
 */
function downloadPackFile(params) {
  window.location.href = lanURL + 'pack/downloadPackFile?id=' + params.id
}

export default {
  getCardRcdTables,
  getCardRcdPage,

  getFilePackPage,
  downloadPackFile
}
