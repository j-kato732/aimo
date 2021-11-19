<template>
  <div>
    <p>中期経営方針</p>
    <textarea class="large"/>
    <br>
    <br>
    <p>21期経営方針</p>
    <textarea class="large"/>
    <br>
    <div id="square">
      <br>
      <p class="pp">面談コメント</p>
      <select name="department">
        <option hidden>部を選択 ▼</option>
        <option value="solution">ソリューション本部</option>
        <option value="asia">エイジア開発事業本部</option>
        <option value="sapporo">札幌開発センター</option>
        <option value="sales">営業本部</option>
        <option value="business">事業推進本部</option>
        <option value="management">管理本部</option>
      </select>
        <br>
        <br>
        <p>部署目標</p>
        <textarea class="large"/>
    </div>
  </div>
</template>

<style scoped>
  textarea.large {
    resize: none;
    width: 600px;
    height: 100px;
    margin: 0 10px;
  }

  #square {
    display: inline-block;
    width: 700px;
    height: 800px;
    padding: 0.5em 1em;
    margin: 2em 0;
    border: solid 1px #000000;
    background: rgba(0, 0, 0, 0.1);;
  }

  textarea, input, select {
    border: 1px solid;
    background: #f7f7f7;
  }
</style>