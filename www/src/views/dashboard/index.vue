<template>
  <el-container>
    <el-main>
      <el-divider />
      <el-header class="pagetab">
        <h4>机器人管理</h4>
      </el-header>
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
          prop="retained_day_1"
          label="是否开启"
          align="center"
        >
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.enable"
              :active-value="1"
              :inactive-value="0"
            />
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
  </el-container>
</template>

<script>
import { getBotList } from '@/api/bot'
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

