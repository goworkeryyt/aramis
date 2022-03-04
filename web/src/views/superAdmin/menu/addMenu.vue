<template name="component-name">
  <div>
    <div class="crumbs-header row-between header-sticky h-70">
      <div class="title">{{ Title }}</div>
      <div>
        <el-button
          class="cancel-btn"
          @click="handleCloseDrawr"
        >取 消</el-button>
        <el-button
          class="submit"
          :loading="btnLoading"
          @click="submitForm('ruleForm')"
        >提 交</el-button>
        <i class="el-icon-close ml-10" @click="handleCloseDrawr" />
      </div>
    </div>

    <el-form
      ref="menuForm"
      :inline="true"
      :model="form"
      :rules="rules"
      label-position="top"
      label-width="85px"
    >
      <div class="theme row-flex">
        <i />
        <span>基本信息</span>
      </div>
      <div class="form_box">
        <el-form-item label="路由Name" prop="routerName" style="width: 30%">
          <el-input
            v-model="form.routerName"
            autocomplete="off"
            placeholder="唯一英文字符串"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item label="路由Path" prop="url" style="width: 30%">
          <el-input
            v-model="form.url"
            :disabled="true"
            autocomplete="off"
            placeholder="唯一英文字符串"
          />
        </el-form-item>
        <el-form-item prop="IsShow" label="是否展示" style="width: 30%">
          <el-select v-model="form.IsShow" placeholder="是否在列表展示">
            <el-option :value="-1" label="否" />
            <el-option :value="1" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="父节点ID" style="width: 30%">
          <el-cascader
            v-model="form.parentId"
            style="width: 100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="optionsProps"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="文件路径" prop="filePath" style="width: 60%">
          <el-input
            v-model="form.filePath"
            autocomplete="off"
            placeholder="请输入文件路径 例：view/routerHolder.vue"
          />
          <!-- <span style="font-size: 12px; margin-right: 12px"
            >如果菜单包含子菜单，请创建router-view二级路由页面或者</span
          ><el-button
            size="mini"
            @click="form.filePath = `view/${parentRouterName}/${form.url}.vue`"
            >点我设置</el-button
          > -->
        </el-form-item>
        <el-form-item label="展示名称" prop="menuName" style="width: 30%">
          <el-input
            v-model="form.menuName"
            autocomplete="off"
            placeholder="请输入菜单展示名称"
          />
        </el-form-item>
        <el-form-item
          v-if="form.parentId == '0' || !form.parentId"
          label="图标"
          prop="menuIcon"
          style="width: 30%"
        >
          <icon
            :menu-icon="form.menuIcon"
            style="width: 100%"
            @getMenuIcon="getMenuIcon"
          />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort" style="width: 30%">
          <!-- <el-input
            v-model.number="form.sort"
            autocomplete="off"
            placeholder="请输入排序标记"
          /> -->
          <el-input-number
            v-model="form.sort"
            controls-position="right"
            :min="0"
          />
        </el-form-item>
      </div>

      <div
        v-show="form.parentId != '0' && form.parentId"
        class="theme row-flex"
      >
        <i />
        <span>功能（按钮）权限设置</span>
        <span
          style="
            color: #696d75;
            margin-left: 50px;
            font-weight: 400;
            font-size: 14px;
          "
        />
      </div>
      <div v-show="form.parentId != '0' && form.parentId" class="form_box">
        <div style="text-align: right">
          <el-button class="add" @click="addButton">新增按钮</el-button>
          <el-button
            v-if="type == 'edit' && buttonList.length == 0"
            class="default-btn"
            @click="getBaseButton"
          >新增基础按钮</el-button>
        </div>
        <el-tag
          v-for="tag in buttonList"
          :key="tag.buttonNo"
          style="margin: 10px 10px 10px 0"
          :closable="!tag.disabled"
          @click="editButton(tag)"
          @close="handleClose(tag)"
        >{{ tag.buttonName }}</el-tag>
      </div>
    </el-form>

    <!-- 添加按钮 -->
    <el-dialog
      width="600px"
      top="20vh"
      :title="btnType == 'edit' ? '编辑按钮' : '新增按钮'"
      :append-to-body="true"
      :visible.sync="dialogFormVisible"
      center
      @close="resetButtonForm('dialogForm')"
    >
      <el-form
        ref="dialogForm"
        :rules="formRolues"
        :model="dialogForm"
        label-width="160px"
      >
        <el-form-item label="按钮名称：" prop="buttonName">
          <el-input
            v-model="dialogForm.buttonName"
            class="input-width"
            placeholder="请输入按钮名称"
          />
        </el-form-item>
        <el-form-item label="按钮标识：" prop="buttonIdentity">
          <el-input
            v-model="dialogForm.buttonIdentity"
            class="input-width"
            placeholder="请输入按钮标识"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button
          class="cancel-btn"
          @click="resetButtonForm('dialogForm')"
        >取 消</el-button>
        <el-button
          :loading="btnloading"
          class="submit"
          @click="submitButtonForm('dialogForm')"
        >确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import icon from './icon.vue'
