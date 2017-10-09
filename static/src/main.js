// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false


import axios from 'axios'
import VueAxios from 'vue-axios'

import VueLazyload from 'vue-lazyload'


import VueTimeago from 'vue-timeago'

Vue.use(VueTimeago, {
    name: 'timeago', // component name, `timeago` by default
    locale: 'zh-CN',
    locales: {
        // you will need json-loader in webpack 1
        'zh-CN': require('vue-timeago/locales/zh-CN.json')
    }
})

Vue.use(VueAxios, axios)
Vue.use(VueLazyload)



/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
