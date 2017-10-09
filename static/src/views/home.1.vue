//index.vue
<template>

    <div class="container">
	<div class="row ">
justify-content-center

<!-- <div class="col-md-8 col-sm-9"> -->

<div>

  <scroller :on-refresh="refresh"
            :on-infinite="infinite">


      <div class="card mt-3" v-for="article in articles" :key="article.id">
        <div class="card-body">

          <h4 class="card-title">
             <a  v-bind:href="[article.URL]"  target="_blank" >{{article.Title}} </a>  
          </h4>

          <div class="row">
              <div class=" mx-auto col-md-3 order-md-2">
                  <!-- <img class="card-img-top"  :src="article.Cover" alt="Card image cap"> -->
                  <img class="card-img-top"  v-lazy="article.Cover" alt="Card image cap">
              </div>
              <div class="col-md-9 order-md-1  text-md-left pr-md-5">

                  <p class="card-text">
                    <small class="text-muted">{{article.AppName}}</small>
                    <small class="text-muted">{{article.PubAt}}</small>
                    <small class="text-muted">4321鲜花</small>
                      <small class="text-muted">10鸡蛋</small>
                  </p>
                
                
                  <p class="card-text">{{article.Intro}}</p>

              </div>
          </div>
        </div>
      </div>
  </scroller>

</div>
</div>


  </div>


      
    <!-- </div> -->
</template>

<script>

// import Page from '../components/page.vue';
import news from '../api/news.js';

