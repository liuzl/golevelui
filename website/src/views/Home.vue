<template>
  <div class="home">
    <el-tabs
      ref="tabs"
      v-model="showTab"
      tab-position="left"
      class="db-tabs"
      @tab-click="handleClick"
    >
      <el-tab-pane
        v-for="(db, index) in dbs"
        :key="index"
        :label="db"
        :name="db"
      >
        <key-list
          v-if="showTab === db"
          :db="db"
          :offset-top.sync="tabsOffsetTop"
        ></key-list>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import KeyList from '@/components/KeyList.vue'
import { dbs } from '@/api/golevelui'
import { Tabs, TabPane } from 'element-ui'

Vue.use(Tabs)
Vue.use(TabPane)

@Component({
  components: {
    KeyList,
  },
})
export default class Home extends Vue {
  private showTab = ''
  private dbs: Array<string> = []
  private tabsOffsetTop = 0

  mounted() {
    this.$nextTick(() => {
      this.tabsOffsetTop = (
        (this.$refs.tabs as Vue).$el as HTMLElement
      ).offsetTop
    })
  }

  created() {
    dbs()
      .then((res) => {
        this.dbs = res.data
        if (res.data.length) {
          this.showTab = res.data[0]
        }
      })
      .catch((error) => {
        console.error('Failed to load databases:', error)
      })
  }

  handleClick(tab: { name: string }) {
    // Handle tab click if needed in the future
    console.log('Switched to tab:', tab.name)
  }
}
</script>

<style scoped>
.home {
  padding: 20px;
  height: calc(100vh - 40px);
  box-sizing: border-box;
}
.db-tabs {
  height: 100%;
}
::v-deep .el-tabs__content {
  height: 100%;
}
::v-deep .el-tab-pane {
  height: 100%;
}
</style>
