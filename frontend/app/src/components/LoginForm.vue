<template>
  <main class="form-login">
    <form>
      <h1 class="h3 mb-3 fw-normal">Login</h1>
      <input
        v-model="email"
        type="email"
        class="form-control"
        placeholder="Email"
        required>
 
      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="Password"
        required>
 
      <button
        class="w-100 btn btn-lg btn-primary"
        @click="handle()">Login</button>
    </form>
  </main>
</template>
<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'LoginForm',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  props: {
    //LoginViewに飛ばす
    login: {
      type: Function,
      required: true
    }
  },
  methods: {
    handle () {
      return this.login({
        'user': {
          'email': this.email,
          'password': this.password,
        }
      })
      .catch(err => { throw err })
    }
  }
});
</script>
<style>
.form-login {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
 
.form-login .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-login .form-control:focus {
  z-index: 2;
}
.form-login input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-login input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
 
</style>

