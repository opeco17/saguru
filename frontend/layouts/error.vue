<template>
  <v-app>
    <div class="container">
      <img
        src="/logo.png"
        alt="logo"
        height="80"
        class="mb-8"
      >
      <h1 class="text-md-h5 font-weight-medium">
        <span v-if="error.statusCode === 404">404 | Not Found</span>
        <span v-else>An error occurred</span>
      </h1>
      <NuxtLink to="/" style="text-decoration: none;">
        <span class="text-md-h6">Home page</span>
      </NuxtLink>
    </div>
  </v-app>
</template>

<script>
export default {
  layout: 'empty',
  props: {
    error: {
      type: Object,
      default: null
    }
  },
  head () {
    const title = this.error.statusCode === 404 ? 'Page not found' : 'An error occurred'
    return {
      title: title
    }
  }
}
</script>

<style scoped>
.container {
  height: 70vh;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}
</style>