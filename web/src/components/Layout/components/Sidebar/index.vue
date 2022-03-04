
<template>
  <div :class="{ 'has-logo': showLogo }">
    <!-- <logo v-if="showLogo"/> -->
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        mode="vertical"
        :collapse-transition="false"
        :show-timeout="200"
        :default-active="$route.path"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :active-text-color="variables.menuActiveText"
        :unique-opened="true"
      >
        <sidebar-item :routes="menuList" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import Logo from './Logo'
import SidebarItem from './SidebarItem'
import variables from '@/styles/variables.scss'
import { mapGetters } from 'vuex'

export default {

  components: {
    SidebarItem,
    Logo
  },
  data() {
    return {
      menuList: []
    }
  },
  computed: {
    showLogo() {
      return this.$store.state.settings.sidebarLogo
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  },
  created() {
    this.getSelfMenuTree()
  },
  methods: {
    async getSelfMenuTree() {
      const res = await this.Api.getSelfMenuTree()
      if (res.data.code !== '1') {
        this.menuList = res.data.content
        return
      }
    }
  }
}
</script>
