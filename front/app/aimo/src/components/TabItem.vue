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
    </v-tabs>

    <v-tabs-items v-model="tab">
      <v-tab-item
        v-for="item in items"
        :key="item.tab"
      >
        <v-card flat>
          <v-card-text class="tabItem">
            <AimSetting_Enter :tab="item.tab"/>
          </v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import AimSetting_Enter from './AimSetting_Enter.vue'

export default {
    components: {
      AimSetting_Enter
    },
    data () {
      return {
        tab: null,
        items: [
          { tab: '1' },
          { tab: '2' },
          { tab: '3' },
          { tab: '4' },
          { tab: '5' },
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
  background: rgba(255, 179, 65, 0.3);
}
</style>