<template>
  <div class="setting_user">
    <h1>個人目標設定シート</h1>
    <br/>
    <ManagementPolicy />
    <br/>
    <DepartmentGoal />
    <br/>
    <Role_of_Person />
    <br/>

    <p>ここにタブ切り替え作りたい</p>

    <br><br>
    <TabItem />
    <!-- <AimSetting_Enter /> -->

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
  import DepartmentGoal from '../components/DepartmentGoal.vue'
  import Role_of_Person from '../components/Role_of_Person.vue'
  // import AimSetting_Enter from '../components/AimSetting_Enter.vue'
  import TabItem from '../components/TabItem.vue'

  export default {
  //   name: 'Home',
    components: {
      ManagementPolicy,
      DepartmentGoal,
      Role_of_Person,
      // AimSetting_Enter,
      TabItem
    }
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  }
</script>

<style scoped>
  .contents {
    position: relative;
    overflow: hidden;
    width: 280px;
    border: 2px solid #000;
  }
  .item {
    box-sizing: border-box;
    padding: 10px;
    width: 100%;
    transition: all 0.8s ease;
  }
</style>