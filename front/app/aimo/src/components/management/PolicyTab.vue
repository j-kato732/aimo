<template>
  <div id="policyTab">
    <p>中期経営方針</p>
    <textarea class="large" v-model="mid_term_policy"/>
    <br>
    <br>
    <p>21期経営方針</p>
    <textarea class="large" v-model="period_policy"/>
    <br>
    <button @click="putPolicy">方針編集</button>
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
        <br>
        <br>
        <TabItemOfEditRole />
    </div>
  </div>
</template>

<script>
import TabItemOfEditRole from '@/components/management/TabItemOfEditRole.vue'
import { getPolicy, postPolicy, putPolicy } from "@/api/AimSettingSheet.js";

export default {
  components: {
    TabItemOfEditRole
  },
  data() {
    return {
      mid_term_policy: "",
      period_policy: "",
      id: 0
    }
  },
  created() {
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async () => {
      await this.getPolicy();
    })();
  },
  methods: {
    async getPolicy() {
      this.$store.commit('setPeriodForManagement', this.$route.params.period)
      const access_token = await this.$auth.getTokenSilently();
      const policy = await getPolicy(this.$store.state.period_management, access_token)
      console.log(policy)

      if(!policy.result) {
        const id = await postPolicy(
          this.$store.state.period_management,
          this.mid_term_policy,
          this.period_policy,
          access_token
        );
        this.id = id.result.id;
      }
      if(policy.result){
        this.mid_term_policy = policy.result.policy.midTermPolicy
        this.period_policy = policy.result.policy.periodPolicy
        this.id = policy.result.policy.id
      }
    },
    async putPolicy() {
      const access_token = await this.$auth.getTokenSilently();
      await putPolicy(
        this.id,
        this.$store.state.period_management,
        this.mid_term_policy,
        this.period_policy,
        access_token
      )
    }
  }
}
</script>


<style scoped>
#policyTab {
  padding: 20px;
}

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