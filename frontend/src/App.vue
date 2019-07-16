<template>
  <div id="app">
    <div v-if="headerShow">
      <Header></Header>
    </div>
    <router-view/>
  </div>

</template>

<script>
import Header from '@/components/Header'

export default {
  name: 'App',
  data () {
    return {
      path: ''
    }
  },
  components: {
    Header
  },
  computed: {
    headerShow: function () {
      return !(this.path === '/admin')
    }
  },
  // admin页面中取消引入Header
  mounted () {
    this.path = this.$route.path
  },
  watch: {
    $route (to, from) {
      this.path = to.path
    }
  }
}
</script>

<style lang="scss">
@import './assets/css/global.scss';

</style>
