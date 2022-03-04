<!--
 * @Description: 商户信息：添加、编辑
-->
<template>
  <div id="merchantAdd">
    <div class="crumbs-header row-between header-sticky h-70">
      <div class="title">
        {{ actionType == "edit" ? "商户/编辑" : "商户/新增" }}
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
        <el-form-item label="商户编号：" prop="merchantNo">
          <el-input v-model="dialogForm.merchantNo" :disabled="actionType == 'edit'" clearable placeholder="请输入商户编号" class="input-width" />
        </el-form-item>
        <el-form-item label="商户名称：" prop="merchantName">
          <el-input v-model="dialogForm.merchantName" clearable placeholder="请输入商户名称" class="input-width" />
        </el-form-item>
        <el-form-item label="地址：" prop="address">
          <el-input v-model="dialogForm.address" clearable placeholder="请输入地址" class="input-width" />
        </el-form-item>
        <!-- <el-form-item label="区域代码：" prop="regionCode">
                    <el-input v-model="dialogForm.regionCode" clearable placeholder="请输入区域代码" class="input-width"/>
                </el-form-item>
                <el-form-item label="区域名称：" prop="regionName">
                    <el-input v-model="dialogForm.regionName" clearable placeholder="请输入区域名称" class="input-width"/>
                </el-form-item> -->
        <el-form-item v-if="actionType == 'edit'" label="状态：" prop="status">
          <el-radio-group v-model="dialogForm.status">
            <el-radio v-for="(label, value) in Constant.merchantStatus" :key="value" :label="value">{{ label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="收单机构代码：" prop="instId">
          <el-input v-model="dialogForm.instId" clearable placeholder="请输入收单机构代码" class="input-width" />
        </el-form-item>
        <el-form-item label="本商户密钥：" prop="selfRsaPriKey">
          <el-input v-model="dialogForm.selfRsaPriKey" clearable placeholder="请输入本商户密钥" class="input-width" />
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'MerchantAdd',
  props: {
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
      dialogForm: {
        merchantNo: '',
        merchantName: '',
        address: '',
        // regionCode: '',
        // regionName: '',
        status: '00',
        instId: '',
        selfRsaPriKey: ''
      },
      // 状态  00：启用  01：停用
      statusList: [
        { label: '00', value: '启用' },
        { label: '01', value: '停用' }
      ],
      formRolues: {
        merchantNo: [
          { required: true, message: '商户编号不能为空！', trigger: 'blur' },
          {
            max: 32,
            message: '商户编号不超过32位字符!'
          }
        ],
        merchantName: [
          { required: true, message: '商户名称不能为空！', trigger: 'blur' },
          {
            pattern: /^[\u4E00-\u9FA5A-Za-z0-9]{2,30}$/,
            message: '商户名称为2-30位中文或字母数字组合!'
          }
        ],
        address: [
          { required: true, message: '地址不能为空！', trigger: 'blur' },
          {
            max: 100,
            message: '地址不超过100位字符!'
          }
        ],
        // regionCode: [
        //     { required: true, message: '区域代码不能为空！', trigger: 'blur' }
        // ],
        // regionName: [
        //     { required: true, message: '区域名称不能为空！', trigger: 'blur' }
        // ],
        status: [
          { required: true, message: '状态不能为空！', trigger: 'change' }
        ],
        // instId: [
        //     { required: true, message: '收单机构代码不能为空！', trigger: 'blur' }
        // ],
        selfRsaPriKey: [
          { required: false, message: '本商户密钥不能为空！', trigger: 'blur' }, {
            pattern: /^[A-Za-z0-9]{32}$/,
            message: '商户密钥为32位数字和大小写字母组合！'
          }
        ]
      }
    }
  },
  computed: {
    getShow(val) {
      return function(val) {
        if (this.actionType == 'edit') {
          return this.actionType == 'edit' && this.dialogForm.level == val
        } else {
          return true
        }
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
      this.dialogForm.merchantNo = this.rowInfo.merchantNo
      this.dialogForm.merchantName = this.rowInfo.merchantName
      this.dialogForm.address = this.rowInfo.address
      // this.dialogForm.regionCode = this.rowInfo.regionCode;
      // this.dialogForm.regionName = this.rowInfo.regionName;
      this.dialogForm.status = this.rowInfo.status || '00'
      this.dialogForm.instId = this.rowInfo.instId || ''
      this.dialogForm.selfRsaPriKey = this.rowInfo.selfRsaPriKey || ''
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
      this.dialogForm.parentId = ''
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
