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

    return axios.request(config)
      .then(res => {
        console.log("res check")
        console.log(res)
        return res
    })
      .catch(error => { throw error })
  },
  signup: (authInfo) => {
    console.log("signup check")
    console.log(authInfo)
    config.method = 'post'
    config.data = authInfo
    config.url = 'http://localhost:5050/signup'

    return axios.request(config)
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