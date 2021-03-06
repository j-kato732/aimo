import axios from 'axios'

const baseUrl = "http://172.25.0.89"
const port = 9002

//promiseの時はasync
export async function getApi(
  access_token
){
  return new Promise((resolve,reject)=>{
    axios
      .get(`${baseUrl}/api/periods`,{ 
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
