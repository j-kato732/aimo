import axios from 'axios'

const baseUrl = "http://localhost"
const port = 9002

//promiseの時はasync
export async function getPolicy(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/policy?period=202105')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getDepartmentGoal(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/departmentGoal?period=202105&department_id=2')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getRole(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/role?period=202105&department_id=2&job_id=6')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getAim(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/aim?period=202105&user_id=1')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

/**
 * 具体的達成手段取得API
 * @param {number} aim_number タブ番号
 * @returns 具体的達成手段
 */
export async function getAchievementMeans(aim_number){
  console.log(aim_number);
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/achievementMeans?period=202105&user_id=1&aim_number=${aim_number}`)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

/** JSDOC
 * function名
 * @description functionの説明
 * @param {String} period 期
 * @param {number} user_id ユーザーID
 * @param {*} aim_number 
 * @param {*} achievement_mean_number 
 * @param {*} achievement_mean 
 * @param {*} first_month 
 * @param {*} second_month 
 * @param {*} third_month 
 * @param {*} fourth_month 
 * @param {*} fifth_month 
 * @param {*} sixth_month 
 * @returns APIの結果
 */
export async function postAchievementMeans(
  period,user_id,aim_number,achievement_mean_number,achievement_mean,
  first_month,second_month,third_month,fourth_month,fifth_month,sixth_month){
  const body = {achievementMeans:[{
    "period": period,
    "user_id": user_id,
    "aim_number": aim_number,
    "achievement_mean_number": achievement_mean_number,
    "achievement_mean": achievement_mean,
    "first_month": first_month,
    "second_month": second_month,
    "third_month": third_month,
    "fourth_month": fourth_month,
    "fifth_month": fifth_month,
    "sixth_month": sixth_month
  }]};
  return new Promise((resolve,reject)=>{
    axios
      .post(`${baseUrl}:${port}/achievementMeans`, body)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function putAchievementMeans(
  id,period,user_id,aim_number,achievement_mean_number,achievement_mean,
  first_month,second_month,third_month,fourth_month,fifth_month,sixth_month){
  const body = {achievementMeans:[{
    "id": id,
    "period": period,
    "user_id": user_id,
    "aim_number": aim_number,
    "achievement_mean_number": achievement_mean_number,
    "achievement_mean": achievement_mean,
    "first_month": first_month,
    "second_month": second_month,
    "third_month": third_month,
    "fourth_month": fourth_month,
    "fifth_month": fifth_month,
    "sixth_month": sixth_month
  }]};
  return new Promise((resolve,reject)=>{
    axios
      .put(`${baseUrl}:${port}/achievementMeans`, body)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getEvaluationBefore(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/evaluationBefore?aim_id=1')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}