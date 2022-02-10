<template>
  <div>
    <div class="aimSetting_Enter">
      {{ tab }}
      <br />
      <div id="form-what">
        <p class="pp">職務目標項目（何を）</p>
        <textarea class="whatAndWhere" v-model="what" />
      </div>
      <div id="form-where">
        <p class="pp">目標水準（どこまで）</p>
        <textarea class="whatAndWhere" v-model="where" />
      </div>

      <br />

      <div id="form-how">
        <p class="pp">具体的達成手順（どのように）</p>
        <input
          type="text"
          class="how"
          name="how1"
          v-model="how1.achievementMean"
        /><br />
        <input
          type="text"
          class="how"
          name="how2"
          v-model="how2.achievementMean"
        /><br />
        <input
          type="text"
          class="how"
          name="how3"
          v-model="how3.achievementMean"
        /><br />
        <input
          type="text"
          class="how"
          name="how4"
          v-model="how4.achievementMean"
        /><br />
        <input
          type="text"
          class="how"
          name="how5"
          v-model="how5.achievementMean"
        /><br />
        <input
          type="text"
          class="how"
          name="how6"
          v-model="how6.achievementMean"
        /><br />
      </div>
      <div id="form-when">
        <p class="pp">スケジュール（いつまで）</p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how1.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how1.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how1.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how1.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how1.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how1.sixthMonth"
          />
        </p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how2.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how2.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how2.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how2.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how2.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how2.sixthMonth"
          />
        </p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how3.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how3.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how3.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how3.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how3.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how3.sixthMonth"
          />
        </p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how4.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how4.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how4.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how4.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how4.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how4.sixthMonth"
          />
        </p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how5.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how5.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how5.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how5.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how5.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how5.sixthMonth"
          />
        </p>
        <p>
          <input
            type="checkbox"
            name="when"
            value="5"
            v-model="how6.firstMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="6"
            v-model="how6.secondMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="7"
            v-model="how6.thirdMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="8"
            v-model="how6.fourthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="9"
            v-model="how6.fifthMonth"
          />
          <input
            type="checkbox"
            name="when"
            value="10"
            v-model="how6.sixthMonth"
          />
        </p>
      </div>

      <br /><br />

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
          <input type="number" class="weight" name="weight" v-model="weight" />
        </div>
        <br />
        <button hidden>評価シミュレーション</button>
        <br />
      </div>
      <div id="weight_graph">
        <PieChart
          :chart-data="datacollection"
          :options="options"
          style="height:150px; width:150px;"
        />

        <!-- <textarea class="weight_graph" placeholder="ここはグラフを表示しますがとりあえず保留します"/> -->
      </div>
      <br /><br />
      <button @click="postAims">目標設定保存</button>
      <button @click="putAims">目標設定編集</button>
      <button @click="postMeans">具体的達成手段保存</button>
      <button @click="putMean">具体的達成手段編集</button>
      <br />
      <button @click="postEB">POST面談コメント</button>

      <br /><br />

      <!-- v-if="EB_first" -->
      <div id="square">
        <p class="pp">面談コメント</p>
        <div id="form-what">
          <p class="pp">一次面談者</p>
          <textarea class="whatAndWhere" v-model="EB_first" />
        </div>
        <div id="form-where">
          <p class="pp">二次面談者</p>
          <textarea class="whatAndWhere" v-model="EB_second" />
        </div>
      </div>

      <br />
    </div>
  </div>
</template>

<script>
import {
  getAims,
  // getAchievementMeans,
  getAchievementMean,
  // postAchievementMeans,
  postAchievementMean,
  // putAchievementMeans,
  putAchievementMean,
  postAim,
  putAim,
  getEvaluationBefore,
  postEvaluationBefore,
} from "@/api/AimSettingSheet.js";
import PieChart from "@/api/PieChart.js";

