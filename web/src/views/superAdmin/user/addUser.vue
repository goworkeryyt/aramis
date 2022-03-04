<!--
 * @Description: 用户信息：添加
-->
<template>
  <div id="busAdd">
    <div class="crumbs-header row-between header-sticky h-70">
      <div class="title">
        用户信息/ {{ this.actionType == 'edit' ? '编辑' :'新增' }}
      </div>
      <div>
        <el-button class="cancel-btn" @click="resetForm('dialogForm')">取 消</el-button>
        <el-button class="submit" @click="submitForm('dialogForm')">确 定</el-button>
        <i class="el-icon-close ml-10" @click="resetForm('dialogForm')" />
      </div>
    </div>
    <el-form ref="dialogForm" :rules="rules" :model="dialogForm" label-width="150px">
      <div class="theme row-flex">
        <i />
        <span>基本信息</span>
      </div>

      <div class="form_box">
        <el-form-item label="用户名：" prop="username">
          <el-input v-model="dialogForm.username" placeholder="请输入用户名" :disabled="actionType == 'edit'" clearable class="input-width" />
        </el-form-item>
        <el-form-item label="昵称：" prop="nickName">
          <el-input v-model="dialogForm.nickName" placeholder="请输入昵称" clearable class="input-width" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="dialogForm.phone" class="input-width" placeholder="请输入手机号" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="dialogForm.email" class="input-width" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="用户角色：" placeholder="请选择用户角色" prop="roleIds">
          <el-select v-model="dialogForm.roleIds" class="input-width" clearable multiple placeholder="请选择用户角色">
            <el-option v-for="(item,index) in userRoles" :key="item.id" :label="item.roleName" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isAdmin" label="用户类型" prop="userType">
          <el-select v-model="dialogForm.userType" class="input-width" :disabled="actionType == 'edit'" clearable placeholder="请选择用户类型" @change="changeUserType">
            <el-option v-for="(item,index) in userTypes" :key="index" :label="item" :value="index" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isAdmin && dialogForm.userType == 'MERT'" label="商户" prop="merchantNo">
          <el-select v-model="dialogForm.merchantNo" class="input-width" clearable placeholder="请选择商户" :disabled="actionType == 'edit'">
            <el-option v-for="(item,index) in merchantList" :key="item.id" :label="item.merchantName" :value="item.merchantNo" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isAdmin && dialogForm.userType == 'CORP'" label="运营单位" prop="corpNo">
          <el-select v-model="dialogForm.corpNo" class="input-width" clearable placeholder="请选择运营单位" :disabled="actionType == 'edit'">
            <el-option v-for="(item,index) in corpList" :key="index" :label="item.corpName" :value="item.corpNo" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isAdmin && dialogForm.userType == 'SUPP'" label="供应商" prop="supplierNo">
          <el-select v-model="dialogForm.supplierNo" class="input-width" clearable placeholder="请选择供应商" :disabled="actionType == 'edit'">
            <el-option v-for="(item,index) in supplierList" :key="index" :label="item.supplierName" :value="item.supplierNo" />
          </el-select>
        </el-form-item>
        <el-form-item label="头像：">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="dialogForm.headerImg" class="header-img-box" :src="headerPath">
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
      </div>
    </el-form>
    <ChooseImg ref="chooseImg" @getImg="getImg" />
  </div>
</template>
<script>
import ChooseImg from '@/components/chooseImg/index.vue'

