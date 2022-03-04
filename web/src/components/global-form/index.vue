<!--
 * @Description: form表单组件封装
-->
<template>
  <div class="global_form" :style="{'padding':layout.padding||'20px'}" style="overflow-y:auto;overflow-x:hidden;height:100%;">
    <el-form
      ref="ruleForm"
      :model="ruleForm"
      :rules="rules"
      :label-width="layout.labelWidth||'100px'"
      class="demo-ruleForm"
    >
      <el-row :gutter="layout.gutter">
        <template v-for="(item,index) in formData">
          <el-col
            v-show="item.show"
            :key="`${item.id}-${item.show}`"
            :span="item.span||layout.span"
          >
            <el-form-item :label="item.label" :prop="item.prop" :class="item.class">
              <!-- 文本输入框 input-text  input-password input-number textarea START-->
              <el-input
                v-if="item.type==='text'||item.type==='password'||item.type==='textarea'||item.type==='number'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                clearable
                :type="item.type"
                :maxlength="item.maxlength||50"
                :size="layout.size"
                :placeholder="item.placeholder"
                @input="inputFn($event,index,item)"
              >
                <!-- 带点击事件的icon（复合型输入框，图标在输入框外部） -->
                <el-button
                  v-if="item.icon && item.icon!==''"
                  slot="append"
                  :icon="item.icon"
                  @click="slotIcon(item,index)"
                />
                <!-- 静态展示的icon（图标在输入框内部） -->
                <template v-if="item.suffixIcon && item.suffixIcon!==''" slot="suffix">
                  <i :class="item.suffixIcon" />
                </template>
              </el-input>
              <!-- input-text  input-password textarea END-->

              <!-- 计数器 input-number START -->
              <el-input-number
                v-if="item.type==='counter'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                :placeholder="item.placeholder"
                :type="item.type"
                :min="item.min"
                :max="item.max"
                :step="item.step"
                :precision="item.precision"
                :controls="item.controls"
                :size="layout.size"
                @change="changeCounter($event,index,item)"
              />
              <!-- input-number END -->

              <!-- 下拉选择框 select START-->
              <el-select
                v-if="item.type==='select'"
                v-model="ruleForm[item.prop]"
                :popper-append-to-body="true"
                :multiple="item.multiple"
                :disabled="item.disabled"
                clearable
                :size="layout.size"
                :placeholder="item.placeholder"
                @change="changeSelect($event,index,item)"
              >
                <el-option
                  v-for="(select,selectIdx) in item.selectData"
                  :key="selectIdx"
                  :label="select.label"
                  :value="select.value"
                />
              </el-select>
              <!-- select END -->

              <!-- 单选框 radio START -->
              <el-radio-group
                v-if="item.type==='radio'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                :size="layout.size"
              >
                <el-radio
                  v-for="(radio,radioIdx) in item.radioData"
                  :key="radioIdx"
                  :label="radio.label"
                >{{ radio.value }}</el-radio>
              </el-radio-group>
              <!-- radio END -->

              <!-- 复选框 checkbox START -->
              <el-checkbox-group
                v-if="item.type==='checkbox'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                :size="layout.size"
              >
                <el-checkbox
                  v-for="(checkbox,checkboxIdx) in item.checkboxData"
                  :key="checkboxIdx"
                  :label="checkbox.label"
                >{{ checkbox.value }}</el-checkbox>
              </el-checkbox-group>
              <!-- checkbox END -->

              <!-- 日期、日期时间选择框 picker START -->
              <!-- 日期 datetime / date / daterange / datetimerange -->
              <el-date-picker
                v-if="item.type==='datePicker'"
                v-model="ruleForm[item.prop]"
                :type="item.dateType||'date'"
                :format="item.dateFormat"
                :value-format="item.dateFormat"
                :disabled="item.disabled"
                :placeholder="item.placeholder"
                :size="layout.size"
                :range-separator="item.rangeSeparator"
                :start-placeholder="item.startPlaceholder"
                :end-placeholder="item.endPlaceholder"
                style="width: 100%;"
              />
              <!-- 时间点 time -->
              <el-time-picker
                v-if="item.type==='timePicker'"
                v-model="ruleForm[item.prop]"
                :size="layout.size"
                :format="item.dateFormat"
                :value-format="item.dateFormat"
                :picker-options="item.pickerOptions"
                :disabled="item.disabled"
                :placeholder="item.placeholder"
                :is-range="item.isRange"
                :range-separator="item.rangeSeparator"
                :start-placeholder="item.startPlaceholder"
                :end-placeholder="item.endPlaceholder"
                style="width: 100%;"
              />
              <!-- picker END -->

              <!-- 级联选择器 cascader START -->
              <el-cascader
                v-if="item.type == 'cascader'"
                ref="myCascader"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                clearable
                :size="layout.size"
                :options="item.cascaderData"
                style="width: 100%;"
                @change="changeCascader($event,index,item)"
                @visible-change="visibleChange($event,index,item)"
              />
              <!-- cascader END -->

              <!-- 颜色选择器 colorPicker START -->
              <el-color-picker
                v-if="item.type==='colorPicker'"
                v-model="ruleForm[item.prop]"
                :size="layout.size"
                @change="changeColor($event,index,item)"
              />
              <!-- colorPicker END -->

              <!-- 上传图片 upload-img START -->
              <global-upload-img
                v-if="item.type==='uploadImg'"
                ref="uploadImg"
                :img-url="imgUrl"
                @uploadImage="uploadImage"
              />
              <!-- upload-img END -->

              <!-- 省市区三级联动选择 city START -->
              <global-city
                v-if="item.type==='region'"
                ref="selectArea"
                :code="item.value"
                :disabled="item.disabled"
                :size="layout.size"
                @code="changeCode($event,item.prop)"
              />
              <!-- city END -->

              <!-- 富文本编辑器 editor START -->
              <global-editor
                v-if="item.type==='editor'"
                ref="wangEditor"
                :value="item.value"
                :disabled="item.disabled"
                :placeholder="item.placeholder"
                @editor="changeEditor($event,item.prop)"
              />
              <!-- editor END -->
            </el-form-item>
          </el-col>
        </template>
      </el-row>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'GlobalForm',
  props: {
    // 表单样式
    layout: {
      type: Object,
      default: () => {
        return {
          padding: '20px',
          labelWidth: '100px',
          gutter: 20,
          span: 6,
          size: ''
        }
      }
    },
    // 表单的所有树形值对象（即：提交时的结果对象）
    formValue: {
      type: Object,
      default: () => {
        return {}
      }
    },
    // 表单组件类型的数据属性
    data: {
      type: Array,
      default: () => {
        return []
      }
    },
    // 表单组件的校验规则
    rule: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data() {
    return {
      formData: [],
      ruleForm: {},
      rules: {},
      formType: {
        editor: 'String',
        city: 'Array',
        uploadImg: 'Array',
        picker: 'String',
        text: 'String',
        select: 'String',
        radio: 'String',
        checkbox: 'Array',
        textarea: 'String',
        cascader: 'Array',
        counter: 'Number'
      },
      dataType: {
        Array: [],
        String: '',
        Object: {},
        Number: 0
      },
      imgUrl: ''
    }
  },
  watch: {
    formValue(val) {
      this.setFormValue(val)
    }
  },
  mounted() {
    const data = this.GlobalFunction.deepcopy(this.data)
    this.formData = data
    this.setFormValue(this.formValue)
    this.rules = this.rule
    /**
             * @description: 解决页面初始化会直接校验多选框的验证的问题
             */
    this.$nextTick(() => {
      this.$refs.ruleForm && this.$refs.ruleForm.resetFields()
    })
  },
  methods: {
    // 赋值：获取表单结果值
    setFormValue(data) {
      const result = {}
      this.formData.map((item) => {
        item.value = data[item.prop]
        result[item.prop] = item.value
      })
      this.ruleForm = result
      this.imgUrl = this.ruleForm.peopleImg
      this.setFormData()
    },
    // 给每个对象赋值（id唯一值 和 show是否展示）
    setFormData() {
      this.formData.map((item) => {
        item.show = this.GlobalFunction.dataType(item.show) === 'Boolean' ? item.show : true
        item.id = !item.id ? this.GlobalFunction.genID(18) : item.id
      })
    },
    // input 输入框实时改变时的方法
    inputFn(data, index, item) {
      this.ruleForm[item.prop] = data
      this.$emit('inputFn', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
    },
    // 点击输入框后的icon触发事件
    slotIcon(data, index) {
      this.$emit('slotIcon', {
        data: data,
        index: index
      })
    },
    // input-number 计数器变化时的方法
    changeCounter(data, index, item) {
      this.ruleForm[item.prop] = data
      this.$emit('changeCounter', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
    },
    // select 下拉框切换事件
    changeSelect(data, index, item) {
      this.$emit('changeSelect', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
    },
    // cascader 级联选择框切换事件
    changeCascader(data, index, item) {
      this.ruleForm[item.prop] = data
      this.$emit('changeCascader', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
      if (item.type == 'cascader' && this.ruleForm[item.prop] && this.ruleForm[item.prop].length == 0) {
        this.$refs.myCascader[0].$refs.panel.checkedValue = [] // 清空选中值（有没有这句代码都行，清空主要是下面三行）
        this.$refs.myCascader[0].$refs.panel.clearCheckedNodes() // 清空级联选择器选中状态
        this.$refs.myCascader[0].$refs.panel.activePath = [] // 清除高亮
        this.$refs.myCascader[0].$refs.panel.syncActivePath() // 初始化（只展示一级节点）
      }
    },
    // cascader 级联下拉框出现/隐藏时触发(第一个参数：出现则为 true，隐藏则为 false)
    visibleChange(data, index, item) {
      this.formData[index].value = this.ruleForm[item.prop]
      // 点击输入框后的清空按钮后，第二次点击弹出下拉框选项时需清空选中值和样式，默认初始化显示
      if (data && this.ruleForm[item.prop].length == 0) {
        this.$refs.myCascader[0].$refs.panel.checkedValue = [] // 清空选中值（有没有这句代码都行，清空主要是下面三行）
        this.$refs.myCascader[0].$refs.panel.clearCheckedNodes() // 清空级联选择器选中状态
        this.$refs.myCascader[0].$refs.panel.activePath = [] // 清除高亮
        this.$refs.myCascader[0].$refs.panel.syncActivePath() // 初始化（只展示一级节点）
      }
    },
    // colorPicker 颜色选择器切换事件
    changeColor(data, index, item) {
      this.ruleForm[item.prop] = data
      this.$emit('changeColor', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
    },
    // 上传照片组件传出来的方法
    uploadImage(img) {
      this.imgUrl = img
      this.formData.map((item) => {
        if (item.type == 'uploadImg') {
          item.value = img
          this.ruleForm[item.prop] = img
        }
      })
      this.$emit('uploadImage', img)
    },
    // 省市区组件传出来的方法
    changeCode(data, prop) {
      this.ruleForm[prop] = data
    },
    // 富文本编辑器组件传出来的方法
    changeEditor(data, prop) {
      this.ruleForm[prop] = data
    },
    // 提交表单并且进行表单校验，返回校验结果（成功：true；失败：false）
    submitForm() {
      let flag = null
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          flag = true
        } else {
          flag = false
        }
      })
      return flag
    },
    // 获取到表单的属性值（返回表单结果的对象：key-value键值对形式）
    getRuleForm() {
      return this.ruleForm
    },
    // 重置表单
    resetForm() {
      if (this.$refsuploadImg) {
        this.$refs.uploadImg[0].deleteIcon() // 清除上传的照片显示框
      }
      this.resetArea() // 清除省市区
      if (this.$refs.wangEditor) {
        this.$refs.wangEditor[0].setHtml() // 清空富文本框的内容
      }
      this.$refs.ruleForm.resetFields() // 重置所有表单项为初始值并且移除所有校验
    },
    // 清除省市区
    resetArea() {
      if (this.$refs.selectArea) {
        this.$refs.selectArea[0].resetArea() // 清除省市区选择框
      }
    }
  }
}
</script>
<style>
    .global_form .el-input, .global_form .el-select, .global_form .el-input-number {
        width: 70%;
    }
    .global_form .el-date-editor.el-input, .global_form .el-date-editor.el-input__inner {
        width: 70% !important;
    }
    .global_form .el-select .el-input {
        width: 100%;
    }
    .global_form .el-input-number .el-input {
        width: 100%;
    }
    .global_form .el-input-number .el-input-number__increase, .global_form .el-input-number .el-input-number__decrease {
        line-height: 30px;
        top: 3px;
    }
    .global_form .el-range-editor.el-input__inner {
        height: 32px;
        border-radius: 2px;
        border-color: #C9C9C9;
    }
    .global_form .el-date-editor .el-range__icon, .global_form .el-date-editor .el-range-separator, .global_form .el-date-editor .el-range__close-icon {
        line-height: 25px;
    }
    /* 日期时间区间后面的清空按钮右移到最右边显示 */
    .global_form .el-date-editor--daterange .el-range__close-icon,
    .global_form .el-date-editor--datetimerange .el-range__close-icon,
    .global_form .el-date-editor--timerange .el-range__close-icon{
        position: absolute;
        right: 3px;
        top: 2px;
    }
    /**
    * 解决el-input设置类型为number时，去掉输入框后面上下箭头
    **/
    .global_form input[type=number]::-webkit-inner-spin-button, .global_form input[type=number]::-webkit-outer-spin-button {
	    -webkit-appearance: none;
	    margin: 0;
	}
    /**
    * 解决el-input设置类型为number时，中文输入法光标上移问题
    **/
    .global_form .el-input__inner{
        line-height: 1px !important;
    }
</style>
