<template>
  <div class="home">
    <div class="header">
      <div class="left">
        <HeaderLeftButton />
      </div>
      <div class="title">
        <HeaderName title="HOME"></HeaderName>
      </div>
      <div class="right">
        <HeaderRightButton />
      </div>
    </div>
    <p>
      HOMEってしたけどここはのちにMY PAGEになる予定
    </p>
    <br/>
    <br/>
    <Announce_board />
    <br/>
    <div v-for="d in data" :key="d.id">
      <router-link to="/aimSettingSheet">
        <button>{{ d.financialYear_YY }} {{ d.mm }}</button>
      </router-link>
    </div>
    <br/>
    <br/>
    <br/>
    <router-link to="/setting_user">
      <button>設定</button>
    </router-link>
    <br/>
    <br>
    <router-link to="/">
      <button>Hello World（ログアウト）</button>
    </router-link>
    <br/>
    <router-view />
  </div>
</template>

<script>
  import Announce_board from '../components/Announce_board.vue'
  import HeaderName from '../components/HeaderName.vue'
  import HeaderLeftButton from '../components/HeaderLeftBotton.vue'
  import HeaderRightButton from '../components/HeaderRightButton.vue'
  import {getApi} from '@/api/MyPage' //axiosでAPI取得する処理をMyPage.jsに切り出し

  export default {
    components: {
      Announce_board,
      HeaderName,
      HeaderLeftButton,
      HeaderRightButton
    },
    data(){
      return {
        data: []
      }
    },
    created(){
      // 下に書いたgetAPI関数をページ遷移時に呼び出す
      (async()=>{
        await this.getAPI();
      })();
    },
    methods:{
      async getAPI(){
        // MyPage.jsでaxiosから取得したAPIを使う
        // asyncとawaitはセット
        const subDatas = await getApi();
        // 自分の使いたい形にデータを入れ直し（period（YYYYMM）をいろんな形に）
        for(let data of subDatas.result.periods){
          let d = {};
          d.id = data.id;
          // data.periodをStringに変換して以下の切り出しを可能にした
          const str = String(data.period);
          d.financialYear = str.replace(/^(\d{4})/, '$1'); //2021
          d.financialYear_YY = str.replace(/^\d{2}(\d{2})\d{2}/, '第$1期'); //21
          d.period = str.replace(/^\d{4}(\d{2})/, '$1'); //05 
          d.mm = this.convertMMToHalfYear(d.period); //convertMMToHalfYear関数で上期or下期判別＆代入
          this.data.push(d);
          console.log(this.data);
        }
      },
      convertMMToHalfYear(period){
        if(parseInt(period) === 5){
          return "上期";
        }else if(parseInt(period) === 11){
          return "下期";
        }
      }
    }
  }
</script>

<style scoped>
  .header {
    display: flex;
    justify-content: center;
  }
  .left { flex: 1; }
  .title {
    flex: 11;
  }
  .right { flex: 1; }

</style>
