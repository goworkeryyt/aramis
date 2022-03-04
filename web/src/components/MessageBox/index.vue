<template>
  <el-dialog
    class="confirm-dialog"
    width="400px"
    :center="true"
    :show-close="false"
    :append-to-body="true"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    :visible.sync="messageVisible"
  >
    <div class="dialog-container">
      <p v-if="messageType=='warning'"><img src="../../images/warning.png"></p>
      <p v-if="messageType=='success'"><img src="../../images/success.png"></p>
      <p>{{ messageContent }}</p>
      <slot />
    </div>
    <span slot="footer" class="dialog-footer">
      <el-button
        type="primary"
        size="small"
        @click="messageRes('yes')"
      >确 定
      </el-button>
    </span>
  </el-dialog>
</template>
<script>
export default {
  name: 'MessageBox',
  props: {
    messageVisible: {
      type: Boolean
    },
    messageContent: {
      type: String
    },
    messageResult: {
      type: Function
    },
    messageType: {
        	type: String,
        	default: 'warning'
    }
  },
  data() {
    return {

    }
  },
  methods: {
    messageRes() {
      this.$emit('closeMessageBox', 'yes')
    }
  }
}
</script>
<style>
	.confirm-dialog .el-dialog{
		border-radius: 6px;
	}
	.confirm-dialog .el-dialog .el-dialog__header{
		padding: 0;
	}
</style>
<style lang="scss" scoped>

    .dialog-container {
        display: flex;
        flex-direction: column;
        p {
        	font-size: 16px;
        	color: #212121;
            line-height: 25px;
            display: flex;
            justify-content: center;
            margin: 5px 0 25px 0;
        }
    }
    .dialog-footer {
        display: flex;
        justify-content: center;
        .el-button {
            min-width: 90px;
            height: 36px;
            background:rgba(37,137,255,1);
			border-radius:4px;
			font-size: 16px;
            &:nth-child(1) {
                margin-left: 0px !important
            }
            &:nth-child(2) {
                margin-left: 30px
            }
        }
    }
</style>
