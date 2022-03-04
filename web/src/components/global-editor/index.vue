<!--
 * @Description: 富文本编辑器
-->
<template>
  <div v-loading="loading" class="editor-wrapper rel">
    <!-- 富文本框 -->
    <div id="editorId" />
    <!-- 富文本内容预览 -->
    <div class="preview" @click="previewStatus = true">富文本预览</div>
    <div v-if="previewStatus" class="preview_view" @click="previewStatus = false">
      <div class="transform_center" :class="{ preview_view_pc: previewType == 'pc', preview_view_phone: previewType == 'phone' }">
        <div :class="{ html_text: previewType == 'phone', html_text_pc: previewType == 'pc' }">
          <div :key="previewType" class="w-e-text preview_text" v-html="previewText" />
        </div>
        <div class="t_c color_f switch cursor fon_13" @click.stop="changeType">
          切换为{{ previewType=='pc' ? '手机' : '浏览器' }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Editor from 'wangeditor'
// import "wangeditor/release/wangEditor.min.css";
export default {
  name: 'WangEditor',
  props: {
    // 传入的富文本内容
    value: {
      type: String,
      default: ''
    },
    // 传入的富文本框的默认提示文本
    placeholder: {
      type: String,
      default: ''
    },
    // 设置change事件触发时间间隔
    changeInterval: {
      type: Number,
      default: 200
    },
    // 是否开启本地存储
    cache: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      previewStatus: false,
      previewType: 'phone',
      previewText: '',
      loading: false,
      fileTypes: ['jpg', 'png'],
      editor: {}
    }
  },
  watch: {
    value(val) {
      this.editor.txt.html(val)
    }
  },
  mounted() {
    this.seteditor()
  },
  methods: {
    // 切换富文本内容的预览方式
    changeType() {
      this.previewType = this.previewType == 'pc' ? 'phone' : 'pc'
    },
    // 将富文本内容清空
    setHtml() {
      this.editor.txt.html('')
    },
    /**
             * @description: 单文件上传
             */
    getImg(e) {
      const files = e.target.files || e.dataTransfer.files

      if (!files.length) {
        return
      }

      const size = (files[0].size / 1024).toFixed(2)

      const maxSize = 256
      if (size > maxSize) {
        const operationMessage = '上传的图片大小为' + size + 'KB,超过了' + maxSize + 'KB'
        this.$message.error(operationMessage)
        return
      }

      if (!/\.(jpg|jpeg|png|JPEG|JPG|PNG)$/.test(files[0].name)) {
        this.operationDialogVisible = true
        const operationMessage = '上传的图片类型必须是jpeg,jpg,png中的一种'
        this.$message.error(operationMessage)
        return
      }

      const reader = new FileReader()
      reader.readAsDataURL(files[0])
      reader.onload = function(e) {
        const imgcode = e.target.result
        this.compressImage(imgcode)
      }
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
      let compressImg = ''
      img.onload = function() { // 利用canvas压缩图片
        canvas.width = width
        canvas.height = width * (img.height / img.width)
        context.drawImage(img, 0, 0, canvas.width, canvas.height)
        compressImg = canvas.toDataURL('image/*', quality) // 转成base64
        // _this.$emit('uploadImage', compressImg);
      }
    },
    /**
             * @description: 多文件上传
             */
    getImgMore(e) {
      const srcList = []
      for (var i = 0; i < e.target.files.length; i++) {
        // eslint-disable-next-line no-async-promise-executor
        srcList[i] = new Promise(async(resolve) => {
          if (e.target.files[i].size > 1024 * 1024 * 2) {
            this.isClear = true
            this.$message.error(
              '图片大小超出限制（建议2M以内）,请重新上传'
            )
            return
          }
          const targetType = e.target.files[i].name.split('.')
          if (
            this.fileTypes.findIndex(
              (val) => val == targetType[targetType.length - 1]
            ) == -1
          ) {
            this.isClear = true
            this.$message.error(
              '不接收此文件类型 请上传jpg或png图片'
            )
            return
          }
          this.loading = true
          const data = await this.$http.form(
            this.$api.upload.request.url,
            {
              uploadFile: e.target.files[i],
              fileType: '1'
            }
          )
          if (data) {
            this.loading = false
            const src = `https://${data.data}`
            resolve(src)
          } else {
            this.loading = false
            this.$message.error('图片上传失败，请重新上传')
          }
        })
      }
      return Promise.all(srcList)
    },
    // 初始化富文本（配置项属性）
    seteditor() {
      this.editor = new Editor(`#editorId`)
      this.editor.config.placeholder = this.placeholder || '请输入内容'
      this.editor.config.showFullScreen = false
      // 配置富文本的工具栏
      this.editor.config.menus = [
        'head',
        'bold',
        'fontSize',
        // "fontName",
        'italic',
        'underline',
        'strikeThrough',
        'indent',
        'lineHeight',
        'foreColor',
        'backColor',
        'link',
        'list',
        'todo',
        'justify',
        'quote',
        // "emoticon",
        'image',
        // "video",
        // "table",
        // "code",
        'splitLine',
        'undo',
        'redo'
      ]
      // 配置富文本中上传图片方法（大小、格式限制等等）
      this.editor.config.customUploadImg = (files, insert) => {
        /* files 是 input 中选中的文件列表 */
        const event = {
          target: {
            files: files
          }
        }
        // let data = this.getImg(event);  // 单文件上传
        // console.log('data', data)
        // let data = await this.getImgMore(event);  // 多文件上传
        // data.map((item) => {
        //     insert(item);
        // });

        const file = event.target.files
        if (!file.length) {
          return
        }
        // 显示图片到富文本框
        for (var i = 0; i < file.length; i++) {
          const size = (file[i].size / 1024).toFixed(2)
          const maxSize = 256
          if (size > maxSize) {
            const operationMessage = '上传的图片大小为' + size + 'KB,超过了' + maxSize + 'KB'
            this.$message.error(operationMessage)
            return
          }
          if (!/\.(jpg|jpeg|png|JPEG|JPG|PNG)$/.test(file[i].name)) {
            const operationMessage = '上传的图片类型必须是jpeg,jpg,png中的一种'
            this.$message.error(operationMessage)
            return
          }
          const reader = new FileReader()
          reader.readAsDataURL(file[i])
          reader.onload = function(e) {
            const data = []
            data.push(e.target.result)
            data.map((item) => {
              insert(item)
            })
          }
        }
      }
      // 配置富文本内容显示
      this.editor.config.onchange = (html) => {
        const text = this.editor.txt.text()
        var reg = new RegExp('&nbsp;', 'g')
        var reg2 = new RegExp(' ', 'g')
        const newhtml = html.replace(reg2, '').replace(reg, '')

        if (html == '<p><br/></p>' || newhtml == '<p></p>') {
          html = ''
        }
        this.$emit('editor', html)

        // if (this.cache) localStorage.editorCache = html
        this.$emit('input', html)
        this.$emit('on-change', html, text)
        this.previewText = html
      }
      // 配置change事件触发间隔
      this.editor.config.onchangeTimeout = this.changeInterval
      // create这个方法一定要在所有配置项之后调用
      this.editor.create()
      // 如果本地有存储加载本地存储内容
      // let html = this.value || localStorage.editorCache
      const html = this.value
      if (html) this.editor.txt.html(html)
    }
  }
}
</script>

