<template>
  <div>
    <div class="managementPolicy">
      <h3>部署目標</h3>
      <br/>
      <p>{{ data[0].department_goal }}</p><br/>
      <br>
      <br/>
      <br/>
    </div>
  </div>
</template>

<script>
import {getDepartmentGoal} from '@/api/AimSettingSheet.js'

export default {
  data(){
    return {
      data: []
      // これ、複数データ返ってくる訳じゃないから配列にする必要ないんだけどねぇ。
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
      console.log(goal.result.department_goal.department);
      let d = {};
      d.id = goal.result.department_goal.department_goal_id;
      const p = String(goal.result.department_goal.period);
      d.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      d.department = goal.result.department_goal.department;
      d.department_goal = goal.result.department_goal.department_goal;
      this.data.push(d);
      console.log(this.data);
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
</style>