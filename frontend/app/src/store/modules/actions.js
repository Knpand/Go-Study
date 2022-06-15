import auth from '@/api/auth'
import * as types from './mutation-types'

export default {
  login({ commit }, data) {
    console.log("login")
    console.log(data.user)
    return auth.login(data.user)
      .then((res) => {
        localStorage.setItem('token', res.data.token)
        commit(types.LOGIN, res.data)
      })
      .catch(error => { throw error })
  },
  logout({ commit }) {
    return auth.logout()
      .then(() => {
        localStorage.removeItem('token')
        commit(types.LOGOUT, { token: null, userId: null })
      })
      .catch(error => { throw error })
  },
  signup({ commit },data) {
    console.log("signup")
    console.log(commit)
    console.log(data.user)
    return auth.signup(data.user)
      .then((res) => res)
      .catch(error => { throw error })
  },
}