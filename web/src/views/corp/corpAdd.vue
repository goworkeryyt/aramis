<!--
 * @Description: 商户信息：添加、编辑
-->
<template>
  <div id="merchantAdd">
    <div class="crumbs-header row-between header-sticky h-70">
      <div class="title">
        {{ actionType == "edit" ? "运营单位/编辑" : "运营单位/新增" }}
      </div>
      <div>
        <el-button class="cancel-btn" @click="resetForm('dialogForm')">取 消</el-button>
        <el-button class="submit" @click="submitForm('dialogForm')">确 定</el-button>
        <i class="el-icon-close ml-10" @click="resetForm('dialogForm')" />
      </div>
    </div>
    <el-form ref="dialogForm" :rules="formRolues" :model="dialogForm" label-width="150px">
      <div class="theme row-flex">
        <i />
        <span>基本信息</span>
      </div>

      <div class="form_box">
        <el-form-item label="运营单位编号：" prop="corpNo">
          <el-input v-model="dialogForm.corpNo" :disabled="actionType == 'edit'" clearable placeholder="请输入运营单位编号" class="input-width" />
        </el-form-item>
        <el-form-item label="运营单位名称：" prop="corpName">
          <el-input v-model="dialogForm.corpName" clearable placeholder="请输入运营单位名称" class="input-width" />
        </el-form-item>
        <el-form-item v-if="userType == 'ADMI' || userType == 'COMM'" label="关联商户：" prop="merchantNo">
          <el-select v-model="dialogForm.merchantNo" placeholder="请选择关联商户" :disabled="actionType == 'edit'" clearable class="input-width">
            <el-option v-for="item in merchantList" :key="item.id" :label="item.merchantName" :value="item.merchantNo" />
          </el-select>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'MerchantAdd',
  props: {
    merchantList: {
      type: Array,
      default: () => { return [] }
    },
    actionType: {
      type: String,
      default: () => { return '' }
    },
    rowInfo: {
      type: Object,
      default: () => { return {} }
    },
    count: {
      type: Number,
      default: () => { return 0 }
    }
  },
  data() {
    return {
      userType: JSON.parse(sessionStorage.getItem('user')).userType,
      dialogForm: {
        corpNo: '',
        corpName: '',
        merchantNo: ''
      },
      formRolues: {
        corpNo: [
          { required: true, message: '运营单位编号不能为空！', trigger: 'blur' },
          {
            max: 32,
            message: '运营单位编号不超过32位字符!'
          }
        ],
        corpName: [
          { required: true, message: '运营单位名称不能为空！', trigger: 'blur' },
          {
            pattern: /^[\u4E00-\u9FA5A-Za-z0-9]{2,30}$/,
            message: '运营单位名称为2-30位中文或字母数字组合!'
          }
        ],
        merchantNo: [
          { required: true, message: '商户名称不能为空！', trigger: 'change' }
        ]
      }
    }
  },
  watch: {
    count: {
      handler(newVal) {
        if (this.actionType == 'edit') {
          this.showInfo()
        }
      }
    }
  },
  mounted() {
    if (this.actionType == 'edit') {
      this.showInfo()
    }
  },
  methods: {
    showInfo() {
      this.dialogForm.corpNo = this.rowInfo.corpNo
      this.dialogForm.corpName = this.rowInfo.corpName
      this.dialogForm.merchantNo = this.rowInfo.merchantNo || ''
    },

    // 等弹窗里的form表单的dom渲染完再执行去除表单验证
    clearValidate() {
      this.$nextTick(() => {
        this.$refs['dialogForm'].clearValidate()
      })
    },

    // 点击表单的“取消”按钮
    resetForm(form) {
      this.$emit('resetForm')
    },

    // 重置表单
    resetFields() {
      this.$refs['dialogForm'].resetFields()
    },

    // 点击表单的“确定”按钮
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit('submitForm', this.dialogForm)
        } else {
          return false
        }
      })
    }
  }
}
</script>
<style lang='scss' scoped>

</style>
