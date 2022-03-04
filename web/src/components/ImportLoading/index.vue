<template>
  <el-dialog
    class="confirm-dialog"
    width="550px"
    :center="true"
    :show-close="false"
    :append-to-body="true"
    :close-on-click-modal="false"
    :modal-append-to-body="true"
    :close-on-press-escape="false"
    :visible.sync="loadingVisible"
  >
    <div class="dialog-container">
      <p><img class="dialog-loading" src="../../images/jiazai.png"></p>
      <p>{{ loadingContent }}</p>
    </div>
  </el-dialog>
</template>
<script>
export default {
  name: 'ImportLoading',
  props: {
    loadingVisible: {
      type: Boolean
    },
    loadingContent: {
      type: String
    }
  },
  data() {
    return {

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
        .dialog-loading{
        	animation: rotate 3s linear infinite;
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
