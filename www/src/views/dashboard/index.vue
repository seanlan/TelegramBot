<template>
  <el-container>
    <el-main>
      <el-header class="pagetab">
        <h4>机器人管理</h4>
      </el-header>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="()=>showEditForm({})">添加BOT</el-button>
      </el-row>
      <el-table
        v-loading="bots.loading"
        :data="bots.list"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="100"
        />
        <el-table-column
          prop="name"
          label="Name"
          align="center"
        />
        <el-table-column
          prop="token"
          label="Token"
          align="center"
        />
        <el-table-column
          prop="enable"
          label="是否开启"
          align="center"
        >
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.enable"
              :active-value="1"
              :inactive-value="0"
              @change="()=>changeBotInfo(scope.row)"
            />
          </template>
        </el-table-column>
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
        :page-size="bots.size"
        :total="bots.totalnum"
        @current-change="currentPageChange"
      />
    </el-main>
    <el-dialog
      :title="!!editForm.id ? '修改BOT': '添加BOT'"
      :visible.sync="editDialogVisible"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
        <el-form-item label="Name" prop="name">
          <el-input
            ref="name"
            v-model="editForm.name"
            tabindex="1"
            placeholder="name"
          />
        </el-form-item>
        <el-form-item label="Token" prop="token">
          <el-input
            ref="token"
            v-model="editForm.token"
            tabindex="2"
            placeholder="token"
          />
        </el-form-item>
        <el-form-item label="是否开启" prop="enable">
          <el-switch
            v-model="editForm.enable"
            :active-value="1"
            :inactive-value="0"
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
import { getBotList, saveBot } from '@/api/bot'
export default {
  components: {
  },
  data() {
    return {
      bots: {
        list: [],
        page: 1,
        size: 20,
        totalnum: 0,
        loading: false
      },
      editDialogVisible: false,
      editForm: {
        id: 0
      },
      editFormRules: {
        name: [
          { required: true, message: '请输入Name', trigger: 'blur' }
        ],
        token: [
          { required: true, message: '请输入Token', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    this.loadBotList()
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
    async loadBotList() {
      this.bots.loading = true
      const data = await getBotList({
        page: this.bots.page,
        size: this.bots.size
      })
      console.log(data)
      this.bots.loading = false
      this.bots.list = data.bots
      this.bots.totalnum = data.total
    },
    showEditForm(row) {
      this.editForm = row
      this.editDialogVisible = true
    },
    doEditFormSubmit() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          saveBot(this.editForm).then(() => {
            this.$message({
              message: (this.editForm.id) ? '编辑成功' : '添加成功',
              type: 'success'
            })
            this.editDialogVisible = false
            this.loadBotList()
          }).catch(() => {
            this.$message({
              message: '操作失败',
              type: 'error'
            })
          })
        }
      })
    },
    changeBotInfo(row) {
      console.log(row)
      this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      saveBot(row).then(_res => {
        this.$message({
          message: '操作成功',
          type: 'success'
        })
      }).catch(_err => {
        this.$message({
          message: '操作失败',
          type: 'error'
        })
      }).finally(() => {
        this.$loading().close()
      })
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

