import Cookies from 'js-cookie'

const API_ENDPOINT = import.meta.env.DEV ? 'http://localhost:8888' : window.location.origin
export default {
    getters: {
        getPlugins(state) {
            return state.plugins
        },
        getExtensions(state) {
            return state.extensions
        },
        getContentByName: (state) => (type, name) => {
            if (type === 'plugin') {
                return state.plugins.get(name)
            }
            return state.extensions.get(name)
        },
    },
    mutations: {
        updateContent(state, data) {
            if (!data) {
                return
            }
            if (data.type === 'plugin') {
                state.plugins.set(data.name, data)
                return
            }
            state.extensions.set(data.name, data)
        },
        deletePlugin(state, data) {
            state.plugins.delete(data)
        },
        addPlugin(state, data) {
            state.plugins.set(data.name, data)
        },
        updateAllContent(state, data) {
            if (!data.content) {
                return
            }
            data.content.forEach((u) => {
                if (u.type === 'plugin') {
                    state.plugins.set(u.name, u)
                } else {
                    state.extensions.set(u.name, u)
                }
            })
        },
        updateExtension(state, data) {
            state.extensions.set(data.name, data)
        },
        deleteExtension(state, data) {
            state.extensions.delete(data)
        },
        addExtension(state, data) {
            state.extensions.set(data.name, data)
        }
    },
    state: {
        plugins: new Map(),
        extensions: new Map()
    },
    actions: {
        async getContentByType(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            const response = await fetch(API_ENDPOINT + '/api/v1/content/' + payload.type, {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                }
            })
            const data = await response.json()
            if (!response.ok) {
                throw data
            }
            context.commit('updateAllContent', data)
            return data
        },
        async getContentByTypeAndName(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            const response = await fetch(API_ENDPOINT + '/api/v1/content/' + payload.type + '/' + payload.name, {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                }
            })
            const data = await response.json()
            if (!response.ok) {
                throw data
            }
            context.commit('updateContent', data)
            return data
        },
        async deleteContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }

            const response = await fetch(API_ENDPOINT + '/api/v1/content/' + payload.type + '/' + payload.name, {
                method: 'DELETE',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                }
            })
            const data = await response.json()
            if (!response.ok) {
                throw data
            }
            if (payload.type === "plugin") {
                context.commit('deletePlugin', data)
                return data
            }
            context.commit("deleteExtension", data)
            return data
        },
        async getContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            const body = JSON.stringify(payload)
            const response = await fetch(API_ENDPOINT + '/api/v1/content/' + payload.type + '/' + payload.name, {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                },
                body: body
            })
            const data = await response.json()
            if (!response.ok) {
                throw data
            }
            if (payload.type === "plugin") {
                context.commit('updatePlugin', data)
                return data
            }
            context.commit("updateExtension", data)
            return data
        },
        async uploadContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            const body = JSON.stringify(payload)
            const response = await fetch(API_ENDPOINT + '/api/v1/content', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                    'Authorization': 'Bearer ' + token
                },
                body: body
            })
            if (!response.ok) {
                throw await response.json()
            }
            const data = await response.json()
            context.commit('updateContent', data)
            return data
        },
        async updateContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            const body = JSON.stringify(payload)
            const response = await fetch(API_ENDPOINT + '/api/v1/content/' + payload.type + '/' + payload.name, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Accept': 'application/json',
                    'Authorization': 'Bearer ' + token
                },
                body: body
            })
            if (!response.ok) {
                throw await response.json()
            }
            const data = await response.json()
            context.commit('updateContent', data)
            return data
        },
    }
}