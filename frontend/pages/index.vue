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
    <v-row justify="center" class="mt-5">
      <v-col cols="12" sm="4" md="4" class="mt-3">
        <v-card class="pa-4"           elevation="0"
          outlined>
        <!-- <v-select
          v-model="temporaryInputs.languages"
          :items="languages"
          label="Languages"
        ></v-select> -->
        <!-- <v-select
          v-model="temporaryInputs.labels"
          :items="labels"
          label="Labels"
        ></v-select> -->
        <div class="grey--text text--darken-2 mb-1">Languages</div>
        <v-autocomplete
          v-model="temporaryInputs.languages"
          :items="languages"
          outlined
          chips
          small-chips
          single-line
          multiple
          clearable
        ></v-autocomplete>
        <div class="grey--text text--darken-2 mb-1">Labels</div>
        <v-autocomplete
          v-model="temporaryInputs.labels"
          :items="labels"
          outlined
          chips
          small-chips
          single-line
          multiple
          clearable
        ></v-autocomplete>
        <div class="grey--text text--darken-2 mb-1">Fork count</div>
        <v-row justify="space-between">
          <v-col
            cols="6"
            sm="6"
          >
            <v-text-field
              outlined
              dense
              single-line
              v-model="message1"
              label="Min"
            ></v-text-field>
          </v-col>
          <v-col
            cols="6"
            sm="6"
          >
            <v-text-field
              v-model="message2"
              outlined
              single-line
              dense
              clearable
              label="Max"
            ></v-text-field>
          </v-col>
        </v-row>
      
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
          @click="search"
        >
          search
        </v-btn>
        </v-card>
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
            {{ repository.name }}
          </v-card-title>
          <v-card-text class="pt-1 pb-2">
            {{ repository.description }}
          </v-card-text>
          <v-card-actions class="py-1 mb-2">
            <v-chip
              color="#f57879"
              class="mr-2"
              label
              small
              outlined
            >
              #{{ repository.language }}
            </v-chip>
            <v-chip
              color="#f57879"
              class="mr-2"
              label
              small
              outlined
            >
              <v-icon small>mdi-star-outline</v-icon>
              {{ repository.starCount }}
            </v-chip>
            <v-chip
              color="#f57879"
              class="mr-2"
              label
              small
              outlined
            >
              <v-icon small>mdi-source-fork</v-icon>
              {{ repository.forkCount }}
            </v-chip>
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
                    <span
                      :href="issue.url"
                      target="_blank"
                      class="text-decoration-none text-subtitle-1 mr-2"
                    >
                      <v-icon left>mdi-label-outline</v-icon>
                      {{ issue.title }}
                    </span>
                    <v-chip
                      v-for="label in issue.labels"
                      :key="label"
                      color="grey darken-2"
                      class="mr-2"
                      label
                      small
                      outlined
                    >
                      #{{ label }}
                    </v-chip>
                  </div>
                </div>
              </v-card-text>
            </div>
          </v-expand-transition>
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
  languages: ['Python', 'Go'],
  labels: ['good first issue', 'help wanted']
}

export default {
  data () {
    return {
      inputs: JSON.parse(JSON.stringify(inputsTemplate)),
      temporaryInputs: JSON.parse(JSON.stringify(inputsTemplate)),
      show: false,
    }
  },
  methods: {
    search (event) {
      this.inputs = JSON.parse(JSON.stringify(this.temporaryInputs))
      this.$store.commit('resetPage')
      console.log(this.getParams())
      this.$store.dispatch('fetchRepositories', this.getParams())
    },
    reset (event) {
      this.temporaryInputs = JSON.parse(JSON.stringify(inputsTemplate))
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
        languages: this.inputs.languages.includes('all') ? '' : this.inputs.languages.join(','),
        labels: this.inputs.labels.includes('all') ? '' : this.inputs.labels.join(','),
      }
    },
    openIssues () {
      
    }
  },
  created () {
    this.$store.dispatch('fetchRepositories', this.getParams())
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
