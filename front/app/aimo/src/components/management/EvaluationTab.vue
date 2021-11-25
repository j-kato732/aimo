<template>
  <div class="setting_user">
    <HeaderName title="個人目標設定シート"></HeaderName>
    <br/>
    <ManagementPolicy />
    <br/>
    <DepartmentGoal />
    <br/>
    <Role_of_Person />
    <br/>
    <select name="department">
      <option hidden>部を選択 ▼</option>
      <option value="solution">ソリューション本部</option>
      <option value="asia">エイジア開発事業本部</option>
      <option value="sapporo">札幌開発センター</option>
      <option value="sales">営業本部</option>
      <option value="business">事業推進本部</option>
      <option value="management">管理本部</option>
    </select>
    <select name="person">
      <option hidden>人を選択 ▼</option>
      <option value="person1">人１</option>
      <option value="person2">人２</option>
      <option value="person3">人３</option>
    </select>
    <br>
    <br>
    <AimSetting_Enter />
  </div>
</template>

<script>
  import ManagementPolicy from '../ManagementPolicy.vue'
  import DepartmentGoal from '../DepartmentGoal.vue'
  import Role_of_Person from '../Role_of_Person.vue'
  import AimSetting_Enter from '../AimSetting_Enter.vue'

  export default {
    components: {
      ManagementPolicy,
      DepartmentGoal,
      Role_of_Person,
      AimSetting_Enter
    }
  }
</script>

<style scoped>
  select {
    margin: 0 10px;
  }

  p.pp {
    padding: 8px;
  }

  textarea.whatAndWhere {
    resize: none;
    height: 60px;
    width: 300px;
    margin: 0 10px;
    background: #f7f7f7;
  }
</style>