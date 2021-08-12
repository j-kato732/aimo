<template>
  <div id="app">
    <Header/>
    <!-- router-viewで囲った範囲が画面遷移の対象範囲（切り替わる範囲） -->
    <router-view>
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Aimo"/>

    </router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/Header.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Header
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 0;
}

/* ヘッダーを作った時に端に余白ができるので指定 */
*{
  margin: 0;
}

html {
  background: #EFEEEA;
}
</style>
