<template>
  <div class="setting_user">
    <HeaderName title="Setting"></HeaderName>
    <br/>
    <br/>
    <div class="login_id">
      氏<input class="input_field" type="text">
    </div>
    <br>
    名<input class="input_field" type="text">
    <br>
    <br>
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

  export default {
  //   name: 'Home',

    components: {
      HeaderName
    },
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  //    }
  }
</script>
