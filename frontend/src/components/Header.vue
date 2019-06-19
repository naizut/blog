<template>
<header id="header">
<div id="logo">{{logo_txt}}</div>
<div>
<div class="nav"></div>
</div>
</header>
</template>

<script>
export default {
  name: 'Header',
  data () {
    return {
      logo_txt: 'Dywoo'
    }
  }
}
</script>
