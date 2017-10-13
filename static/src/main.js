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



// // main.js
// Vue.prototype.SDKRegister = (that, callback) => {
//     // 接入微信JSSDK
//     // 获取微信JSSDK配置
//     let url = that.webUrl
//     that.$http.get(url).then(res => {
//       res = res.body
//       that.$wechat.config({
//         debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
//         appId: res.appId, // 必填，公众号的唯一标识
//         timestamp: res.timestamp, // 必填，生成签名的时间戳
//         nonceStr: res.nonceStr, // 必填，生成签名的随机串
//         signature: res.signature, // 必填，签名，见附录1
//         jsApiList: res.jsApiList
//         // jsApiList: [
//         //   'checkJsApi',
//         //   'onMenuShareAppMessage', // 获取“分享给朋友”按钮点击状态及自定义分享内容接口
//         //   'onMenuShareTimeline', // 获取“分享到朋友圈”按钮点击状态及自定义分享内容接口
//         //   'onMenuShareQQ', // 获取“分享到QQ”按钮点击状态及自定义分享内容接口
//         //   'onMenuShareWeibo' // 获取“分享到腾讯微博”按钮点击状态及自定义分享内容接口
//         // ] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
//       })
//     })
//     that.$wechat.ready((res) => {
//       that.$wechat.showAllNonBaseMenuItem()
//       // 分享到朋友圈
//       let link = ''
//       let title = ''
//       let imgUrl = ''
//       let desc = ''
//     //   ...删掉了我业务相关变量的内容定义
//       that.$wechat.onMenuShareTimeline({
//         title: title, // 分享标题
//         link: link, // 分享链接
//         imgUrl: imgUrl, // 分享图标
//         success () {
//           // 用户确认分享后执行的回调函数
//           that.$alert('分享成功', 'success')
//         },
//         cancel () {
//           // 用户取消分享后执行的回调函数
//         }
//       })
//       // 分享给朋友
//       that.$wechat.onMenuShareAppMessage({
//         title: title, // 分享标题
//         desc: desc, // 分享描述
//         link: link, // 分享链接
//         imgUrl: imgUrl, // 分享图标
//         success: function () {
//           // 用户确认分享后执行的回调函数
//           that.$alert('分享成功', 'success')
//         },
//         cancel: function () {
//           // 用户取消分享后执行的回调函数
//         }
//       })
//       // 如果需要定制ready回调方法
//       if(callback){
//         callback.call(that)
//       }
//     })
//   }
  
//   // 要用到定位的页面
//   beforeCreate () {
//     this.SDKRegister(this, () => {
//       this.$wechat.getLocation({
//         success (res) {
//           let citylocation = new qqMap.CityService({
//             complete (results) {
//               let location = results['detail']['detail'].split(',')
//               that.$store.commit('city', location[1])
//               window.localStorage.city = location[1]
//             }
//           })
//           let latitude = res.latitude // 纬度，浮点数，范围为90 ~ -90
//           let longitude = res.longitude // 经度，浮点数，范围为180 ~ -180。
//           // 设置经纬度信息
//           let latLng = new qqMap.LatLng(latitude, longitude)
//           // 调用城市经纬度查询接口实现经纬查询
//           citylocation.searchCityByLatLng(latLng)
//         },
//         cancel: function () {
//           // 这个地方是用户拒绝获取地理位置
//         }
//       })
//     })
//   }

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
