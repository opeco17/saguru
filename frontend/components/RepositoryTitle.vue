<template>
  <a 
    :href="url"
    target="_blank"
    class="repository-title"
  >
    <slot></slot>
  </a>
</template>

<script>
export default {
  props: {
    url: String
  }
}
</script>

<style scoped>
.repository-title {
  text-decoration: none;
  color: black;
  transition: color 0.1s;
}
.repository-title:hover {
  color: #F85758;
}
</style>