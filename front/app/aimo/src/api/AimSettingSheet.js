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
      .get('http://localhost:8000/departmentGoal?period=202105&departmentId=2')
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
      .get('http://localhost:8000/role?period=202105&departmentId=2&jobId=6')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getAims(access_token){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/aims?period=202105&userId=1`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      })
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function postAim(
  period,
  user_id,
  aim_item,
  achievement_level,
  achievement_weight,
  achievement_difficulty_before,
  aim_number,
  access_token
){
  const body = {
    "period": period,
    "user_id": user_id,
    "aim_item": aim_item,
    "achievement_level": achievement_level,
    "achievement_weight": achievement_weight,
    "achievement_difficulty_before": achievement_difficulty_before,
    "aim_number": aim_number
  };
  return new Promise((resolve,reject)=>{
    axios
      .post(`${baseUrl}:${port}/aim`, 
      body,
      { headers: {
        Authorization: `Bearer ${access_token}`,
      },})
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
export async function getAchievementMeans(aim_number, access_token){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/achievementMeans?period=202105&userId=1&aimNumber=${aim_number}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      })
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getAchievementMean(aim_number, achievement_mean_number, access_token){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/achievementMean?period=202105&userId=1&aimNumber=${aim_number}&achievementMeanNumber=${achievement_mean_number}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      })
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
  period1,
  user_id1,
  aim_number1,
  achievement_mean_number1,
  achievement_mean1,
  first_month1,
  second_month1,
  third_month1,
  fourth_month1,
  fifth_month1,
  sixth_month1,

  period2,
  user_id2,
  aim_number2,
  achievement_mean_number2,
  achievement_mean2,
  first_month2,
  second_month2,
  third_month2,
  fourth_month2,
  fifth_month2,
  sixth_month2,

  period3,
  user_id3,
  aim_number3,
  achievement_mean_number3,
  achievement_mean3,
  first_month3,
  second_month3,
  third_month3,
  fourth_month3,
  fifth_month3,
  sixth_month3,

  period4,
  user_id4,
  aim_number4,
  achievement_mean_number4,
  achievement_mean4,
  first_month4,
  second_month4,
  third_month4,
  fourth_month4,
  fifth_month4,
  sixth_month4,

  period5,
  user_id5,
  aim_number5,
  achievement_mean_number5,
  achievement_mean5,
  first_month5,
  second_month5,
  third_month5,
  fourth_month5,
  fifth_month5,
  sixth_month5,

  period6,
  user_id6,
  aim_number6,
  achievement_mean_number6,
  achievement_mean6,
  first_month6,
  second_month6,
  third_month6,
  fourth_month6,
  fifth_month6,
  sixth_month6,
){
  const body = {achievementMeans:[{
    "period": period1,
    "user_id": user_id1,
    "aim_number": aim_number1,
    "achievement_mean_number": achievement_mean_number1,
    "achievement_mean": achievement_mean1,
    "first_month": first_month1,
    "second_month": second_month1,
    "third_month": third_month1,
    "fourth_month": fourth_month1,
    "fifth_month": fifth_month1,
    "sixth_month": sixth_month1
  },
  {
    "period": period2,
    "user_id": user_id2,
    "aim_number": aim_number2,
    "achievement_mean_number": achievement_mean_number2,
    "achievement_mean": achievement_mean2,
    "first_month": first_month2,
    "second_month": second_month2,
    "third_month": third_month2,
    "fourth_month": fourth_month2,
    "fifth_month": fifth_month2,
    "sixth_month": sixth_month2
  },
  {
    "period": period3,
    "user_id": user_id3,
    "aim_number": aim_number3,
    "achievement_mean_number": achievement_mean_number3,
    "achievement_mean": achievement_mean3,
    "first_month": first_month3,
    "second_month": second_month3,
    "third_month": third_month3,
    "fourth_month": fourth_month3,
    "fifth_month": fifth_month3,
    "sixth_month": sixth_month3
  },
  {
    "period": period4,
    "user_id": user_id4,
    "aim_number": aim_number4,
    "achievement_mean_number": achievement_mean_number4,
    "achievement_mean": achievement_mean4,
    "first_month": first_month4,
    "second_month": second_month4,
    "third_month": third_month4,
    "fourth_month": fourth_month4,
    "fifth_month": fifth_month4,
    "sixth_month": sixth_month4
  },
  {
    "period": period5,
    "user_id": user_id5,
    "aim_number": aim_number5,
    "achievement_mean_number": achievement_mean_number5,
    "achievement_mean": achievement_mean5,
    "first_month": first_month5,
    "second_month": second_month5,
    "third_month": third_month5,
    "fourth_month": fourth_month5,
    "fifth_month": fifth_month5,
    "sixth_month": sixth_month5
  },
  {
    "period": period6,
    "user_id": user_id6,
    "aim_number": aim_number6,
    "achievement_mean_number": achievement_mean_number6,
    "achievement_mean": achievement_mean6,
    "first_month": first_month6,
    "second_month": second_month6,
    "third_month": third_month6,
    "fourth_month": fourth_month6,
    "fifth_month": fifth_month6,
    "sixth_month": sixth_month6,
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

export async function postAchievementMean(
  period,
  user_id,
  aim_number,
  achievement_mean_number,
  achievement_mean,
  first_month,
  second_month,
  third_month,
  fourth_month,
  fifth_month,
  sixth_month,
){
  const body = {
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
  };
  return new Promise((resolve,reject)=>{
    axios
      .post(`${baseUrl}:${port}/achievementMean`, body)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}


export async function putAchievementMean(
  id,
  period,
  user_id,
  aim_number,
  achievement_mean_number,
  achievement_mean,
  first_month,
  second_month,
  third_month,
  fourth_month,
  fifth_month,
  sixth_month,
  ){
  const body = {
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
    "sixth_month": sixth_month,
  };
  console.log(body);
  return new Promise((resolve,reject)=>{
    axios
      .put(`${baseUrl}:${port}/achievementMean`, body)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function putAim(
  id,
  period,
  userId,
  aimItem,
  achievementLevel,
  achievementWeight,
  achievementDifficultyBefore,
  aimNumber,
  access_token
  ){
  const body = {
    "id": id,
    "period": period,
    "userId": userId,
    "aimItem": aimItem,
    "achievementLevel": achievementLevel,
    "achievementWeight": achievementWeight,
    "achievementDifficultyBefore": achievementDifficultyBefore,
    "aimNumber": aimNumber,
  };
  return new Promise((resolve,reject)=>{
    axios
      .put(`${baseUrl}:${port}/aim`, 
      body,{
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      })
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getEvaluationBefore(aim_id, evaluatorNumber, access_token){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/evaluationBefore?aimId=${aim_id}&evaluatorNumber=${evaluatorNumber}`, {
      headers: {
        Authorization: `Bearer ${access_token}`,
        },
      })
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function postEvaluationBefore(
  aimId,
  comment,
  evaluatorNumber,
  commentUserId
){
  const body = {
    "aimId": aimId,
    "comment": comment,
    "evaluatorNumber": evaluatorNumber,
    "commentUserId": commentUserId
  };
  return new Promise((resolve,reject)=>{
    axios
      .post(`${baseUrl}:${port}/evaluationBefore`, body)
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}