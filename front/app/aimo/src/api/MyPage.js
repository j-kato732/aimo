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