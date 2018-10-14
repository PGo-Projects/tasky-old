import Vue from 'vue';
import Vuetify from 'vuetify';
import Index from 'Index/Index.vue';

import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);
new Vue({
    el: 'app',
    render: h => h(Index),
});
