<template>
<div id="main">
<button type="button" v-on:click="getArticles">Go</button>
</div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Index',
  data () {
    return {
      logo_txt: 'Duiwu'
    }
  },
  methods: {
    getArticles: function () {
      axios.get('/api/api/v1/articles', {
        params: {
          token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbjEyMyIsImV4cCI6MTU2MjU4MDUyOSwiaXNzIjoiZ2luLWJsb2cifQ.xkLvAEhMGDP9SRKS1Nc87C_1cXgohFvW4IAikrbC9Ck'
        }
      })
        .then(function (res) {
          let $txtRes = JSON.stringify(res['data']['data'])
          console.log($txtRes)
        })
        .catch(function (error) {
          console.log('An error occurs:' + error)
        })
    }
  }
}
</script>
