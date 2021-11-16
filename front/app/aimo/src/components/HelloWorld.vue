<template>
  <div class="hello">
    <HeaderName title="Aimo"></HeaderName>
    <br/>
    <div class="login_id">
      <input class="input_field" type="text" placeholder="  ID">
    </div>
    <br>
    <input class="input_field" type="text" placeholder="  Password">
    <br>
    <br>
    <input id="keepLogin" type="checkbox">
    <label for="keepLogin">  ログイン状態を保持する</label>
    <br>
    <br>
    <router-link to="/mypage">
      <button>Login</button>
    </router-link>
    <br>
    <br>
    <br>
    <br>
    <h1>{{ msg }}</h1>
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
import HeaderName from '../components/HeaderName.vue'

export default {
  components: {
    HeaderName
  },
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
  color: #222222;
}

.input_field {
  width: 304px;
  height: 56px;
  background: #F7F7F7;
  border-radius: 90px;
}

button {
  width: 304px;
  height: 56px;
  background: rgba(255, 179, 65, 0.3);
  border: 6px solid #FFB341;
  border-radius: 90px;
}

/* .login_id i {
  position: absolute;
	top: 8px;
	left: 0;
	padding: 9px 8px;
	transition: 0.3s;
	color: #aaaaaa;
} */
</style>
