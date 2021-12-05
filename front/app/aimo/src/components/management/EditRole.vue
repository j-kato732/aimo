<template>
  <div>
    <div>
      {{ tab }}
    </div>
    <br>
    <p>本人の役割</p>
    <textarea />
    <br>
    <br>
    <input id="acd-check1" class="acd-check" type="checkbox">
    <label class="acd-label" for="acd-check1">個人設定</label>
    <div class="acd-content">
      <select name="department">
        <option hidden>人を選択 ▼</option>
        <option value="AAA">AAA</option>
        <option value="BBB">BBB</option>
        <option value="CCC">CCC</option>
        <option value="DDD">DDD</option>
      </select>
      <br>
      <br>
      <textarea />
    </div>
  </div>
</template>

<script>

export default {
  props: ["tab"]
}
</script>

<style scoped>
  select {
    border: solid 1px #000000;
  }

  textarea {
    resize: none;
    width: 80%;
    height: 100px;
    margin: 0 10px;
    border: solid 1px #000000;
  }

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
    height: 0;
    opacity: 0;
    padding: 0 10px;
    transition: .5s;
    visibility: hidden;
  }
  .acd-check:checked + .acd-label + .acd-content{
    height: 190px;
    opacity: 1;
    padding: 10px;
    visibility: visible;
  }
</style>