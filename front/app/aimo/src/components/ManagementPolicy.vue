<template>
  <div>
    <div class="managementPolicy">
      <h3>中期経営方針</h3>
      <br/>
      <!-- <p>{{ data[0].mid_term_policy }}</p> -->
      <textarea :value="mid_term_policy" readonly disabled /><br/>
      <br>
      <br/>
      <br/>
    </div>
    <br/>
    <div class="managementPolicy">
      <h3>{{this.financialYear_YY}}経営方針</h3>
      <br/>
      <textarea :value="period_policy" readonly disabled /><br/>
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
      id:"",
      financialYear_YY:"",
      mid_term_policy:"",
      period_policy:""
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
      this.id = policy.result.policy.id;
      const p = String(policy.result.policy.period);
      this.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      this.mid_term_policy = policy.result.policy.midTermPolicy;
      this.period_policy = policy.result.policy.periodPolicy;
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