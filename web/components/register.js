import Vue from 'vue';
import Vuetify from 'vuetify';
import Register from 'Register/Register.vue';

import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);
new Vue({
    el: 'app',
    render: h => h(Register),
});
