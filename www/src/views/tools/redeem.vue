<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">兑换码</h4>
    </el-header>
    <el-main>
      <el-row type="flex" justify="end">
        <el-button type="success" icon="el-icon-s-ticket" @click="showCodeModal">生成兑换码</el-button>
      </el-row>
      <el-divider />
      <el-table
        ref="multipleTable"
        v-loading="logs_table.loading"
        :data="logs_table.list"
        tooltip-effect="dark"
        style="width: 100%"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="160"
        />
        <el-table-column
          prop="code"
          label="兑换码"
        />
        <el-table-column
          prop="diamond"
          label="钻石奖励"
        />
        <el-table-column
          prop="flag"
          label="是否使用"
        >
          <template slot-scope="scope">
            <span> {{ scope.row.flag > 0 ? "已兑换":"未使用" }} </span>
          </template>
        </el-table-column>
        <el-table-column
          prop="create_at"
          label="创建时间"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.create_at | utcTimeFormat }} </span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="160"
        >
          <template slot-scope="scope">
            <el-button type="primary" size="mini" @click="()=>{copyRedeemCodeLink(scope.row)}">复制链接</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :current-page="logs_table.page"
        :page-size="logs_table.pagesize"
        :hide-on-single-page="false"
        :total="logs_table.totalnum"
        @current-change="LogsPageChange"
      />
    </el-main>
    <el-dialog
      title="生成兑换码"
      :visible.sync="codeDialogVisible"
      width="50%"
    >
      <el-form class="generateForm" label-width="120px">
        <el-form-item label="钻石奖励">
          <el-input v-model.number="generateForm.diamond" type="number" placeholder="钻石数量" />
        </el-form-item>
        <el-form-item label="生成数量">
          <el-input v-model.number="generateForm.count" type="number" placeholder="生成数量" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="generateRedeemCode">确定</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </el-container>
</template>

<script>
import { makeAjaxParamData } from '@/utils/util'
import { redeemList, redeemAdd } from '@/api/redeem'

export default {
  components: {
  },
  data() {
    return {
      redeemEditDialogVisible: false,
      redeemDialogVisible: false,
      codeDialogVisible: false,
      searchForm: {},
      redeemForm: {},
      generateForm: {},
      logs_table: {
        loading: false,
        list: [],
        pagesize: 20,
        totalnum: 0,
        page: 1
      }
    }
  },
  computed: {
  },
  async mounted() {
    this.getRedeemLogs()
  },
  methods: {
    showRedeemModal() {
      this.redeemDialogVisible = true
      this.getRedeemList()
    },
    showRedeemEditModal(row) {
      this.redeemEditDialogVisible = true
      this.redeemForm = {
        id: row ? row.id : 0,
        name: row ? row.name : '',
        diamonds: row ? row.diamonds : 0,
        energy: row ? row.energy : 0
      }
    },
    showCodeModal() {
      this.codeDialogVisible = true
    },
    async getRedeemLogs() {
      this.logs_table.loading = true
      var params = makeAjaxParamData(this.searchForm, {
        page: this.logs_table.page,
        size: this.logs_table.pagesize
      })
      const res = await redeemList(params)
      this.logs_table.list = res.list
      this.logs_table.totalnum = res.total
      this.logs_table.loading = false
    },
    async generateRedeemCode() {
      await redeemAdd(this.generateForm)
      this.$message.success('生成成功')
      this.codeDialogVisible = false
      this.getRedeemLogs()
    },
    LogsPageChange(e) {
      this.logs_table.page = e
      this.getRedeemLogs()
    },
    redeemPageChange(e) {
      this.redeem_table.page = e
      this.getRedeemList()
    },
    onRedeemChange() {
      var redeem = this.redeem_table.list.find(item => item.id === this.generateForm.id)
      this.generateForm.name = redeem.name
      this.generateForm.diamonds = redeem.diamonds
      this.generateForm.energy = redeem.energy
    },
    copyRedeemCodeLink(row) {
      console.log(row.link)
      navigator.clipboard.writeText(row.link).then(() => {
        this.$message.success('复制成功')
      }, () => {
        this.$message.error('复制失败')
      })
    },
    doSearch() {
      this.logs_table.page = 1
      this.getRedeemLogs()
    },
    doReset() {
      this.searchForm.code = ''
      this.searchForm.id = null
      this.searchForm.used = -1
      this.logs_table.page = 1
      this.getRedeemLogs()
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
