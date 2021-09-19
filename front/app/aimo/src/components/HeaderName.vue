<template>
  <header class="header">
    <h1 class="title">{{ title }}</h1>
  </header>
</template>

<script>
export default {
    // 親コンポーネントからデータを受け取る
    props: ["title"]
}
</script>

<style scoped>
  header {
    background-color: #F7F7F7;
    height: 80px;
    color: #2C3E50;
  }

  .title {
    line-height: 80px;
  }
</style>