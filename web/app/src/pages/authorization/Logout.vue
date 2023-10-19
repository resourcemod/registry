<template>
  <div></div>
</template>
<script lang="ts">
import Cookies from 'js-cookie'
export default {
  mounted() {
    Cookies.remove('rmod_auth')
    location.reload()
  }
}
</script>