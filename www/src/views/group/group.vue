<template>
  <el-container>
    <el-main>
      <el-header class="pagetab">
        <h4>群组管理</h4>
      </el-header>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="()=>showEditForm({})">添加群组</el-button>
      </el-row>
      <el-table
        v-loading="groups.loading"
        :data="groups.list"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="100"
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
          prop="group_tg_name"
          label="群组名称"
        />
        <el-table-column
          prop="group_tg_id"
          label="TG ID"
          align="center"
        />
        <el-table-column
          label="操作"
          align="center"
        >
          <template slot-scope="scope">
            <el-button type="success" size="mini" @click="()=>showEditForm(scope.row)">编辑</el-button>
            <el-button type="danger" size="mini" @click="()=>doDeleteGroup(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :page-size="groups.size"
        :total="groups.totalnum"
        @current-change="currentPageChange"
      />
    </el-main>
    <el-dialog
      :title="!!editForm.id ? '修改群组': '添加群组'"
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
        <el-form-item label="群组名称" prop="group_tg_name">
          <el-input
            ref="token"
            v-model="editForm.group_tg_name"
            tabindex="2"
            placeholder="群组名称"
          />
        </el-form-item>
        <el-form-item label="Telegram ID" prop="group_tg_id">
          <el-input
            ref="token"
            v-model.number="editForm.group_tg_id"
            tabindex="2"
            placeholder="群组名称"
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
import { getGroupList, saveGroup, deleteGroup } from '@/api/group'
import { s3config } from '@/settings'
export default {
  data() {
    return {
      s3config,
      allBots: [],
      groups: {
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
        bot_id: [
          { required: true, message: 'Please select a bot', trigger: 'blur' }
        ],
        group_tg_name: [
          { required: true, message: 'Please input group name', trigger: 'blur' }
        ],
        group_tg_id: [
          { required: true, message: 'Please input group telegram id', trigger: 'blur' }
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
    this.loadGroupList()
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
    async loadGroupList() {
      this.groups.loading = true
      const data = await getGroupList({
        page: this.groups.page,
        size: this.groups.size
      })
      console.log(data)
      this.groups.loading = false
      this.groups.list = data.list
      this.groups.totalnum = data.total
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
    },
    doEditFormSubmit() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          console.log('editForm:', this.editForm)
          saveGroup(this.editForm).then(() => {
            this.$message({
              message: (this.editForm.id) ? 'Edit success' : 'Add success',
              type: 'success'
            })
            this.editDialogVisible = false
            this.loadGroupList()
          }).finally(() => {
            this.$refs.editForm.resetFields()
          })
        }
      })
    },
    doDeleteGroup(row) {
      this.$confirm('Are you sure to delete this group?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        deleteGroup(row).then(() => {
          this.$message({
            type: 'success',
            message: 'Delete success'
          })
          this.loadGroupList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: 'Delete canceled'
        })
      })
    },
    getBotInfo(bot_id) {
      return this.allBots.find(item => item.id === bot_id)
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

