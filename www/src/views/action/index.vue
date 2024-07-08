<template>
  <el-container>
    <el-main>
      <el-header class="pagetab">
        <h4>Action管理</h4>
      </el-header>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="()=>showEditForm({})">添加Action</el-button>
      </el-row>
      <el-table
        v-loading="actions.loading"
        :data="actions.list"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="100"
        />
        <el-table-column
          prop="bot_name"
          label="Bot Name"
          align="center"
        />
        <el-table-column
          prop="command"
          label="Command"
          align="center"
        />
        <el-table-column
          prop="image"
          label="Image"
          align="center"
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
          label="Content"
        />
        <el-table-column
          prop="extension"
          label="Extension"
        />
        <el-table-column
          label="操作"
          align="center"
        >
          <template slot-scope="scope">
            <el-button type="success" size="mini" @click="()=>showEditForm(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :page-size="actions.size"
        :total="actions.totalnum"
        @current-change="currentPageChange"
      />
    </el-main>
    <el-dialog
      :title="!!editForm.id ? '修改Action': '添加Action'"
      :visible.sync="editDialogVisible"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
        <el-form-item label="Bot" prop="bot_name">
          <el-select
            v-model="editForm.bot_name"
            placeholder="Select a Bot"
          >
            <el-option
              v-for="item in allBots"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Command" prop="command">
          <el-input
            ref="token"
            v-model="editForm.command"
            tabindex="2"
            placeholder="command"
          />
        </el-form-item>
        <el-form-item label="图片(建议640*320)" prop="image">
          <S3Upload
            class="avatar-uploader"
            accept=".jpg,.png"
            :single="true"
            :config="s3config"
            :file-list="editFormImageFiles"
          />
          <el-input
            ref="icon_url"
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
            :rows="5"
          />
        </el-form-item>
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
                  v-for="item in ['url', 'webapp']"
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
import { getActionList, saveAction } from '@/api/action'
import S3Upload from '@/components/S3Upload'
import { s3config } from '@/settings'
export default {
  components: {
    S3Upload
  },
  data() {
    return {
      s3config,
      allBots: [],
      actions: {
        list: [],
        page: 1,
        size: 20,
        totalnum: 0,
        loading: false
      },
      editForm: {},
      editFormImageFiles: [],
      editFormExtends: [],
      editFormRules: {
        bot_name: [
          { required: true, message: 'Please select a bot', trigger: 'blur' }
        ],
        command: [
          { required: true, message: 'Please input command', trigger: 'blur' }
        ]
      },
      editDialogVisible: false
    }
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
  mounted() {
    this.loadAllBots()
    this.loadActionList()
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
    async loadActionList() {
      this.actions.loading = true
      const data = await getActionList({
        page: this.actions.page,
        size: this.actions.size
      })
      console.log(data)
      this.actions.loading = false
      this.actions.list = data.actions
      this.actions.totalnum = data.total
    },
    async loadAllBots() {
      const data = await getAllBots()
      console.log('loadAllBots:', data)
      this.allBots = data.bots || []
      console.log('this.allBots:', this.allBots)
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
    doEditFormSubmit() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          console.log('editForm:', this.editForm)
          saveAction(this.editForm).then(() => {
            this.$message({
              message: (this.editForm.id) ? 'Edit success' : 'Add success',
              type: 'success'
            })
            this.editDialogVisible = false
            this.loadActionList()
          }).finally(() => {
            this.$refs.editForm.resetFields()
          })
        }
      })
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
      const extList = _.cloneDeep(this.editFormExtends) || []
      extList.splice(index, 1)
      this.editFormExtends = extList
      this.editForm.extension = JSON.stringify(this.editFormExtends)
    },
    changeExt() {
      this.editForm.extension = JSON.stringify(this.editFormExtends)
    }
  }
}
</script>

<style lang="scss" scoped>
.card-panel{
  text-align: center;
}
.card-panel span{
  font-size: 12px;
  color: #999;
}
.mean {
  color: rgb(78, 136, 224);
}
.span-red {
  color: red;
}
.span-green {
  color: green;
}
</style>

