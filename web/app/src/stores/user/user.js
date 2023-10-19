import Cookies from 'js-cookie'
const API_ENDPOINT = import.meta.env.DEV ? 'http://localhost:8888' : window.location.origin
export default {
    getters: {
        getUser(state) {
            return state.user
        },
        getUsers(state) {
            return state.usersList
        },
        getUserByName: (state) => (name) => {
            return state.usersList.get(name)
        },
    },
    mutations: {
        updateUser(state, data){
            state.user.name = data.name
            state.user.access_token = data.access_token
            state.user.created_at = data.created_at
            state.user.updated_at = data.updated_at
            state.user.is_owner = data.is_owner
            Cookies.set('rmod_auth', data.access_token, { expires: 30, path: '/' })
        },
        updateUsers(state, data) {
            data.users.forEach((u) => {
                state.usersList.set(u.name, u)
            })
        },
        deleteUser(state, data) {
          state.usersList.delete(data.name)
        },
        putUserIntoMap(state, data) {
            state.usersList.set(data.name, data)
        }
    },
    state: {
        user: {
            name: '',
            access_token: '',
            created_at: '',
            updated_at: '',
            is_owner: false
        },
        usersList: new Map()
    },
    actions: {
        async getUsers(context) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/users', {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateUsers', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteUser(context, payload) {
            const token = Cookies.get('rmod_auth')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }

            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/users/'+payload.name, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': 'Bearer '+token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('deleteUser', payload.name)
                return data
            } catch (e) {
                throw e
            }
        },
        async login(context, payload) {
            try {
                const body = JSON.stringify(payload)
                const response = await fetch(API_ENDPOINT+'/api/v1/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                if (!response.ok) {
                    throw response.message
                }
                context.commit('updateUser', data)
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
        async checkIsSetupRequired(context) {
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/setup', {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json'
                    }
                })
                return await response.json()
            } catch (e) {
                throw e
            }
        },
    }
}