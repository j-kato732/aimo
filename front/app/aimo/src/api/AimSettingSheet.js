import axios from 'axios'

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
      .get('http://localhost:8000/departmentGoal?period=202105&department=%E3%82%BD%E3%83%AA%E3%83%A5%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E6%9C%AC%E9%83%A8%20%E9%96%8B%E7%99%BA%EF%BC%91%E9%83%A8')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}