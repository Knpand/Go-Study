<template>
  <main class="form-register">
    <form>
      <h1 class="h3 mb-3 fw-normal">Please Register</h1>
      <input
        v-model="name"
        class="form-control"
        placeholder="Name"
        required>
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
         @click="handle1()" >Submit</button>
    </form>
  </main>
</template>
<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'SignupForm',
  data () {
    return {
      name:'',
      email: '',
      password: ''
    }
  },
  props: {
    //LoginViewに飛ばす
    signup: {
      type: Function,
      required: true
    }
  },
  methods: {
    handle1 () {
      return this.signup({
        'user': {
          'name':this.name,
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
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
 
.form-register .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-register .form-control:focus {
  z-index: 2;
}
.form-register input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-register input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
 
</style>




