<template>
<div id="main">
<button v-on:click="getArticles">Go</button>
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
          token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbjEyMyIsImV4cCI6MTU2MDk0NjQ3MCwiaXNzIjoiZ2luLWJsb2cifQ.7DVN62IB4ekGFVWkPVqJtLQt-idQb284IwcZbP4EVCU'
        }
      })
        .then(function (res) {
          let $txt_res = JSON.stringify(res["data"]["data"]["lists"][0]["title"])
          console.log($txt_res)
        })
        .catch(function (error) {
          console.log("An error occurs:" + error)
        })
    }
  }
}
</script>
