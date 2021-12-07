<template>
  <div>
    <input id="acd-check2" class="acd-check" type="checkbox">
    <label class="acd-label" for="acd-check2">クリックで開く2</label>
    <div class="acd-content">
      <p>hello.world2!</p>
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
  height: 40px;
  opacity: 1;
  padding: 10px;
  visibility: visible;
}
</style>