<template>
  <div>
    <div class="roleOfPerson">
      <h3>本人の役割</h3>
      <br/>
      <textarea :value="role" readonly disabled /><br/>
      <br>
      <br/>
      <br/>
    </div>
  </div>
</template>

<script>
import {getUser, getRole} from '@/api/AimSettingSheet.js'
import {replaceIndention} from '@/utils/StringUtil.js'

export default {
  data(){
    return {
      id:"",
      financialYear_YY:"",
      department_id:"",
      job_id:"",
      role:""
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
      const access_token = await this.$auth.getTokenSilently();
      const user = await getUser(this.$auth.user.email, this.$route.params.period, access_token)
      const role = await getRole(this.$route.params.period, user.result.user.departmentId, user.result.user.jobId, access_token);
      this.id = role.result.role.id;
      const p = String(role.result.role.period);
      this.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
      this.department_id = role.result.role.departmentId;
      this.job_id = role.result.role.jobId;
      this.role = replaceIndention(role.result.role.role);
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

  textarea {
    resize: none;
    width: 100%;
    height: 140px;
    text-align: center;
  }
</style>