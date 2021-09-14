<template>
  <div>
    <div class="aimSetting_Enter">
      {{ tab }}
      <br>
      <div id="form-what">
        <p class="pp">職務目標項目（何を）</p>
        <textarea class="whatAndWhere" v-model="data[0].what" />
      </div>
      <div id="form-where">
        <p class="pp">目標水準（どこまで）</p>
        <textarea class="whatAndWhere" v-model="data[0].where" />
      </div>

      <br>
      
      <div id="form-how">
        <p class="pp">具体的達成手順（どのように）</p>
        <input type=text class="how" name="how1" v-model="data[0].how1" /><br>
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
            <option value="6">6</option>
            <option value="5" v-if="data[0].level === 5" selected>5です</option>
            <option value="5" v-else>5</option>
            <option value="4">4</option>
            <option value="3">3</option>
            <option value="2">2</option>
            <option value="1">1</option>
          </select>
        </div>
        <div id="weight">
          <p class="pp">ウエイト</p>
          <input type=number class="weight" name="weight" v-model="data[0].weight" />
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
import {getAim} from '@/api/AimSettingSheet.js'

export default{
  props: ["tab"],
  data(){
    return {
      data: [],
      what: "",
      where:"",
      when:"",
      weight:"",
      how1:"",
      how2:"",
      how3:"",
      how4:"",
      how5:"",
      how6:""
    }
  },
  created(){
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async()=>{
      await this.getAim();
      })();
  },
  methods:{
    async getAim(){
      const aim = await getAim();
      console.log(aim.result.aim[0].aim_item)
      let d = {};
      d.what = aim.result.aim[0].aim_item;
      d.where = aim.result.aim[0].achievement_level;
      d.how1 = aim.result.aim[0].achievement_means[0].achievement_mean;
      d.level = aim.result.aim[0].achievement_diffculty_before;
      d.weight = aim.result.aim[0].achievement_weight;
      this.data.push(d);
      console.log(this.data);
    }
  }
}
// import {getRole} from '@/api/AimSettingSheet.js'
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
//}
</script>

<style scoped>
  .aimSetting_Enter {
    /* width: 100%; */
    height: 750px;
    background: rgba(255, 179, 65, 0.3);
    /* margin: 0 20%; */
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

  textarea, input, select {
    border: 1px solid;
  }

  button {
    border: 3px solid #FFB341;
    border-radius: 100px;
    padding: 4px;
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