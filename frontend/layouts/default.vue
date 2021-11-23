<template>
  <v-app>
    <v-app-bar
      fixed
      app
      color="white"
    >
      <img
        src="/logo.png"
        alt="Vuetify.js"
        height="43"
        class="ml-2"
      >
      <v-spacer />
      <v-btn
        icon
        class="mr-2"
        href="https://github.com/opeco17/saguru"
      >
        <v-icon large>mdi-github</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container class="my-5">
        <Nuxt />
      </v-container>
    </v-main>
    <!-- <v-footer
      :absolute="true"
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer> -->
  </v-app>
</template>

<script>
export default {
  data () {
    return {
    }
  }
}
</script>
