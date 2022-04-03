import axios from 'axios'

const baseUrl = "http://172.25.0.89"
const port = 9002

export async function postUser(
  authId,
  period,
  lastName,
  firstName,
  departmentId,
  jobId,
  enrollmentflg,
  adminFlg,
  access_token
){
  const body = {
    "authId": authId, 
    "period": period,
    "lastName": lastName,
    "firstName": firstName,
    "departmentId": departmentId,
    "jobId": jobId,
    "enrollmentflg": enrollmentflg,
    "adminFlg": adminFlg
  };
  return new Promise((resolve,reject)=>{
    axios
      .post(`${baseUrl}:${port}/user`, 
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