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
      <label class="acd-label" for="acd-check3">エイジア開発事業本部</label>
      <div class="acd-content">
        <p>西野英樹</p>
      </div>
      <input id="acd-check4" class="acd-check" type="checkbox">
      <label class="acd-label" for="acd-check4">札幌開発センター</label>
      <div class="acd-content">
        <p>神農渉</p>
      </div>
    </div>
    <div id="right-box">
      氏<input type=text /><br><br>
      名<input type=text /><br><br>
      職位
      <select name="job">
        <option value="s" selected>ジュニア</option>
        <option value="a">サブリーダー</option>
        <option value="sa">リーダー</option>
        <option value="sal">マネージャー</option>
        <option value="b">スペシャリスト</option>
        <option value="m">シニアスペシャリスト</option>
      </select>
      <br><br>
      部署
      <select name="department">
        <option value="solution" selected>ソリューション本部</option>
        <option value="asia">エイジア開発事業本部</option>
        <option value="sapporo">札幌開発センター</option>
        <option value="sales">営業本部</option>
        <option value="business">事業推進本部</option>
        <option value="management">管理本部</option>
      </select>
      <br>
      <br>
      <select name="department_detail">
        <option value="no1" selected>開発１部</option>
        <option value="no2">開発２部</option>
        <option value="no3">開発３部</option>
        <option value="no4">開発４部</option>
      </select>
      <br>
      <br>
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
  height: 1000px;
}
select, input {
  border: solid 1px #000000;
}
</style>
<!-- 高さ1000pxにしたけど動的にしたい -->