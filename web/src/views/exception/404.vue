<template>
  <div class="page-container">
    <div class="inner-container">
      <span>Ooops......</span>
      <span>抱歉，页面君去洗澡了，或者去洗澡了......</span>
      <span>您还可以...</span>
      <div class="btn-container">
        <div class="first-page" @click="backToFirst">返回首页</div>
        <div class="last-page" @click="backToLast">返回上一页</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {}
  },
  methods: {
    backToFirst() {
      this.$router.push({ path: '/' })
    },
    backToLast() {
      this.$router.go(-1)
    }
  }
}
</script>
<style lang="scss" scoped>
.page-container {
  position: relative;
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  background-image: url(../../images/bg404.png);
  background-repeat: no-repeat;
  background-size: 100% 100%;
  -moz-background-size: 100% 100%;
  box-sizing: border-box;
  .inner-container {
    padding-top: 50px;
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    span {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      font-size: 24px;
      font-weight: 400;
      color: rgba(51, 51, 51, 1);
      height: 50px;
      line-height: 50px;
      &:nth-child(3) {
        font-size: 16px;
        font-family: Microsoft YaHei;
        font-weight: 400;
        color: rgba(102, 102, 102, 1);
        margin-bottom: 10px;
      }
    }
    .btn-container {
      display: flex;
      flex-direction: row;
      justify-content: center;
      font-size: 18px;
      color: rgba(255, 255, 255, 1);
      .first-page {
        width: 150px;
        height: 40px;
        line-height: 40px;
        text-align: center;
        background: rgba(37, 137, 255, 1);
        border-radius: 10px;
        margin-right: 35px;
        cursor: pointer;
      }
      .last-page {
        width: 150px;
        height: 40px;
        line-height: 40px;
        text-align: center;
        border-radius: 10px;
        color: rgba(37, 137, 255, 1);
        border: 1px solid rgba(37, 137, 255, 1);
        cursor: pointer;
      }
    }
  }
}
</style>
