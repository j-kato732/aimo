<template>
  <div class="periodSeparatePage">
    <h1>期別ページ</h1>
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
  // //import HelloWorld from '../components/HelloWorld'

  //export default {
  //   name: 'Home',

  //   components: {
  //     //HelloWorld,
  //   },
  //   methods:{
  //     checkTest(){
  //       alert(this.$store.state.store_test);
  //     }
  //   }
  //}
</script>
