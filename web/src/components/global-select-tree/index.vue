<!--
 * @Description: 封装组件 el-select & el-tree 下拉树形选择
-->
<template>
  <el-select ref="selectTree" v-model="valueName" :value="value" :multiple="multiple" :clearable="clearable" @clear="clearHandle" @change="changeValue">
    <el-option :value="valueName" class="options">
      <el-tree
        id="tree-option"
        ref="selectTree"
        :accordion="accordion"
        :data="options"
        :props="props"
        :node-key="props.value"
        @node-click="handleNodeClick"
      >
        <span slot-scope="{ data }">
          <i :class="[data.color!=null?'ification_col':'']" :style="{'background-color':data.color}" />&nbsp;&nbsp;{{ data.name }}
        </span>
      </el-tree>
    </el-option>
  </el-select>
</template>
<script>
export default {
  name: 'ElTreeSelect',
  props: {
    // 配置项（对应的属性值）
    props: {
      type: Object,
      default: () => {
        return {
          value: 'id',
          children: 'children',
          label: 'name'
        }
      }
    },
    // 选项列表数据(树形结构的对象数组)
    options: {
      type: Array,
      default: () => {
        return []
      }
    },
    // 初始值（单选）
    value: {
      type: Object,
      default: () => {
        return {}
      }
    },
    // 初始值（多选）
    valueMultiple: {
      type: Array,
      default: () => {
        return []
      }
    },
    // 可清空选项
    clearable: {
      type: Boolean,
      default: true
    },
    // 自动收起
    accordion: {
      type: Boolean,
      default: false
    },
    // 是否多选
    multiple: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      resultValue: [], // 传给父组件的数组对象值
      valueName: this.multiple ? [] : '' // 输入框显示值
    }
  },
  watch: {
    value() {
      this.resultValue = this.multiple ? this.valueMultiple : this.value // 初始值
      this.initHandle()
    }
  },
  mounted() {
    this.resultValue = this.multiple ? this.valueMultiple : this.value // 初始值
    this.initHandle()
  },
  methods: {
    // 初始化显示
    initHandle() {
      if (this.resultValue) {
        if (this.multiple) {
          // 多选
          this.resultValue.forEach(item => this.valueName.push(item.name))
        } else {
          // 单选
          this.valueName = this.resultValue.name
        }
      }
      this.initScroll()
    },
    // 初始化滚动条
    initScroll() {
      this.$nextTick(() => {
        const scrollWrap = document.querySelectorAll('.el-scrollbar .el-select-dropdown__wrap')[0]
        const scrollBar = document.querySelectorAll('.el-scrollbar .el-scrollbar__bar')
        scrollWrap.style.cssText = 'margin: 0px; max-height: none; overflow: hidden;'
        scrollBar.forEach(ele => ele.style.width = 0)
      })
    },
    // 切换选项
    handleNodeClick(node) {
      // 点击叶子节点后被选中
      if (node.children == null || node.children == undefined) {
        if (this.multiple) {
          // 多选（判重后添加）
          let num = 0
          this.valueName.forEach(item => {
            item == node[this.props.label] ? num++ : num
          })
          if (num == 0) {
            this.valueName.push(node[this.props.label]) // 输入框显示值
            this.resultValue.push(node)
          }
        } else {
          // 单选
          this.$refs.selectTree.blur()
          this.valueName = node[this.props.label]
          this.resultValue = node
        }

        this.$emit('getValue', this.resultValue)
      }
    },
    // 从输入框中直接删除某个值时
    changeValue(val) {
      if (this.multiple) {
        // 多选（同时删掉传给父组件中多余的值，再传给父组件）
        this.resultValue.map((item, index) => {
          const i = val.indexOf(item.name)
          if (i == -1) {
            this.resultValue.splice(index, 1)
          }
        })
        this.$emit('getValue', this.resultValue)
      } else {
        // 单选
        this.$emit('getValue', val)
      }
    },
    // 清除选中
    clearHandle() {
      this.valueName = this.multiple ? [] : ''
      this.resultValue = []
      this.clearSelected()
      this.$emit('getValue', this.resultValue)
    },
    // 清空选中样式
    clearSelected() {
      const allNode = document.querySelectorAll('#tree-option .el-tree-node')
      allNode.forEach((element) => element.classList.remove('is-current'))
    }
  }
}
</script>
<style scoped>
    .el-scrollbar .el-scrollbar__view .el-select-dropdown__item {
        height: auto;
        max-height: 300px;
        padding: 0;
        overflow: hidden;
        overflow-y: auto;
    }
    .el-select-dropdown__item.selected {
        font-weight: normal;
    }
    ul li >>> .el-tree .el-tree-node__content {
        height: auto;
        padding: 0 20px;
    }
    .el-tree-node__label {
        font-weight: normal;
    }
    .el-tree >>> .is-current .el-tree-node__label {
        color: #409EFF;
        font-weight: 700;
    }
    .el-tree >>> .is-current .el-tree-node__children .el-tree-node__label {
        color: #606266;
        font-weight: normal;
    }
    .el-popper {
        z-index: 9999;
    }
    .ification_col {
        width: 20px;
        height: 10px;
        display: inline-block;
    }
</style>
