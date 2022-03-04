<template>
  <div class="app-wrapper" :class="{hideSidebar:!sidebar.opened}">
    <navbar />
    <div class="main-container">
      <sidebar class="sidebar-container" />
      <tags-view />
    </div>
  </div>
</template>

<script>
import { Navbar, Sidebar, TagsView } from './components'
export default {
  name: 'Layout',
  components: {
    Navbar,
    Sidebar,
    TagsView
  },
  data() {
    return {
    }
  },
  computed: {
    sidebar() {
		    return this.$store.state.app.sidebar
    }
  }
}
</script>

<style lang="scss" scoped>
    @import "~@/styles/mixin.scss";
    @import "~@/styles/variables.scss";
    .footer {
        height: 60px;
        position: relative;
        background: #2a579a;
        box-shadow: 0 1px 4px rgba(0,21,41,.08);
    }
    .copyright{
        text-align:center;
        line-height:60px;
        font-size:14px;
        color:#f0f0f0;
    }
    .app-wrapper {
        @include clearfix;
        position: relative;
        height: 90%;
        width: 100%;
    }
</style>
