<!--
 * @Description: 上传文件
-->
<template>
  <div id="upload-file">
    <el-input v-model="uploadFileName" readonly="readonly" placeholder="请上传文件" />
    <el-button class="addClient" @click="uploadFile"><i class="el-icon-upload2" /> <span>点击上传</span></el-button>
    <input id="uploadFile" ref="uploadFile" type="file" accept="*" hidden class="uploadInput" @change="fileChange">
  </div>
</template>
<script>
export default {
  name: 'UploadFile',
  props: {
    fileName: {
      type: String,
      default: () => { return '' }
    }
  },
  data() {
    return {
      uploadFileName: ''
    }
  },
  watch: {
    fileName: {
      handler(newVal) {
        this.uploadFileName = newVal
      }
    }
  },
  mounted() {
    this.uploadFileName = this.fileName
  },
  methods: {
    // 清空显示文件输入框
    clearFileInput() {
      this.uploadFileName = ''
    },
    // 点击“上传文件”按钮
    uploadFile() {
      document.getElementById('uploadFile').click()
    },
    // 选择上传的文件
    fileChange(e) {
      const files = e.target.files || e.dataTransfer.files

      if (!files.length) {
        this.uploadFileName = ''
        this.$emit('uploadFileNull')
        return
      }

      const file = files[0]
      this.uploadFileName = file.name
      // 将file文件传给父组件
      this.$emit('uploadFile', file)
      e.target.value = '' // 解决了 input上传第二次不能选择同一文件 问题
    }
  }
}
</script>
<style lang='scss'>
    .addClient {
		position: absolute;
		top: 0;
        left: 299px;
    }

	.uploadInput {
		width: 300px !important;
		position: relative !important;
		top: -1px !important;
        left: -1px !important;
	}
</style>
