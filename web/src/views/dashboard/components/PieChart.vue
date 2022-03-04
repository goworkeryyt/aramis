<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons')
import { debounce } from '@/utils'
const animationDuration = 6000

export default {
  name: 'SalasPie',
  props: {
    className: {
      type: String,
      default: 'pie-chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '130px'
    },
    chartType: {
      type: String
    },
    dataPie: {
      type: Array
    }
  },
  data() {
    return {
      chart: null,
      pData: []
    }
  },
  watch: {
    dataPie(val) {
      this.pData = val

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
        title: {
          x: 'left',
          textStyle: {
            color: '#409EFF',
            fontSize: '16'
          }
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        calculable: true,
        legend: {
          orient: 'vertical',
          right: 'right',
          data: ['消费订单量', '充值订单量', '退款订单量']
        },
        series: [
          {
            name: '本月订单统计',
            type: 'pie',
            radius: '55%',
            center: ['50%', '50%'],
            color: ['#2589FF', '#FB7171', '#FFCC55'],
            label: {
              position: 'outside'
            },
            data: this.pData
          }
        ]
      })
    }
  }
}
</script>
