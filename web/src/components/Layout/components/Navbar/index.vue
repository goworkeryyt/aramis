<template>
  <el-menu class="navbar" mode="horizontal">
    <logo />
    <div class="right-menu">
      <div class="avatar-container right-menu-item user-info-pane">
        <el-dropdown
          class="avatar-container right-menu-item hover-effect"
          trigger="click"
          placement="bottom-start"
        >
          <div class="avatar-wrapper">
            <img class="user-avatar" :src="getHeaderImg(headerImg)">
          </div>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item class="drop-border">
              <span
                style="display: block"
                @click="updatePersonal('edit')"
              >修改个人资料</span>
            </el-dropdown-item>
            <el-dropdown-item class="drop-border">
              <span
                style="display: block"
                @click="updatePassWord"
              >修改登录密码</span>
            </el-dropdown-item>
            <el-dropdown-item>
              <span style="display: block" @click="logout">退出</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <el-tooltip :content="nickName" placement="bottom-start" effect="dark">
          <span class="user-name">{{ nickName }}</span>
        </el-tooltip>
      </div>
    </div>
    <message-box
      :message-content="messageContent"
      :message-visible="messageVisible"
      @closeMessageBox="closeMessageBox"
    />
    <update-password
      :pass-word-form-visible="pwFormVisible"
      @operatePassword="updatePassWord"
      @showOperation="showOperation"
    />
    <update-personal-info
      ref="updatePersonal"
      :personal-form-visible="pnFormVisible"
      @operatePersonal="updatePersonal"
      @showOperation="showOperation"
    />
  </el-menu>
</template>

<script>
import UpdatePassword from './updatePassword.vue'
import UpdatePersonalInfo from './updatePersonalInfo.vue'
import MessageBox from '@/components/MessageBox/index'
import Logo from '../Sidebar/Logo'
export default {
  components: {
    UpdatePassword,
    UpdatePersonalInfo,
    MessageBox,
    Logo
  },
  data() {
    return {
      pwFormVisible: false,
      pnFormVisible: false,
      headerImg: 'boy',
      nickName: '',

      messageContent: '',
      messageVisible: false
    }
  },
  computed: {
    getHeaderImg(headerImg) {
      return function(headerImg) {
        return headerImg
          ? headerImg.slice(0, 4) == 'http'
            ? headerImg
            : this.Constant.headerImgMap[headerImg]
          : 'boy'
      }
    }
  },
  created() {
    const user = JSON.parse(sessionStorage.getItem('user'))
    this.headerImg = user && user.headerImg ? user.headerImg : 'boy'
    this.nickName = user && user.nickName ? user.nickName : ''
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    updatePersonal(type) {
      if (type) {
        this.$refs.updatePersonal.setFormValue()
      }
      this.pnFormVisible = !this.pnFormVisible
    },
    updatePassWord() {
      this.pwFormVisible = !this.pwFormVisible
    },
    showOperation(msg) {
      this.messageVisible = true
      this.messageContent = msg
    },
    closeMessageBox() {
      this.messageVisible = false
    },
    logout() {
      sessionStorage.clear()
      this.$router.push('/')
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>
.navbar {
  position: relative;
  width: 100% !important;
  height: 60px;
  line-height: 60px;
  border-radius: 0px !important;
  background: #1292ff;
  border: none !important;

  .right-menu {
    position: absolute;
    top: 0px;
    right: 0px;
    height: 100%;
    line-height: 70px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      margin: 0 8px;
      padding-top: -30px;
      height: 100%;
      font-size: 18px;
      color: #f0f0f0;
      vertical-align: text-bottom;
      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;
        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }
    .user-info-pane {
      height: 40px;
      min-width: 150px;
      margin-bottom: 10px;
    }
    .avatar-container {
      margin-right: 30px;
      height: 40px;
      .avatar-wrapper {
        margin: 0 0 0 2px;
        position: relative;
        .user-avatar {
          cursor: pointer;
          width: 36px;
          height: 36px;
          border-radius: 18px;
        }
        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
      .user-name {
        position: absolute;
        line-height: 60px;
        font-size: 16px;
        // color: #ffffff;
        margin-left: -20px;
        width: 90px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
      }
    }
  }
}
</style>
