<template>
  <v-app>
    <v-app-bar
      fixed
      app
      color="white"
    >
      <img
        src="/logo.png"
        alt="logo"
        height="43"
        class="ml-2"
      >
      <v-spacer />
      <v-btn
        icon
        class="mr-2"
        href="https://github.com/opeco17/gitnavi"
      >
        <v-icon large class="github-icon">mdi-github</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container class="my-5">
        <Nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
}
</script>
