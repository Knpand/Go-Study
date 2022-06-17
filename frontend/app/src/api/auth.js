import axios from 'axios'

const headers = {}
headers['Content-type'] = 'application/json'

const config = {
  method: null,
  url: 'http://localhost:5050', // APIサーバー
  headers,
  data: null
}

export default {
  login: (authInfo) => {
    console.log("log check")
    console.log(authInfo)
    config.method = 'post'
    config.data = authInfo
    var params = new URLSearchParams();
    params.append("email",authInfo.email)
    params.append("password",authInfo.password)
    return axios.post(`http://localhost:5050`,params,{ withCredentials: true })
      .then(res => {
        console.log("res check")
        console.log(res)
        return res
    })
      .catch(error => { throw error })
  },
  signup: (authInfo) => {
    console.log("signup check1")
    console.log(authInfo.name)
    var params = new URLSearchParams();
    params.append("name",authInfo.name)
    params.append("email",authInfo.email)
    params.append("password",authInfo.password)
    console.log(params)


    return axios.post(`http://localhost:5050/signup`,params,{ withCredentials: true })
      .then(res => {
        console.log("res check")
        console.log(res)
        return res
    })
      .catch(error => { throw error })
  },


  logout: () => {
    config.method = 'post'
    return axios.request(config)
      .then(res => res)
      .catch(error => { throw error })
  }
}