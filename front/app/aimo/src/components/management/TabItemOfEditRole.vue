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
            <editRole :tab="item.tab"/>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import EditRole from '@/components/management/EditRole.vue'

export default {
    components: {
      EditRole
    },
    data () {
      return {
        tab: null,
        items: [
          { tab: 'Jr' },
          { tab: 'SL' },
          { tab: 'L' },
          { tab: 'M' },
          { tab: 'SP' },
          { tab: 'SSP' },
        ],
      }
    },
  }
</script>

<style scoped>
.tab {
  width: 600px;
  margin: 0 auto;
}

.v-tab {
  margin: 0 auto;
  width: 20%;
}

.v-tab--active {
  color: orange;
}
</style>