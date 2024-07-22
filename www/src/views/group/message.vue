<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">消息推送</h4>
    </el-header>
    <el-main>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="()=>showEditForm({})">添加消息</el-button>
      </el-row>
      <el-table
        ref="multipleTable"
        v-loading="table_loading"
        :data="list"
        tooltip-effect="dark"
        style="width: 100%"
      >
        <el-table-column
          prop="id"
          label="ID"
        />
        <el-table-column
          prop="bot_id"
          label="Bot Name"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ getBotInfo(scope.row.bot_id).name }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="image"
          label="图片"
        >
          <template slot-scope="scope">
            <div class="tablecell">
              <CustomImage
                style="width: 50px; height: 50px"
                :src="scope.row.image ? scope.row.image : ''"
                fit="contain"
                class="shoplogo"
                :preview-src-list="[scope.row.image ? scope.row.image : '']"
              />
            </div>
          </template>
        </el-table-column>
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
          prop="extension"
          label="附加选项"
        />
        <el-table-column
          prop="start_at"
          label="开始时间（北京时间）"
          width="160"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.start_at | utcTimeFormat }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="state"
          label="状态"
          width="100"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.state | messagePushState }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="160"
        >
          <template slot-scope="scope">
            <el-button type="success" size="mini" @click="showEditForm(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
    <el-dialog
      :title="!!editForm.id ? '修改消息推送': '添加消息推送'"
      :visible.sync="editDialogVisible"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
        <el-form-item label="Bot" prop="bot_id">
          <el-select
            v-model="editForm.bot_id"
            placeholder="Select a Bot"
          >
            <el-option
              v-for="item in allBots"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="图片" prop="image">
          <S3Upload
            class="avatar-uploader"
            accept=".jpg,.png"
            :single="true"
            :config="s3config"
            :file-list="editFormImageFiles"
            @onSuccess="uploadSuccess"
          />
          <el-input
            ref="image"
            v-model="editForm.image"
            tabindex="1"
            placeholder="请输入图片地址"
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

        <!-- <el-form-item label="附加选项" prop="ext">
          <el-input
            ref="ext"
            v-model="editForm.ext"
            tabindex="2"
            placeholder="请输入附加选项"
            type="textarea"
            :rows="4"
          />
        </el-form-item> -->

        <el-form-item label="附加选项" prop="ext">
          <el-button type="primary" icon="el-icon-plus" size="small" @click="addExt">添加</el-button>
        </el-form-item>

        <el-form-item
          v-for="(ext, index) in editFormExtends"
          :key="index"
        >
          <el-row>
            <el-col :span="8">
              <el-input
                v-model="ext.text"
                placeholder="请输入文本"
                @change="changeExt"
              />
            </el-col>
            <el-col :span="8">
              <el-select
                v-model="ext.type"
                placeholder="请选择类型"
                @change="changeExt"
              >
                <el-option
                  v-for="item in ['url']"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-col>
            <el-col :span="8">
              <el-input
                v-model="ext.url"
                placeholder="请输入链接"
                @change="changeExt"
              />
            </el-col>
            <el-col :span="4">
              <el-button @click.prevent="removeExt(index)">删除</el-button>
            </el-col>
          </el-row>
        </el-form-item>

        <el-form-item label="开始时间" prop="start_at">
          <el-date-picker
            v-model="editForm.start_at"
            type="datetime"
            placeholder="选择日期时间"
            value-format="timestamp"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button v-if="!!editForm.id" type="primary" @click="doEditFormSubmit">编 辑</el-button>
        <el-button v-else type="primary" @click="doEditFormSubmit">添 加</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import _ from 'lodash'
import { getAllBots } from '@/api/bot'
import S3Upload from '@/components/S3Upload'
import { getMessagePushList, saveMessagePush } from '@/api/group'
import { s3config } from '@/settings'

export default {
  components: {
    S3Upload
  },
  data() {
    return {
      s3config,
      allBots: [],
      table_loading: false,
      editDialogVisible: false,
      modaltype: 'add',
      editForm: {
        ext: '[]',
        start: new Date().getTime(),
        image: '123123',
        image_files: []
      },
      editFormImageFiles: [],
      editFormExtends: [],
      list: [],
      pagesize: 20,
      totalnum: 0,
      page: 1,
      editFormRules: {
        content: [{ required: true, trigger: 'blur', message: 'VALUE不能为空' }],
        bot_id: [{ required: true, trigger: 'blur', message: 'bot不能为空' }],
        start_at: [{ required: true, trigger: 'blur', message: '开始时间不能为空' }]
      }
    }
  },
  computed: {
  },
  watch: {
    editFormImageFiles: {
      handler(newVal) {
        console.log('newVal', newVal)
        if (!newVal || newVal.length === 0) {
          this.editForm.image = ''
        } else {
          this.editForm.image = newVal[0].url
        }
      }
    },
    editFormExtends: {
      handler(newVal) {
        this.editForm.extension = JSON.stringify(newVal)
      }
    }
  },
  async mounted() {
    this.loadMessageList()
    this.loadAllBots()
  },
  methods: {
    async loadAllBots() {
      const data = await getAllBots()
      console.log('loadAllBots:', data)
      this.allBots = data.bots || []
      console.log('this.allBots:', this.allBots)
    },
    async loadMessageList() {
      this.table_loading = true
      var params = {
        page: this.page,
        size: this.pagesize
      }
      const res = await getMessagePushList(params)
      this.list = res.list || []
      this.table_loading = false
    },
    doEditFormSubmit() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          console.log('editForm:', this.editForm)
          saveMessagePush(this.editForm).then(() => {
            this.$message({
              message: (this.editForm.id) ? 'Edit success' : 'Add success',
              type: 'success'
            })
            this.editDialogVisible = false
            this.loadMessageList()
          }).finally(() => {
            this.$refs.editForm.resetFields()
          })
        }
      })
    },
    showEditForm(row) {
      this.editForm = _.cloneDeep(row)
      this.editDialogVisible = true
      if (row.image) {
        this.editFormImageFiles = [{
          url: row.image
        }]
      } else {
        this.editFormImageFiles = []
      }
      if (row.extension) {
        this.editFormExtends = JSON.parse(row.extension)
      } else {
        this.editFormExtends = []
      }
    },
    addExt() {
      const extList = _.cloneDeep(this.editFormExtends) || []
      extList.push({
        text: '',
        type: 'url',
        url: ''
      })
      this.editFormExtends = extList
      this.editForm.extension = JSON.stringify(this.editFormExtends)
    },
    removeExt(index) {
      const extList = _.cloneDeep(this.editForm.extList) || []
      extList.splice(index, 1)
      this.editForm.extList = extList
      this.editForm.ext = JSON.stringify(this.editForm.extList)
    },
    changeExt() {
      this.editForm.extension = JSON.stringify(this.editFormExtends)
    },
    uploadSuccess(file) {
      console.log('file', file)
      console.log('this.editForm', this.editForm)
      this.editForm.image = file.url
    },
    getBotInfo(bot_id) {
      return this.allBots.find(item => item.id === bot_id)
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
