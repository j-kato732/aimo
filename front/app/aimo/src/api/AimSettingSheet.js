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