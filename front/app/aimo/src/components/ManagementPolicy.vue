<template>
  <div>
    <div class="managementPolicy">
      <h3>中期経営方針</h3>
      <br/>
      <p>{{ data[0].mid_term_policy }}</p><br/>
      <br>
      <br/>
      <br/>
    </div>
    <br/>
    <div class="managementPolicy">
      <h3>{{ data[0].financialYear_YY }}経営方針</h3>
      <br/>
      <p>{{ data[0].period_policy }}</p><br/>
      <br>
      <br/>
      <br/>
    </div>
  </div>
</template>

<script>
import {getPolicy} from '@/api/AimSettingSheet.js'

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
      await this.getPolicy();
      })();
  },
  methods:{
    async getPolicy(){
      const policy = await getPolicy();
      console.log(policy.result.policy.mid_term_policy);
      let d = {};
      d.id = policy.result.policy.policy_id;
      const p = String(policy.result.policy.period);
      d.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      d.mid_term_policy = policy.result.policy.mid_term_policy;
      d.period_policy = policy.result.policy.period_policy;
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