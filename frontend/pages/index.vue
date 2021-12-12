<template>
  <div>
    <v-row justify="center" align-content="center" class="my-1 mx-1">
      <h1 class="text-md-h3 text-sm-h4 font-weight-medium">
        {{ $t('title') }}
      </h1>
    </v-row>
    <v-row justify="center" align-content="center" class="my-2 mx-1">
      <h3 class="text-md-h5 text-sm-h6 grey--text text--darken-2 font-weight-medium"> 
        <a
          style="color: #F85758; text-decoration: none;"
          href="https://github.com/opeco17/gitnavi"
          target="_blank"
        >
          gitnavi
        </a>
        {{ $t('subTitle') }}
      </h3>
    </v-row>
    <v-row style="height: 8px;">
      <v-progress-linear
        indeterminate
        color="#F85758"
        v-show="initLoading || searchLoading"
      ></v-progress-linear>
    </v-row>
    <v-row justify="center" class="mt-1">
      <v-col cols="12" sm="5" md="4" class="mt-3">
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
            <form-label :label="$t('languageLabel')" />
            <multiple-chips-complete
              v-model="temporaryInputs.languages" 
              :items="formatedLanguages" 
              @close="removeLanguage"
            />
            <form-label :label="$t('labelLabel')" />
            <multiple-chips-complete
              v-model="temporaryInputs.labels" 
              :items="formatedLabels" 
              @close="removeLabel"
            />
            <form-label :label="$t('assignStatusLabel')" />
            <single-text-select v-model="temporaryInputs.assigned" :items="formatedAssignStatuses"/>
            <form-label :label="$t('orderByLabel')" />
            <single-text-select v-model="temporaryInputs.ordermetric" :items="formatedOrderMetrics"/>
            <v-row justify="end" class="mb-1 mr-1">
              <v-btn @click="showDetail=!showDetail" text small class="px-1">
                {{ $t('detailLabel') }}
                <v-icon>{{ showDetail ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
              </v-btn>
            </v-row>
            <v-divider></v-divider>
            <v-expand-transition>
              <div v-show="showDetail" class="mt-3">
                <form-label :label="$t('starCountLabel')" />
                <v-row justify="space-between">
                  <v-col cols="6" sm="6">
                    <single-integer-field v-model="temporaryInputs.star_count_lower" :label="$t('min')" />
                  </v-col>
                  <v-col cols="6" sm="6">
                    <single-integer-field v-model="temporaryInputs.star_count_upper" :label="$t('max')" />
                  </v-col>
                </v-row>
                <form-label :label="$t('forkCountLabel')" />
                <v-row justify="space-between">
                  <v-col cols="6" sm="6">
                    <single-integer-field 
                      v-model="temporaryInputs.fork_count_lower" 
                      :label="$t('min')"
                    />
                  </v-col>
                  <v-col cols="6" sm="6">
                    <single-integer-field 
                      v-model="temporaryInputs.fork_count_upper" 
                      :label="$t('max')"
                    />
                  </v-col>
                </v-row>
                <form-label :label="$t('licenseLabel')" />
                <single-text-select v-model="temporaryInputs.license" :items="formatedLicenses"/>
              </div>
            </v-expand-transition>
            <v-row justify="center" class="mb-1 mt-4">
              <outlined-button width="88px" @click="reset">
                {{ $t('reset') }}
              </outlined-button>
              <basic-button width="88px" @click="search">
                {{ $t('search') }}
                <basic-button-circular v-show="searchLoading" />
              </basic-button>
            </v-row>
          </v-card>
        </v-form>
      </v-col>
      <v-col cols="12" sm="7" md="8">
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
              {{ repository.starCount | kilo }}
            </repository-chip>
            <repository-chip class="hidden-xs-only">
              <v-icon small>mdi-source-fork</v-icon>
              {{ repository.forkCount | kilo }}
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
                      {{ label }}
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
            {{ $t('showmoreLabel') }}
            <basic-button-circular v-show="showmoreLoading" />
          </basic-button>
        </v-row>
      </v-col>
    </v-row>
    <go-top-button />
  </div>
</template>

<script>
export default {
  data () {
    return {
      inputs: {},
      temporaryInputs: {},
      valid: true,
      show: false,
      showDetail: false,
    }
  },
  created () {
    this.initializeInputs()
    this.$store.dispatch('fetchRepositories', { params: this.getParams(), type: 'init'})
    this.$store.dispatch('fetchLanguages')
    this.$store.dispatch('fetchLicenses')
    this.$store.dispatch('fetchLabels')
    this.$store.dispatch('fetchOrdermetrics')
  },
  computed: {
    page () {
      return this.$store.state.page
    },
    hasNext () {
      return this.$store.state.hasNext
    },
    repositories () {
      return this.$store.state.repositories
    },
    languages () {
      return this.$store.state.languages
    },
    licenses () {
      return this.$store.state.licenses
    },
    labels () {
      return this.$store.state.labels
    },
    assignStatuses () {
      return this.$store.state.assignStatuses
    },
    ordermetrics () {
      return this.$store.state.ordermetrics
    },
    initLoading () {
      return this.$store.state.initLoading
    },
    searchLoading () {
      return this.$store.state.searchLoading
    },
    showmoreLoading () {
      return this.$store.state.showmoreLoading
    },
    formatedLanguages () {
      return this.languages.map(x => { return x === 'all' ? this.$t(x) : x })
    },
    formatedLicenses () {
      return this.licenses.map(x => { return x === 'all' ? this.$t(x) : x })
    },
    formatedLabels () {
      return this.labels.map(x => { return x === 'all' ? this.$t(x) : x })
    },
    formatedAssignStatuses () {
      return this.assignStatuses.map(x => this.$t(x))
    },
    formatedOrderMetrics () {
      return this.ordermetrics.map(x => this.$t(x))
    },
  },
  methods: {
    initializeInputs () {
      let defaultInputs = {
        languages: [this.$t('all')],
        labels: this.$labelsDefault,
        star_count_lower: '',
        star_count_upper: '',
        fork_count_lower: '',
        fork_count_upper: '',
        license: this.$t('all'),
        assigned: this.$t('all'),
        ordermetric: this.$t(this.$ordermetricDefault),
      }
      this.inputs = JSON.parse(JSON.stringify(defaultInputs))
      this.temporaryInputs = JSON.parse(JSON.stringify(defaultInputs))
    },
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
      this.initializeInputs()
    },
    showmore (event) {
      this.$store.commit('incrementPage')
      this.$store.dispatch('fetchRepositories', { params: this.getParams(), type: 'showmore' })
    },
    getParams () {
      let i = this.inputs

      let params = {}
      params.page = this.page
      if (i.languages !== '' && !i.languages.includes(this.$t('all'))) params.languages = i.languages.join(',')
      if (i.labels !== '' && !i.labels.includes(this.$t('all'))) params.labels = i.labels.join(',')
      if (i.star_count_lower !== '') params.star_count_lower = i.star_count_lower
      if (i.star_count_upper !== '') params.star_count_upper = i.star_count_upper
      if (i.fork_count_lower !== '') params.fork_count_lower = i.fork_count_lower
      if (i.fork_count_upper !== '') params.fork_count_upper = i.fork_count_upper
      if (i.license !== '' && i.license !== this.$t('all')) params.license = i.license
      if (i.ordermetric === this.$t(this.$ordermetricDefault)) {
        params.orderby = this.$camelToSnake(this.$ordermetricDefault)
      } else {
        for (const ordermetric of this.$store.state.ordermetrics) {
          if (i.ordermetric === this.$t(ordermetric)) {
            params.orderby = this.$camelToSnake(ordermetric)
            break
          }
        }
      }
      if (this.inputs.assigned == this.$t('assigned')) {
        params.assigned = true
      } else if (this.inputs.assigned == this.$t('unassigned')) {
        params.assigned = false
      }
      return params
    },
    removeLanguage (language) {
      const index = this.temporaryInputs.languages.indexOf(language)
      if (index >= 0) this.temporaryInputs.languages.splice(index, 1)
    },
    removeLabel (label) {
      const index = this.temporaryInputs.labels.indexOf(label)
      if (index >= 0) this.temporaryInputs.labels.splice(index, 1)
    },
  },
  filters: {
    kilo (value) {
      return value > 999 ? (value / 1000).toFixed(1) + 'k' : value
    }
  },
  head () {
    return {
      htmlAttrs: {
        lang: this.$t('headerLang')
      },
      title: this.$t('headerTitle'),
      meta: [
        { hid: 'description', name: 'description', content: this.$t('headerDescription') },
        { hid: 'og:title', property: 'og:title', content: 'gitnavi - ' + this.$t('headerTitle') },
        { hid: 'og:description', property: 'og:description', content: this.$t('headerDescription') },
      ],
    }
  }
}
</script>