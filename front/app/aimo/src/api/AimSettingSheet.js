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