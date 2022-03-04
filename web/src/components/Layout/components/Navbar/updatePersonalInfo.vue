<template>
  <el-dialog
    class="dialog-form"
    :title="showUserName ? '修改用户名' : '修改个人资料'"
    width="654px"
    :append-to-body="true"
    :close-on-click-modal="false"
    :modal-append-to-body="true"
    :close-on-press-escape="false"
    :before-close="closeDialog"
    :visible.sync="personalFormVisible"
  >
    <el-form v-show="!showUserName" ref="dialogPersonalForm" :model="dialogPersonalForm" :rules="formRolues" label-width="160px">
      <el-form-item label="用户名" prop="username">
        <el-input id="username" v-model="dialogPersonalForm.username" class="input-width" placeholder="请输入用户名" disabled />
        <!-- 暂不允许修改用户名
            <el-button type="text" @click.prevent="updateUserName">修改用户名</el-button> -->
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input id="phone" v-model="dialogPersonalForm.phone" class="input-width" placeholder="请输入手机号" />
      </el-form-item>
      <el-form-item label="昵称" prop="nickName">
        <el-input id="nickName" v-model="dialogPersonalForm.nickName" class="input-width" placeholder="请输入昵称" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input id="email" v-model="dialogPersonalForm.email" class="input-width" placeholder="请输入邮箱" />
      </el-form-item>
    </el-form>
    <div v-show="!showUserName" slot="footer" class="dialog-footer">
      <el-button class="cancel-btn" @click="resetPersonalForm('dialogPersonalForm')">取 消</el-button>
      <el-button :loading="btnloading" class="submit" @click="submitPersonalForm('dialogPersonalForm')">
        确 定
      </el-button>
    </div>
    <el-form v-show="showUserName" ref="dialogUserNameForm" :model="dialogUserNameForm" :rules="formRolues2" label-width="180px">
      <el-form-item label="旧用户名">
        <el-input v-model="dialogUserNameForm.oldUserName" class="input-width" disabled />
      </el-form-item>
      <el-form-item label="旧密码" prop="oldPassWd">
        <el-input v-model="dialogUserNameForm.oldPassWd" class="input-width" placeholder="请输入旧密码" />
      </el-form-item>
      <el-form-item label="新用户名" prop="newUserName">
        <el-input v-model="dialogUserNameForm.newUserName" class="input-width" placeholder="请输入更改后的用户名" />
      </el-form-item>
      <el-form-item label="新密码" prop="newPassWd">
        <el-input v-model="dialogUserNameForm.newPassWd" class="input-width" placeholder="请输入新密码" />
      </el-form-item>
    </el-form>
    <div v-show="showUserName" slot="footer" class="dialog-footer">
      <el-button class="cancel-btn" @click="resetUserNameForm('dialogUserNameForm')">取 消</el-button>
      <el-button :loading="btnloading" class="submit" @click="submitUserNameForm('dialogUserNameForm')">
        确 定
      </el-button>
    </div>

  </el-dialog>
</template>
<script>
export default {
  name: 'UpdatePersonalInfo',
  props: {
    personalFormVisible: {
      type: Boolean
    },
    operatePersonal: {
      type: Function
    }
  },
  data() {
    return {
      showUserName: false,
      btnloading: false,
      dialogPersonalForm: {
        username: '',
        phone: '',
        email: '',
        nickName: ''
      },
      formRolues: {
        phone: [{
          required: true,
          message: '手机号不能为空',
          trigger: 'blur'
        }, {
					  pattern: /^1[0-9]{10}$/,
          message: '请输入正确的手机号'
        }],
        nickName: [{
          required: true,
          message: '昵称不能为空！',
          trigger: 'blur'
        }],
        email: [{
          required: true,
          message: '邮箱不能为空',
          trigger: 'blur'
        }, {
          pattern: /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/,
          message: '请输入正确的邮箱'
        }]
      },
      dialogUserNameForm: {
        oldUserName: '',
        oldPassWd: '',
        newUserName: '',
        newPassWd: ''
      },
      formRolues2: {
        oldPassWd: [{
          required: true,
          message: '旧密码不能为空',
          trigger: 'blur'
        }],
        newUserName: [{ required: true, message: '新用户名不能为空', trigger: 'blur' },
          { min: 5, message: '最低5位字符', trigger: 'blur' }
        ],
        newPassWd: [{
          required: false
        }, {
          pattern: /^[A-Za-z0-9]{6,13}$/,
          message: '密码为6-13位数字和字母的组合'
        }]
      }
    }
  },
  methods: {
    setFormValue() {
      const user = JSON.parse(sessionStorage.getItem('user'))
      this.dialogPersonalForm.id = user.id ? user.id : ''
      this.dialogPersonalForm.username = user.username ? user.username : ''
      this.dialogPersonalForm.phone = user.phone ? user.phone : ''
      this.dialogPersonalForm.nickName = user.nickName ? user.nickName : ''
      this.dialogPersonalForm.email = user.email ? user.email : ''

      this.dialogUserNameForm.oldUserName = user.username ? user.username : ''
    },
    closeDialog() {
      this.showUserName = false
      this.$emit('operatePersonal')
      this.$refs['dialogPersonalForm'].resetFields()
      this.$refs['dialogUserNameForm'].resetFields()
    },
    resetPersonalForm(form) {
      this.$emit('operatePersonal')
      this.$refs[form].resetFields()
    },
    resetUserNameForm(form) {
      this.$refs[form].resetFields()
      this.showUserName = false
    },
    updateUserName() {
      this.showUserName = true
    },
    submitPersonalForm(form) {
      this.$refs[form].validate(async(valid) => {
        if (valid) {
          this.btnloading = true
          const res = await this.Api.updateUser(this.dialogPersonalForm)
          const { code, message, content } = res.data
          this.btnloading = false

          if (code === '0') {
            this.$emit('operatePersonal')
            this.$message({
              message,
              center: true,
              type: 'success'
            })

							  this.getSelfUserInfo()
          } else {
            this.$emit('showOperation', message)
          }
        }
      })
    },
    // 暂不允许修改用户名 --暂无接口
    submitUserNameForm(form) {
      this.$refs[form].validate(async(valid) => {
        if (valid) {
          this.btnloading = true
          const data = {
            oldUserName: this.dialogUserNameForm.oldUserName,
            oldPassWd: this.dialogUserNameForm.oldPassWd,
            newUserName: this.dialogUserNameForm.newUserName,
            newPassWd: this.dialogUserNameForm.newPassWd
          }

          const res = await this.Api.modifyPhone(data)
          const { code, message, content } = res.data

          this.btnloading = false
          if (code === '0') {
            this.$refs['dialogUserNameForm'].resetFields()
            this.showUserName = false
            this.$emit('operatePersonal')
            this.$message({
              message,
              center: true,
              type: 'success'
            })

            this.getSelfUserInfo()
          } else {
            this.$emit('showOperation', message)
          }
        }
      })
    },

    async getSelfUserInfo() {
      const res = await this.Api.getSelfUserInfo()
      const { code, message, content } = res.data

      if (code == '0') {
        sessionStorage.setItem('user', JSON.stringify(content.userInfo))
      }
    }
  }
}
</script>
