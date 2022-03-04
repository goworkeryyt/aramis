<template>
  <el-dialog
    :title="titleText"
    :center="true"
    :show-close="false"
    :append-to-body="true"
    :close-on-click-modal="false"
    :modal-append-to-body="true"
    :close-on-press-escape="false"
    :visible.sync="operationVisible"
    width="30%"
    class="delete-confirm"
  >
    <div class="el-icon-close close-icon" @click="operationRes('no')" />
    <div class="message-container">
      <span class="el-icon-warning delete-img" />
      <span>{{ operationContent }}</span>
    </div>
    <span slot="footer" class="dialog-footer">
      <el-button class="cancel-btn" @click="operationRes('no')">取 消</el-button>
      <el-button class="submit" :loading="confirmLoading" @click="operationRes('yes')">确 定</el-button>
    </span>
  </el-dialog>
</template>
<script>
export default {
  name: 'OperationReminder',
  props: {
    operationVisible: {
      type: Boolean
    },
    operationContent: {
      type: String
    },
    operationResult: {
      type: Function
    },
    titleText: {
      type: String,
      default: '删除确认'
    },
    confirmLoading: {
      type: Boolean
    }
  },
  data() {
    return {

    }
  },
  methods: {
    operationRes(value) {
      this.$emit('operationResult', value)
    }
  }
}
</script>
<style scoped>
  .message-container {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
  }

  .delete-img {
    display: flex;
    font-size: 28px;
    color: #E6A23C;
    margin-right: 15px;
  }

  .close-icon {
    font-size: 24px;
    position: absolute;
    top: 20px;
    right: 20px;
    cursor: pointer;
  }

  .delete-confirm>>>.dialog-footer{
    display: flex;
    justify-content: flex-end;
  }

  .delete-confirm>>>.el-dialog__header {
    text-align: left;
  }
</style>
