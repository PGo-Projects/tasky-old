import Vue from 'vue';
import Vuetify from 'vuetify';
import Login from 'Login/Login.vue';

import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);
new Vue({
    el: 'app',
    render: h => h(Login),
});
