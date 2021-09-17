<template>
  <div class="management">
    <HeaderName title="Management"></HeaderName>
    <br/>
    <router-link to="/periodSeparatePage">
      <button>期別ページ</button>
    </router-link>
    <br/>
    <br/>
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

  export default {
  //   name: 'Home',

    components: {
      HeaderName
    },
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  //   }
  }
</script>