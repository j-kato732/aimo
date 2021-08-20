import axios from 'axios'

//promiseの時はasync
export async function getApi(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/period')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}

export async function getInfo(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/aimoInfo/20210501')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}