export default {
  name: 'indexP',
  
  
    data () {
      return {
        counter:1,
        articles: [],
        // articles : [{"ID":1,"Title":"《秋之歌》，美极了！！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdWRpY3hYU1BSdHNVellFWEVBYmJpYm8yWldJR3R3NVhmeXZGbFZCRmRIYzZoOWNCTnRMV1p2cGlhdy8wP3d4X2ZtdD1qcGVn","Intro":"太美了..","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=1\u0026sn=1a0820525abea356c62808748ec8f337\u0026chksm=8483843cb3f40d2aa86ef1e0265e692ce001b6f617a7e62a28ef13799094d2123f9e18efc10c#rd","CreatedAt":"2017-09-28T20:27:34.7191501+08:00","UpdatedAt":"2017-10-01T23:28:39.3294006+08:00","DeletedAt":null},{"ID":2,"Title":"经典音乐：古筝钢琴版《南山南》，静静地听，很美！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdVVWYnBUb0lxZGhaVWliT1lEbVcwUzhuQ0djaG9FeGlhTUlQSjhvYU1zWEI3S1N5S05jc1ZqaWJCZy8wP3d4X2ZtdD1qcGVn","Intro":"..","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=5\u0026sn=30f6645ff4cf1044efd32ab591ed595e\u0026chksm=8483843cb3f40d2a83f30369e2b27c240c15b1f06b3275a2713c0cf14112a47145ed8e3c656c#rd","CreatedAt":"2017-09-28T20:28:06.8909902+08:00","UpdatedAt":"2017-10-01T11:31:16.517349+08:00","DeletedAt":null},{"ID":3,"Title":"乡下小伙一首《雨花石》，惊呆全场！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdWpkczE4T0xpYUMwWGpYdU96MzdRM1VRZmsyQmc4NEdMaE5yQjJUZzYzd2lhaWFVOFBSM2d1bVhwUS8wP3d4X2ZtdD1qcGVn","Intro":"◇▼ 乡下小伙唱《雨花石》喜欢本文的朋友们，记得转发点哦","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=3\u0026sn=71734f3864c4774a2c364b4f518eb3fb\u0026chksm=8483843cb3f40d2a69bf1c09350262b6ee6695a81e668919ba941998fc15531a71a4e83ff584#rd","CreatedAt":"2017-09-28T22:27:44.3305163+08:00","UpdatedAt":"2017-09-28T22:27:44.3365166+08:00","DeletedAt":null},{"ID":4,"Title":"听雨，原来那么美！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFmaWNWMmI3Njl6aWJDbEtaU2pCcUJVVkk4ZDdtZWljTkJ0T1NQeHVyUzRXRVVvREVNMERpYTdaeGlidUZRb0FpY2ljYWNQalJIS2Y0VzRTaEU5US8wP3d4X2ZtdD1qcGVn","Intro":"..","PubAt":"2017-09-27","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648386\u0026idx=2\u0026sn=99ca38053bbe67e2661fbdf7bb0aef72\u0026chksm=8483841eb3f40d083d1c468a08cbaafa209f401a512c16815dc9239d1202eeb0b132f4a24f84#rd","CreatedAt":"2017-09-28T22:27:44.6285333+08:00","UpdatedAt":"2017-09-28T22:27:44.6325336+08:00","DeletedAt":null},{"ID":5,"Title":"11点后睡的人，白开水里一定要加一点它！让肝越来越好","Author":"","AppName":"新华保险微店","AppID":"gh_eca184ec3dc5","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL25NU2htVm50SGV0NjFCZ1k0b0N5V1ZwZU5pYU94TU1lSGNRM1VYcWJOWWljeGliaWNpYWdwVURKazMzS0E2WGpqaGFkNWtjaWJsajBQeEdYQ3lKZkZqRUlnb1hBLzA_d3hfZm10PWpwZWc=","Intro":"不知有多少人，晚上常常熬夜，或者大吃大喝。我们都知道这样很伤身体，但是到底伤害在哪里呢？真的应该好好看看！","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MTU3NjE1NA==\u0026mid=2650030645\u0026idx=1\u0026sn=77651df410bec95899ae709d23cc0107\u0026chksm=872b9d00b05c1416279a74e59bd7d9d2b2974c38b4a08c4593134f75c19e53ad465206acde93#rd","CreatedAt":"2017-09-28T22:37:20.9994999+08:00","UpdatedAt":"2017-09-28T22:37:21.0065003+08:00","DeletedAt":null},{"ID":6,"Title":"它陪伴了中国人的一生","Author":"中国精致生活►","AppName":"物道","AppID":"gh_f27ca3831aea","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL0R0a2ljUHM4d0x0aWNWMVdpYzlsN3pNTlhpY2FlQk5vVXByTzd6aWJiN2JIVVBkaEVheUM2NE1HZnZpYU1zclpvSVNtU2pQOVZBZHh3VUF5bDdTSTA1MW9vRjNBLzA_d3hfZm10PWpwZWc=","Intro":"生活，不过一碗人间烟火。","PubAt":"1506564000","URL":"http://mp.weixin.qq.com/s?__biz=MzA4MzI1Mzc2MA==\u0026mid=2653818101\u0026idx=1\u0026sn=b88229172064f71ab9c7d21ef680c722\u0026chksm=8423e39ab3546a8c9fddff0dc2473e6b00946b1d6291d539f469dc3f3c13f4aa935def916b59#rd","CreatedAt":"2017-09-29T00:12:31.8671428+08:00","UpdatedAt":"2017-10-01T23:30:05.4893286+08:00","DeletedAt":null},{"ID":7,"Title":"世界上最惊艳的12张照片，绝对震撼！你见过吗？","Author":"","AppName":"无罪","AppID":"gh_b1a7224ce2f4","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1VOQmlhcFlDZm00bzdKM0V5aWN6WEFpY3l0dXNpYkVmWFhpYjBqMEFBZkJSTktpY2NibjJucGliZVlhSHR1ckdhRVNLVDFBd2hrSE1kQnA3VjZLRmZPbUFzaWFraWNnLzA_d3hfZm10PWpwZWc=","Intro":"保持对世界的好奇心，让生活充满欢乐。","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MjM5NzgzNTUyNA==\u0026mid=2650373542\u0026idx=1\u0026sn=a419c3bf6461ecaf4c6fc43525cd8204\u0026chksm=bede3bcd89a9b2dba5e293bc6815b3a685b7a20169fadcf7899ce157c15a7172fd7ad966e1dd#rd","CreatedAt":"2017-09-29T00:20:17.0817516+08:00","UpdatedAt":"2017-09-29T00:20:17.0927522+08:00","DeletedAt":null},{"ID":8,"Title":"关于新书，已经有了18万字。","Author":"梦入神机","AppName":"梦入神机","AppID":"gh_af6d690a9029","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1dVaWFid2Q1TTNlcTNpYm9UYXQwM3RyaFpJcm1YQWh0NUxTN2I1RmZCZmJsTjk3Z080d0J6d0pGVWJraWNpY3J2cW5MN1NSOGliZEJrT2J2bURuRWljbmFLdlR3LzA_d3hfZm10PWpwZWc=","Intro":"我以前说过8月份开新书的事情，实际上","PubAt":"2017-08-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3NDA4MjYxNQ==\u0026mid=2454021024\u0026idx=1\u0026sn=8bb3f5c901b8a17781ce87765c438ae0\u0026chksm=88b4f3f8bfc37aeec7a0afdea9de96b770b3a7064ed43104604bc0daeaa14605e7ee964a0c6c#rd","CreatedAt":"2017-09-29T00:24:50.8014074+08:00","UpdatedAt":"2017-09-29T00:24:50.811408+08:00","DeletedAt":null},{"ID":9,"Title":"降央卓玛一曲《国家》喜迎国庆，好听！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkR0h1dTk5WDE3UW42WW5kRTZBaFR3UVh4TE1hZEF2SllKTHVGbGlhNkFQbU1jaWJPeW9KNDVjNzNMc2xIZjZRTGVSbmVxVFdNUkdBTWcvMD93eF9mbXQ9anBlZw==","Intro":"..","PubAt":"2017-10-01","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648485\u0026idx=1\u0026sn=27f1299a3d7650302fa2dd84fb6cfede\u0026chksm=84838479b3f40d6f4d78a5822efed106c67d21380a5fb19bdefcdc2a5224c7dcba133d6c508e#rd","CreatedAt":"2017-10-01T21:06:42.22425+08:00","UpdatedAt":"2017-10-01T21:06:42.2352506+08:00","DeletedAt":null}],

        items: [],
      }
    },
    mounted() {
      this.counter = 1
      var api = 'http://localhost:8888/hot?limit=2';
      var site = this
      // site.articles = [{"ID":1,"Title":"《秋之歌》，美极了！！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdWRpY3hYU1BSdHNVellFWEVBYmJpYm8yWldJR3R3NVhmeXZGbFZCRmRIYzZoOWNCTnRMV1p2cGlhdy8wP3d4X2ZtdD1qcGVn","Intro":"太美了..","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=1\u0026sn=1a0820525abea356c62808748ec8f337\u0026chksm=8483843cb3f40d2aa86ef1e0265e692ce001b6f617a7e62a28ef13799094d2123f9e18efc10c#rd","CreatedAt":"2017-09-28T20:27:34.7191501+08:00","UpdatedAt":"2017-10-01T23:28:39.3294006+08:00","DeletedAt":null},{"ID":2,"Title":"经典音乐：古筝钢琴版《南山南》，静静地听，很美！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdVVWYnBUb0lxZGhaVWliT1lEbVcwUzhuQ0djaG9FeGlhTUlQSjhvYU1zWEI3S1N5S05jc1ZqaWJCZy8wP3d4X2ZtdD1qcGVn","Intro":"..","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=5\u0026sn=30f6645ff4cf1044efd32ab591ed595e\u0026chksm=8483843cb3f40d2a83f30369e2b27c240c15b1f06b3275a2713c0cf14112a47145ed8e3c656c#rd","CreatedAt":"2017-09-28T20:28:06.8909902+08:00","UpdatedAt":"2017-10-01T11:31:16.517349+08:00","DeletedAt":null},{"ID":3,"Title":"乡下小伙一首《雨花石》，惊呆全场！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdWpkczE4T0xpYUMwWGpYdU96MzdRM1VRZmsyQmc4NEdMaE5yQjJUZzYzd2lhaWFVOFBSM2d1bVhwUS8wP3d4X2ZtdD1qcGVn","Intro":"◇▼ 乡下小伙唱《雨花石》喜欢本文的朋友们，记得转发点哦","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648416\u0026idx=3\u0026sn=71734f3864c4774a2c364b4f518eb3fb\u0026chksm=8483843cb3f40d2a69bf1c09350262b6ee6695a81e668919ba941998fc15531a71a4e83ff584#rd","CreatedAt":"2017-09-28T22:27:44.3305163+08:00","UpdatedAt":"2017-09-28T22:27:44.3365166+08:00","DeletedAt":null},{"ID":4,"Title":"听雨，原来那么美！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFmaWNWMmI3Njl6aWJDbEtaU2pCcUJVVkk4ZDdtZWljTkJ0T1NQeHVyUzRXRVVvREVNMERpYTdaeGlidUZRb0FpY2ljYWNQalJIS2Y0VzRTaEU5US8wP3d4X2ZtdD1qcGVn","Intro":"..","PubAt":"2017-09-27","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648386\u0026idx=2\u0026sn=99ca38053bbe67e2661fbdf7bb0aef72\u0026chksm=8483841eb3f40d083d1c468a08cbaafa209f401a512c16815dc9239d1202eeb0b132f4a24f84#rd","CreatedAt":"2017-09-28T22:27:44.6285333+08:00","UpdatedAt":"2017-09-28T22:27:44.6325336+08:00","DeletedAt":null},{"ID":5,"Title":"11点后睡的人，白开水里一定要加一点它！让肝越来越好","Author":"","AppName":"新华保险微店","AppID":"gh_eca184ec3dc5","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL25NU2htVm50SGV0NjFCZ1k0b0N5V1ZwZU5pYU94TU1lSGNRM1VYcWJOWWljeGliaWNpYWdwVURKazMzS0E2WGpqaGFkNWtjaWJsajBQeEdYQ3lKZkZqRUlnb1hBLzA_d3hfZm10PWpwZWc=","Intro":"不知有多少人，晚上常常熬夜，或者大吃大喝。我们都知道这样很伤身体，但是到底伤害在哪里呢？真的应该好好看看！","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MTU3NjE1NA==\u0026mid=2650030645\u0026idx=1\u0026sn=77651df410bec95899ae709d23cc0107\u0026chksm=872b9d00b05c1416279a74e59bd7d9d2b2974c38b4a08c4593134f75c19e53ad465206acde93#rd","CreatedAt":"2017-09-28T22:37:20.9994999+08:00","UpdatedAt":"2017-09-28T22:37:21.0065003+08:00","DeletedAt":null},{"ID":6,"Title":"它陪伴了中国人的一生","Author":"中国精致生活►","AppName":"物道","AppID":"gh_f27ca3831aea","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL0R0a2ljUHM4d0x0aWNWMVdpYzlsN3pNTlhpY2FlQk5vVXByTzd6aWJiN2JIVVBkaEVheUM2NE1HZnZpYU1zclpvSVNtU2pQOVZBZHh3VUF5bDdTSTA1MW9vRjNBLzA_d3hfZm10PWpwZWc=","Intro":"生活，不过一碗人间烟火。","PubAt":"1506564000","URL":"http://mp.weixin.qq.com/s?__biz=MzA4MzI1Mzc2MA==\u0026mid=2653818101\u0026idx=1\u0026sn=b88229172064f71ab9c7d21ef680c722\u0026chksm=8423e39ab3546a8c9fddff0dc2473e6b00946b1d6291d539f469dc3f3c13f4aa935def916b59#rd","CreatedAt":"2017-09-29T00:12:31.8671428+08:00","UpdatedAt":"2017-10-01T23:30:05.4893286+08:00","DeletedAt":null},{"ID":7,"Title":"世界上最惊艳的12张照片，绝对震撼！你见过吗？","Author":"","AppName":"无罪","AppID":"gh_b1a7224ce2f4","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1VOQmlhcFlDZm00bzdKM0V5aWN6WEFpY3l0dXNpYkVmWFhpYjBqMEFBZkJSTktpY2NibjJucGliZVlhSHR1ckdhRVNLVDFBd2hrSE1kQnA3VjZLRmZPbUFzaWFraWNnLzA_d3hfZm10PWpwZWc=","Intro":"保持对世界的好奇心，让生活充满欢乐。","PubAt":"2017-09-28","URL":"http://mp.weixin.qq.com/s?__biz=MjM5NzgzNTUyNA==\u0026mid=2650373542\u0026idx=1\u0026sn=a419c3bf6461ecaf4c6fc43525cd8204\u0026chksm=bede3bcd89a9b2dba5e293bc6815b3a685b7a20169fadcf7899ce157c15a7172fd7ad966e1dd#rd","CreatedAt":"2017-09-29T00:20:17.0817516+08:00","UpdatedAt":"2017-09-29T00:20:17.0927522+08:00","DeletedAt":null},{"ID":8,"Title":"关于新书，已经有了18万字。","Author":"梦入神机","AppName":"梦入神机","AppID":"gh_af6d690a9029","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1dVaWFid2Q1TTNlcTNpYm9UYXQwM3RyaFpJcm1YQWh0NUxTN2I1RmZCZmJsTjk3Z080d0J6d0pGVWJraWNpY3J2cW5MN1NSOGliZEJrT2J2bURuRWljbmFLdlR3LzA_d3hfZm10PWpwZWc=","Intro":"我以前说过8月份开新书的事情，实际上","PubAt":"2017-08-28","URL":"http://mp.weixin.qq.com/s?__biz=MzA3NDA4MjYxNQ==\u0026mid=2454021024\u0026idx=1\u0026sn=8bb3f5c901b8a17781ce87765c438ae0\u0026chksm=88b4f3f8bfc37aeec7a0afdea9de96b770b3a7064ed43104604bc0daeaa14605e7ee964a0c6c#rd","CreatedAt":"2017-09-29T00:24:50.8014074+08:00","UpdatedAt":"2017-09-29T00:24:50.811408+08:00","DeletedAt":null},{"ID":9,"Title":"降央卓玛一曲《国家》喜迎国庆，好听！","Author":"","AppName":"每日音乐","AppID":"gh_4924dfc59a65","Cover":"http://localhost:1323/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkR0h1dTk5WDE3UW42WW5kRTZBaFR3UVh4TE1hZEF2SllKTHVGbGlhNkFQbU1jaWJPeW9KNDVjNzNMc2xIZjZRTGVSbmVxVFdNUkdBTWcvMD93eF9mbXQ9anBlZw==","Intro":"..","PubAt":"2017-10-01","URL":"http://mp.weixin.qq.com/s?__biz=MzA3MDMwNjgwMw==\u0026mid=2655648485\u0026idx=1\u0026sn=27f1299a3d7650302fa2dd84fb6cfede\u0026chksm=84838479b3f40d6f4d78a5822efed106c67d21380a5fb19bdefcdc2a5224c7dcba133d6c508e#rd","CreatedAt":"2017-10-01T21:06:42.22425+08:00","UpdatedAt":"2017-10-01T21:06:42.2352506+08:00","DeletedAt":null}]

      news.getNew(api,function(err,data){
            site.articles = data
      })
       for (let i = 1; i <= 20; i++) {
        this.items.push(i + ' - keep walking, be 2 with you.')
      }
      this.top = 1
      this.bottom = 20
    },
    methods: {
      refresh(done) {
        console.log(' refresh..')

        setTimeout(() => {
           this.counter = 1
            var api = 'http://localhost:8888/hot';
            var site = this
            news.getNew(api,function(err,data){
                  site.articles = data
            })
          // done()
           let start = this.top - 1
          for (let i = start; i > start - 10; i--) {
            this.items.splice(0, 0, i + ' - keep walking, be 2 with you.')
          }
          this.top = this.top - 10;
          done()
        }, 1500)
      },
      infinite(done) {
        console.log('infinite called..')
        
        setTimeout(() => {
            var api = 'http://localhost:8888/hot';
            var tempdata 
            var site = this
            news.getNew(api,function(err,data){
                for(var t=0;t<data.length;t++){
                  site.articles.push(data[t])
                }
            })
          // done()
          let start = this.bottom + 1
          for (let i = start; i < start + 10; i++) {
            this.items.push(i + ' - keep walking, be 2 with you.')
          }
          this.bottom = this.bottom + 10
          done()
        }, 1500)
      },
      onItemClick(index, item) {
        console.log(index)
      }
    },

}
</script>



