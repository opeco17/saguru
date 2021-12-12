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
        :class="{ 'logo': !xs, 'ml-2': !xs, 'logo-xs': xs, 'ml-1': xs }"
      >
      <v-spacer />

      <v-menu offset-y transition="slide-y-transition" :open-on-hover="!mobile">
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            :class="{ 'mr-4': !xs, 'mr-2': xs }"
            v-bind="attrs"
            v-on="on"
          >
            <v-icon :class="{ 'translate-icon': !xs, 'translate-icon-xs': xs }">mdi-translate</v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-item
            v-for="(language, index) in languages"
            :key="index"
          >
            <v-list-item-title>
              <NuxtLink :to="switchLocalePath(language.code)" style="text-decoration: none; color: black;">
                <span color="black">{{ language.label }}</span>
              </NuxtLink>
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

    <v-menu>
      <template v-slot:activator="{ on: menu, attrs }">
        <v-tooltip bottom :disabled="mobile">
          <template v-slot:activator="{ on: tooltip }">
            <v-btn
              icon
              :class="{ 'mr-2': !xs, 'mr-1': xs }"
              href="https://github.com/opeco17/gitnavi"
              v-bind="attrs"
              v-on="{ ...tooltip, ...menu }"
            >
              <v-icon :class="{ 'github-icon': !xs, 'github-icon-xs': xs }">mdi-github</v-icon>
            </v-btn>
          </template>
          <span>{{ $t('githubIconMessage') }}</span>
        </v-tooltip>
      </template>
    </v-menu>

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
  computed: {
    xs () {
      return this.$vuetify.breakpoint.xs
    },
    mobile () {
      return this.$vuetify.mobile
    },
    languages () {
      return [
        { label: 'English', code: 'en' },
        { label: 'Japanese', code: 'ja' },
      ]
    }
  }
}
</script>

<style scoped>
.logo {
  height: 43px;
}
.logo-xs {
  height: 40px;
}
.translate-icon {
  font-size: 31px !important;
}
.translate-icon-xs {
  font-size: 28px !important;
}
.github-icon {
  font-size: 40px !important;
}
.github-icon-xs {
  font-size: 37px !important;
}
</style>