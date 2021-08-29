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
    <AimSetting_Enter />

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
  import AimSetting_Enter from '../components/AimSetting_Enter.vue'

  export default {
  //   name: 'Home',
    components: {
      ManagementPolicy,
      DepartmentGoal,
      Role_of_Person,
      AimSetting_Enter
    },
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  }
</script>