<style lang="scss">
    .editor-wrapper {
        z-index: 0;
    }
    .preview {
        width: 100px;
        height: 30px;
        border-radius: 30px;
        line-height: 30px;
        text-align: center;
        background: #409eff;
        color: #fff;
        font-size: 12px;
        cursor: pointer;
        margin-top: 10px;
        float: right;
    }
    .preview_view {
        position: fixed;
        width: 100vw;
        height: 100vh;
        top: 0;
        left: 0;
        z-index: 1000000;
        background: rgba(0, 0, 0, 0.5);
    }
    .preview_view_pc {
        padding: 10px 10px 0 10px;
        width: 60%;
        height: 650px;
        border-radius: 5px;
        left: 20%;
    }
    .preview_view_phone {
        // background: url("../../assets/images/phone.png") no-repeat;
        // background-size: 100%;
        // background: #fff;
        width: 300px;
        height: 650px;
    }
    .html_text {
        background: #fff;
        border-radius: 5px;
        padding: 5px 10px;
        width: 245px;
        height: 459px;
        margin-top: 72px;
        margin-left: 20px;
        overflow-y: auto;
    }
    .html_text_pc {
        background: #fff;
        border-radius: 5px;
        padding: 5px 10px;
        height: 520px;
        overflow-y: auto;
    }
    .preview_text {
        width: 100%;
        overflow-x: hidden;
        padding: 0;
        overflow-y: auto;
        height: 100%;
    }
    .switch {
        background: #409eff;
        position: absolute;
        bottom: 54px;
        left: 20%;
        border-radius: 50px;
        height: 30px;
        width: 60%;
        line-height: 30px;
    }
    .w-e-text ol li {
        list-style-type: decimal;
    }
    .w-e-text ul li {
        list-style-type: disc;
    }
    .preview_view .w-e-text ul, .preview_view .w-e-text ol {
        margin: 0;
    }
    .w-e-panel-tab-title li:last-child {
        display: none;
    }
</style>
