<template>
  <el-button
    v-if="hasPerms(perms)"
    :size="size"
    :type="type"
    :icon="icon"
    :loading="loading"
    :disabled="disabled"
    @click="handleClick"
  >
    {{ permissionLabel(perms) }}
  </el-button>
</template>

<script>
import { hasPermission } from '@/permission/index.js'
export default {
  name: 'KtButton',
  props: {
    label: { // 按钮显示文本
      type: String,
      default: 'Button'
    },
    icon: { // 按钮显示图标
      type: String,
      default: ''
    },
    size: { // 按钮尺寸
      type: String,
      default: 'medium'
    },
    type: { // 按钮类型
      type: String,
      default: null
    },
    loading: { // 按钮加载标识
      type: Boolean,
      default: false
    },
    disabled: { // 按钮是否禁用
      type: Boolean,
      default: false
    },
    perms: { // 按钮权限标识，外部使用者传入
      type: String,
      default: null
    }
  },
  data() {
    return {
    }
  },

  mounted() {
  },
  methods: {
    handleClick: function() {
      // 按钮操作处理函数
      this.$emit('click', {})
    },
    // 是否有权限
    hasPerms: function(perms) {
      // 根据权限标识和外部指示状态进行权限判断
      if (perms && perms.show == 'true') {
        return true
      } else {
        return hasPermission(perms).show
      }
    },
    // 按钮名称
    permissionLabel(perms) {
  		if (perms && perms.label) {
        return perms.label
      } else {
        return hasPermission(perms).label
      }
  	}
  }
}
</script>

<style scoped>

</style>
