<template>
  <div id="loginForm" @keyup.13="userLogin">
    <div class="form-group">
      <input type="text" name="username" v-model="user.username">
      <input type="password" name="password" v-model="user.password">
      <button class="btn-login btn" @click="userLogin">Log In</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Login',
  data () {
    return {
      logo_txt: 'Dywoo',
      txt: '',
      user: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    userLogin: function (event) {
      console.log(this.user)
      axios.post('/api/auth', JSON.stringify(this.user))
        .then(function (res) {
          let $token = JSON.stringify(res)
          console.log($token)
        })
        .catch(function (error) {
          console.log('An error occurs:' + error)
        })
    }
  }
}
</script>
