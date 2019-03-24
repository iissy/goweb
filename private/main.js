import Vue from 'vue';

import App from './App';
import router from './router';

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next(true);
});
router.afterEach((to, from) => {
    if (from.fullPath == '/loginRegister') {}
});

Vue.config.productionTip = false

new Vue({
    el: '#webpack',
    router,
    components: { App },
    template: '<App/>'
})