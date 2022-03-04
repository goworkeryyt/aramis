<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons')
import { debounce } from '@/utils'
const animationDuration = 6000

export default {
  name: 'SalasStack',
  props: {
    className: {
      type: String,
      default: 'rect-chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '360px'
    },
    dataRect: {
      type: Array
    },
    itemKey: {
      type: String
    }
  },
  data() {
    return {
      chart: null,
      xDatas: [],
      moneyDatas: [],
      numberDatas: [],
      moneyMax: 0,
      numberMax: 0
    }
  },
  watch: {
    dataRect(val) {
      val.map(item => {
        this.xDatas.push(item.date)
        this.moneyDatas.push((Number(item.data.realAmt) / 100).toFixed(2))
        this.numberDatas.push(Number(item.data.deviceNum))
      })

      this.moneyMax = Math.max.apply(null, this.moneyDatas) === 0 ? 1 : Math.max.apply(null, this.moneyDatas)
      this.numberMax = Math.max.apply(null, this.numberDatas) === 0 ? 1 : Math.max.apply(null, this.numberDatas)

      this.initChart()
      this.__resizeHanlder = debounce(() => {
        if (this.chart) {
          this.chart.resize()
        }
      }, 100)

      window.addEventListener('resize', this.__resizeHanlder)
    }
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }

    window.removeEventListener('resize', this.__resizeHanlder)

    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')

      this.chart.setOption({
        color: ['#61A0F0', '#F64468'],
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            crossStyle: {
              color: '#999'
            }
          }
        },
        legend: {
          top: '5%',
          data: ['设备消费金额', '设备数量']
        },
        xAxis: {
          type: 'category',
          data: this.xDatas,
          axisPointer: {
            type: 'shadow'
          },
          axisLine: {
            lineStyle: {
              color: '#333'
            }
          }
        },
        yAxis: [
          {
            type: 'value',
            scale: true,
            name: '金额（元）',
            min: 0,
            max: this.moneyMax,
            // interval: 50,
            // axisLabel: {
            //     formatter: '{value} 元'
            // },
            axisLine: {
              lineStyle: {
                color: '#333'
              }
            },
            axisTick: {
              show: false
            }
          },
          {
            type: 'value',
            scale: true,
            name: '数量（台）',
            min: 0,
            max: this.numberMax,
            // interval: 5,
            // axisLabel: {
            //     formatter: '{value} 台'
            // },
            axisLine: {
              lineStyle: {
                color: '#333'
              }
            },
            axisTick: {
              show: false
            }
          }
        ],
        series: [
          {
            name: '设备消费金额',
            type: 'bar',
            data: this.moneyDatas
          },
          {
            name: '设备数量',
            type: 'bar',
            data: this.numberDatas
          }
        ]
      })
    }
  }
}
</script>