export default {
  name: 'AddUser',
  components: {
    ChooseImg
  },
  props: {
    userTypes: {
      type: Object,
      default: () => { return [] }
    },
    userRoles: {
      type: Array,
      default: () => { return {} }
    },
    count: {
      type: Number,
      default: () => { return 0 }
    },
    actionType: {
      type: String,
      default: () => { return '' }
    },
    rowInfo: {
      type: Object,
      default: () => { return {} }
    }
  },

  data() {
    return {
      showMert: false,
      userRoleList: [],
      merchantList: [],
      corpList: [],
      supplierList: [],
      headerPath: '',
      dialogForm: {
        username: '',
        nickName: '',
        phone: '',
        email: '',
        nickName: '',
        headerImg: '',
        roleIds: [],
        userType: '',
        merchantNo: '',
        corpNo: '',
        supplierNo: ''
      },

      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 5, message: '用户名最低5位字符', trigger: 'blur' }
        ],
        phone: [{
          required: false,
          message: '手机号码不能为空！',
          trigger: 'blur'
        }, {
          pattern: /^1[0-9]{10}$/,
          message: '非法的手机号码！'
        }],
        email: [{
          required: false,
          message: '请输入用户邮箱',
          trigger: 'blur'
        }, {
          min: 6,
          max: 32,
          message: '邮箱地址为6-32个字符!'
        },
        {
          pattern: /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$/,
          message: '非法的邮箱地址！'
        }
        ],
        nickName: [
          { required: false, message: '请输入用户昵称', trigger: 'blur' }
        ],
        roleIds: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ],
        userType: [{
          required: true,
          message: '请选择用户类型',
          trigger: 'change'
        }],
        merchantNo: [{
          required: true,
          message: '请选择商户',
          trigger: 'change'
        }],
        corpNo: [{
          required: true,
          message: '请选择运营单位',
          trigger: 'change'
        }],
        supplierNo: [{
          required: true,
          message: '请选择供应商',
          trigger: 'change'
        }]
      }
    }
  },
  computed: {
    isAdmin() {
      const user = JSON.parse(sessionStorage.getItem('user'))
      return user.userType === 'ADMI' || user.userType === 'MERT'
    }
  },
  watch: {
    count: {
      handler(newVal) {
        this.getAllMerchants()
        this.findAllCorps()
        this.showInfo()
      }
    }
  },
  mounted() {
    this.getAllMerchants()
    this.findAllCorps()
    this.showInfo()
  },

  methods: {
    showInfo() {
      if (this.actionType == 'edit') {
        this.dialogForm.username = this.rowInfo.username
        this.dialogForm.nickName = this.rowInfo.nickName
        this.dialogForm.phone = this.rowInfo.phone
        this.dialogForm.email = this.rowInfo.email
        this.dialogForm.roleIds = this.rowInfo.roleIds
        this.dialogForm.userType = this.rowInfo.userType ? this.rowInfo.userType : ''
        this.dialogForm.headerImg = this.rowInfo.headerImg
        this.headerPath = this.rowInfo.headerImg
          ? this.rowInfo.headerImg.slice(0, 4) == 'http'
            ? this.rowInfo.headerImg
            : this.Constant.headerImgMap[this.rowInfo.headerImg] : ''
      }
    },

    // 获取所有已添加的商户列表（不分页）
    async getAllMerchants() {
      const res = await this.Api.getAllMerchants()
      const { code, message, content } = res.data

      if (code == 0) {
        this.merchantList = res.data.content
        this.dialogForm.merchantNo = this.rowInfo.merchantNo
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    // 获取所有运营信息列表（不分页）
    async findAllCorps() {
      const res = await this.Api.findAllCorps()
      const { code, message, content } = res.data

      if (code == 0) {
        this.corpList = res.data.content
        this.dialogForm.corpNo = this.rowInfo.corpNo
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    // 获取所有运营信息列表（不分页）---暂无接口
    // async findAllSuppliers() {
    //   const res = await this.Api.findAllSuppliers()
    //   const { code, message, content } = res.data;

    //   if (code == 0) {
    //     this.supplierList = res.data.content;
    //     this.dialogForm.supplierNo = this.rowInfo.supplierNo;
    //   }else{
    //     this.$message({
    //       type: "warning",
    //       center: true,
    //       message,
    //     });
    //   }
    // },

    changeUserType(value) {
      if (value == 'ADMI') {
        this.showMert = false
      } else {
        this.showMert = true
      }

      this.resetNo()
    },

    resetNo() {
      this.dialogForm.merchantNo = ''
      this.dialogForm.corpNo = ''
      this.dialogForm.supplierNo = ''
    },

    openHeaderChange() {
      this.$refs.chooseImg.open()
    },

    getImg(item) {
      this.headerPath = item.url
      this.dialogForm.headerImg = item.urlType
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

    resetFields() {
      this.headerPath = ''
      this.dialogForm.headerImg = ''
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
<style lang="scss" scoped>
.header-img-box {
  width: 150px;
  height: 150px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 150px;
  cursor: pointer;
  color: #9B9E9F;
}
</style>
