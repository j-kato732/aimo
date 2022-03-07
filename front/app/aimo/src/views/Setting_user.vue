<template>
  <div class="setting_user">
    <HeaderName title="Setting"></HeaderName>
    <br/>
    <div id="background">
      <br>
      <h3>プロフィールの変更</h3>
      <br/>
      氏<input class="input_field" type="text" v-model="last_name">
      名<input class="input_field" type="text" v-model="first_name">
      <br>
      <br>
      職位
      <select name="job" v-model="job_id">
        <option hidden>選択してください</option>
        <option value=1 selected>ジュニア</option>
        <option value=2>サブリーダー</option>
        <option value=3>リーダー</option>
        <option value=4>マネージャー</option>
        <option value=5>スペシャリスト</option>
        <option value=6>シニアスペシャリスト</option>
      </select>
      <br>
      <br>
      部署
      <select name="department" v-model="department_id">
        <option hidden>選択してください</option>
        <option value=1 selected>ソリューション本部</option>
        <option value=2>札幌開発センター</option>
        <option value=3>エイジア開発事業本部</option>
        <option value=4>営業本部</option>
        <option value=5>事業推進本部</option>
        <option value=6>管理本部</option>
        <option value=7>情報システム室</option>
        <option value=8>モバイルライフジャパン</option>
      </select>
      <select name="department_detail" hidden>
        <option value="no1" selected>開発１部</option>
        <option value="no2">開発２部</option>
        <option value="no3">開発３部</option>
        <option value="no4">開発４部</option>
      </select>
      <br>
      <br>
    </div>
    <br>
    <button @click="putUser">Submit</button>
    <router-link to="/mypage">
      <button>Back</button>
    </router-link>
    <br/>
    <br/>
    <router-view />
  </div>
</template>

<script>
  import HeaderName from '../components/HeaderName.vue'
  import { getUser, putUser } from "@/api/AimSettingSheet.js";

  export default {
  //   name: 'Home',

    components: {
      HeaderName
    },
    data() {
      return {
        year: "",
        month: "",
        YYYYMM: "",
        last_name: "",
        first_name: "",
        department_id: 0,
        job_id: 0,
        id: 0,
        auth_id: "",
        enrollment_flg: true,
        admin_flg: false,
      }
    },
    created() {
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async () => {
      await this.getUser();
    })();
    },
    methods:{
      async getUser(){
        const access_token = await this.$auth.getTokenSilently();

        //今日の日付をとってくる
        const current_date = new Date()
        //今日の月から現在は上期か下期かを判別
        const current_month = current_date.getMonth() + 1
        if( 5 <= current_month && current_month <= 10 ){
          //5~10月は05（下期）
          this.year = String(current_date.getFullYear())
          this.month = "05"
        } else if( current_month == 11 || current_month == 12 || 1 <= current_month || current_month <= 4 ){
          //11~4月は11（上期）
          //11~12月はyear+1してあげる
          if( current_month == 11 || current_month == 12 ){
            this.year = String(current_date.getFullYear() + 1)
          } else if ( 1 <= current_month || current_month <= 4 ){
            this.year = String(current_date.getFullYear())
          }
          this.month = "11"
        }
        this.YYYYMM = this.year + this.month

        const user = await getUser(this.$auth.user.email, this.YYYYMM, access_token)
        this.id = user.result.user.id
        this.enrollment_flg = user.result.user.enrollmentFlg
        this.admin_flg = user.result.user.adminFlg
        console.log(user.result)
        
        //データ反映
        this.last_name = user.result.user.lastName
        this.first_name = user.result.user.firstName
        this.department_id = user.result.user.departmentId
        this.job_id = user.result.user.jobId
      },
      async putUser() {
        const access_token = await this.$auth.getTokenSilently();
        await putUser(
          this.id,
          this.auth_id,
          this.YYYYMM,
          this.last_name,
          this.first_name,
          this.department_id,
          this.job_id,
          this.enrollment_flg,
          this.admin_flg,
          access_token
        )
        this.$router.push('/mypage')
      },
    }
  }
</script>

<style scoped>
input, select {
  border: 1px solid;
  margin: 0 10px;
}

#background {
  background: rgba(255, 179, 65, 0.3);
  width: 60%;
  margin: auto;
}

button {
  width: 160px;
  height: 44px;
  background: rgba(255, 179, 65, 0.3);
  border: 6px solid #FFB341;
  border-radius: 90px;
}
</style>
