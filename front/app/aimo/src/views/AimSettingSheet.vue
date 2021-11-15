<template>
  <div class="setting_user">
    <HeaderName title="個人目標設定シート"></HeaderName>
    <br/>
    <ManagementPolicy />
    <br/>
    <DepartmentGoal />
    <br/>
    <Role_of_Person />
    <br>
    <br>
    <!-- タブ切り替え -->
    <TabItem />
    <br/>
    <router-link to="/mypage">
      <button>保存する</button>
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
  import TabItem from '../components/TabItem.vue'
  import HeaderName from '../components/HeaderName.vue'

  export default {
    components: {
      ManagementPolicy,
      DepartmentGoal,
      Role_of_Person,
      TabItem,
      HeaderName,
    }
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