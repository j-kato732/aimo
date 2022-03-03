<template>
  <div class="mypage">
    <div class="header">
      <div class="left">
        <HeaderLeftButton />
      </div>
      <div class="title">
        <HeaderName title="MY PAGE"></HeaderName>
      </div>
      <div class="right">
        <HeaderRightButton />
      </div>
    </div>
    <br />
    <Announce_board />
    <br />
    <div v-for="d in data" :key="d.id">
      <router-link to="/aimSettingSheet">
        <button>{{ d.financialYear_YY }} {{ d.mm }}</button>
      </router-link>
    </div>
    <br />
    <br />
    <br />
    <router-link v-if="$auth.isAuthenticated" to="/setting_user">
      <button>設定</button>
    </router-link>
    <br />
    <br />
    <router-link to="/">
      <button>Hello World（ログアウト）</button>
    </router-link>
    <div class="home">
      <div v-if="!$auth.loading">
        <button v-if="!$auth.isAuthenticated" @click="login">Log in</button>
        <button v-if="$auth.isAuthenticated" @click="logout">Log out</button>
      </div>
      <div>
        <h1>以下に取得したユーザ情報を記載</h1>
        <pre>{{ JSON.stringify($auth.user, null, 2) }}</pre>
      </div>
    </div>
    <br />
    <router-view />
  </div>
</template>

<script>
import Announce_board from "../components/Announce_board.vue";
import HeaderName from "../components/HeaderName.vue";
import HeaderLeftButton from "../components/HeaderLeftBotton.vue";
import HeaderRightButton from "../components/HeaderRightButton.vue";
import { getApi } from "@/api/MyPage"; //axiosでAPI取得する処理をMyPage.jsに切り出し
import { getUser } from "@/api/AimSettingSheet.js";
import { mapGetters } from 'vuex'

export default {
  components: {
    Announce_board,
    HeaderName,
    HeaderLeftButton,
    HeaderRightButton,
  },
  data() {
    return {
      data: [],
      year: "",
      month: "",
      YYYYMM: "",
    };
  },
  computed:{
    ...mapGetters(['setUserId'])
  },
  created() {
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async () => {
      await this.getAPI();
    })();
  },
  methods: {
    //ページ遷移時に呼び出す
    async getAPI() {
      // MyPage.jsでaxiosから取得したAPIを使う
      // asyncとawaitはセット
      const access_token = await this.$auth.getTokenSilently();
      const subDatas = await getApi(access_token);
      // 自分の使いたい形にデータを入れ直し（period（YYYYMM）をいろんな形に）
      for (let data of subDatas.result.periods) {
        let d = {};
        d.id = data.id;
        // data.periodをStringに変換して以下の切り出しを可能にした
        const str = String(data.period);
        d.financialYear = str.replace(/^(\d{4})/, "$1"); //2021
        d.financialYear_int = parseInt(data.period.replace(/^\d{2}(\d{2})\d{2}/, "$1")); //21
        d.period = str.replace(/^\d{4}(\d{2})/, "$1"); //05
        if ( d.period == 11 ){
          d.financialYear_int += 1
        }
        d.financialYear_YY = "第" + String(d.financialYear_int) + "期"; //21
        d.mm = this.convertMMToHalfYear(d.period); //convertMMToHalfYear関数で上期or下期判別＆代入
        this.data.push(d);
        //console.log(this.data);
      }

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
      console.log(this.YYYYMM)
      const user = await getUser(this.$auth.user.email, this.YYYYMM, access_token)
      
      //const user = "amnos.39ra38@gmail.com"
      console.log(user)
      if(this.user == null){
        this.$router.push('/registration')
      }
      //console.log(user.id)
      //this.user.idをvuexに格納する
      this.$store.commit('setUserId', user.id)
      //this.$store.commit('setUserId', user)
      console.log(this.$store.state.userId)
    },
    convertMMToHalfYear(period) {
      if (parseInt(period) === 5) {
        return "下期";
      } else if (parseInt(period) === 11) {
        return "上期";
      }
    },
    login() {
      this.$auth.loginWithRedirect();
    },
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin,
      });
    },
  },
};
</script>

<style scoped>
.header {
  display: flex;
  justify-content: center;
}
.left {
  flex: 1;
}
.title {
  flex: 11;
}
.right {
  flex: 1;
}
</style>
