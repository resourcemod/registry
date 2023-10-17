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
        getUserById: (state) => (id) => {
            return state.usersList.get(parseInt(id))
        },
    },
    mutations: {
        updateUser(state, data){
            state.user.Name = data.user.Name
            state.user.Token = data.user.Token
            Cookies.set('rmod_auth', data.user.Token, { expires: 7, path: '/' })
        },
        updateUsers(state, data) {
            data.users.forEach((u) => {
                state.usersList.set(u.ID, u)
            })
        }
    },
    state: {
        user: {
            Name: '',
            Token: '',
        },
        usersList: new Map()
    },
    actions: {
        async getUsers(context) {
            const token = Cookies.get('rmod_token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/users', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
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
        async login(context, payload) {
            try {
                const body = JSON.stringify({
                    name: payload.name,
                    password: payload.password,
                })
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
            const token = Cookies.get('rmod_token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(API_ENDPOINT+'/api/v1/user', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
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
                const body = JSON.stringify({
                    name: payload.name,
                    password: payload.password
                })
                const response = await fetch(API_ENDPOINT+'/api/v1/users', {
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