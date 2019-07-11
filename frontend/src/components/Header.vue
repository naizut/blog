<template>
<header id="header">
  <div class="navbar-brand">
    <router-link to="/">
      {{logo_txt}}
    </router-link>
  </div>
  <div class="navbar-navlist">
    <ul class="nav">
      <li>
        <router-link to="/inspire">
          <a class="btn-link">Inspire</a>
        </router-link>
      </li>
      <li>
        <router-link to="/tech">
          <a class="btn-link">Devlopment</a>
        </router-link>
      </li>
      <li>
        <router-link to="/contact">
          <a class="btn-link">Contact</a>
        </router-link>
      </li>
    </ul>
  </div>
  <div class="navbar-status">
    <ul class="nav">
      <li>
        <router-link to="/login">
          <a class="btn-link nav-login">Login</a>
        </router-link>
      </li>
    </ul>
  </div>
</header>
</template>

<script>
export default {
  name: 'Header',
  data () {
    return {
      logo_txt: "Kai's Blog"
    }
  }
}
</script>
