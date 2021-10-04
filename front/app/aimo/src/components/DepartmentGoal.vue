<template>
  <div>
    <div class="managementPolicy">
      <h3>部署目標</h3>
      <br/>
      <!-- <p>{{ department_goal }}</p><br/> -->
      <!-- <v-textarea :value="department_goal" /> -->
      <textarea v-model="department_goal" readonly disabled="disabled"/><br/>
      <br>
      <br/>
      <br/>
    </div>
  </div>
</template>

<script>
import {getDepartmentGoal} from '@/api/AimSettingSheet.js'
import {replaceIndention} from '@/utils/StringUtil.js'

export default {
  data(){
    return {
      id:"",
      financialYear_YY:"",
      department_id:"",
      department_goal:""
    }
  },
  created(){
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async()=>{
      await this.getDepartmentGoal();
      })();
  },
  methods:{
    async getDepartmentGoal(){
      const goal = await getDepartmentGoal();
      console.log(goal.result.departmentGoal.departmentGoal);
      this.id = goal.result.departmentGoal.id;
      const p = String(goal.result.departmentGoal.period);
      this.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      this.department_id = goal.result.departmentGoal.departmentId;
      this.department_goal = replaceIndention(goal.result.departmentGoal.departmentGoal);
    }
  }
}
</script>

<style scoped>
  .managementPolicy {
    width: 60%;
    height: 140px;
    background: #F7F7F7;
    margin: 0 20%;
  }

  textarea {
    resize: none;
    width: 100%;
    height: 140px;
    text-align: center;
  }
</style>