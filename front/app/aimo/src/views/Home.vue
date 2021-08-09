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
    <router-link to="/aimSettingSheet">
      <button>目標設定シート</button>
    </router-link>
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

  export default {
     components: {
       Announce_board
     }
    //  methods:{
    //    checkTest(){
    //      alert(this.$store.state.store_test);
    //    }
    //  }
  }
</script>
