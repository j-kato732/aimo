<template>
  <div class="management">
    <HeaderName title="Management"></HeaderName>
    <br />
    <br />
    <Announce_board />
    <br />
    <div v-for="d in data" :key="d.id">
      <button @click="move(d.period2)">{{ d.financialYear_YY }} {{ d.mm }}</button>
    </div>
    <br />
    <br />
    <router-link to="/">
      <button>Back</button>
    </router-link>
    <br />
    <br />
    <router-view />
  </div>
</template>

<script>
import HeaderName from "../components/HeaderName.vue";
import Announce_board from "../components/Announce_board.vue";
import { getApi } from "@/api/MyPage"; //axiosでAPI取得する処理をMyPage.jsに切り出し

export default {
  //   name: 'Home',

  components: {
    Announce_board,
    HeaderName,
  },
  data() {
    return {
      data: [],
    };
  },
  created() {
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async () => {
      await this.getAPI();
    })();
  },
  methods: {
    async getAPI() {
      // access_tokenの取得
      const access_token = await this.$auth.getTokenSilently();
      // MyPage.jsでaxiosから取得したAPIを使う
      // asyncとawaitはセット
      const subDatas = await getApi(access_token);
      // 自分の使いたい形にデータを入れ直し（period（YYYYMM）をいろんな形に）
      for (let data of subDatas.result.periods) {
        let d = {};
        d.id = data.id;
        d.period2 = data.period; //YYYYMM
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
      }
    },
    convertMMToHalfYear(period) {
      if (parseInt(period) === 5) {
        return "下期";
      } else if (parseInt(period) === 11) {
        return "上期";
      }
    },
    move(period){
      this.$router.push({
        name: 'PeriodSeparatePage',
        params: {
          period: period,
        },
      });
    },
  },
};
</script>

<style scoped>
button {
  width: 160px;
  height: 44px;
  background: rgba(65, 100, 255, 0.3);
  border: 6px solid #416dff;
  border-radius: 90px;
}

</style>
