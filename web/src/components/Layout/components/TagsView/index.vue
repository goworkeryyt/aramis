<template>
  <div>
    <app-main :class="visitedViews.length > 0 ? 'tag-main' : 'no-tag-main'" />

    <!-- <el-card class="tag-box-card">
      <div class="tabs-container" :style="{display: visitedViews.length > 0 ? 'block': 'none'}">
        <el-tabs v-show="visitedViews.length > 0" v-model="activePath" type="card" closable @tab-click="tabPaneClick"
          @tab-remove="removeCurrentTab">
          <el-tab-pane v-for="item in visitedViews" :key="item.name" :name="item.path" :label="item.meta.title">
          </el-tab-pane>
        </el-tabs>
        <el-dropdown class="close-container" trigger="click">
          <span>
            关闭操作<i class="el-icon-caret-bottom el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item icon="el-icon-circle-close">
              <span @click="closeAllTabs">关闭所有</span>
            </el-dropdown-item>
            <el-dropdown-item icon="el-icon-close">
              <span @click="closeOtherTabs">关闭其他</span>
            </el-dropdown-item>
            <el-dropdown-item icon="el-icon-close">
              <span @click="removeCurrentTab(activePath)">关闭当前</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </el-card> -->
  </div>
</template>
<script>
import AppMain from './AppMain'
export default {
  name: 'TagsView',
  components: {
    AppMain
  },
  data() {
    return {

    }
  },
  computed: {
    visitedViews() {
      return this.$store.state.tagsView.visitedViews
    },
    activePath: {
      get() {
        return this.$store.state.tagsView.activePath
      },
      set(val) {
        this.$store.dispatch('tagsView/setActivePath', val)
      }
    }
  },
  watch: {
    $route() {
      this.addTabs()
    }
  },
  mounted() {
    this.addTabs()
  },
  methods: {
    addTabs() {
      const {
        name,
        path,
        meta
      } = this.$route
      const noTag = this.$route.meta.noTag
      let flag = false

      for (const view of this.visitedViews) {
        if (view.path === path) {
          flag = true
          this.$store.dispatch('tagsView/setActivePath', path)

          break
        }
      }

      if (!flag && !noTag) {
        if (this.visitedViews.length > 0) {
          this.$store.dispatch('tagsView/setActivePath', path)
          this.$store.dispatch('tagsView/addView', {
            name: name,
            path: path,
            meta: meta
          })
        } else {
          if (name !== 'dashboard') {
            this.$store.dispatch('tagsView/setActivePath', path)
            this.$store.dispatch('tagsView/addView', {
              name: name,
              path: path,
              meta: meta
            })
          }
        }
      }
    },
    tabPaneClick(e) {
      const currentPath = e.$options.propsData.name

      this.$router.replace({
        path: currentPath
      })
    },
    removeCurrentTab(currentTargetPath) {
      this.$store.dispatch('tagsView/delView', currentTargetPath).then(({
        visitedViews
      }) => {
        const isActive = currentTargetPath === this.$route.path

        if (isActive) {
          const latestView = visitedViews.slice(-1)[0]

          if (latestView) {
            this.$router.replace({
              path: latestView.path
            })
          } else {
            this.$router.replace({
              path: '/home/dashboard'
            })
          }
        }
      })
    },
    closeOtherTabs(currentTargetPath) {
      const spliceLength = this.visitedViews.length - 1
      this.$store.dispatch('tagsView/delOtherVisitedViews', spliceLength)
    },
    closeAllTabs() {
      this.$store.dispatch('tagsView/delAllVisitedViews').then(() => {
        this.$router.replace({
          path: '/home/dashboard'
        })
      })
    }
  }
}

</script>
<style lang="scss">
  .el-card{
      border-radius:10px;
      background:#f9f8fc;
    }

  .tag-box-card {
    .el-card__body {
      padding: 0px !important;
      height: 100%;

      .el-tabs {
        position: relative;
        width: 100%;
      }

      .el-tabs--card>.el-tabs__header {
        width: 100%;
        background-color:#F9F8FC !important;
        border-bottom: none;
        padding:8px 105px 0 15px;

        .el-tabs__nav-wrap.is-scrollable {
          padding: 0 40px;
          box-sizing: border-box;
        }

        .el-tabs__item{
           background-color: #F9F8FC;
           border:1px solid #E8E8E8;
           border-radius: 8px 8px 0px 0px;
           color:#212121;
        }

        .el-tabs__item:hover {
          color: #3F6BFF!important;
          cursor: default;
        }

         .el-tabs__item .el-icon-close {
          position: relative;
          font-size: 12px;
          height: 14px;
          vertical-align: middle;
          line-height: 15px;
          overflow: visible;
          top: -1px;
          right: -2px;
        }

        .el-tabs__item .el-icon-close:hover{
             background:#1292FF;
        }

      }

      .el-tabs__item.is-active {
        color: #1292FF!important;
        border-bottom:1px solid #FFF !important;
        background-color: #FFF !important;
      }
    }

    .el-icon-arrow-left:before {
      content: "<<" !important;
      font-size: 16px;
      text-align:center;
    }

    .el-icon-arrow-right:before {
      content: ">>" !important;
      font-size: 16px;
        text-align:center;
    }
  }

</style>
<style lang="scss" scoped>
  // .tag-box-card {
  //   position: relative;
  //   margin: 10px;
  //   margin-bottom: 0;
  //   padding: 0px;
  //   height: calc(100vh - 72px);
  //   box-sizing: border-box;

  //   .tabs-container {
  //     position: relative;
  //     height: 42px;
  //     margin-bottom: 15px;

  //     .close-container {
  //       position: absolute;
  //       right: 15px;
  //       top: 8px;
  //       height: 15px;
  //       padding-left:5px;
  //       cursor: default;
  //       width: 85px;
  //       height:40px;
  //       line-height:40px;
  //       border-radius: 6px !important;
  //       color:#1292FF !important;
  //       border:1px solid  #1292FF !important;
  //       font-size: 14px;
  //       display: inline-block;
  //     }
  //   }

  // }
.tag-main {
      overflow-y: auto;
      box-sizing: border-box;
      height: 100%;
    }

    .no-tag-main {
      overflow-y: auto;
      box-sizing: border-box;
      height: 100%;
    }
.tag-box-card >>>.el-tabs--card>.el-tabs__header .el-tabs__nav{
    border:none;
}
</style>
