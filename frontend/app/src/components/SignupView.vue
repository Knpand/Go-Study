<template>
  <div class="signup-view">
    <SignupForm :signup="handleSignup" />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import SignupForm from '@/components/SignupForm.vue'

export default defineComponent({
  name: 'SignupView',
  components: {
    SignupForm
  },

  methods: {
    handleSignup (authInfo) {
        console.log("handleSignup")
        console.log(authInfo)
      return this.$store.dispatch('signup', authInfo)
        .then(() => {
          this.$router.push({ path: '/login' })
        })
        .catch(err => { throw err })
    }
  }
});
</script>

<style scoped>
.login-view {
  width: 400px;
  margin: auto;
}
</style>