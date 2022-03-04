<!--
 * @Description: 搜索区组件
-->
<template>
  <div id="globalSearch">
    <el-form class="wid" :inline="true">
      <div
        v-for="(item,index) in formData"
        :key="`${item.id}-${index}`"
        class="dis_in_center mb-10"
        :style="{'width':item.width || '220px', 'min-width': item.min_width || '200px'}"
        :v-if="item.isShow"
      >
        <div class="flex">
          <!-- 搜索条件名称 label -->
          <div v-if="item.isShowLabel" class="flex-grow-0 t_r plr-10" :style="{width: labelWidth}">
            <span class="fon_13 label dis_in_center color_6">{{ item.label }}</span>
          </div>
          <!-- 搜索条件形式 -->
          <div class="flex-grow-1 flex-y-center mr-20" :style="{width: contentWidth}">
            <!-- 普通input框 text -->
            <el-input
              v-if="item.type=='text'"
              v-model="item.value"
              size="small"
              clearable
              :placeholder="item.placeholder"
              :value="item.value"
            />
            <!-- 普通input框 number -->
            <el-input
              v-if="item.type=='number'"
              v-model.number="item.value"
              size="small"
              clearable
              :placeholder="item.placeholder"
              :value="item.value"
            />
            <!-- 文本框 textarea -->
            <el-input
              v-if="item.type=='textarea'"
              v-model="item.value"
              size="small"
              autosize
              :maxlength="item.maxLength"
              type="textarea"
              :placeholder="item.placeholder"
            />
            <!-- 单选框 radio -->
            <div v-if="item.type=='radio'">
              <el-radio-group v-model="item.value" size="small" @change="changeRadio(index)">
                <el-radio v-for="(radio,radioIndex) in item.radioData" :key="`radio${radioIndex}`" :label="radio.label">
                  {{ radio.value }}
                </el-radio>
              </el-radio-group>
            </div>
            <!-- 复选框 checkBox -->
            <div v-if="item.type=='checkBox'">
              <el-checkbox-group v-model="item.value" size="small">
                <el-checkbox v-for="(checkBox,checkBoxIndex) in item.checkBoxData" :key="`checkBox${checkBoxIndex}`" :label="checkBox.label">
                  {{ checkBox.value }}
                </el-checkbox>
              </el-checkbox-group>
            </div>
            <!-- 选择器 select -->
            <div v-if="item.type=='select'">
              <el-select
                v-model="item.value"
                :filterable="item.filterable"
                size="small"
                clearable
                :multiple="item.multiple"
                :placeholder="item.placeholder"
                :disabled="item.disabled"
                @change="changeSelect($event,item,index)"
              >
                <el-option
                  v-for="(select,selectIndex) in item.selectData"
                  :key="`select${selectIndex}`"
                  :label="select.label"
                  :value="select.value"
                />
              </el-select>
            </div>
            <!-- 选择日期 date || 日期时间区间 datetimerange -->
            <div v-if="item.type=='date' || item.type=='datetimerange'" class="wid">
              <el-date-picker
                v-model="item.value"
                align="right"
                :type="item.type"
                :start-placeholder="item.start_placeholder"
                :end-placeholder="item.end_placeholder"
                :placeholder="item.placeholder"
                :value-format="item.format"
                :picker-options="item.pickerOptions"
              />
            </div>
            <!-- 省市区三级联动 -->
            <div v-if="item.type=='region'" class="wid">
              <global-city ref="selectArea" :value="item.value" @code="changeCode($event,index)" />
            </div>
            <!-- 选择输入框：select和input联用 -->
            <div v-if="item.type=='select-input'" class="wid">
              <el-input v-model="item.value" size="small" :placeholder="item.placeholder" :value="item.value" class="input-with-select" clearable>
                <el-select slot="prepend" v-model="item.inputSelectValue" placeholder="请选择" :disabled="item.disabled" @change="changeInputSelect($event,item,index)">
                  <el-option
                    v-for="(select,selectIndex) in item.inputSelectData"
                    :key="`select${selectIndex}`"
                    :label="select.label"
                    :value="select.value"
                  />
                </el-select>
              </el-input>
            </div>
          </div>
        </div>
      </div>
      <!-- 按钮（查询、刷新） -->
      <div class="dis_in_center mb-10">
        <el-button class="pan-btn light-blue-btn" type="primary" @click="search">查询</el-button>
        <el-button class="light-blue-btn-text" type="text" @click="reset">刷新</el-button>
      </div>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'GlobalSearch',
  props: {
    searchStyle: {
      type: Object,
      default: () => {
        return {}
      }
    },
    labelWidth: {
      type: String,
      default: ''
    },
    contentWidth: {
      type: String,
      default: '200px'
    }
  },
  data() {
    return {
      formData: []
    }
  },
  watch: {
    searchStyle: {
      handler(val) {
        if (val.searchForm.length > 1) {
          let flag = false
          val.searchForm.forEach(element => {
            if ((element.selectData && element.selectData.length == 0) || (element.inputSelectData && element.inputSelectData.length == 0)) {
              flag = true
            }
          })
          if (!flag) {
            this.formData = val.searchForm
          }
        }
      },
      deep: true
    }
  },
  mounted() {
    this.formData = this.GlobalFunction.deepcopy(this.searchStyle.searchForm)
  },
  methods: {
    // 省市区三级联动选择切换时触发的方法
    changeCode(data, index) {
      this.formData[index].value = data
      this.change(index)
    },
    // select下拉框切换时触发
    changeSelect(data, item, index) {
      this.change(index)
    },
    // 选择输入框的下拉框切换时触发
    changeInputSelect(data, item, index) {
      this.change(index)
    },
    // 搜索框的结值果：属性和值
    getResult() {
      const result = {}
      this.formData.map(item => {
        // 普通 属性和值
        result[item.prop] = item.value || ''
        // 选择输入框 的属性和值
        item.inputSelectProp ? (result[item.inputSelectProp] = item.inputSelectValue || '') : ''
      })
      return result
    },
    // 下拉框切换时，传给父组件的方法
    change(index) {
      this.$emit('searchChange', {
        data: this.formData,
        result: this.getResult(),
        index: index
      })
    },
    // 单选框切换时，传给父组件的方法
    changeRadio(index) {
      this.$emit('changeRadio', {
        data: this.formData,
        result: this.getResult(),
        index: index
      })
    },
    // 点击“查询”按钮，传给父组件的方法
    search() {
      this.$emit('confim', {
        result: this.getResult()
      })
    },
    // 点击“刷新”按钮，传给父组件的方法
    reset() {
      this.$message.closeAll()
      this.$message.success('重置成功')
      this.formData.map(item => {
        item.value = ''
      })
      this.$refs.selectArea[0].resetArea() // 清除省市区
      this.$emit('reset', {
        result: this.getResult()
      })
    }
  }
}
</script>
<style>
    #globalSearch .label {
        line-height: 32px;
    }
    #globalSearch .el-date-editor--date .el-input__inner {
        width: 200px !important;
    }
    #globalSearch .el-date-editor--datetimerange {
        width: 380px !important;
        height: 32px;
        line-height: 32px;
        border: 1px solid #C9C9C9;
        border-radius: 2px;
    }
    #globalSearch .el-date-editor .el-range-separator {
        line-height: 25px;
    }
    #globalSearch .el-date-editor--datetimerange .el-icon-time {
        line-height: 26px;
    }
    #globalSearch .el-date-editor--datetimerange .el-range__close-icon {
        line-height: 26px;
    }
    #globalSearch .el-date-editor {
        width: 200px;
    }
    #globalSearch .el-date-editor .el-range-input {
        width: 50%;
    }
    #globalSearch .input-with-select .el-select {
        width: 100px;
    }
    .el-icon-arrow-left:before {
        content: "\e6de";
    }
    .el-icon-arrow-right:before {
        content: "\e6e0";
    }
    .el-date-range-picker__time-header>.el-icon-arrow-right {
        font-size: 20px !important;
    }
    #globalSearch .el-input-group__append, .el-input-group__prepend {
        border-radius: 2px;
        border: 1px solid #C9C9C9;
    }
</style>
