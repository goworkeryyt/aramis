<!--
 * @Description: 上传图片组件
-->
<template>
  <div class="upload-image">
    <div v-if="imgSrc" class="selected-icon-panel">
      <el-image v-if="imgSrc" :src="imgSrc">
        <div slot="placeholder" class="image-slot">
          <i class="el-icon-loading" />
        </div>
      </el-image>
      <el-tooltip class="item" effect="dark" content="删除" placement="top">
        <img class="delete-img" src="../../images/guanbitubiao.png" @click="deleteIcon()">
      </el-tooltip>
    </div>
    <div class="icon-upload" @click="uploadImage">
      <i class="el-icon-plus avatar-uploader-icon" />
    </div>
    <!-- 上传 -->
    <input id="fileImage" ref="upFileImage" type="file" hidden="hidden" accept="image/png,image/jpeg,image/jpg" @change="onChange">
  </div>
</template>

<script>
export default {
  name: 'UploadImage',
  props: {
    imgUrl: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      imgSrc: ''
    }
  },
  watch: {
    imgUrl: {
      handler(newValue, oldValue) {
        this.imgSrc = newValue
      },
      immediate: true
    }
  },
  methods: {
    setDefaultImg(img) {
      this.imgSrc = img
    },
    uploadImage() {
      document.getElementById('fileImage').click()
    },
    // 选择文件时触发
    onChange(e) {
      const files = e.target.files || e.dataTransfer.files

      if (!files.length) {
        return
      }

      const size = (files[0].size / 1024).toFixed(2)

      const maxSize = 256
      if (size > maxSize) {
        this.$refs.upFileImage.value = ''
        const operationMessage = '上传的图片大小为' + size + 'KB,超过了' + maxSize + 'KB'
        // this.$emit('showOperationDialog', operationMessage);
        this.$message.error(operationMessage)
        return
      }

      if (!/\.(jpg|jpeg|png|JPEG|JPG|PNG)$/.test(files[0].name)) {
        this.$refs.upFileImage.value = ''
        this.operationDialogVisible = true
        const operationMessage = '上传的图片类型必须是jpeg,jpg,png中的一种'
        // that.$emit('showOperationDialog', operationMessage);
        this.$message.error(operationMessage)
        return
      }

      const that = this
      const reader = new FileReader()
      reader.readAsDataURL(files[0])

      reader.onload = function(e) {
        const imgcode = e.target.result
        that.type = 'change'
        that.imgSrc = imgcode
        that.compressImage(imgcode)
      }

      // e.target.value = ''; // 解决了 input上传第二次不能选择同一图片 问题
    },
    // 图片压缩
    compressImage(file) {
      const _this = this
      const width = 512 // 压缩后的宽
      const quality = 0.5 // 压缩质量
      const canvas = document.createElement('canvas')
      const context = canvas.getContext('2d')
      const img = new Image()
      img.src = file
      img.onload = function() { // 利用canvas压缩图片
        canvas.width = width
        canvas.height = width * (img.height / img.width)
        context.drawImage(img, 0, 0, canvas.width, canvas.height)
        const compressImg = canvas.toDataURL('image/*', quality) // 转成base64
        _this.$emit('uploadImage', compressImg)
      }
    },
    // 删除图片
    deleteIcon() {
      this.imgSrc = '' // 清空图片显示
      document.getElementById('fileImage').value = '' // 清空input的value值
      this.$emit('deleteIcon')
    }
  }
}
</script>
<style lang="scss">
	.upload-image{
		.el-image{
			width: 100px;
			height: 100px;
			.image-slot{
				line-height: 100px;
			}
		}
	}

</style>
<style scoped="scoped" lang="scss">
	// .upload-image {
	// 	margin-left: 16px;
	// }

	.icon-upload {
		border: 1px dashed rgba(201, 201, 201, 1);
		background: rgba(255, 255, 255, 1);
		cursor: pointer;
		position: relative;
		display: inline-block;
		text-align: center;
		outline: none;
		width: 100px;
		height: 100px;
	}

	.selected-icon-panel {
		width: 100px;
		height: 100px;
		margin-right: 30px;
		position: relative;
		float: left;
		text-align: center;
		outline: none;
		border: 1px solid #F2F2F6;

		img:first-child{
			vertical-align: middle;
			width: 100%;
		}
		.delete-img {
			position: absolute;
			top: -8px;
			right: -8px;
			cursor: pointer;
		}
	}

	.avatar-uploader-icon {
		font-size: 25px;
		color: #8c939d;
		vertical-align: middle;
		width: 100px;
		height: 100px;
		line-height: 100px;
		text-align: center;
	}

	.icon-panel {
		li {
			position: relative;
			display: inline-block;
			width: 100px;
			height: 100px;
			background: #CCCCCC;
			margin-right: 20px;
			margin-bottom: 20px;
			text-align: center;
			line-height: 60px;
			cursor: pointer;
			border: 1px solid #CCCCCC;
		}
		li:hover {
			border: 1px solid #2589FF;
		}
		li.checked {
			border: 1px solid #2589FF;
			position: relative;
			transition: all 0.5s ease;
			&::after {
				content: "\2714";
				display: block;
				height: 0px;
				width: 0px;
				position: absolute;
				bottom: 0;
				right: 0;
				color: #fff;
				font-size: 12px;
				line-height: 8px;
				border: 10px solid;
				border-color: transparent #2589FF #2589FF transparent;
			}
		}
	}
</style>
