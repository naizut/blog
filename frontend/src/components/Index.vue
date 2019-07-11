<template>
<div id="main">
  <div class="banner home-banner"></div>
  <div class="whatsnew">{{articleText}}</div>
</div>
</template>

<script>
export default {
  name: 'Index',
  data () {
    return {
      logo_txt: 'Duiwu',
      articleText: ''
    }
  },
  created () {
    this.getArticles()
  },
  computed: {

  },
  methods: {
    getArticles: function () {
      this.axios.get('/api/api/v1/articles')
        .then((res) => {
          let $con = JSON.stringify(res['data']['data']['lists'][0]['content'])
          this.articleText = $con
          console.log($con)
        })
        .catch(function (error) {
          console.log('An error occurs:' + error)
        })
    }
  }
}
</script>
