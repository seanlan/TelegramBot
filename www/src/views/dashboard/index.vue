<template>
  <el-container>
    <el-main>
      <el-header class="pagetab">
        <h4 class="links">数据概览</h4>
      </el-header>
      <el-row>
        <el-col :span="6">
          <div class="card-panel">
            <div class="card-panel-content">
              <h1>{{ today.register }}</h1>
              <span>新增用户</span>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="card-panel">
            <div class="card-panel-content">
              <h1>{{ today.online }}</h1>
              <span>活跃用户</span>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="card-panel">
            <div class="card-panel-content">
              <h1>{{ today.retained_day_1 }} / {{ today.retained_rate_1 | percent }}%</h1>
              <span>次留</span>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="card-panel">
            <div class="card-panel-content">
              <h1>{{ today.invite }} / {{ today.inviter }}</h1>
              <span>邀请</span>
            </div>
          </div>
        </el-col>
      </el-row>
      <el-divider />
      <el-header class="pagetab">
        <h4>每日统计</h4>
      </el-header>
      <el-table
        v-loading="stats.loading"
        :data="stats.list"
      >
        <el-table-column
          prop="date_stamp"
          label="日期"
          width="100"
        />
        <el-table-column
          prop="register"
          label="注册"
          align="center"
        />
        <el-table-column
          prop="online"
          label="活跃"
          align="center"
        />
        <el-table-column
          prop="retained_day_1"
          label="次留"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag v-if="scope.row.retained_day_1 > 0">
              {{ scope.row.retained_day_1 }} ({{ scope.row.retained_rate_1 | percent }}%)
            </el-tag>
            <span v-else> - </span>
          </template>
        </el-table-column>
        <el-table-column
          prop="retained_day_1"
          label="3日留存"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag v-if="scope.row.retained_day_3 > 0">
              {{ scope.row.retained_day_3 }} ({{ scope.row.retained_rate_3 | percent }}%)
            </el-tag>
            <span v-else> - </span>
          </template>
        </el-table-column>
        <el-table-column
          prop="retained_day_7"
          label="7日留存"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag v-if="scope.row.retained_day_7 > 0">
              {{ scope.row.retained_day_7 }} ({{ scope.row.retained_rate_7 | percent }}%)
            </el-tag>
            <span v-else> - </span>
          </template>
        </el-table-column>
        <el-table-column
          prop="invite"
          label="邀请人数/邀请人"
          align="center"
          width="200"
        >
          <template slot-scope="scope">
            {{ scope.row.invite }} / {{ scope.row.inviter }}
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, jumper, ->, total"
        :page-size="stats.pagesize"
        :total="stats.totalnum"
        @current-change="statsPageChange"
      />
    </el-main>
  </el-container>
</template>

<script>
import { statsList, statsToday } from '@/api/stats'
export default {
  components: {
  },
  data() {
    return {
      today: {
        register: 0,
        online: 0,
        retained_day_1: 0,
        retained_rate_1: 0,
        invite: 0,
        inviter: 0
      },
      stats: {
        list: [],
        pagesize: 20,
        totalnum: 0,
        loading: false
      }
    }
  },
  mounted() {
    this.getStatsToday()
    this.getStatsList()
  },
  methods: {
    async getStatsToday(pagenum = 1) {
      const res = await statsToday()
      this.today = res.record
    },
    async getStatsList(pagenum = 1) {
      this.stats.loading = true
      const res = await statsList({
        page: pagenum,
        size: this.stats.pagesize
      })
      this.stats.list = res.list.map(item => {
        item.assets = []
        return item
      })
      console.log(res.list)
      this.stats.totalnum = res.total
      this.stats.loading = false
    },
    statsPageChange(e) {
      this.getStatsList(e)
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

