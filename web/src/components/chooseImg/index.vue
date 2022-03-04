<template>
  <el-drawer class="img-drawer" :visible.sync="drawer" title="媒体库" :append-to-body="true" size="440px">
    <div class="media">
      <el-image
        v-for="(value,label) in Constant.headerImgMap"
        :key="label"
        class="header-img-box-list"
        :src="value"
        @click="chooseImg({urlType:label,url:value})"
      >
        <template #error>
          <div class="header-img-box-list">
            <i class="el-icon-picture-outline" />
          </div>
        </template>
      </el-image>
    </div>
  </el-drawer>
</template>

<script>
export default {
  data() {
    return {
      drawer: false
    }
  },
  methods: {
    chooseImg(item) {
      this.$emit('getImg', item)
      this.drawer = false
    },

    open() {
      this.drawer = true
    }
  }
}
</script>

<style scoped>
.media{
  display:flex;
  flex-wrap:wrap;
}

.header-img-box-list {
  margin-top: 20px;
  width: 120px;
  margin-left: 20px;
  height: 120px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 120px;
  cursor: pointer;
}

.img-drawer>>>.el-drawer__header {
  padding: 20px
}

</style>
