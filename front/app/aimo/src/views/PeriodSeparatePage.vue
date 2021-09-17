<template>
  <div class="periodSeparatePage">
    <HeaderName title="期別ページ"></HeaderName>
    <br/>
    <br/>
    <router-link to="/management">
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
