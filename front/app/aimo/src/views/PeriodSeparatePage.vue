<template>
  <div class="periodSeparatePage">
    <HeaderName title="Management"></HeaderName>
    <br/>
    <h3>21期上期</h3>
    <br/>
    目標設定期間  <input type='date' /> 〜 <input type='date' />
    <br/>
    <br/>
    自己評価期間  <input type='date' /> 〜 <input type='date' />
    <br/>
    <br/>
    <!-- タブ切り替え -->
    <ManagementTab />
    <br/>
    <br/>
    <router-link to="/management">
      <button>Back</button>
    </router-link>
    <br/>
    <br/>
    <router-view />
  </div>
</template>

<script>
  import HeaderName from '../components/HeaderName.vue'
  import ManagementTab from '../components/management/ManagementTab.vue'

  export default {
  //   name: 'Home',

    components: {
      HeaderName,
      ManagementTab
    },
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  //   }
  }
</script>

<style scoped>
  input {
    border: 1px solid;
    width: 140px;
  }
</style>
