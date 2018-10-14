import Vue from 'vue';
import Vuetify from 'vuetify';
import Tasky from 'Tasky/Tasky.vue';

import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);
new Vue({
    el: 'app',
    render: h => h(Tasky),
});
