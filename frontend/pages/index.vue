<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
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
        color="success"
        @click="submit"
      >
        submit
      </v-btn>
      <v-btn
        class="mr-4"
        color="success"
        @click="previous"
        :disabled="page === 0"
      >
        previous
      </v-btn>
      <v-btn
        class="mr-4"
        color="success"
        @click="next"
        :disabled="!hasNext"
      >
        next
      </v-btn>
      <p></p>
      <p>language: {{ inputs.language }}</p>
      <p>{{ repositories }}</p>
    </v-col>
  </v-row>
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
