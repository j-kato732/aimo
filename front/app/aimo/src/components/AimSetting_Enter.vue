<template>
  <div>
    <button @click="$emit('input', id)" :class="[active, 'tab']">
    {{ label }}
    </button>
    <div class="aimSetting_Enter">
      <br>
      <div id="form-what">
        <p class="pp">職務目標項目（何を）</p>
        <textarea class="whatAndWhere" v-model="what" />
      </div>
      <div id="form-where">
        <p class="pp">目標水準（どこまで）</p>
        <textarea class="whatAndWhere" v-model="where" />
      </div>

      <br>
      
      <div id="form-how">
        <p class="pp">具体的達成手順（どのように）</p>
        <input type=text class="how" name="how1" v-model="how1" /><br>
        <input type=text class="how" name="how2" v-model="how2" /><br>
        <input type=text class="how" name="how3" v-model="how3" /><br>
        <input type=text class="how" name="how4" v-model="how4" /><br>
        <input type=text class="how" name="how5" v-model="how5" /><br>
        <input type=text class="how" name="how6" v-model="how6" /><br>
      </div>
      <div id="form-when">
        <p class="pp">スケジュール（いつまで）</p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
        <p>
          <input type="checkbox" name="when" value="5">
          <input type="checkbox" name="when" value="6">
          <input type="checkbox" name="when" value="7">
          <input type="checkbox" name="when" value="8">
          <input type="checkbox" name="when" value="9">
          <input type="checkbox" name="when" value="10">
        </p>
      </div>
      
      <br><br>

      <div id="levelAndWeight">
        <div id="level">
          <p>難易度</p>
          <select class="level">
            <option value="5">5</option>
            <option value="4">4</option>
            <option value="3">3</option>
            <option value="2">4</option>
            <option value="1">3</option>
          </select>
        </div>
        <div id="weight">
          <p class="pp">ウエイト</p>
          <input type=number class="weight" name="weight" v-model="weight" />
        </div>
        <br><br>
        <button>評価シミュレーション</button>
      </div>
      <div id="weight_graph">
        <textarea class="weight_graph" placeholder="ここはグラフを表示しますがとりあえず保留します"/>
      </div>

      <br><br><br>

      <div id="square">
        <p class="pp">評価面談コメント</p>
        <div id="form-what">
          <p class="pp">一次面談者</p>
          <textarea class="whatAndWhere" v-model="what" />
        </div>
        <div id="form-where">
          <p class="pp">二次面談者</p>
          <textarea class="whatAndWhere" v-model="where" />
        </div>
      </div>

      <br>
    </div>
  </div>
</template>

<script>
// import {getRole} from '@/api/AimSettingSheet.js'

 export default {
   props: {
    id: Number,
    label: String,
    value: Number
  },
  computed: {
    active() {
      return this.value === this.id ? 'active' : false
    }
  }
//   data(){
//     return {
//       data: []
//       // これ、複数データ返ってくる訳じゃないから配列にする必要ないんだけどねぇ。
//     }
//   },
//   created(){
//     // 下に書いたgetAPI関数をページ遷移時に呼び出す
//     (async()=>{
//       await this.getRole();
//       })();
//   },
//   methods:{
//     async getRole(){
//       const role = await getRole();
//       console.log(role.result.role.role);
//       let d = {};
//       d.id = role.result.role.role_id;
//       const p = String(role.result.role.period);
//       d.financialYear_YY = p.replace(/^\d{2}(\d{2})\d{2}/, '$1期');
//       d.department_id = role.result.role.department_id;
//       d.job_id = role.result.role.job_id;
//       d.role = role.result.role.role;
//       this.data.push(d);
//       console.log(this.data);
//     }
//   }
}
</script>

<style scoped>
  .aimSetting_Enter {
    width: 60%;
    height: 750px;
    background: rgba(255, 179, 65, 0.3);
    margin: 0 20%;
  }

  p.pp {
    padding: 8px;
  }

  textarea.whatAndWhere {
    resize: none;
    height: 60px;
    width: 300px;
    margin: 0 10px;
  }

  input.how {
    width: 300px;
    margin: 0 10px;
  }

  input.weight{
    width: 140px;
    height: 24px;
    margin: 0 10px;
  }

  .level {
    width: 140px;
    height: 30px;
    margin: 0 10px;
  }

  textarea.when, .weight_graph {
    resize: none;
    height: 124px;
    width: 300px;
    margin: 0 10px;
  }

  #square {
    display: inline-block;
    padding: 0.5em 1em;
    margin: 2em 0;
    border: solid 1px #000000;
  }

/* 横並び調整 */
  #form-what, #form-where, #form-how, #form-when, #level, #weight, #levelAndWeight, #weight_graph {
    display: inline-block;
    vertical-align:  middle;
  }

  .tab {
  border-radius: 2px 2px 0 0;
  background: #fff;
  color: #311d0a;
  line-height: 24px;
  }
  .tab:hover {
    background: #eeeeee;
  }
  .active {
    background: #f7c9c9;
  }

  .button {
    resize: none;
    height: 60px;
    width: 100px;
  }

  .checkbox {
    resize: none;
    width: 30px;
    height: 30px;
    padding: 2px;
  }
</style>