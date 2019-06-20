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
      axios.get('/api/v1/articles', {
        params: {
          token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbjEyMyIsImV4cCI6MTU2MTAwNDUyNywiaXNzIjoiZ2luLWJsb2cifQ.CnoWN6OYsN-tXbCQ-h5R1xSrlJKYxrgLzENwgnXIJ0w'
        }
      })
        .then(function (res) {
          let $txt_res = JSON.stringify(res['data']['data'])
          console.log($txt_res)
        })
        .catch(function (error) {
          console.log('An error occurs:' + error)
        })
    }
  }
}
</script>
