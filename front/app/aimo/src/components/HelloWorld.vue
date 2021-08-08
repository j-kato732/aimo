<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <router-link to="/home">
      <button>Home</button>
    </router-link>
    <p>
      ↑<br/>
      MY PAGE<br/>
      <br/>
      ここはのちにログインページにする予定<br/>
      ログインするときのID次第で遷移先がMYPAGE（ユーザー）かManagement（管理者）か識別するけど<br/>
      今はとりあえず画面遷移がしたいので別々のボタンにします<br/>
      <br/>
      Management<br/>
      ↓
    </p>
    <router-link to="/management">
      <button>Management</button>
    </router-link>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
