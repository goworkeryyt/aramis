
<template>
  <div class="menu-wrapper">
    <template v-for="item in routes">
      <el-submenu
        v-if="item.IsShow === 1"
        :key="item.url"
        :index="'/' + item.url"
      >
        <template slot="title">
          <!-- <div v-if="item && item.icon" :class="item.icon"></div> -->
          <i v-if="item && item.menuIcon" :class="'el-icon-' + item.menuIcon" />
          <span
            v-if="item && item.menuName"
            slot="title"
            style="margin-left: 8px"
          >{{ item.menuName }}
          </span>
        </template>

        <template v-for="child in item.children">
          <!-- <sidebar-item
            style="width: 100px"
            class="nest-menu"
            :is-nest="true"
            :routes="[child]"
            :key="child.url"
            v-if="child.children && child.children.length > 0"
          >
          </sidebar-item> -->
          <router-link :key="child.url" :to="'/' + item.url + '/' + child.url">
            <el-menu-item
              v-if="child.IsShow === 1"
              mode="vertical"
              :index="'/' + item.url + '/' + child.url"
              collapse="true"
            >
              <div @mouseenter="enter" @mouseleave="leave">
                <span
                  v-if="child && child.menuName"
                  slot="title"
                  :class="getPadding(item.url, child.url)"
                >
                  <span :class="getPointBg(item.url, child.url)" />
                  {{ child.menuName }}
                </span>
              </div>
            </el-menu-item>
          </router-link>
        </template>
      </el-submenu>
    </template>
  </div>
</template>

<script>
export default {
  name: 'SidebarItem',
  props: {
    isNest: {
      type: Boolean,
      default: false
    },
    routes: {
      type: Array
    }
  },
  data() {
    return {
      userNo: sessionStorage.getItem('userNo'),
      jurisdiction: sessionStorage.getItem('jurisdiction')
    }
  },
  computed: {
    getPointBg(itemPath, childPath) {
      return function(itemPath, childPath) {
        return this.$route.path == '/' + itemPath + '/' + childPath
          ? 'child-active-point'
          : 'child-point'
      }
    },
    getPadding(itemPath, childPath) {
      return function(itemPath, childPath) {
        return this.$route.path == '/' + itemPath + '/' + childPath
          ? 'child-active-padding'
          : 'child-padding'
      }
    }
  },
  mounted() {},
  methods: {
    enter() {
      const wrapper = document.getElementsByClassName('app-wrapper')[0]

      if (wrapper.classList.contains('hideSidebar')) {
        wrapper.classList.add('hoverItem')
      }
    },
    leave() {
      const wrapper = document.getElementsByClassName('app-wrapper')[0]

      if (wrapper.classList.contains('hideSidebar')) {
        wrapper.classList.remove('hoverItem')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.child-active-padding {
  padding-left: 24px;
}

.child-padding {
  padding-left: 30px;
}

.child-point {
  width: 8px;
  height: 8px;
  border-radius: 8px;
  border: 1px solid #393e46;
  background: #ffffff;
  display: inline-block;
  margin-right: 14px;
}

.child-active-point {
  width: 8px;
  height: 8px;
  border-radius: 8px;
  border: 1px solid #1292ff;
  background: #1292ff;
  display: inline-block;
  margin-right: 14px;
}
</style>

