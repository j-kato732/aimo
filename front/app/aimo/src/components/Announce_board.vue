<template>
  <div class="announce_board">
    <h3>お知らせ</h3>
    <br>
    <p v-for="d in data" :key="d.id">
      2021/05/01 ▼ {{ d.content }}
    </p>
    <br>
    <br/>
    <br/>
  </div>
</template>

<script>
import {getInfo} from '@/api/MyPage'

export default {
  data(){
    return {
      data: []
    }
  },
  created(){
    // 下に書いたgetAPI関数をページ遷移時に呼び出す
    (async()=>{
      await this.getAimoInfo();
      })();
    },
    methods:{
      async getAimoInfo(){
        const subDatas = await getInfo();
        //console.log(subDatas.result);
        for(let data of subDatas.result.aimoInfo){
          let d = {};
          d.id = data.aimo_info_id;
          d.content = data.aimo_info_content;
          this.data.push(d);
          console.log(this.data);
      }
    }
  }
}
</script>

<style scoped>
  .announce_board {
    width: 60%;
    height: 120px;
    background: #F7F7F7;
    margin: 0 20%;
  }
</style>