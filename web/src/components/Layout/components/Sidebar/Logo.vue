<template>
  <div class="sidebar-logo-container" :class="{'collapse':collapse}">
    <div v-if="collapse" class="sidebar-logo-link">
      <img v-if="logo" :src="logo" class="sidebar-logo">
      <h1 v-else class="sidebar-title">{{ title }} </h1>
    </div>
    <div v-else class="sidebar-logo-link">
      <img v-if="logo" :src="logo" class="sidebar-logo">
      <h1 class="sidebar-title">{{ title }} </h1>
    </div>
  </div>
</template>

<script>
// import {
//  getLogo
// } from '@/http/Login.js';

export default {
  name: 'SidebarLogo',
  props: {
    collapse: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      title: 'Aramis'
      // logo: 'https://chuangwork.oss-cn-beijing.aliyuncs.com/2f0dbbcddea948bfb12fa97ddf722448_logo_1.png',
      // logo: require('@/images/aramis120_120.png')
      // logo: ''
    }
  },
  mounted() {
    //   this.getLogo();
  },
  methods: {
    getLogo() {
      const userType = sessionStorage.getItem('userType')
      const json = JSON.parse(this.userTypeLo)
      const merchantNo = sessionStorage.getItem('merchantNo')
      if (json[userType] != 'A') {
        getLogo({ merchantNo }).then(res => {
          if (res.data.code === '1' && res.data.content && res.data.content.imageSrc) {
            this.logo = res.data.content.imageSrc
          }
        })
          .catch((err) => {

          })
      }
    }
  }
}

</script>

<style lang="scss" scoped>
  .sidebarLogoFade-enter-active {
    transition: opacity 1.5s;
  }

  .sidebarLogoFade-enter,
  .sidebarLogoFade-leave-to {
    opacity: 0;
  }

  .sidebar-logo-container {
    position: relative;
    width: 200px;
    height: 60px;
    background: #1292FF;

    & .sidebar-logo-link {
      height:100%;
      width:200px;

      & .sidebar-logo {
          width: 110px;
          margin-left:30px;
          vertical-align: middle;
      }

      & .sidebar-title {
            width: 100%;
            display: inline-block;
            margin: 0;
            color: #fff;
            font-weight: 600;
            line-height: 50px;
            font-size: 18px;
            font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
            vertical-align: middle;
            text-align: center;
      }
    }

    &.collapse {
      .sidebar-logo {
        margin-right: 0px;
      }
    }
  }

</style>
