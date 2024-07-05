<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">配置管理</h4>
    </el-header>
    <el-main>
      <el-row type="flex" justify="end">
        <el-button type="warning" icon="el-icon-refresh" @click="refreshConfig">刷新缓存</el-button>
        <el-button type="primary" icon="el-icon-plus" @click="showAddModal">添加配置</el-button>
      </el-row>
      <el-table
        ref="multipleTable"
        v-loading="configs.loading"
        :data="configs.list"
        tooltip-effect="dark"
        style="width: 100%"
      >
        <el-table-column
          prop="key"
          label="KEY"
        />
        <el-table-column
          prop="value"
          label="VALUE"
        />
        <el-table-column
          prop="comment"
          label="备注"
          width="300"
        />
        <el-table-column
          label="操作"
          width="160"
        >
          <template slot-scope="scope">
            <el-button type="success" size="mini" @click="showEditModal(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :page-size="configs.size"
        :total="configs.totalnum"
        @current-change="currentPageChange"
      />
    </el-main>
    <el-dialog
      :title="modaltype=='edit'? '修改系统配置': '添加系统配置'"
      :visible.sync="addDialogVisible"
      width="50%"
    >
      <el-form ref="configForm" :model="configForm" :rules="groupRules" label-width="140px">
        <el-form-item label="KEY" prop="key">
          <el-input
            ref="key"
            v-model="configForm.key"
            tabindex="1"
            placeholder="KEY"
            :disabled="modaltype=='edit'? true: false"
          />
        </el-form-item>
        <el-form-item label="VALUE" prop="value">
          <el-input
            ref="value"
            v-model="configForm.value"
            tabindex="2"
            placeholder="VALUE"
            type="textarea"
            :rows="4"
          />
        </el-form-item>
        <el-form-item label="备注" prop="comment">
          <el-input
            ref="comment"
            v-model="configForm.comment"
            tabindex="2"
            placeholder="备注"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button v-if="modaltype === 'edit'" type="primary" @click="doAddGroup">编 辑</el-button>
        <el-button v-else type="primary" @click="doAddGroup">添 加</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import _ from 'lodash'
import { Message } from 'element-ui'
import { getConfigList, saveConfig, refreshConfig } from '@/api/config'

export default {
  components: {
  },
  data() {
    return {
      table_loading: false,
      addDialogVisible: false,
      modaltype: 'add',
      configForm: {},
      configs: {
        list: [],
        page: 1,
        size: 20,
        totalnum: 0,
        loading: false
      },
      groupRules: {
        key: [{ required: true, trigger: 'blur', message: 'KEY不能为空' }],
        remarks: [{ required: true, trigger: 'blur', message: '备注不能为空' }]
      }
    }
  },
  computed: {
  },
  async mounted() {
    this.loadConfigList()
  },
  methods: {
    handleSizeChange(val) {
      this.size = val
      this.page = 1
      this.loadData()
    },
    currentPageChange(e) {
      this.page = e
      this.loadData()
    },
    async loadConfigList() {
      this.table_loading = true
      const data = await getConfigList()
      this.configs.list = data.list
      this.configs.loading = false
    },
    doAddGroup() {
      this.$refs.configForm.validate(valid => {
        if (valid) {
          saveConfig(this.configForm).then(() => {
            Message({
              message: this.modaltype === 'add' ? '添加成功' : '修改成功',
              type: 'success'
            })
            this.addDialogVisible = false
            this.loadConfigList()
          }).catch((e) => {
            this.loading = false
            console.log(e)
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    showAddModal() {
      this.modaltype = 'add'
      this.addDialogVisible = true
      this.configForm = {
        key: '',
        value: '',
        comment: ''
      }
    },
    showEditModal(row) {
      this.modaltype = 'edit'
      this.addDialogVisible = true
      this.configForm = _.cloneDeep(row)
    },
    refreshConfig() {
      refreshConfig().then(() => {
        Message({
          message: '刷新成功',
          type: 'success'
        })
      }).catch((e) => {
        console.log(e)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
