<template>
  <div id="user-list-tab">
    <div id="left-box">
      <input id="acd-check2" class="acd-check" type="checkbox">
      <label class="acd-label" for="acd-check2">ソリューション本部</label>
      <div class="acd-content">
        <p class="name">加藤烈</p>
        <p class="name">山室雄一郎</p>
        <p class="name">古山篤</p>
        <p class="name">東幸平</p>
      </div>
      <input id="acd-check3" class="acd-check" type="checkbox">
      <label class="acd-label" for="acd-check3">クリックで開く3</label>
      <div class="acd-content">
        <p>hello.world3!</p>
      </div>
      <input id="acd-check4" class="acd-check" type="checkbox">
      <label class="acd-label" for="acd-check4">クリックで開く4</label>
      <div class="acd-content">
        <p>hello.world4!</p>
      </div>
    </div>
    <div id="right-box">
      <p>ああああああああああ</p>
    </div>
  </div>
</template>

<style scoped>
.acd-check{
  display: none;
}
.acd-label{
  background: #333;
  color: #fff;
  display: block;
  margin-bottom: 1px;
  padding: 10px;
}
.acd-content{
  border: 1px solid #333;
  height: 0;
  opacity: 0;
  padding: 0 10px;
  transition: .5s;
  visibility: hidden;
}
.acd-check:checked + .acd-label + .acd-content{
  height: 200px;
  opacity: 1;
  padding: 10px;
  visibility: visible;
}
#left-box {
  float: left;
  width: 50%
}
#right-box {
  float: right;
  width: 50%;
}
.name {
  border-bottom: 1px solid #333;
  padding: 10px 0;
}
#user-list-tab {
  height: 90%;
}
</style>