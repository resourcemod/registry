import {createStore} from 'vuex'
import user from './user/user'
import content from './content/content'

const store = createStore({
    modules: [
        user,
        content,
    ]
})

export default store