<template>
  <v-card class="tab">
    <v-tabs
      v-model="tab"
      background-color="primary"
      dark
    >
      <v-tab
        v-for="item in items"
        :key="item.tab"
      >
        {{ item.tab }}
      </v-tab>

      <v-tab-item>
        <v-card>
          <v-card-text class="tabItem">
            <PolicyTab />
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card>
          <v-card-text class="tabItem">
            <EvaluationTab />
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card>
          <v-card-text class="tabItem">
            <UserListTab />
          </v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-card>
</template>

<script>
import EvaluationTab from '../management/EvaluationTab.vue'
import PolicyTab from '../management/PolicyTab.vue'
import UserListTab from '../management/UserListTab.vue'

export default {
    components: {
      EvaluationTab,
      PolicyTab,
      UserListTab
    },
    data () {
      return {
        tab: null,
        items: [
          { tab: '方針入力' },
          { tab: '評価入力' },
          { tab: 'ユーザー一覧' },
        ],
      }
    },
  }
</script>

<style scoped>
.tab {
  width: 60%;
  margin: 0 auto;
}

.v-tab {
  margin: 0 auto;
  width: 20%;
}

.v-tab--active {
  color: orange;
}

.tabItem {
  background: rgba(53, 122, 255, 0.3);
  padding: 0;
}
</style>