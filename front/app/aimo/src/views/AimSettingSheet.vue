<template>
  <div class="setting_user">
    <h1>個人目標設定シート</h1>
    <br/>
    <ManagementPolicy />
    <br/>
    <router-link to="/home">
      <button>Back</button>
    </router-link>
    <br/>
    <br/>
    <router-view />
  </div>
</template>

<script>
  import ManagementPolicy from '../components/ManagementPolicy.vue'

  export default {
  //   name: 'Home',

     components: {
       ManagementPolicy
     }
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  }
</script>
