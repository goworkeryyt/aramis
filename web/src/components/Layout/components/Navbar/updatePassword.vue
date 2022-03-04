<template lang="">
  <el-dialog
    class="dialog-form"
    title="修改登录密码"
    width="654px"
    :append-to-body="true"
    :close-on-click-modal="false"
    :modal-append-to-body="true"
    :close-on-press-escape="false"
    :before-close="closeDialog"
    :visible.sync="passWordFormVisible"
  >
    <el-form ref="dialogPassWordForm" label-width="180px" :model="dialogPassWordForm" :rules="formRolues">
      <el-form-item label="旧密码" prop="password">
        <el-input v-model="dialogPassWordForm.password" type="password" placeholder="请输入旧密码" class="input-width" />
      </el-form-item>
      <el-form-item label="新密码" prop="newPassword">
        <el-input v-model="dialogPassWordForm.newPassword" placeholder="请输入新密码" type="password" class="input-width" />
      </el-form-item>
      <el-form-item label="确认新密码" prop="confirmPassWord">
        <el-input v-model="dialogPassWordForm.confirmPassWord" placeholder="再次输入新密码" type="password"class="input-width" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button class="cancel-btn" @click="resetPassWordForm('dialogPassWordForm')">取 消
      </el-button>
      <el-button :loading="btnloading" class="submit" @click="submitPassWordForm('dialogPassWordForm')">
        确 定
      </el-button>
    </div>
  </el-dialog>
</template>
<script>
export default {
  name: 'UpdatePassword',
  props: {
    passWordFormVisible: {
      type: Boolean
    },
    operatePassword: {
      type: Function
    }
  },
  data() {
    return {
      btnloading: false,
      dialogPassWordForm: {
        password: '',
        newPassword: '',
        confirmPassWord: ''
      },
      formRolues: {
        password: [
          {
            required: true,
            message: '旧密码不能为空',
            trigger: 'blur'
          },
          {
            pattern: /^[A-Za-z0-9]{6,13}$/,
            message: '密码为6-13位数字和字母的组合'
          }
        ],
        newPassword: [
          {
            required: true,
            message: '新密码不能为空',
            trigger: 'blur'
          },
          {
            pattern: /^[A-Za-z0-9]{6,13}$/,
            message: '密码为6-13位数字和字母的组合'
          }
        ],
        confirmPassWord: [
          {
            required: true,
            message: '新密码不能为空',
            trigger: 'blur'
          },
          {
            pattern: /^[A-Za-z0-9]{6,13}$/,
            message: '密码为6-13位数字和字母的组合'
          }
        ]
      }
    }
  },
  methods: {
    closeDialog() {
      this.$emit('operatePassword')
      this.$refs['dialogPassWordForm'].resetFields()
    },
    resetPassWordForm(form) {
      this.$emit('operatePassword')
      this.$refs[form].resetFields()
    },

    submitPassWordForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (
            this.dialogPassWordForm.newPassword ===
            this.dialogPassWordForm.confirmPassWord
          ) {
            this.btnloading = true
            const params = {
              password: this.dialogPassWordForm.password,
              newPassword: this.dialogPassWordForm.newPassword,
              username: JSON.parse(sessionStorage.getItem('user')).username
            }
            this.changePassword(params)
          } else {
            this.$message({
              message: '两次输入的新密码不一致！',
              center: true,
              customClass: 'message-error'
            })
          }
        }
      })
    },
    // 修改密码请求
    async changePassword(params) {
      const res = await this.Api.changePassword(params)
      this.btnloading = false
      const { code, message } = res.data
      if (code !== '0') {
        this.$emit('showOperation', message)
        return
      }
      this.$message({
        type: 'success',
        message
      })
      sessionStorage.clear()
      this.$router.push('/')
    }
  }
}
</script>
