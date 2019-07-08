<template>
<header id="header">
  <div class="navbar-brand">
    <router-link to="/">
      {{logo_txt}}
    </router-link>
  </div>
  <div class="navbar-navlist">
    <ul class="nav">
      <li>Inspire</li>
      <li>Design</li>
      <li>Development</li>
      <li>Contact</li>
    </ul>
    <router-link to="/login">
      <a class="nav-login">Login</a>
    </router-link>
  </div>
</header>
</template>

<script>
export default {
  name: 'Header',
  data () {
    return {
      logo_txt: "Kenny's Zion"
    }
  }
}
</script>
