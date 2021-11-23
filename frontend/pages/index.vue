<template>
  <div>
    <v-row justify="center" align-content="center" class="my-2">
      <h1 class="text-md-h3 text-xs-h4 font-weight-medium">
        Grab a favorite issue and commit there
      </h1>
    </v-row>
    <v-row justify="center" align-content="center" class="my-2">
      <h2 class="text-md-h5 text-xs-h6 grey--text text--darken-2 font-weight-medium">
        Explore GitHub issues using flexible query with <span style="color: #F85758">saguru</span>
      </h2>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" sm="4" md="4">
        <v-select
          v-model="temporaryInputs.language"
          :items="languages"
          label="Language"
        ></v-select>
        <v-select
          v-model="temporaryInputs.labels"
          :items="labels"
          label="Labels"
        ></v-select>
        <v-btn
          class="mr-4"
          color="#F85758"
          outlined
          @click="reset"
        >
          reset
        </v-btn>
        <v-btn
          class="mr-4 white--text"
          color="#F85758"
          depressed
          @click="submit"
        >
          submit
        </v-btn>
      </v-col>
      <v-col cols="12" sm="8" md="8">
        <v-card
          class="my-3 pa-3"
          elevation="0"
          outlined
          v-for="repository in repositories"
          :key="repository.id"
        >
          {{ repository.name }}
        </v-card>
        <v-btn
          class="mr-4 white--text"
          color="#F85758"
          depressed
          @click="previous"
          :disabled="page === 0"
        >
          <v-icon>mdi-chevron-left</v-icon>
          previous
        </v-btn>
        <v-btn
          class="mr-4 white--text"
          color="#F85758"
          depressed
          @click="next"
          :disabled="!hasNext"
        >
          next
          <v-icon>mdi-chevron-right</v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// define class and constructor?
let inputsTemplate = {
  language: '',
  labels: ''
}

export default {
  data () {
    return {
      inputs: JSON.parse(JSON.stringify(inputsTemplate)),
      temporaryInputs: JSON.parse(JSON.stringify(inputsTemplate)),
    }
  },
  methods: {
    submit (event) {
      this.inputs = JSON.parse(JSON.stringify(this.temporaryInputs))
      this.$store.commit('resetPage')
      this.$store.dispatch('fetchRepositories', this.getParams())
    },
    reset (event) {
      // reset inputs
    },
    previous (event) {
      this.$store.commit('decrimentPage')
      this.$store.dispatch('fetchRepositories', this.getParams())
    },
    next (event) {
      this.$store.commit('incrementPage')
      this.$store.dispatch('fetchRepositories', this.getParams())
    },
    getParams () {
      return {
        page: this.page,
        language: this.inputs.language === 'all' ? '' : this.inputs.language,
        labels: this.inputs.labels === 'all' ? '' : this.inputs.labels,
      }
    },
  },
  created () {
    this.$store.dispatch('fetchRepositories', {page: 0})
    this.$store.dispatch('fetchLanguages')
    this.$store.dispatch('fetchLicenses')
    this.$store.dispatch('fetchLabels')
  },
  computed: {
    page() {
      return this.$store.state.page
    },
    hasNext() {
      return this.$store.state.hasNext
    },
    repositories() {
      return this.$store.state.repositories
    },
    languages() {
      return this.$store.state.languages
    },
    licenses() {
      return this.$store.state.licenses
    },
    labels() {
      return this.$store.state.labels
    }
  }
}
</script>
