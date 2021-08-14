<template>
  <div class="home">
    <h1>HOME</h1>
    <p>
      HOMEってしたけどここはのちにMY PAGEになる予定
    </p>
    <br/>
    <br/>
    <Announce_board />
    <br/>
    第21期
    <router-link to="/aimSettingSheet">
      <button>上期</button>
    </router-link>
    <router-link to="/aimSettingSheet">
      <button>下期</button>
    </router-link>
    <br/>
    <button @click="getAPI">get api</button>
    <br/>
    <div>
      <ul>
        <li v-for="financialYear in financialYears" :key="financialYear.id" v-cloak>{{ financialYear.period }}</li>
      </ul>
    </div>
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
  import axios from 'axios'

  export default {
    components: {
      Announce_board
    },
    data(){
      return {
        financialYears : null
      }
    },
    created(){
      axios
        .get('http://localhost:8000/period')
        // .get('https://api.coindesk.com/v1/bpi/currentprice.json')
        .then(res =>{
          this.financialYears = res.data
          console.log(res);
        })
        .catch(err =>{
          console.error(err);
        })
    },
    methods:{
      getAPI(){
        axios
          .get('http://localhost:8000/period')
          // .get('https://api.coindesk.com/v1/bpi/currentprice.json')
          .then(res =>{
            console.log(res);
          })
          .catch(err =>{
            console.error(err);
          })
      }
    }
  }
</script>
