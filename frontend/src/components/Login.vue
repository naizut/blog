<template>
  <div id='loginForm'>
    <form>
      <input type="text" name="username" v-model="user.username">
      <input type="password" name="password" v-model="user.password">
      <input type="submit" value="Submit">
    </form>
  </div>
</template>

<script>
new Vue({
  el: '#loginForm',
  data: {
    user: {
      username: '',
      password: ''
    }
  },
  methods: {
    submit: function(event){
      var $qs = JSON.stringify(this.user);
      console.log($qs)
    }
  }
})
export default {
  name: 'Login',
  data () {
    return {
      logo_txt: 'Dywoo'
    }
  }
}
</script>