import infoList from '@/mixins/infoList'
export default {
  components: {
    icon
  },
  mixins: [infoList],
  props: {
    title: {
      type: String
    },
    form: {
      type: Object
    },
    type: {
      type: String
    },
    isEdit: {
      type: Boolean
    }
  },
  data() {
    return {
      btnLoading: false,
      menuOption: [
        {
          id: '0',
          title: '根菜单'
        }
      ],
      optionsProps: {
        checkStrictly: true,
        label: 'title',
        value: 'id',
        disabled: 'disabled',
        emitPath: false
      },
      // parentRouterName: "",
      form: {
        url: '',
        routerName: '',
        IsShow: 1,
        parentId: '',
        filePath: '',
        menuName: '',
        menuIcon: '',
        sort: 1
      },
      rules: {
        routerName: [
          { required: true, message: '请输入菜单name', trigger: 'blur' }
        ],
        url: [{ required: true, message: '请输入路由Path', trigger: 'blur' }],
        filePath: [
          { required: true, message: '请输入文件路径', trigger: 'blur' }
        ],
        menuName: [
          { required: true, message: '请输入菜单展示名称', trigger: 'blur' },
          {
            pattern: /^[\u4E00-\u9FA5A-Za-z]{2,20}$/,
            message: '菜单展示名称为2-20位中文或字母组合'
          }
        ],
        menuIcon: [
          { required: true, message: '请选择图标', trigger: 'blur' }
        ]
      },

      buttonList: [],
      buttonNo: 1,
      btnType: 'add',
      dialogFormVisible: false,
      dialogForm: {
        buttonName: '',
        buttonIdentity: ''
      },
      formRolues: {
        buttonName: [
          {
            required: true,
            message: '按钮名称不能为空！',
            trigger: 'blur'
          },
          {
            pattern: /^[\u4E00-\u9FA5A-Za-z]{2,20}$/,
            message: '按钮名称为2-20位中文、字母'
          }
        ],
        buttonIdentity: [
          {
            required: true,
            message: '按钮标识不能为空！',
            trigger: 'blur'
          },
          {
            pattern: /^[A-Za-z]{2,20}$/,
            message: '按钮标识为2-20位字母组合'
          }
        ]
      },
      btnloading: false
    }
  },
  watch: {
    Form(newVal) {
      this.form = newVal
    }
  },
  beforeMount() {},
  methods: {
    setOptions(arr, val) {
      if (this.type === 'add') {
        this.form.parentId = val === '0' ? '0' : val.id
        // this.parentRouterName = val.routerName;

        this.buttonList = val === '0' ? [] : this.Constant.baseButtons
      }
      this.menuOption = [
        {
          id: '0',
          title: '根目录'
        }
      ]
      this.setMenuOptions(arr, this.menuOption, false)
    },
    setMenuOptions(menuData, optionsData, disabled) {
      menuData &&
        menuData.forEach((item) => {
          if (!item.parentId) {
            item.parentId = '0'
          }
          if (item.children && item.children.length > 0) {
            const option = {
              title: item.menuName,
              id: String(item.id),
              disabled: disabled || item.id === this.form.id,
              children: []
            }
            this.setMenuOptions(
              item.children,
              option.children,
              disabled || item.id === this.form.id
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.menuName,
              id: String(item.id),
              disabled: disabled || item.id === this.form.id
            }
            optionsData.push(option)
          }
        })
    },

    // 获取菜单详情(按钮详情)
    async findMenu(id) {
      this.buttonList = []
      const res = await this.Api.findMenu({ id })
      this.btnLoading = false
      const { code, message, content } = res.data

      if (code === '0') {
        if (content.buttons) {
          content.buttons.forEach((item) => {
            this.buttonNo += 1
            item.buttonNo = this.buttonNo
          })

          this.buttonList = content.buttons
        } else {
          this.buttonList = []
        }
      } else {
        this.$message({
          type: 'warning',
          center: true,
          message
        })
      }
    },

    // 初始化弹窗内表格方法
    initForm() {
      this.$refs.menuForm.resetFields()
      this.form = {
        url: '',
        routerName: '',
        IsShow: 1,
        ParentId: '',
        filePath: '',
        menuName: '',
        menuIcon: '',
        sort: 1,
        buttons: []
      }
      this.buttonNo = 1
      this.buttonList = []
    },
    changeName() {
      this.form.url = this.form.routerName
    },
    // 获取图标
    getMenuIcon(val) {
      this.form.menuIcon = val
    },
    // 关闭抽屉
    handleCloseDrawr() {
      this.$emit('closeDrawer')
    },

    // 新增按钮
    addButton() {
      this.dialogFormVisible = true
      this.btnType = 'add'
      this.buttonNo += 1
      this.dialogForm.buttonNo = this.buttonNo
    },

    // 添加基础按钮（增删改查）
    getBaseButton() {
      this.buttonList = this.buttonList.concat(this.Constant.baseButtons)
    },

    // 编辑按钮
    editButton(tag) {
      this.dialogFormVisible = true
      this.btnType = 'edit'
      this.dialogForm = { ...tag }
    },

    // 删除按钮
    handleClose(tag) {
      this.buttonList.splice(this.buttonList.indexOf(tag), 1)
    },

    // 取消新建或修改按钮
    resetButtonForm(form) {
      this.dialogFormVisible = false
      this.$refs[form].resetFields()
    },

    // 新建或修改按钮
    submitButtonForm(formName) {
      this.$refs[formName].validate(async(valid) => {
        if (valid) {
          const data = { ...this.dialogForm }

          if (this.btnType === 'add') {
            this.buttonNo += 1;
            (data.buttonNo = this.buttonNo), this.buttonList.push(data)
          } else {
            this.buttonList.forEach((item) => {
              if (item.buttonNo == data.buttonNo) {
                item.buttonName = data.buttonName
                item.buttonIdentity = data.buttonIdentity
              }
            })
          }

          this.dialogFormVisible = false
        } else {
          return false
        }
      })
    },

    // 提交
    submitForm() {
      this.$refs.menuForm.validate((valid) => {
        if (valid) {
          this.btnLoading = true
          const data = { ...this.form }
          data.buttons = this.buttonList

          if (data.parentId === '0') {
            data.parentId = ''
            data.buttons = []
          } else {
            data.buttons.forEach((item) => {
              delete item.buttonNo
            })
          }

          switch (this.type) {
            case 'add':
              this.createMenu(data)
              break
            case 'edit':
              this.updateMenu(data)
              break
            default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true,
                  center: true
                })
              }
              break
          }
        } else {
          return false
        }
      })
    },
    // 创建菜单请求
    async createMenu(params) {
      const res = await this.Api.createMenu(params)
      this.btnLoading = false
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.$emit('closeDrawer', true)
        this.$store.dispatch('router/getSelfMenuTree')
      }
    },
    // 编辑菜单请求
    async updateMenu(params) {
      const res = await this.Api.updateMenu(params)
      this.btnLoading = false
      const { code, message } = res.data
      this.$message({
        type: code !== '0' ? 'warning' : 'success',
        center: true,
        message
      })
      if (code === '0') {
        this.$emit('closeDrawer', true)
        this.$store.dispatch('router/getSelfMenuTree')
      }
    }
  }
}
</script>

