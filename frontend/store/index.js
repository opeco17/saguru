import Vuex from 'vuex'

export default () => (new Vuex.Store({
    state: {
        page: 0,
        hasNext: true,
        repositories: [],
        languages: [],
        licenses: [],
        labels: [],
        ordermetrics: [],
        assignStatuses: ["all", "assigned", "unassigned"],
        initLoading: false,
        searchLoading: false,
        showmoreLoading: false,
    },
    getters: {
        page: (state) => state.page,
        hasNext: (state) => state.hasNext,
        repositories: (state) => state.repositories,
        languages: (state) => state.languages,
        licenses: (state) => state.licenses,
        labels: (state) => state.labels,
        ordermetrics: (state) => state.ordermetrics,
        initLoading: (state) => state.initLoading,
        searchLoading: (state) => state.searchLoading,
        showmoreLoading: (state) => state.showmoreLoading,
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
        setOrdermetrics (state, ordermetrics) {
            state.ordermetrics = ordermetrics
        },
        setInitLoading (state, initLoading) {
            state.initLoading = initLoading
        },
        setSearchLoading (state, searchLoading) {
            state.searchLoading = searchLoading
        },
        setShowmoreLoading (state, showmoreLoading) {
            state.showmoreLoading = showmoreLoading
        }
    },
    actions: {        
        async fetchRepositories(ctx, {params, type}) {
            let loadingMutation = ''
            if (type == 'init') {
                loadingMutation = 'setInitLoading'
            } else if (type == 'search') {
                loadingMutation = 'setSearchLoading'
            } else if (type == 'showmore') {
                loadingMutation = 'setShowmoreLoading'
            }
            ctx.commit(loadingMutation, true)

            const res = await this.$axios.$get('repositories', { params: params})
            for (let item of res.items) {
                item.show = false
            }
            if (type == 'showmore') {
                ctx.commit('addRepositories', res.items)
            } else {
                ctx.commit('setRepositories', res.items)
            }
            ctx.commit(loadingMutation, false)
            ctx.commit('setHasNext', res.hasNext)
        },
        async fetchLanguages(ctx) {
            const res = await this.$axios.$get('languages')
            res.items.unshift('all')
            ctx.commit('setLanguages', res.items)
        },
        async fetchLicenses(ctx) {
            const res = await this.$axios.$get('licenses')
            res.items.unshift('all')
            ctx.commit('setLicenses', res.items)
        },
        async fetchLabels(ctx) {
            const res = await this.$axios.$get('labels')
            res.items.unshift('all')
            ctx.commit('setLabels', res.items)
        },
        async fetchOrdermetrics(ctx) {
            const res = await this.$axios.$get('ordermetrics')
            ctx.commit('setOrdermetrics', res.items.map(this.$snakeToCamel))
        }
    }
}))