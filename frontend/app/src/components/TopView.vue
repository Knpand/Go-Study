<template>
  <h1>Top page</h1>
  <div class="container">
  <h1>{{ message }}</h1>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import axios from 'axios'
export default defineComponent({
  name: 'TopView',
  data () {
    return {
      message: "You are not logged in!"
    }
  },
  methods: {
    logout () {
      return this.$store.dispatch('logout')
        .then(() => {
          this.$router.push('/login')
        })
        .catch(error => { throw error })
    }
  },
  mounted:function(){
    // user情報を取得
    // ログイン情報は、Cookieに保存してあるので、
    // リクエストするだけでOK
    axios.get(`http://localhost:5050/jwt`,{ withCredentials: true })
      .then(res => {
        console.log("res check")
        console.log(res)
        this.message = `Hi ${res.name} `
    })
      .catch(error => { throw error })

  }

});
</script>