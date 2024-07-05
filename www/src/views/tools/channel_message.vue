<template>
  <el-container>
    <el-header class="pagetab">
      <h4 class="links">频道消息</h4>
    </el-header>
    <el-main>
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="140px">
        <el-form-item label="图片" prop="image">
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
          v-for="(ext, index) in editForm.extList"
          :key="index"
        >
          <el-row>
            <el-col :span="10">
              <el-input
                v-model="ext.text"
                placeholder="请输入文本"
                @change="changeExt"
              />
            </el-col>
            <el-col :span="10">
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
        <el-form-item>
          <el-button type="primary" @click="sendMessage">发送</el-button>
          <el-button @click="doReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
// import Needaddshop from '@/components/Needaddshop'
import _ from 'lodash'
import { sendChannelMessage } from '@/api/message'

export default {
  components: {
  },
  data() {
    return {
      editForm: {
        ext: '[]',
        image: '',
        content: '',
        extList: []
      },
      editFormRules: {
        content: [{ required: true, trigger: 'blur', message: 'VALUE不能为空' }]
      }
    }
  },
  computed: {
  },
  async mounted() {
  },
  methods: {
    doReset() {
      this.editForm = {
        ext: '[]',
        image: '',
        content: '',
        extList: []
      }
    },
    addExt() {
      const extList = _.cloneDeep(this.editForm.extList) || []
      extList.push({
        text: '',
        url: ''
      })
      this.editForm.extList = extList
      this.editForm.ext = JSON.stringify(this.editForm.extList)
    },
    removeExt(index) {
      const extList = _.cloneDeep(this.editForm.extList) || []
      extList.splice(index, 1)
      this.editForm.extList = extList
      this.editForm.ext = JSON.stringify(this.editForm.extList)
    },
    changeExt() {
      this.editForm.ext = JSON.stringify(this.editForm.extList)
    },
    sendMessage() {
      this.$refs.editForm.validate(valid => {
        console.log(this.editForm)
        if (valid) {
          const configInfo = _.cloneDeep(this.editForm)
          // 过滤掉extList中的空项
          configInfo.extList = configInfo.extList.filter(item => item.text && item.url)
          configInfo.ext = JSON.stringify(configInfo.extList)
          const loading = this.$loading()
          sendChannelMessage(configInfo).then(res => {
            this.$message({
              message: '发送成功',
              type: 'success'
            })
            this.doReset()
          }).finally(() => {
            loading.close()
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
