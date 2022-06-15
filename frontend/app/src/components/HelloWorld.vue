<template>
<h3>Hello world</h3>
{{response}}
</template>

<script>
import axios from 'axios';
const URL = 'http://127.0.0.1:5050/'
export default {
  data(){
    return{
      response: "",
    }
  },
  mounted(){
    axios.get(`${URL}tasks`).then(response =>{
        
        this.response = response
        
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
