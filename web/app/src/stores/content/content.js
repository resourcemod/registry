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
        getPluginByName: (state) => (name) => {
            return state.plugins.get(name)
        },
        getExtensionByName: (state) => (name) => {
            return state.extensions.get(name)
        },
    },
    mutations: {
        updatePlugins(state, data) {
            if (!data.content) {
                return
            }
            data.content.forEach((u) => {
                state.plugins.set(u.name, u)
            })
        },
        updatePlugin(state, data) {
            state.plugins.set(data.content.name, data)
        },
        deletePlugin(state, data) {
          state.plugins.delete(data)
        },
        addPlugin(state, data) {
            state.plugins.set(data.name, data)
        },
        updateExtensions(state, data) {
            if (!data.content) {
                return
            }
            data.content.forEach((u) => {
                state.extensions.set(u.name, u)
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
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/content/'+payload.type, {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                if (payload.type === "plugin") {
                    context.commit('updatePlugins', data)
                    return data
                }
                context.commit("updateExtensions", data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }

            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/content/'+payload.type+'/'+payload.name, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                if (payload.type === "plugin") {
                    context.commit('deletePlugin', data)
                    return data
                }
                context.commit("deleteExtension", data)
                return data
            } catch (e) {
                throw e
            }
        },
        async getContent(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const body = JSON.stringify(payload)
                const response = await fetch(API_ENDPOINT+'/api/v1/content/'+payload.type+'/'+payload.name, {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                if (!response.ok) {
                    throw response.message
                }
                if (payload.type === "plugin") {
                    context.commit('updatePlugin', data)
                    return data
                }
                context.commit("updateExtension", data)
                return data
            } catch(e) {
                throw e
            }
        },
        async getUser(context) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/user', {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateUser', data)
                return data
            } catch (e) {
                throw e
            }
        },

        async createFirstAccount(context, payload) {
            try {
                const body = JSON.stringify(payload)
                const response = await fetch(API_ENDPOINT+'/api/v1/setup/create', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                return await response.json()
            } catch(e) {
                throw e
            }
        },
        async register(context, payload) {
            try {
                const body = JSON.stringify(payload)
                const response = await fetch(API_ENDPOINT+'/api/v1/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                context.commit('updateUser', data)
                return data
            } catch(e) {
                throw e
            }
        },
        async createUser(context, payload) {
            try {
                const token = Cookies.get('rmod_auth')
                if (!token || token.length === 0) {
                    throw "Token undefined."
                }
                const body = JSON.stringify(payload)
                const response = await fetch(API_ENDPOINT+'/api/v1/users', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                        'Authorization': 'Bearer '+token
                    },
                    body: body
                })
                const data = await response.json()
                context.commit('putUserIntoMap', data)
                return data
            } catch(e) {
                throw e
            }
        },
    }
}