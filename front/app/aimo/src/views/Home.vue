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
