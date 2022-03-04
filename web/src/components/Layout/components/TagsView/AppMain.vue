<template>
  <section class="app-main" style="min-height: 100%;">
    <transition name="fade-transform" mode="out-in">
      <keep-alive :include="cachedViews">
        <router-view :key="key" />
      </keep-alive>
    </transition>
  </section>
</template>

<script>
export default {
  name: 'AppMain',
  computed: {
    cachedViews() {
      const cachedViews = []
      const visitedViews = this.$store.state.tagsView.visitedViews

      for (const item of visitedViews) {
        cachedViews.push(item.name)
      }

      return cachedViews
    },
    key() {
      return this.$route.fullPath
    }
  }
}
</script>

<style lang="scss" scoped>
.app-main {
    position: relative;
    box-sizing: border-box;
}
</style>

