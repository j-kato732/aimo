<template>
  <div id="app">
    <!-- router-viewで囲った範囲が画面遷移の対象範囲（切り替わる範囲） -->
    <router-view>
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Aimo"/>

    </router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
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
  margin-top: 60px;
}
</style>
