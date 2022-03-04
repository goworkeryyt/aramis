<template>
  <div class="dashboard-container">
    <div class="top-container">
      <h4 style="font-size: 18px; font-weight: normal">数据管理</h4>
      <div class="container-title">
        <span class="title">今日概况统计</span>
        <span class="update-time">更新时间：{{ orderDatas.updateTime }}</span>
      </div>
      <ul>
        <li>
          <span>消费订单数</span>
          <span><span class="number">{{ orderDatas.dayConsumeNums }}</span>单</span>
        </li>
        <li>
          <span>消费应收金额</span>
          <span><span class="number">{{
            (orderDatas.dayConsumeAmt / 100).toFixed(2)
          }}</span>元</span>
        </li>
        <li>
          <span>消费实收金额</span>
          <span><span class="number">{{
            (orderDatas.dayConsumeRealAmt / 100).toFixed(2)
          }}</span>元</span>
        </li>
        <li>
          <span>充值订单数</span>
          <span><span class="number">{{ orderDatas.dayRechargeNums }}</span>单</span>
        </li>
        <li>
          <span>充值总金额</span>
          <span><span class="number">{{
            (orderDatas.dayRechargeAmt / 100).toFixed(2)
          }}</span>元</span>
        </li>
      </ul>
    </div>
    <div class="middle-container">
      <div class="inner left">
        <div class="container-title" style="margin-left: -10px">
          <span class="title">设备统计</span>
        </div>
        <div class="device-container">
          <div class="device-content">
            <div class="device-left">
              <img src="../../images/icon-shebei.png">
              <span>设备数量</span>
            </div>
            <div class="device-right">
              <span>{{ orderDatas.deviceNums }}</span>
              <span style="font-size: 16px">台</span>
            </div>
          </div>
        </div>
        <ul>
          <li>各分组设备数</li>
          <li v-for="(item, index) in orderDatas.deviceAreas" :key="index">
            {{ item.name }}：{{ item.deviceCount }}台
          </li>
        </ul>
      </div>
      <div class="inner right">
        <div class="container-title">
          <span class="title">本月订单统计</span>
        </div>
        <div id="pieChart">
          <pie-chart :data-pie="pieDataList" />
        </div>
      </div>
    </div>
    <div class="bottom-container">
      <div class="container-title">
        <span class="title">近7天设备消费金额统计</span>
      </div>
      <div id="rectChart">
        <rect-chart :data-rect="deviceDatas" />
      </div>
    </div>
  </div>
</template>
<script>
import PieChart from './components/PieChart.vue'
import RectChart from './components/RectChart.vue'

export default {
  components: {
    PieChart,
    RectChart
  },
  data() {
    return {
      orderDatas: {},
      deviceDatas: []
    }
  },
  computed: {
    pieDataList() {
      const tempList = [
        { name: '消费订单量', value: this.orderDatas.mouthConsumeNums },
        { name: '充值订单量', value: this.orderDatas.mouthRechargeNums },
        { name: '退款订单量', value: this.orderDatas.revokeOrderNums }
      ]

      return tempList
    }
  },
  watch: {
    $route(newValue, oldValue) {
      const { name } = newValue

      if (name === 'dashboard') {
        this.statisticsOrder()
      }
    }
  },
  created() {
    // this.statisticsOrder();
    // this.statisticsDevice();
  },
  methods: {
    // async statisticsOrder() {
    //     // const res = await this.Api.statisticsOrder();
    //     // if(res.data.code === '0') {
    //     //     this.orderDatas = res.data.content;
    //     // }
    // },
    // async statisticsDevice() {
    //     const res = await this.Api.statisticsDevice();
    //     if(res.data.code === '0') {
    //         this.deviceDatas = res.data.content;
    //     }
    // }
  }
}
</script>
<style lang="scss" scoped>
.dashboard-container {
  padding: 0 15px;
  .container-title {
    width: 100%;
    height: 18px;
    line-height: 18px;
    margin: 10px 0;
    border-left: 3px solid #2589ff;
    .title {
      font-size: 16px;
      font-weight: bold;
      margin-left: 5px;
      margin-right: 10px;
    }
    .update-time {
      font-size: 14px;
    }
  }
  .top-container {
    ul {
      display: flex;
      justify-content: space-between;
      height: 100px;
      margin-top: 15px;
      margin-bottom: 25px;
      margin-left: -43px;
      li {
        list-style-type: none;
        width: 17%;
        height: 100%;
        color: #fff;
        border-radius: 2px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        -moz-background-size: 100% 100%;
        span {
          font-size: 16px;
          text-align: center;
          &:nth-child(1) {
            margin-bottom: 7px;
          }
          .number {
            font-size: 32px;
            margin-right: 3px;
          }
        }
        &:nth-child(1) {
          background-image: url("../../images/xiaofeidingdan.png");
        }
        &:nth-child(2) {
          background-image: url("../../images/xiaofeiyingfijine.png");
        }
        &:nth-child(3) {
          background-image: url("../../images/xiaofeishishoujine.png");
        }
        &:nth-child(4) {
          background-image: url("../../images/chongzhidingdanliang.png");
        }
        &:nth-child(5) {
          background-image: url("../../images/chongzhijine.png");
        }
      }
    }
  }
  .middle-container {
    position: relative;
    width: 100%;
    height: 200px;
    padding: 5px 0;
    box-sizing: border-box;
    background-color: #f2f2f5;
    .inner {
      float: left;
      height: 190px;
      padding: 0 10px;
      background-color: #fff;
      box-sizing: border-box;
    }
    .left {
      position: relative;
      width: 52.5%;
      display: flex;
      flex-direction: column;
      .device-container {
        width: 100%;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        .device-content {
          width: 35%;
          height: 100px;
          margin-top: 10px;
          margin-left: -40%;
          background: rgba(255, 255, 255, 1);
          border-radius: 2px;
          box-shadow: 0px 0px 3px 1px rgba(180, 144, 232, 0.36);
          display: flex;
          flex-direction: row;
          justify-content: space-around;
          .device-left {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            img {
              width: 35px;
              height: 35px;
              margin-bottom: 10px;
            }
            span {
              font-size: 16px;
              font-family: Microsoft YaHei;
              font-weight: 400;
              color: rgba(153, 153, 153, 1);
              line-height: 13px;
              opacity: 0.87;
            }
          }
          .device-right {
            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;
            font-size: 36px;
            font-family: Microsoft YaHei;
            font-weight: bold;
            color: rgba(48, 143, 255, 1);
            line-height: 27px;
            opacity: 0.87;
          }
        }
      }
      ul {
        position: absolute;
        top: 0;
        right: 10px;
        height: 140px;
        width: 43%;
        overflow-y: auto;
        li {
          list-style-type: square;
          width: 100%;
          height: 24px;
          line-height: 24px;
          font-size: 14px;
          margin-left: -20px;
          color: #333;
          &:nth-child(1) {
            list-style-type: none;
            width: 100%;
            font-size: 16px;
            height: 26px;
            line-height: 26px;
            margin-left: -40px;
            background: rgba(244, 244, 244, 1);
            border-radius: 2px;
            text-align: center;
            margin-bottom: 5px;
          }
        }
      }
    }
    .right {
      margin-left: 0.5%;
      width: 47%;
      #pieChart {
        width: 90%;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
  .bottom-container {
    width: 100%;
    height: 300px;
  }
}
</style>
