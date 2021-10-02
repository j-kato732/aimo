<template>
  <div>
    <div class="aimSetting_Enter">
      {{ tab }}
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
        <input type=text class="how" name="how1" v-model="how1.achievementMean" /><br>
        <input type=text class="how" name="how2" v-model="how2.achievementMean" /><br>
        <input type=text class="how" name="how3" v-model="how3.achievementMean" /><br>
        <input type=text class="how" name="how4" v-model="how4.achievementMean" /><br>
        <input type=text class="how" name="how5" v-model="how5.achievementMean" /><br>
        <input type=text class="how" name="how6" v-model="how6.achievementMean" /><br>
      </div>
      <div id="form-when">
        <p class="pp">スケジュール（いつまで）</p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how1.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how1.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how1.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how1.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how1.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how1.sixthMonth">
        </p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how2.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how2.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how2.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how2.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how2.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how2.sixthMonth">
        </p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how3.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how3.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how3.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how3.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how3.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how3.sixthMonth">
        </p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how4.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how4.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how4.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how4.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how4.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how4.sixthMonth">
        </p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how5.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how5.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how5.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how5.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how5.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how5.sixthMonth">
        </p>
        <p>
          <input type="checkbox" name="when" value="5" v-model="how6.firstMonth">
          <input type="checkbox" name="when" value="6" v-model="how6.secondMonth">
          <input type="checkbox" name="when" value="7" v-model="how6.thirdMonth">
          <input type="checkbox" name="when" value="8" v-model="how6.fourthMonth">
          <input type="checkbox" name="when" value="9" v-model="how6.fifthMonth">
          <input type="checkbox" name="when" value="10" v-model="how6.sixthMonth">
        </p>
      </div>
      
      <br><br>

      <div id="levelAndWeight">
        <div id="level">
          <p>難易度</p>
          <select class="level" v-model="level">
            <option value="6">6</option>
            <!-- <option value="5" v-if="level === 5" selected>5です</option> -->
            <option value="5">5</option>
            <option value="4">4</option>
            <option value="3">3</option>
            <option value="2">2</option>
            <option value="1">1</option>
          </select>
        </div>
        <div id="weight">
          <p class="pp">ウエイト</p>
          <input type=number class="weight" name="weight" v-model="weight" />
        </div>
        <br>
        <button hidden>評価シミュレーション</button>
        <br>
      </div>
      <div id="weight_graph">
        <textarea class="weight_graph" placeholder="ここはグラフを表示しますがとりあえず保留します"/>
      </div>
      <br><br>
      <button @click="postAims">目標設定保存</button>
      <button @click="postMeans">具体的達成手段保存</button>
      <button @click="putMeans">具体的達成手段編集</button>

      <br><br>

      <!-- <div id="square" v-if="what">
        <p class="pp">評価面談コメント</p>
        <div id="form-what">
          <p class="pp">一次面談者</p>
          <textarea class="whatAndWhere" v-model="what" />
        </div>
        <div id="form-where">
          <p class="pp">二次面談者</p>
          <textarea class="whatAndWhere" v-model="where" />
        </div>
      </div> -->

      <br>
    </div>
  </div>
</template>

<script>
import {
  getAim,
  getAchievementMeans,
  postAchievementMeans,
  putAchievementMeans,
  postAim
} from '@/api/AimSettingSheet.js'

