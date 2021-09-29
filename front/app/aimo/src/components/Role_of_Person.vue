<template>
  <div>
    <div class="roleOfPerson">
      <h3>本人の役割</h3>
      <br/>
      <p>{{ data[0].role }}</p><br/>
      <br>
      <br/>
      <br/>
    </div>
  </div>
</template>

<script>
import {getRole} from '@/api/AimSettingSheet.js'

export default {
  data(){
    return {
      data: [],
      role: {}
      // これ、複数データ返ってくる訳じゃないから配列にする必要ないんだけどねぇ。
    }
  },
  created(){
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async()=>{
      await this.getRole();
      })();
  },
  methods:{
    async getRole(){
      const role = await getRole();
      console.log(role.result.role.role);
      let d = {};
      d.id = role.result.role.role_id;
      const p = String(role.result.role.period);
      d.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      d.department_id = role.result.role.department_id;
      d.job_id = role.result.role.job_id;
      d.role = role.result.role.role;
      this.data.push(d);
      console.log(this.data);
    }
  }
}
</script>

<style scoped>
  .roleOfPerson {
    width: 60%;
    height: 140px;
    background: #F7F7F7;
    margin: 0 20%;
  }
</style>