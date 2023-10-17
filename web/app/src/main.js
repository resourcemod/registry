import {createApp} from 'vue';
import App from './App.vue';
import './app.css';
import Cookies from 'js-cookie'
import router from './routes/router.js'
import store from './stores/index'

const preload = async () => {
    const token = Cookies.get('rmod_auth')
    if (token && token.length > 0) {
        try {
            const data = await store.dispatch('getUser')
            if (!data) {
                throw "Failed to request user"
            }
        } catch (e) {
            Cookies.remove('rmod_auth')
        }
    }
    const app = createApp(App)
    app.use(router)
    router.app = app
    app.use(store)
    app.mount('#app');
}

preload()
