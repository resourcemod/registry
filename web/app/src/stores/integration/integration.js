import Cookies from 'js-cookie'

const API_ENDPOINT = import.meta.env.DEV ? 'http://localhost:8888' : window.location.origin
export default {
    getters: {
        getIntegrations(state) {
            return state.integrations
        },
        getIntegrationByName: (state) => (name) => {
            return state.integrations.get(name)
        },
    },
    mutations: {
        updateIntegrations(state, data) {
            if (data.integrations === null) {
                return;
            }
            data.integrations.forEach((u) => {
                state.integrations.set(u.name, u)
            })
        },
        deleteIntegration(state, data) {
            state.integrations.delete(data)
        },
        putIntegrationIntoMap(state, data) {
            state.integrations.set(data.name, data)
        }
    },
    state: {
        integrations: new Map()
    },
    actions: {
        async getIntegrations(context) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT + '/api/v1/integrations', {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer ' + token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateIntegrations', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteIntegration(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }

            try {
                const response = await fetch(API_ENDPOINT + '/api/v1/integrations/' + payload.name, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': 'Bearer ' + token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('deleteIntegration', payload.name)
                return data
            } catch (e) {
                throw e
            }
        },
        async createIntegration(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }

            const body = JSON.stringify(payload)
            const response = await fetch(API_ENDPOINT + '/api/v1/integrations', {
                method: 'POST',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                },
                body: body
            })
            const data = await response.json()
            if (!response.ok) {
                throw data.message
            }
            context.commit('putIntegrationIntoMap', data)
            return data
        },
        async getRepositories(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT + '/api/v1/integrations/'+payload.name+'/repositories', {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer ' + token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                return await response.json()
            } catch (e) {
                throw e
            }
        },
    }
}