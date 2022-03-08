<template>
  <div class="login-container">
    <div class="login-content">
      <!-- <img class="login-container-title" src="../../images/login/title.png" /> -->
      <div class="login">
        <div class="login-title">
          <div class="title">平台登录</div>
        </div>
        <el-form ref="loginForm" :rules="rules" :model="loginForm">
          <el-form-item prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入账号">
              <template slot="prepend">
                <img src="../../images/login/zhanghao.png">
              </template>
            </el-input>
            <div class="clear-container">
              <img
                src="../../images/login/guanbi.png"
                @click="loginForm.username = ''"
              >
            </div>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              :type="passwordType"
              auto-complete="on"
              placeholder="请输入密码"
            >
              <template slot="prepend">
                <img src="../../images/login/mima.png">
              </template>
            </el-input>
            <div class="clear-container">
              <img
                src="../../images/login/guanbi.png"
                @click="loginForm.password = ''"
              >
            </div>
          </el-form-item>
          <el-form-item prop="captcha">
            <el-input
              v-model="loginForm.captcha"
              auto-complete="on"
              placeholder="请输入验证码"
              @keyup.enter.native="handleSubmit"
            >
              <template slot="prepend">
                <img src="../../images/login/yanzhengma.png">
              </template>
            </el-input>
            <div class="checkcode-container" @click="getCaptcha">
              <img class="check-code-image" :src="codeImage">
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              :loading="loading"
              size="large"
              @click="handleSubmit('loginForm')"
            >
              登录
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    const validateUsername = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入账号'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入6位密码'))
      } else {
        callback()
      }
    }
    const validateCaptcha = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入6位验证码'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      codeImage: '',
      passwordType: 'password',
      loginForm: {
        username: '',
        password: '',
        captcha: '',
        captchaId: ''
      },
      rules: {
        username: [{ validator: validateUsername, trigger: 'blur' }],
        password: [{ validator: validatePassword, trigger: 'blur' }],
        captcha: [{ validator: validateCaptcha, trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.getCaptcha()
  },
  methods: {
    // 请求验证码
    async getCaptcha() {
      const res = await this.Api.captcha()
      this.codeImage = res.data.content.picPath
      this.loginForm.captchaId = res.data.content.captchaId
    },
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
    },
    resetForm(form) {
      this.$refs[form].resetFields()
    },
    async handleSubmit() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true
          this.login(this.loginForm)
        } else {
          return false
        }
      })
    },
    // 登录请求
    async login(params) {
      this.loading = false
      const res = await this.Api.login(params)
      const { content, code, message } = res.data
      if (code !== '0') {
        this.$message({
          type: 'warning',
          message,
          center: true
        })
        this.getCaptcha()
        this.loginForm.captcha = ''
        this.loginForm.captchaId = ''
        return
      }

      sessionStorage.setItem('user', JSON.stringify(content.user))
      sessionStorage.setItem('token', content.token)
      sessionStorage.setItem('userType', content.user.userType)
      this.$store.dispatch('router/getSelfMenuTree')
      this.$store.dispatch('user/getSelfButtons')

      this.$router.push('/home/dashboard')

      // 根据角色权限存储商户、运营单位、线路信息
      if (content.user.userType == 'ADMI') {
        this.getAllMerchants()
      } else if (content.user.userType == 'MERT') {
        sessionStorage.setItem('merchantNo', content.user.merchantNo)
        this.getCorpList(content.user.merchantNo)
      } else if (content.user.userType == 'CORP') {
        sessionStorage.setItem('merchantNo', content.user.merchantNo)
        sessionStorage.setItem('corpNo', content.user.corpNo)
        this.getLineList(content.user.merchantNo, content.user.corpNo)
      }
    },
    // 获取所有商户列表（不分页）
    async getAllMerchants() {
      const res = await this.Api.getAllMerchants()
      if (res.data.code === '0') {
        sessionStorage.setItem('merchantListData', JSON.stringify(res.data.content))
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },
    // 根据商户获取所有运营单位列表
    async getCorpList(merchantNo) {
      const res = await this.Api.findAllCorps({ merchantNo: merchantNo })
      if (res.data.code === '0') {
        sessionStorage.setItem('corpListData', JSON.stringify(res.data.content))
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },
    // 根据商户和运营单位获取线路编号列表
    async getLineList(merchantNo, corpId) {
      const res = await this.Api.findAllLines({
        merchantNo: merchantNo,
        corpId: corpId
      })
      if (res.data.code === '0') {
        sessionStorage.setItem('lineListData', JSON.stringify(res.data.content))
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    }
  }
}
</script>
<style lang="scss">
.login-container {
  .el-form-item {
    .el-input {
      .el-input-group__prepend {
        border: none;
        border-bottom: 1px solid #dcdfe6;
        background-color: #ffffff;
      }
      .el-input__inner {
        border: none;
        border-bottom: 1px solid #dcdfe6;
      }
    }
    .el-form-item__error {
      left: 78px;
    }
    .el-button {
      border: none;
      border-radius: 30px;
      width: 300px;
      height: 60px;
      font-size: 18px;
    }
    .el-button--primary {
      background-color: #248aff;
    }
    .el-button-large {
      padding: 6px 15px 6px 15px;
      font-size: 14px;
      border-radius: 4px;
    }
  }
}
</style>

<style lang="scss" scoped>
.login-container {
  position: relative;
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  background-image: url(../../images/login/bg.png);
  background-repeat: no-repeat;
  background-size: 100% 100%;
  -moz-background-size: 100% 100%;
}
.login-container .login-content {
  position: absolute;
  right: 10%;
  top: 50%;
  width: 400px;
  margin-top: -200px;
  padding: 35px 50px 15px 50px;
  border-radius: 5px;
  background-color: #fff;
  box-shadow: 0px 3px 3px 3px rgba(0, 0, 0, 0.1);
}
.login-container .login-content .login-container-title {
  position: absolute;
  left: 20px;
  top: -65px;
}
.login-container .login-content .login-title {
  display: flex;
  justify-content: center;
  margin-bottom: 55px;
}
.login-container .login-content .login-title .title {
  height: 50px;
  font-size: 30px;
  color: #303133;
  border-bottom: 2px solid #248aff;
}
.login-container .login-content .login img {
  width: 25px;
  height: 28px;
}
.login-container .login-content .login .clear-container {
  position: absolute;
  right: 5px;
  top: 0;
}
.login-container .login-content .login .clear-container img {
  width: 18px;
  height: 18px;
}
.login-container .login-content .login .checkcode-container {
  position: absolute;
  right: 0px;
  top: -15px;
  width: 120px;
  height: 40px;
  cursor: pointer;
  img {
    width: 120px;
    height: 40px;
  }
}
</style>