export default {
  components: {
    PieChart,
  },
  props: ["tab"],
  data() {
    return {
      datacollection: null,
      options: {
        legend: {
          display: false,
        },
      },
      data: [],
      what: "",
      where: "",
      when: "",
      weight: "",
      level: "5",
      period: "",
      aim_id: "",
      how1: {},
      how2: {},
      how3: {},
      how4: {},
      how5: {},
      how6: {},
      how1_flag: false,
      how2_flag: false,
      how3_flag: false,
      how4_flag: false,
      how5_flag: false,
      how6_flag: false,
      how1_counter: 0,
      how2_counter: 0,
      how3_counter: 0,
      how4_counter: 0,
      how5_counter: 0,
      how6_counter: 0,
      EB_first: "",
      EB_second: "",
      weight1: 0,
      weight2: 0,
      weight3: 0,
      weight4: 0,
      weight5: 0,
      postMeans: "",
    };
  },
  created() {
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async () => {
      await this.getAim();
    })();
    this.fillData();
  },
  methods: {
    async getAim() {
      // アクセストークンの取得
      const access_token = await this.$auth.getTokenSilently();
      const aim = await getAims(access_token);
      //const achievementMeans = await getAchievementMeans(parseInt(this.tab));
      const am1 = await getAchievementMean(parseInt(this.tab), 1, access_token);
      const am2 = await getAchievementMean(parseInt(this.tab), 2, access_token);
      const am3 = await getAchievementMean(parseInt(this.tab), 3, access_token);
      const am4 = await getAchievementMean(parseInt(this.tab), 4, access_token);
      const am5 = await getAchievementMean(parseInt(this.tab), 5, access_token);
      const am6 = await getAchievementMean(parseInt(this.tab), 6, access_token);
      console.log(am1);

      if (
        !am1.result &&
        !am2.result &&
        !am3.result &&
        !am4.result &&
        !am5.result &&
        !am6.result
      ) {
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          1,
          this.how1.achievementMean,
          this.how1.firstMonth,
          this.how1.secondMonth,
          this.how1.thirdMonth,
          this.how1.fourthMonth,
          this.how1.fifthMonth,
          this.how1.sixthMonth,
          access_token
        );
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          2,
          this.how2.achievementMean,
          this.how2.firstMonth,
          this.how2.secondMonth,
          this.how2.thirdMonth,
          this.how2.fourthMonth,
          this.how2.fifthMonth,
          this.how2.sixthMonth,
          access_token
        );
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          3,
          this.how3.achievementMean,
          this.how3.firstMonth,
          this.how3.secondMonth,
          this.how3.thirdMonth,
          this.how3.fourthMonth,
          this.how3.fifthMonth,
          this.how3.sixthMonth,
          access_token
        );
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          4,
          this.how4.achievementMean,
          this.how4.firstMonth,
          this.how4.secondMonth,
          this.how4.thirdMonth,
          this.how4.fourthMonth,
          this.how4.fifthMonth,
          this.how4.sixthMonth,
          access_token
        );
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          5,
          this.how5.achievementMean,
          this.how5.firstMonth,
          this.how5.secondMonth,
          this.how5.thirdMonth,
          this.how5.fourthMonth,
          this.how5.fifthMonth,
          this.how5.sixthMonth,
          access_token
        );
        await postAchievementMean(
          "202105",
          1,
          this.tab,
          6,
          this.how6.achievementMean,
          this.how6.firstMonth,
          this.how6.secondMonth,
          this.how6.thirdMonth,
          this.how6.fourthMonth,
          this.how6.fifthMonth,
          this.how6.sixthMonth,
          access_token
        );
      } else {
        console.log("postAchievementMeanしなかった");
      }

      //const evaluationBefore = await getEvaluationBefore();
      //let d = {};
      const aim_target = aim.result.aims.find((v) => v.aimNumber === this.tab);
      //const achievement_target = achievementMeans.result.achievementMeans.find((v) => v.aim_number === parseInt(this.tab))
      if (aim_target) {
        this.what = aim_target.aimItem;
        this.where = aim_target.achievementLevel;
        this.level = aim_target.achievementDifficultyBefore;
        this.weight = aim_target.achievementWeight;
        this.period = aim_target.period;
        this.aim_id = aim_target.id;
      }

      // この書き方性能悪そう
      const chartData1 = aim.result.aims.find(
        (v) => parseInt(v.aimNumber) === 1
      );
      if (chartData1) {
        this.weight1 = chartData1.achievementWeight;
      }
      console.log(this.weight1);
      const chartData2 = aim.result.aims.find(
        (v) => parseInt(v.aimNumber) === 2
      );
      if (chartData2) {
        this.weight2 = aim_target.achievementWeight;
      }
      const chartData3 = aim.result.aims.find(
        (v) => parseInt(v.aimNumber) === 3
      );
      if (chartData3) {
        this.weight3 = aim_target.achievementWeight;
      }
      const chartData4 = aim.result.aims.find(
        (v) => parseInt(v.aimNumber) === 4
      );
      if (chartData4) {
        this.weight4 = aim_target.achievementWeight;
      }
      const chartData5 = aim.result.aims.find(
        (v) => parseInt(v.aimNumber) === 5
      );
      if (chartData5) {
        this.weight5 = aim_target.achievementWeight;
      }

      if (am1.result) {
        this.how1 = am1.result.achievementMean;
      }
      if (am2.result) {
        this.how2 = am2.result.achievementMean;
      }
      if (am3.result) {
        this.how3 = am3.result.achievementMean;
      }
      if (am4.result) {
        this.how4 = am4.result.achievementMean;
      }
      if (am5.result) {
        this.how5 = am5.result.achievementMean;
      }
      if (am6.result) {
        this.how6 = am6.result.achievementMean;
      }

      // 面談コメント（getEvaluiationBefore)の呼び出し
      if (aim_target) {
        const access_token = await this.$auth.getTokenSilently();
        const EB1 = await getEvaluationBefore(this.aim_id, 1, access_token);
        if (EB1.result) {
          this.EB_first = EB1.result.evaluationBefore.comment;
        }
        const EB2 = await getEvaluationBefore(this.aim_id, 2, access_token);
        if (EB2.result) {
          this.EB_second = EB2.result.evaluationBefore.comment;
        }
      }
    },
    // 具体的達成手段編集のボタンを押した時、変更があった具体的達成手段（flag=true）のものだけputする
    // putMeansが動いて欲しい本当のタイミングは①保存ボタンを押した時②タブが変わった時
    async putMean() {
      const access_token = await this.$auth.getTokenSilently();
      if (this.how1_flag == true) {
        await putAchievementMean(
          this.how1.id,
          "202105",
          this.how1.userId,
          this.how1.aimNumber,
          this.how1.achievementMeanNumber,
          this.how1.achievementMean,
          this.how1.firstMonth,
          this.how1.secondMonth,
          this.how1.thirdMonth,
          this.how1.fourthMonth,
          this.how1.fifthMonth,
          this.how1.sixthMonth,
          access_token
        );
        this.how1_flag = false;
      }
      if (this.how2_flag == true) {
        await putAchievementMean(
          this.how2.id,
          "202105",
          this.how2.userId,
          this.how2.aimNumber,
          this.how2.achievementMeanNumber,
          this.how2.achievementMean,
          this.how2.firstMonth,
          this.how2.secondMonth,
          this.how2.thirdMonth,
          this.how2.fourthMonth,
          this.how2.fifthMonth,
          this.how2.sixthMonth,
          access_token
        );
        this.how2_flag = false;
      }
      if (this.how3_flag == true) {
        await putAchievementMean(
          this.how3.id,
          "202105",
          this.how3.userId,
          this.how3.aimNumber,
          this.how3.achievementMeanNumber,
          this.how3.achievementMean,
          this.how3.firstMonth,
          this.how3.secondMonth,
          this.how3.thirdMonth,
          this.how3.fourthMonth,
          this.how3.fifthMonth,
          this.how3.sixthMonth,
          access_token
        );
        this.how3_flag = false;
      }
      if (this.how4_flag == true) {
        await putAchievementMean(
          this.how4.id,
          "202105",
          this.how4.userId,
          this.how4.aimNumber,
          this.how4.achievementMeanNumber,
          this.how4.achievementMean,
          this.how4.firstMonth,
          this.how4.secondMonth,
          this.how4.thirdMonth,
          this.how4.fourthMonth,
          this.how4.fifthMonth,
          this.how4.sixthMonth,
          access_token
        );
        this.how4_flag = false;
      }
      if (this.how5_flag == true) {
        await putAchievementMean(
          this.how5.id,
          "202105",
          this.how5.userId,
          this.how5.aimNumber,
          this.how5.achievementMeanNumber,
          this.how5.achievementMean,
          this.how5.firstMonth,
          this.how5.secondMonth,
          this.how5.thirdMonth,
          this.how5.fourthMonth,
          this.how5.fifthMonth,
          this.how5.sixthMonth,
          access_token
        );
        this.how5_flag = false;
      }
      if (this.how6_flag == true) {
        await putAchievementMean(
          this.how6.id,
          "202105",
          this.how6.userId,
          this.how6.aimNumber,
          this.how6.achievementMeanNumber,
          this.how6.achievementMean,
          this.how6.firstMonth,
          this.how6.secondMonth,
          this.how6.thirdMonth,
          this.how6.fourthMonth,
          this.how6.fifthMonth,
          this.how6.sixthMonth,
          access_token
        );
        this.how6_flag = false;
      }
    },
    async postAims() {
      const access_token = await this.$auth.getTokenSilently();
      await postAim(
        // periodがベタ書きになってるよ
        // -> vuexにperiodってのを作ってそこから持ってくる（String(this.period)）
        "202105",
        1,
        this.what,
        this.where,
        parseInt(this.weight),
        parseInt(this.level),
        parseInt(this.tab),
        access_token
      );
    },
    async putAims() {
      const access_token = await this.$auth.getTokenSilently();
      await putAim(
        parseInt(this.aim_id),
        "202105",
        1,
        this.what,
        this.where,
        parseInt(this.weight),
        parseInt(this.level),
        parseInt(this.tab),
        access_token
      );
    },
    async postEB() {
      // 本当は管理者画面？にあるべき
      await postEvaluationBefore(1, "頑張れ", 2, 1);
    },
    fillData() {
      console.log(this.weight1);
      console.log("↑mountedで動かしてるからweightの値が初期値のまま...泣");
      this.datacollection = {
        labels: ["1", "2", "3", "4", "5"],
        datasets: [
          {
            label: "データセットラベルA",
            // data: [parseInt(this.weight1), parseInt(this.weight2), parseInt(this.weight3), parseInt(this.weight4), parseInt(this.weight5)],
            data: [20, 20, 20, 20, 20],
            backgroundColor: [
              "#ff7b00",
              "#ff9500",
              "#ffaa00",
              "#ffc300",
              "#ffdd00",
            ],
          },
        ],
      };
    },
  },
  // 具体的達成手段とスケジュールに変更があるか監視
  // 変更なし：false（初期値） 変更あり：true
  watch: {
    how1: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how1_counter++;
        if (this.how1_counter > 1) {
          this.how1_flag = true;
        }
      },
      deep: true,
    },
    how2: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how2_counter++;
        if (this.how2_counter > 1) {
          this.how2_flag = true;
        }
      },
      deep: true,
    },
    how3: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how3_counter++;
        if (this.how3_counter > 1) {
          this.how3_flag = true;
        }
      },
      deep: true,
    },
    how4: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how4_counter++;
        if (this.how4_counter > 1) {
          this.how4_flag = true;
        }
      },
      deep: true,
    },
    how5: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how5_counter++;
        if (this.how5_counter > 1) {
          this.how5_flag = true;
        }
      },
      deep: true,
    },
    how6: {
      // データの値が変化した時にやりたいこと
      handler: function() {
        this.how6_counter++;
        if (this.how6_counter > 1) {
          this.how6_flag = true;
        }
      },
      deep: true,
    },
  },
};
</script>

<style scoped>
.aimSetting_Enter {
  /* width: 100%; */
  height: 750px;
  /* height: 550px; */
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

input.weight {
  width: 140px;
  height: 24px;
  margin: 0 10px;
}

.level {
  width: 140px;
  height: 30px;
  margin: 0 10px;
}

textarea.when,
.weight_graph {
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

textarea,
input,
select {
  border: 1px solid;
}

button {
  border: 3px solid #ffb341;
  border-radius: 100px;
  padding: 4px;
}

/* 横並び調整 */
#form-what,
#form-where,
#form-how,
#form-when,
#level,
#weight,
#levelAndWeight,
#weight_graph {
  display: inline-block;
  vertical-align: middle;
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
