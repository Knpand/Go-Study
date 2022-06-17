<template>
  
  <div class="container">
  <font size="7" style="margin-top:40px; color:white">{{ message }}</font>


<div class="row" style="margin-top:45px; ">
    <div class="flex md6 lg4" >
      <img  v-bind:src="image13"  style="margin-right:15px; width:330px" />
    </div>
   <div class="flex md6 lg4" >
      <img  v-bind:src="image9"  style="margin-right:15px; margin-top:20px; width:330px" />
    </div>
       <div class="flex md6 lg4" >
      <img  v-bind:src="image10"  style="margin-right:15px; width:330px" />
    </div>
       <div class="flex md6 lg4" >
      <img  v-bind:src="image7"  style="margin-right:15px; width:330px" />
    </div>
       <div class="flex md6 lg4" >
      <img  v-bind:src="image4"  style="margin-right:15px; margin-top:20px; width:330px" />
    </div>
       <div class="flex md6 lg4" >
      <img  v-bind:src="image3"  style="margin-right:15px; width:330px" />
    </div>
       <div class="flex md6 lg4" >
      <img  v-bind:src="image2"  style="margin-right:15px; width:330px" />
    </div>
      <div class="flex md6 lg4" >
      <img  v-bind:src="image12"  style="margin-right:15px; margin-top:20px; width:330px" />
    </div>
      <div class="flex md6 lg4" >
      <img  v-bind:src="image11"  style="margin-right:15px; width:330px" />
    </div>
</div>


  </div>
</template>

<script>
import { defineComponent } from 'vue';
import axios from 'axios'

export default defineComponent({
  name: 'TopView',
  data () {
    return {
      image1:"image/a1.jpg",
      image2:"image/a2.jpg",
      image3:"image/a3.jpg",
      image4:"image/a4.jpg",
      image5:"image/a5.jpg",
      image6:"image/a6.jpg",
      image7:"image/a7.jpg",
      image8:"image/a8.jpg",
      image9:"image/a9.jpg",
      image10:"image/a10.jpg",
      image11:"image/a11.jpg",
      image12:"image/a12.jpg",
      image13:"image/a13.jpg",



      images:[],
      message: "You are not logged in!",
      logo:'',
    }
  },
  methods: {
    logout () {
      return this.$store.dispatch('logout')
        .then(() => {
          this.$router.push('/login')
        })
        .catch(error => { throw error })
    },

  },
  mounted:function(){
    // user情報を取得
    // ログイン情報は、Cookieに保存してあるので、
    // リクエストするだけでOK
  
    this.logo = require('../assets/images/a1.jpg')
    // for (let i = 1; i < 18; i++) {
    //     str='../assets/images/a'+String(i)+'.jpg'
    //     images[i-1] = require(str)
    // }
    axios.get(`http://localhost:5050/jwt`,{ withCredentials: true })
      .then(res => {
        console.log("res check")
        console.log(res)
        this.message = `${res.data.name} Photo Gallery`
    })
      .catch(error => { throw error })

  }

});
</script>

<style>

    body{/* 45度に傾いたグラデーションで、赤から黒に */

    background:linear-gradient(60deg, darkgray,slategray );
    }


</style>