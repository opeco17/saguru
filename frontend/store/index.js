import Vuex from 'vuex'

export default () => (new Vuex.Store({
    state: {
        page: 0,
        hasNext: true,
        repositories: [],
        languages: [],
        licenses: [],
        labels: [],
    },
    getters: {
        page: (state) => state.page,
        hasNext: (state) => state.hasNext,
        repositories: (state) => state.repositories,
        languages: (state) => state.languages,
        licenses: (state) => state.licenses,
        labels: (state) => state.labels,
    },
    mutations: {
        resetPage (state) {
            state.page = 0
        },
        incrementPage (state) {
            state.page += 1
        },
        decrimentPage (state) {
            state.page -= 1
        },
        setHasNext (state, hasNext) {
            state.hasNext = hasNext
        },
        setRepositories (state, repositories) {
            state.repositories = repositories
        },
        addRepositories (state, repositories) {
            state.repositories = state.repositories.concat(repositories)
        },
        setLanguages (state, languages) {
            state.languages = languages
        },
        setLicenses (state, licenses) {
            state.licenses = licenses
        },
        setLabels (state, labels) {
            state.labels = labels
        },
    },
    actions: {        
        async fetchRepositories(ctx, {params, add}) {
            const res = await this.$axios.$get(`${this.$API_BASE_URL}/repositories`, { params: params})
            for (let item of res.items) {
                item.show = false
            }
            if (add) {
                ctx.commit('addRepositories', res.items)
            } else {
                ctx.commit('setRepositories', res.items)
            }
            ctx.commit('setHasNext', res.hasNext)
        },
        async fetchLanguages(ctx) {
            const res = await this.$axios.$get(`${this.$API_BASE_URL}/languages`)
            res.items.unshift('all')
            ctx.commit('setLanguages', res.items)
        },
        async fetchLicenses(ctx) {
            const res = await this.$axios.$get(`${this.$API_BASE_URL}/licenses`)
            res.items.unshift('all')
            ctx.commit('setLicenses', res.items)
        },
        async fetchLabels(ctx) {
            const res = await this.$axios.$get(`${this.$API_BASE_URL}/labels`)
            res.items.unshift('all')
            ctx.commit('setLabels', res.items)
        }
    }
}))