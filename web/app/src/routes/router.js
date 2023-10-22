import {createRouter, createWebHistory} from 'vue-router'
import Login from "../pages/authorization/Login.vue";
import Dashboard from "../pages/Dashboard.vue";
import Users from "../pages/users/Users.vue";
import store from "../stores";
import Logout from "../pages/authorization/Logout.vue";
import UserDetails from "../pages/users/UserDetails.vue";
import Setup from "../pages/setup/Setup.vue";
import Register from "../pages/authorization/Register.vue";
import UserCreate from "../pages/users/UserCreate.vue";
import PluginsList from "../pages/plugins/PluginsList.vue";
import PluginDetails from "../pages/plugins/PluginDetails.vue";
import ExtensionsList from "../pages/extensions/ExtensionsList.vue";
import ExtensionDetails from "../pages/extensions/ExtensionDetails.vue";
import ExtensionCreate from "../pages/extensions/ExtensionCreate.vue";
import UserDelete from "../pages/users/UserDelete.vue";
import UserFeed from "../pages/users/UserFeed.vue";
import ContentCreate from "../pages/ContentCreate.vue";
import Integrations from "../pages/integrations/Integrations.vue";

const router = createRouter({
    routes: [
        {
            path: '/',
            name: 'Dashboard',
            component: Dashboard
        },
        {
            path: '/setup',
            name: 'Setup',
            component: Setup
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
        {
            path: '/register',
            name: 'Register',
            component: Register
        },
        {
            path: '/logout',
            name: 'Logout',
            component: Logout
        },
        {
            path: '/users',
            name: 'Users',
            component: Users
        },
        {
            path: '/users/:name',
            name: 'Users.Details',
            component: UserDetails
        },
        {
            path: '/users/create',
            name: 'Users.Create',
            component: UserCreate
        },
        {
            path: '/users/feed',
            name: 'Users.Feed',
            component: UserFeed
        },
        {
            path: '/users/:name/delete',
            name: 'Users.Delete',
            component: UserDelete
        },
        {
            path: '/content/plugins',
            name: 'Plugins.List',
            component: PluginsList
        },
        {
            path: '/content/plugins/:name',
            name: 'Plugins.Details',
            component: PluginDetails
        },
        {
            path: '/content/create',
            name: 'Content.Create',
            component: ContentCreate
        },
        {
            path: '/content/extensions',
            name: 'Extensions.List',
            component: ExtensionsList
        },
        {
            path: '/content/extensions/:name',
            name: 'Extensions.Details',
            component: ExtensionDetails
        },
        {
            path: '/content/extensions/create',
            name: 'Extensions.Create',
            component: ExtensionCreate
        },
        {
            path: '/integrations',
            name: 'Integrations.List',
            component: Integrations
        }
    ],
    history: createWebHistory()
})

function authMiddleware() {
    const user = store.getters.getUser

    router.beforeEach(async (to, from) => {
        if (!user.access_token || user.access_token.length === 0) {
            const data = await store.dispatch('checkIsSetupRequired')

            if (to.name !== 'Setup' && data.required === true) {
                return {name:'Setup'}
            } else if (data.required === false) {
                if (to.name !== 'Login' && to.name !== 'Register') {
                    return {name: 'Login'}
                }
                if (to.name === 'Logout') {
                    return {name: 'Login'}
                }
            }
        }

        if (
            user.access_token &&
            (to.name === 'Login' || to.name === 'Register')
        ) {
            if (user.is_owner) {
                return {name: 'Dashboard'}
            } else {
                return {name: 'Logout'}
            }
        }
    })
}
authMiddleware()

export default router