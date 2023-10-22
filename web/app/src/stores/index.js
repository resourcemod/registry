import {createStore} from 'vuex'
import user from './user/user'
import content from './content/content'
import integration from "./integration/integration";

const store = createStore({
    modules: [
        user,
        content,
        integration
    ]
})

export default store