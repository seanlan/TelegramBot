<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">任务管理</h4>
    </el-header>
    <el-main>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="showEditModal">添加任务</el-button>
      </el-row>
      <el-divider />
      <el-table
        ref="multipleTable"
        v-loading="table.loading"
        :data="table.list"
        tooltip-effect="dark"
        style="width: 100%"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="100"
        />
        <el-table-column label="图标" width="100">
          <template slot-scope="scope">
            <div class="tablecell">
              <CustomImage
                style="width: 50px; height: 50px"
                :src="scope.row.icon_url ? scope.row.icon_url : ''"
                fit="contain"
                class="shoplogo"
                :preview-src-list="[scope.row.icon_url ? scope.row.icon_url : '']"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column
          prop="title"
          label="标题"
          width="160"
        />
        <el-table-column
          prop="content"
          label="内容"
        >
          <template slot-scope="scope">
            <div style="white-space: pre-line;">
              <span>{{ scope.row.content }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          prop="goto_url"
          label="链接"
        />
        <el-table-column
          prop="req_count"
          label="请求次数"
        />
        <el-table-column
          prop="reward"
          label="奖励（Satoshi ）"
        />
        <el-table-column
          prop="task_type"
          label="任务类型"
        >
          <template slot-scope="scope">
            <span>{{ filterTaskType(scope.row.task_type).label }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="is_tg_open" label="在tg中打开">
          <template slot-scope="scope">
            <span>{{ scope.row.is_tg_open > 0 ? '是' : '否' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="ranking"
          label="排序"
          width="100"
        />
        <el-table-column prop="is_enable" label="是否开启">
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.is_enable"
              :active-value="1"
              :inactive-value="0"
              @change="onlineChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="160"
        >
          <template slot-scope="scope">
            <el-button type="primary" size="mini" @click="()=>{showEditModal(scope.row)}">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :current-page="table.page"
        :page-size="table.pagesize"
        :hide-on-single-page="false"
        :total="table.totalnum"
        @current-change="PageChange"
      />
    </el-main>
    <el-dialog
      :title="editForm.id > 0? '修改任务': '添加任务'"
      :visible.sync="editDialogVisible"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
        <el-form-item label="图标" prop="icon_url">
          <S3Upload
            class="avatar-uploader"
            accept=".jpg,.png"
            :single="true"
            :config="s3config"
            :file-list="editFormIconFiles"
          />
          <el-input
            ref="icon_url"
            v-model="editForm.icon_url"
            tabindex="1"
            placeholder="请输入图片地址"
          />
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input
            ref="title"
            v-model="editForm.title"
            tabindex="1"
            placeholder="请输入标题"
          />
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input
            ref="content"
            v-model="editForm.content"
            tabindex="2"
            placeholder="请输入内容"
            type="textarea"
            :rows="4"
          />
        </el-form-item>
        <el-form-item label="跳转链接" prop="goto_url">
          <el-input
            ref="goto_url"
            v-model="editForm.goto_url"
            tabindex="1"
            placeholder="请输入链接"
          />
        </el-form-item>
        <el-form-item label="需要完成的次数" prop="req_count">
          <el-input
            ref="req_count"
            v-model.number="editForm.req_count"
            placeholder="需要完成的次数"
          />
        </el-form-item>
        <el-form-item label="任务奖励" prop="reward">
          <el-input
            ref="reward"
            v-model.number="editForm.reward"
            placeholder="任务奖励(Satoshi)"
          />
        </el-form-item>
        <el-form-item label="任务类型" prop="task_type">
          <el-select v-model="editForm.task_type" placeholder="请选择任务类型">
            <el-option
              v-for="item in taskTypeOption"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="editForm.task_type==2" label="等待时间（秒）" prop="keep_seconds">
          <el-input
            ref="keep_seconds"
            v-model.number="editForm.keep_seconds"
            placeholder="等待时间（秒）"
          />
        </el-form-item>
        <el-form-item label="展示类型" prop="show_type">
          <el-select v-model="editForm.show_type" placeholder="请选择任务类型">
            <el-option
              v-for="item in taskShowOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="主任务筛选条件" prop="main_condition">
          <el-select v-model="editForm.isRepeat" placeholder="请选择任务类型">
            <el-option label="不可重复" :value="false" />
            <el-option label="可重复" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="ranking">
          <el-input
            ref="ranking"
            v-model.number="editForm.ranking"
            placeholder="排序值越大越靠前"
          />
        </el-form-item>
        <el-form-item label="是否在TG中打开" prop="is_tg_open">
          <el-switch
            v-model="editForm.is_tg_open"
            :active-value="1"
            :inactive-value="0"
          />
        </el-form-item>
        <el-form-item label="是否开启" prop="is_open">
          <el-switch
            v-model="editForm.is_enable"
            :active-value="1"
            :inactive-value="0"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button v-if="editForm.promotion_id > 0" type="primary" @click="submitEditForm">编 辑</el-button>
        <el-button v-else type="primary" @click="submitEditForm">添 加</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import { taskList, taskSave } from '@/api/task'
import { taskTypeOption, taskShowOptions } from '@/utils/config'
import S3Upload from '@/components/S3Upload'
import { s3config } from '@/settings'
export default {
  components: {
    S3Upload
  },
  data() {
    return {
      table: {
        loading: false,
        list: [],
        pagesize: 20,
        totalnum: 0,
        page: 1
      },
      editDialogVisible: false,
      editForm: {
        id: 0
      },
      editFormIconFiles: [],
      editFormRules: {
        icon_url: [{ required: true, trigger: 'blur', message: '图标不能为空' }],
        title: [{ required: true, trigger: 'blur', message: '标题不能为空' }],
        content: [{ required: true, trigger: 'blur', message: '内容不能为空' }]
      },
      taskTypeOption,
      taskShowOptions,
      s3config
    }
  },
  computed: {
  },
  watch: {
    editFormIconFiles: {
      handler(newVal) {
        console.log('newVal', newVal)
        if (!newVal || newVal.length === 0) {
          this.editForm.icon_url = ''
        } else {
          this.editForm.icon_url = newVal[0].url
        }
      }
    }
  },
  async mounted() {
    this.loadData()
  },
  methods: {
    async loadData() {
      this.table.loading = true
      var params = {
        page: this.table.page,
        size: this.table.pagesize
      }
      const res = await taskList(params)
      this.table.list = res.list || []
      this.table.totalnum = res.total || 0
      this.table.loading = false
    },
    PageChange(e) {
      this.table.page = e
      this.loadData()
    },
    showEditModal(row) {
      const main_condition_str = row.main_condition || '{}'
      const main_condition = JSON.parse(main_condition_str)
      const isRepeat = !(main_condition.conds)
      this.editForm = {
        id: row ? row.id : 0,
        title: row ? row.title : '',
        content: row ? row.content : '',
        icon_url: row ? row.icon_url : '',
        ranking: row ? row.ranking : 0,
        goto_url: row ? row.goto_url : '',
        req_count: row ? row.req_count : 0,
        reward: row ? row.reward : 0,
        is_tg_open: row ? row.is_tg_open : 0,
        show_type: row ? row.show_type : '',
        task_type: row ? row.task_type : 0,
        is_enable: row ? row.is_enable : 0,
        is_by_day: row ? row.is_by_day : 0,
        keep_seconds: row ? row.keep_seconds : 0,
        isRepeat
      }
      this.editFormIconFiles = row?.icon_url ? [{ url: row.icon_url }] : []
      this.editDialogVisible = true
    },
    submitEditForm() {
      this.$refs.editForm.validate(async valid => {
        if (valid) {
          const params = {
            ...this.editForm
          }
          if (this.editForm.isRepeat) {
            params.main_condition = '{}'
          } else {
            params.main_condition = '{"conds": [{"type": "no_in_life"}]}'
          }
          await taskSave(params)
          this.$message.success('提交成功')
          this.editDialogVisible = false
          this.loadData()
          return true
        } else {
          return false
        }
      })
    },
    async onlineChange(row) {
      const update = {
        ...row
      }
      await taskSave(update)
      this.loadData()
    },
    filterTaskType(taskType) {
      return taskTypeOption.find(item => item.value === taskType)
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
