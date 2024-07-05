<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">红包管理</h4>
    </el-header>
    <el-main>
      <el-row type="flex" justify="end">
        <el-button type="primary" icon="el-icon-plus" @click="()=>showEditModal()">发放红包</el-button>
      </el-row>
      <el-divider />
      <el-table
        ref="multipleTable"
        v-loading="packet_table.loading"
        :data="packet_table.list"
        tooltip-effect="dark"
        style="width: 100%"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="160"
        />
        <el-table-column
          prop="title"
          label="标题"
          width="160"
        />
        <el-table-column
          prop="content"
          label="内容"
        />
        <el-table-column
          prop="diamonds"
          label="钻石"
        />
        <el-table-column
          prop="number"
          label="红包数量"
        />
        <el-table-column
          prop="used"
          label="已领取"
        />
        <el-table-column
          prop="code"
          label="红包码"
          width="300"
        />
        <el-table-column
          prop="create_at"
          label="创建时间"
          width="160"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.create_at | utcTimeFormat }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="160"
        >
          <template slot-scope="scope">
            <el-button type="primary" size="mini" @click="()=>{copyLink(scope.row)}">复制链接</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :current-page="packet_table.page"
        :page-size="packet_table.pagesize"
        :hide-on-single-page="false"
        :total="packet_table.totalnum"
        @current-change="PageChange"
      />
    </el-main>
    <el-dialog
      :title="editForm.packet_id > 0? '修改红包': '发放红包'"
      :visible.sync="editDialogVisible"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
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
        <el-form-item label="钻石数量" prop="diamonds">
          <el-input
            ref="url"
            v-model.number="editForm.diamonds"
            placeholder="钻石数量"
          />
        </el-form-item>
        <el-form-item label="红包数量" prop="number">
          <el-input
            ref="url"
            v-model.number="editForm.number"
            placeholder="红包数量"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button v-if="editForm.packet_id > 0" type="primary" @click="submitEditForm">编 辑</el-button>
        <el-button v-else type="primary" @click="submitEditForm">添 加</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import { redpacketList, redpacketAdd } from '@/api/repacket'
export default {
  components: {
  },
  data() {
    return {
      packet_table: {
        loading: false,
        list: [],
        pagesize: 100,
        totalnum: 0,
        page: 1
      },
      editDialogVisible: false,
      editForm: {
        packet_id: 0
      },
      editFormRules: {
        title: [{ required: true, trigger: 'blur', message: '标题不能为空' }],
        diamonds: [{ required: true, trigger: 'blur', message: '钻石数量不能为空' }],
        number: [{ required: true, trigger: 'blur', message: '红包数量不能为空' }]
      }
    }
  },
  computed: {
  },
  async mounted() {
    this.loadData()
  },
  methods: {
    async loadData() {
      this.packet_table.loading = true
      var params = {
        page: this.packet_table.page,
        size: this.packet_table.pagesize
      }
      const res = await redpacketList(params)
      this.packet_table.list = res.list || []
      this.packet_table.totalnum = res.total || 0
      this.packet_table.loading = false
    },
    PageChange(e) {
      this.packet_table.page = e
      this.loadData()
    },
    showEditModal(row) {
      this.editDialogVisible = true
      this.editForm = {
        packet_id: row ? row.id : 0,
        title: row ? row.title : 'Try your luck',
        content: row ? row.content : 'Winer Winer Chicken Dinner',
        diamonds: row ? row.diamonds : 0,
        number: row ? row.number : 0
      }
    },
    submitEditForm() {
      this.$refs.editForm.validate(async valid => {
        if (valid) {
          await redpacketAdd({
            packet_id: this.editForm.packet_id,
            title: this.editForm.title,
            content: this.editForm.content,
            diamonds: this.editForm.diamonds,
            number: this.editForm.number
          })
          this.$message.success('提交成功')
          this.editDialogVisible = false
          this.loadData()
          return true
        } else {
          return false
        }
      })
    },
    copyLink(row) {
      console.log(row.link)
      if (row.code === '8ba941f90f2c4e74bb679bd910ac7d12') {
        navigator.clipboard.writeText('🎁 Newcomer Benefits 🎁 \n Click to claim 💎 reward \n\n' + row.link).then(() => {
          this.$message.success('复制成功')
        }, () => {
          this.$message.error('复制失败')
        })
      } else {
        navigator.clipboard.writeText('Click to receive a lucky draw red envelope. \n\n' + row.link).then(() => {
          this.$message.success('复制成功')
        }, () => {
          this.$message.error('复制失败')
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>

