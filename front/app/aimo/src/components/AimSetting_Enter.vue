<template>
  <div>
    <div class="aimSetting_Enter">
      <div id="form-what">
        <p>職務目標項目（何を）</p>
        <textarea class="whatAndWhere" v-model="what" />
      </div>
      <div id="form-where">
        <p>目標水準（どこまで）</p>
        <textarea class="whatAndWhere" v-model="where" />
      </div>

      <br>
      
      <div id="form-how">
        <p>具体的達成手順（どのように）</p>
        <input type=text class="how" name="how1" v-model="how1" /><br>
        <input type=text class="how" name="how2" v-model="how2" /><br>
        <input type=text class="how" name="how3" v-model="how3" /><br>
        <input type=text class="how" name="how4" v-model="how4" /><br>
        <input type=text class="how" name="how5" v-model="how5" /><br>
        <input type=text class="how" name="how6" v-model="how6" /><br>
      </div>
      <div id="form-when">
        <p>スケジュール（いつまで）</p>
        <textarea class="when" v-model="when" placeholder="ここはとりあえず保留にします"/>
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
          <p>ウエイト</p>
          <input type=number class="weight" name="weight" v-model="weight" />
        </div>
        <br><br>
        <button>評価シミュレーション</button>
      </div>
      <div id="weight_graph">
        <textarea class="weight_graph" placeholder="ここはグラフを表示しますがとりあえず保留します"/>
      </div>

      <br>
    </div>
  </div>
</template>

<script>
// import {getRole} from '@/api/AimSettingSheet.js'

// export default {
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
// }
</script>

<style scoped>
  .aimSetting_Enter {
    width: 60%;
    height: 1000px;
    background: rgba(255, 179, 65, 0.3);
    margin: 0 20%;
  }

  p {
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

/* 横並び調整 */
  #form-what, #form-where, #form-how, #form-when, #level, #weight, #levelAndWeight, #weight_graph {
    display: inline-block;
    vertical-align:  middle;
  }
</style>