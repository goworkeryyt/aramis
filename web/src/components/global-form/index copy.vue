<!--
 * @Description: form表单组件封装
-->
<template>
  <div class="global_form" :style="{'padding':layout.padding||'20px'}" style="overflow-y:auto;overflow-x:hidden;height:100%">
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
              <!-- 输入框 input-text  input-password textarea START-->
              <el-input
                v-if="item.type==='text'||item.type==='password'||item.type==='textarea'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                clearable
                :type="item.type"
                :maxlength="item.maxlength||50"
                :size="layout.size"
                :placeholder="item.placeholder"
              >
                <!-- 带点击事件的icon（复合型输入框，图标在输入框外部） -->
                <el-button
                  v-if="item.icon && item.icon!==''"
                  slot="append"
                  :icon="item.icon"
                  @click="slotIcon(item,index)"
                />
                <!-- 静态展示的icon（图标在输入框内部） -->
                <template v-if="item.icon && item.suffixIcon!==''" slot="suffix">
                  <i :class="item.suffixIcon" />
                </template>
              </el-input>
              <!-- input-text  input-password textarea END-->

              <!-- 数字框 input-integer START -->
              <el-input
                v-if="item.type==='integer'"
                v-model="ruleForm[item.prop]"
                :disabled="item.disabled"
                clearable
                :type="item.type"
                :maxlength="item.maxlength||50"
                :size="layout.size"
                :placeholder="item.placeholder"
                @input="inputFn(index,item)"
              />
              <!-- input-integer END -->

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
                @change="change(index,item)"
              >
                <el-option
                  v-for="(select,selectIdx) in item.selectData"
                  :key="selectIdx"
                  :label="select.label"
                  :value="select.value"
                />
              </el-select>
              <!-- select END -->

              <!-- 日期、日期时间选择框 picker START -->
              <!-- 日期 datetime / date -->
              <el-date-picker
                v-if="item.type==='datePicker'"
                v-model="ruleForm[item.prop]"
                :type="item.dateType||'date'"
                :format="item.dateFormat"
                :value-format="item.dateFormat"
                :disabled="item.disabled"
                :placeholder="item.placeholder"
                :size="layout.size"
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
                style="width: 100%;"
              />
              <!-- picker END -->

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

              <!-- 颜色选择器 colorPicker START -->
              <el-color-picker
                v-if="item.type==='colorPicker'"
                v-model="ruleForm[item.prop]"
                :size="layout.size"
                @change="change($event,item.prop)"
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
    layout: {
      type: Object,
      default: () => {
        return {
          gutter: 20,
          span: 6,
          size: ''
        }
      }
    },
    formValue: Object,
    data: Array,
    rule: Object
  },
  data() {
    return {
      checkBoxList: [],
      visible: false,
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
        textarea: 'String'
      },
      dataType: {
        Array: [],
        String: '',
        Object: {}
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
             * @description: 解决设置multiple属性后，页面初始化会直接校验多选框的验证的问题
             */
    this.$nextTick(() => {
      this.$refs.ruleForm && this.$refs.ruleForm.resetFields()
    })
  },
  methods: {
    // 上传照片
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
    // 赋值：获取表单结果值
    setFormValue(data) {
      const result = {}
      const isDefaultData = this.GlobalFunction.dataType(data) === 'Object' && JSON.stringify(data) !== '{}'
      this.formData.map((item) => {
        item.value = isDefaultData ? data[item.prop] : this.setValue(item)
        result[item.prop] = item.value
      })
      this.ruleForm = result
      this.imgUrl = this.ruleForm.peopleImg
      this.setFormData()
    },
    setFormData() {
      this.formData.map((item) => {
        item.show = this.GlobalFunction.dataType(item.show) === 'Boolean' ? item.show : true
        item.id = !item.id ? this.GlobalFunction.genID(18) : item.id
      })
    },
    setValue(data) {
      const valType = this.formType[data.type]
      if (this.GlobalFunction.dataType(data.value) !== valType) {
        const value = this.dataType[valType]
        return value
      } else {
        return data.value
      }
    },
    // 设置单个select组件 数据
    setSelectData(data, prop) {
      const index = this.formData.findIndex((val) => val.prop === prop)
      this.formData[index].selectData = data
      this.$forceUpdate()
    },
    // 获取到表单的属性值（返回对象：key-value键值对形式）
    getRuleForm() {
      return this.ruleForm
    },
    // 提交表单并且进行表单校验，返回校验结果
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
    // 省市区组件传出来的方法
    changeCode(data, prop) {
      this.ruleForm[prop] = data
    },
    // 富文本编辑器组件传出来的方法
    changeEditor(data, prop) {
      this.ruleForm[prop] = data
    },
    // 下拉框切换事件
    change(index, item) {
      this.$emit('change', {
        value: this.ruleForm[item.prop],
        index: index,
        prop: item.prop
      })
    },
    // 数字输入框实时变化时的方法
    inputFn(index, item) {
      this.$emit('input', {
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
    // 重置表单
    resetForm() {
      if (this.$refsuploadImg) {
        this.$refs.uploadImg[0].deleteIcon() // 清除上传的照片框
      }
      this.resetArea()
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
    .global_form .el-input, .global_form .el-select {
        width: 70%;
    }
    .global_form .el-date-editor.el-input, .global_form .el-date-editor.el-input__inner {
        width: 70% !important;
    }
    .global_form .el-select .el-input {
        width: 100%;
    }
    .mapinput .el-input-group__append {
        background-color: #7cb9f8;
        color: #fff;
        border-color: #7cb9f8;
    }
    .mapinput .el-input-group__append:hover {
        background-color: #66b1ff;
    }
</style>
