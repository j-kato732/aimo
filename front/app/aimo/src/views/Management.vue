<template>
  <div class="management">
    <HeaderName title="Management"></HeaderName>
    <br/>
    <br/>
    <Announce_board />
    <br/>
    <div v-for="d in data" :key="d.id">
      <router-link to="/aimSettingSheet">
        <button>{{ d.financialYear_YY }} {{ d.mm }}</button>
      </router-link>
    </div>
    <br>
    <router-link to="/periodSeparatePage">
      <button>期別ページ</button>
    </router-link>
    <br>
    <br>
    <router-link to="/">
      <button>Back</button>
    </router-link>
    <br/>
    <br/>
    <router-view />
  </div>
</template>

<script>
  import HeaderName from '../components/HeaderName.vue'
  import Announce_board from '../components/Announce_board.vue'
  import {getApi} from '@/api/MyPage' //axiosでAPI取得する処理をMyPage.jsに切り出し

  export default {
  //   name: 'Home',

    components: {
      Announce_board,
      HeaderName
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

