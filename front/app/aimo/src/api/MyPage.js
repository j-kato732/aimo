import axios from 'axios'

const baseUrl = "http://localhost"
const port = 9002

//promiseの時はasync
export async function getApi(
  access_token
){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}:${port}/periods`,{ 
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

export async function getInfo(){
  return new Promise((resolve,reject)=>{
    axios
      .get('http://localhost:8000/aimoInfos/20210501')
      .then(res =>{
        resolve(res.data);
      })
      .catch(err =>{
        reject(err);
    })
  })
}
