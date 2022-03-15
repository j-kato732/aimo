<template>
  <div class="setting_user">
    <HeaderName title="個人目標設定シート"></HeaderName>
    <br>
    <input class="input_field" type="text" v-model="last_name1" readonly disabled="disabled">
    <input class="input_field" type="text" v-model="first_name1" readonly disabled="disabled">
    <br>
    職位
    <select name="job" v-model="job_id1" readonly disabled="disabled">
      <option hidden>選択してください</option>
      <option value=1 selected>ジュニア</option>
      <option value=2>サブリーダー</option>
      <option value=3>リーダー</option>
      <option value=4>マネージャー</option>
      <option value=5>スペシャリスト</option>
      <option value=6>シニアスペシャリスト</option>
    </select>
    部署
    <select name="department" v-model="department_id1" readonly disabled="disabled">
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
    <br>
    <br>
    <ManagementPolicy />
    <br>
    <DepartmentGoal />
    <br>
    <Role_of_Person />
    <br>
    <br>
    <!-- タブ切り替え -->
    <TabItem />
    <br>
    <router-link to="/mypage">
      <button>保存する</button>
    </router-link>
    <br>
    <br>
    <router-view />
  </div>
</template>

<script>
  import ManagementPolicy from '../components/ManagementPolicy.vue'
  import DepartmentGoal from '../components/DepartmentGoal.vue'
  import Role_of_Person from '../components/Role_of_Person.vue'
  import TabItem from '../components/TabItem.vue'
  import HeaderName from '../components/HeaderName.vue'
  import { getUser } from "@/api/AimSettingSheet.js";

  export default {
    components: {
      ManagementPolicy,
      DepartmentGoal,
      Role_of_Person,
      TabItem,
      HeaderName,
    },
    data() {
      return {
        last_name1: "",
        first_name1: "",
        department_id1: 0,
        job_id1: 0
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

        const user = await getUser(this.$auth.user.email, this.$store.state.period, access_token)
        //this.id = user.result.user.id
        // console.log(user.result)
        
        //データ反映
        this.last_name1 = user.result.user.lastName
        this.first_name1 = user.result.user.firstName
        this.department_id1 = user.result.user.departmentId
        this.job_id1 = user.result.user.jobId
      },
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

  input {
    width: 50px
  }

  select {
    width: 150px
  }
</style>