<style scoped lang="scss">
.form_box /deep/ .icon {
  line-height: 28px !important;
}

.function-checkall {
  margin-bottom: 20px;

  .el-checkbox__inner {
    width: 18px;
    height: 18px;
    border: 1px solid rgba(153, 153, 153, 1);

    &::after {
      height: 9px;
      left: 5px;
      top: 0px;
      width: 4px;
    }
  }
}

.function-btn {
  margin-left: 50px;

  label {
    margin-bottom: 20px;
    width: 85px;

    .el-checkbox__inner {
      width: 18px;
      height: 18px;
      border: 1px solid rgba(153, 153, 153, 1);

      &::after {
        height: 9px;
        left: 5px;
        top: 0px;
        width: 4px;
      }
    }
  }
}
</style>
<style  scoped="scoped">
.function-btnPanel {
  margin: 20px 20px;
  width: 688px;
  border: 1px dashed #979797;
  padding: 22px 20px;
  animation-name: mymove;
  animation-duration: 0.5s;
  /* 5s表示执行动画的时间，0或不写则无动画效果 */
  /* 兼容调试如果animation-name执行，那么-wekit-animation-name则不执行 */
  -webkit-animation-name: mymove;
  -webkit-animation-duration: 0.5s;
  /* 5s表示执行动画的时间，0或不写则无动画效果 */
}

@keyframes mymove

  /* 对应animation-name，里面为执行的动画*/ {
  from {
    left: 200px;
  }

  /*执行动画的初始位置*/
  to {
    left: 0px;
  }

  /*动画结束位置*/
  0% {
    opacity: 0.1;
    /*初始状态 透明度为10%*/
  }

  25% {
    opacity: 0.3;
  }

  50% {
    opacity: 0.6;
    /*中间状态 透明度为50%*/
  }

  75% {
    opacity: 0.8;
  }

  85% {
    opacity: 0.9;
  }

  100% {
    opacity: 1;
    /*结尾状态 不透明*/
  }
}
</style>
