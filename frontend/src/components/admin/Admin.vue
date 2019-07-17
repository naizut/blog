<template>
<div class="admin">
  <router-link to='/admin/user'>UserAdmin</router-link>
  <router-view></router-view>
</div>
</template>

<script>
export default {
  name: 'Admin',
  data () {
    return {

    }
  }
}
</script>