export default{
  props: ["tab"],
  data(){
    return {
      data: [],
      what: "",
      where:"",
      when:"",
      weight:"",
      how1:{},
      how2:{},
      how3:{},
      how4:{},
      how5:{},
      how6:{},
      level:"5",
      period:""
    }
  },
  mounted(){
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async()=>{
      await this.getAim();
      })();
  },
  methods:{
    async getAim(){
      const aim = await getAim();
      const achievementMeans = await getAchievementMeans(parseInt(this.tab));
      console.log(this.tab);
      console.log(aim);
      console.log(achievementMeans);
      //let d = {};
      const aim_target = aim.result.aim.find((v) => v.aimNumber === this.tab)
      //const achievement_target = achievementMeans.result.achievementMeans.find((v) => v.aim_number === parseInt(this.tab))
      console.log(aim_target);
      if(aim_target){
        this.what = aim_target.aimItem;
        this.where = aim_target.achievementLevel;
        this.level = aim_target.achievementDifficultyBefore;
        this.weight = aim_target.achievementWeight;
        this.period = aim_target.period;
      } 
      if(achievementMeans){
        const sub1 = achievementMeans.result.achievementMeans.find((v) => parseInt(v.achievementMeanNumber) === 1);
        console.log("サブ１");
        console.log(sub1);
        console.log(achievementMeans.result.achievementMeans);
        this.how1 = sub1 ? sub1 : {};
        const sub2 = achievementMeans.result.achievementMeans.find((v) => parseInt(v.achievementMeanNumber) === 2);
        this.how2 = sub2 ? sub2 : {};
        console.log("サブ2");
        console.log(sub2);
        const sub3 = achievementMeans.result.achievementMeans.find((v) => parseInt(v.achievementMeanNumber) === 3);
        this.how3 = sub3 ? sub3 : {};
        console.log("サブ3");
        console.log(sub3);
        const sub4 = achievementMeans.result.achievementMeans.find((v) => v.achievementMeanNumber === "4");
        this.how4 = sub4 ? sub4 : {};
        const sub5 = achievementMeans.result.achievementMeans.find((v) => v.achievementMeanNumber === "5");
        this.how5 = sub5 ? sub5 : {};
        const sub6 = achievementMeans.result.achievementMeans.find((v) => v.achievementMeanNumber === "6");
        this.how6 = sub6 ? sub6 : {};
        console.log("サブ6");
        console.log(sub6);
      }
      console.log("＝＝＝＝＝HOW＝＝＝＝＝");
      console.log(this.how1);
      console.log(this.how2);
      console.log(this.how3);
      console.log(this.how6);
      // this.data.push(d);
      // console.log(this.data);
    },
    async postMeans(){
      const achievementMeans = await postAchievementMeans(
        // これだと開いてるタブの内容しか保存できないよ
        // -> タブが切り替わるたびにputかpostするようにする
        // ここ将来的に１行ずつになるよ
        // -> JSONが配列じゃなくなったら１行ずつgetした時のstatusを見てputするかpostするか処理を変える
        // userIDがベタ書きになってるよ（periodも）
        // -> vuexにuserIdってのを作ってそこから持ってくる（vuexに入ってくる値はとりあえず適当でOK -> 後に認証APIからとってくることになる）
        // periodがベタ書きになってるよ
        // -> vuexにperiodってのを作ってそこから持ってくる
        // achievement_mean_numberがベタ書きになってるよ
        // -> 1行ずつgetするようにするときにどこの行かがわかるような処理を作らないといけないからそこから参照するよ
        "202105", 1, this.tab, 1, this.how1.achievementMean, this.how1.firstMonth, this.how1.secondMonth, this.how1.thirdMonth, this.how1.fourthMonth, this.how1.fifthMonth, this.how1.sixthMonth,
        "202105", 1, this.tab, 2, this.how2.achievementMean, this.how2.firstMonth, this.how2.secondMonth, this.how2.thirdMonth, this.how2.fourthMonth, this.how2.fifthMonth, this.how2.sixthMonth,
        "202105", 1, this.tab, 3, this.how3.achievementMean, this.how3.firstMonth, this.how3.secondMonth, this.how3.thirdMonth, this.how3.fourthMonth, this.how3.fifthMonth, this.how3.sixthMonth,
        "202105", 1, this.tab, 4, this.how4.achievementMean, this.how4.firstMonth, this.how4.secondMonth, this.how4.thirdMonth, this.how4.fourthMonth, this.how4.fifthMonth, this.how4.sixthMonth,
        "202105", 1, this.tab, 5, this.how5.achievementMean, this.how5.firstMonth, this.how5.secondMonth, this.how5.thirdMonth, this.how5.fourthMonth, this.how5.fifthMonth, this.how5.sixthMonth,
        "202105", 1, this.tab, 6, this.how6.achievementMean, this.how6.firstMonth, this.how6.secondMonth, this.how6.thirdMonth, this.how6.fourthMonth, this.how6.fifthMonth, this.how6.sixthMonth,
      );
      console.log(achievementMeans);
    },
    async putMeans(){
      const achievementMeans = await putAchievementMeans(
        this.how1.id, "202105", this.how1.userId, this.how1.aimNumber, this.how1.achievementMeanNumber, this.how1.achievementMean, this.how1.firstMonth, this.how1.secondMonth, this.how1.thirdMonth, this.how1.fourthMonth, this.how1.fifthMonth, this.how1.sixthMonth,
        this.how2.id, "202105", this.how2.userId, this.how2.aimNumber, this.how2.achievementMeanNumber, this.how2.achievementMean, this.how2.firstMonth, this.how2.secondMonth, this.how2.thirdMonth, this.how2.fourthMonth, this.how2.fifthMonth, this.how2.sixthMonth,
        this.how3.id, "202105", this.how3.userId, this.how3.aimNumber, this.how3.achievementMeanNumber, this.how3.achievementMean, this.how3.firstMonth, this.how3.secondMonth, this.how3.thirdMonth, this.how3.fourthMonth, this.how3.fifthMonth, this.how3.sixthMonth,
        this.how4.id, "202105", this.how4.userId, this.how4.aimNumber, this.how4.achievementMeanNumber, this.how4.achievementMean, this.how4.firstMonth, this.how4.secondMonth, this.how4.thirdMonth, this.how4.fourthMonth, this.how4.fifthMonth, this.how4.sixthMonth,
        this.how5.id, "202105", this.how5.userId, this.how5.aimNumber, this.how5.achievementMeanNumber, this.how5.achievementMean, this.how5.firstMonth, this.how5.secondMonth, this.how5.thirdMonth, this.how5.fourthMonth, this.how5.fifthMonth, this.how5.sixthMonth,
        this.how6.id, "202105", this.how6.userId, this.how6.aimNumber, this.how6.achievementMeanNumber, this.how6.achievementMean, this.how6.firstMonth, this.how6.secondMonth, this.how6.thirdMonth, this.how6.fourthMonth, this.how6.fifthMonth, this.how6.sixthMonth,
      );
      console.log(achievementMeans);
      console.log("=====putMeans=====");
      console.log(this.how1.id, "202105", this.how1.userId, this.how1.aimNumber, this.how1.achievementMeanNumber, this.how1.achievementMean, this.how1.firstMonth, this.how1.secondMonth, this.how1.thirdMonth, this.how1.fourthMonth, this.how1.fifthMonth, this.how1.sixthMonth); 
    },
    async postAims(){
      const aim = await postAim(
        // periodがベタ書きになってるよ
        // -> vuexにperiodってのを作ってそこから持ってくる（String(this.period)）
        "202105", 1, this.what, this.where, this.weight, this.level, this.tab
      );
      console.log(aim);
    }
    // async putMeans(){
    //   const achievementMeans = await putAchievementMeans(
    //     4,String(this.period), 1, 1, 1, "put!!!!!!!!", true, true, true, true, true, true
    //   );
    //   console.log(achievementMeans);
    // },
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
    /* height: 750px; */
    height: 550px;
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