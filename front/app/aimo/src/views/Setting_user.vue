<template>
  <div class="setting_user">
    <HeaderName title="Setting"></HeaderName>
    <br/>
    <br/>
    氏<input class="input_field" type="text">
    名<input class="input_field" type="text">
    <br>
    <br>
    職位
    <select name="job">
      <option value="s" selected>ジュニア</option>
      <option value="a">サブリーダー</option>
      <option value="sa">リーダー</option>
      <option value="sal">マネージャー</option>
      <option value="b">スペシャリスト</option>
      <option value="m">シニアスペシャリスト</option>
    </select>
    <br>
    <br>
    部署
    <select name="department">
      <option value="solution" selected>ソリューション本部</option>
      <option value="asia">エイジア開発事業本部</option>
      <option value="sapporo">札幌開発センター</option>
      <option value="sales">営業本部</option>
      <option value="business">事業推進本部</option>
      <option value="management">管理本部</option>
    </select>
    <select name="department_detail">
      <option value="no1" selected>開発１部</option>
      <option value="no2">開発２部</option>
      <option value="no3">開発３部</option>
      <option value="no4">開発４部</option>
    </select>
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
