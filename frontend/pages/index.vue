<template>
  <div>
    <v-row justify="center" align-content="center" class="my-2 mx-1">
      <h1 class="text-md-h3 text-xs-h4 font-weight-medium">
        Grab a favorite issue and commit there
      </h1>
    </v-row>
    <v-row justify="center" align-content="center" class="my-2 mx-1">
      <h2 class="text-md-h5 text-xs-h6 grey--text text--darken-2 font-weight-medium">
        Explore GitHub issues using flexible query with 
        <a
          style="color: #F85758; text-decoration: none;"
          href="https://github.com/opeco17/saguru"
          target="_blank"
        >
          saguru
        </a>
      </h2>
    </v-row>
    <v-row style="height: 10px;">
      <v-progress-linear
        indeterminate
        color="#F85758"
        v-show="initLoading"
      ></v-progress-linear>
    </v-row>
    <v-row justify="center" class="mt-5">
      <v-col cols="12" sm="4" md="4" class="mt-3">
        <v-form
          ref="form"
          v-model="valid"
          lazy-validation
        >
          <v-card
            class="pa-4"
            elevation="0"
            outlined
          >
            <form-label label="Languages" />
            <multiple-chips-select
              v-model="temporaryInputs.languages" 
              :items="languages" 
              @close="removeLanguage"
            />
            <form-label label="Labels" />
            <multiple-chips-select
              v-model="temporaryInputs.labels" 
              :items="labels" 
              @close="removeLabel" 
            />
            <form-label label="Star count" />
            <v-row justify="space-between">
              <v-col cols="6" sm="6">
                <single-integer-field v-model="temporaryInputs.star_count_lower" label="Min" />
              </v-col>
              <v-col cols="6" sm="6">
                <single-integer-field v-model="temporaryInputs.star_count_upper" label="Max" />
              </v-col>
            </v-row>
            <form-label label="Order by" />
            <orderby-select v-model="temporaryInputs.ordermetric" :items="ordermetrics" />
            <v-row justify="end" class="mb-1 mr-1">
              <v-btn @click="showDetail=!showDetail" text small class="px-1">
                detail
                <v-icon>{{ showDetail ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
              </v-btn>
            </v-row>
            <v-divider></v-divider>
            <v-expand-transition>
              <div v-show="showDetail" class="mt-3">
                <form-label label="Fork count" />
                <v-row justify="space-between">
                  <v-col cols="6" sm="6">
                    <single-integer-field 
                      v-model="temporaryInputs.fork_count_lower" 
                      label="Min"
                    />
                  </v-col>
                  <v-col cols="6" sm="6">
                    <single-integer-field 
                      v-model="temporaryInputs.fork_count_upper" 
                      label="Max"
                    />
                  </v-col>
                </v-row>
                <form-label label="License" />
                <single-text-select v-model="temporaryInputs.license" :items="licenses"/>
                <form-label label="Assign status" />
                <single-text-select v-model="temporaryInputs.assigned" :items="assignStatuses"/>
              </div>
            </v-expand-transition>
            <v-row justify="center" class="mb-1 mt-4">
              <v-btn
                class="mr-4"
                color="#F85758"
                outlined
                @click="reset"
              >
                reset
              </v-btn>
              <basic-button @click="search">
                search
                <basic-button-circular v-show="searchLoading" />
              </basic-button>
            </v-row>
          </v-card>
        </v-form>
      </v-col>
      <v-col cols="12" sm="8" md="8">
        <v-card
          class="my-3"
          elevation="0"
          outlined
          v-for="repository in repositories"
          :key="repository.id"
        >
          <v-card-title class="py-2">
            <repository-title :url="repository.url">{{ repository.name }}</repository-title>
          </v-card-title>
          <v-card-text class="pt-1 pb-2">
            {{ repository.description }}
          </v-card-text>
          <v-card-actions class="py-1 mb-2">
            <repository-chip>
              #{{ repository.language }}
            </repository-chip>
            <repository-chip>
              <v-icon small>mdi-star-outline</v-icon>
              {{ repository.starCount | killo }}
            </repository-chip>
            <repository-chip class="hidden-xs-only">
              <v-icon small>mdi-source-fork</v-icon>
              {{ repository.forkCount | killo }}
            </repository-chip>
            <v-spacer></v-spacer>
            <v-btn
              @click="repository.show = !repository.show"
              outlined
              small
              color="purple"
            >
              {{ repository.issues.length }} Issues
              <v-icon>{{ repository.show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
            </v-btn>
          </v-card-actions>
          <v-expand-transition>
            <div v-show="repository.show">
              <v-card-text>
                <div
                  v-for="issue in repository.issues"
                  :key="issue.id"
                >
                  <v-divider></v-divider>
                  <div class="py-3">
                    <v-icon left>mdi-label-outline</v-icon>
                      <issue-title :url="issue.url">{{ issue.title }}</issue-title>
                    <issue-chip v-for="label in issue.labels" :key="label">
                      #{{ label }}
                    </issue-chip>
                  </div>
                </div>
              </v-card-text>
            </div>
          </v-expand-transition>
        </v-card>
        <v-row justify="center" class="mt-4">
          <basic-button @click="showmore" :disabled="!hasNext" v-show="!initLoading">
            <v-icon>mdi-chevron-down</v-icon>
            show more
            <basic-button-circular v-show="showmoreLoading" />
          </basic-button>
        </v-row>
      </v-col>
    </v-row>
  </div>
</template>

<script>
let inputsTemplate = {
  languages: ['all'],
  labels: ['good first issue'],
  star_count_lower: '',
  star_count_upper: '',
  fork_count_lower: '',
  fork_count_upper: '',
  license: 'all',
  assigned: 'all',
  ordermetric: 'star_count_desc',
}

export default {
  data () {
    return {
      inputs: JSON.parse(JSON.stringify(inputsTemplate)),
      temporaryInputs: JSON.parse(JSON.stringify(inputsTemplate)),
      valid: true,
      show: false,
      showDetail: false,
    }
  },
  created () {
    this.$store.dispatch('fetchRepositories', { params: this.getParams(), type: 'init'})
    this.$store.dispatch('fetchLanguages')
    this.$store.dispatch('fetchLicenses')
    this.$store.dispatch('fetchLabels')
    this.$store.dispatch('fetchOrdermetrics')
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
    },
    assignStatuses() {
      return this.$store.state.assignStatuses
    },
    ordermetrics() {
      return this.$store.state.ordermetrics
    },
    initLoading() {
      return this.$store.state.initLoading
    },
    searchLoading() {
      return this.$store.state.searchLoading
    },
    showmoreLoading() {
      return this.$store.state.showmoreLoading
    },
  },
  methods: {
    search (event) {
      this.$refs.form.validate()
      if (!this.valid) {
        return
      }
      this.inputs = JSON.parse(JSON.stringify(this.temporaryInputs))
      this.$store.commit('resetPage')
      this.$store.dispatch('fetchRepositories', { params: this.getParams(), type: 'search' })
    },
    reset (event) {
      this.temporaryInputs = JSON.parse(JSON.stringify(inputsTemplate))
    },
    showmore (event) {
      this.$store.commit('incrementPage')
      this.$store.dispatch('fetchRepositories', { params: this.getParams(), type: 'showmore' })
    },
    getParams () {
      let params = {}
      params.page = this.page
      if (this.inputs.languages !== '' && !this.inputs.languages.includes('all')) params.languages = this.inputs.languages.join(',')
      if (this.inputs.labels !== '' && !this.inputs.labels.includes('all')) params.labels = this.inputs.labels.join(',')
      if (this.inputs.star_count_lower !== '') params.star_count_lower = this.inputs.star_count_lower
      if (this.inputs.star_count_upper !== '') params.star_count_upper = this.inputs.star_count_upper
      if (this.inputs.fork_count_lower !== '') params.fork_count_lower = this.inputs.fork_count_lower
      if (this.inputs.fork_count_upper !== '') params.fork_count_upper = this.inputs.fork_count_upper
      if (this.inputs.license !== '' && this.inputs.license !== 'all') params.license = this.inputs.license
      if (this.inputs.ordermetric !== '') params.orderby = this.inputs.ordermetric
      if (this.inputs.assigned == 'assigned') {
        params.assigned = true
      } else if (this.inputs.assigned == 'unassigned') {
        params.assigned = false
      }
      return params
    },
    removeLanguage(language) {
      const index = this.temporaryInputs.languages.indexOf(language)
      if (index >= 0) this.temporaryInputs.languages.splice(index, 1)
    },
    removeLabel(label) {
      const index = this.temporaryInputs.labels.indexOf(label)
      if (index >= 0) this.temporaryInputs.labels.splice(index, 1)
    },
  },
  filters: {
    killo (value) {
      return value > 999 ? (value / 1000).toFixed(1) + 'k' : value
    }
  }
}
</script>
