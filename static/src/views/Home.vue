//index.vue
<template>

    <div class="container">
	<div class="row justify-content-center">


<div class="col-md-9">


<!-- <br>
<div class="alert alert-secondary" role="alert">
  只提供信息跳转服务,文章观点仅代表来源公众号所属者,不代本平台立场
  <router-link class="alert-link" to="/post">分享文章</router-link>
</div> -->

<div>

      <div class="card mt-3" v-for="article in articles" :key="article.id">
        <div class="card-body">

          <h4 class="card-title">
             <a  v-bind:href="[article.URL]" v-on:click="cliLink(article)" >{{article.Title}}</a>  
          </h4>

          <div class="row">
              <div class=" mx-auto col-md-3 order-md-2">
                  <!-- <img class="card-img-top"  :src="article.Cover" alt="Card image cap"> -->
                  <a  v-bind:href="[article.URL]" v-on:click="cliLink(article)" >
                    <img class="card-img-top"  v-lazy="article.Cover" >
                  </a>  
              </div>
              <div class="col-md-9 order-md-1  text-md-left pr-md-5">

                  <p class="card-text">
                    <small class="text-muted">{{article.AppName}}</small>
                    <small class="text-muted">{{article.Author}}</small>
                    <small class="text-muted"><timeago :since="article.PubAt" class="text-muted" locale="zh-CN"></timeago></small>
                    
                  </p>
                
                
                  <p class="card-text">{{article.Intro}}</p>
                  <p class="card-text">
                    <small class="text-muted">
                      <a v-on:click.once="like(article)">({{article.Like}})喜欢</a>
                    </small>
                    <small class="text-muted">
                      <a v-on:click.once="hate(article)">({{article.Hate}})不喜欢</a>
                    </small>
                  </p>

              </div>
          </div>
        </div>
      </div>
      <infinite-loading @infinite="infiniteHandler" :distance="distance" ref="infiniteLoading">
      <span slot="no-results">数据都不见啦，赶紧把程序员哥哥叫回来！</span> 
      <span slot="no-more">到底啦！</span> 
      </infinite-loading> 
</div>
</div>

<div class="col-md-3">
  <br>
  <!--以后放东西-->

</div>

  </div>


      
    </div>
</template>

<script>

import InfiniteLoading from 'vue-infinite-loading'
import news from '../api/news.js';

export default {
  
  name: 'indexP',
  
  
    components: {
      InfiniteLoading
    },
    data () {
      return {
        articles: [],
        distance: 200,
        page:0,
      }
    },
    mounted() {
      // var api = 'http://localhost:8888/hot?limit=2';
      // var site = this
      // news.getNew(api,function(err,data){
      //       site.articles = data
      // })
      this.articles = JSON.parse(window.localStorage.getItem("hot_articles"))||[]
      this.page = parseInt(window.localStorage.getItem("hot_page")) || 0
    },
    methods: {

        cliLink:function(article){
          console.log(article)
          // alert(article.ID)
          return false
        },

        like:function(article){
          news.like(article.ID,function(err,data){
            article.Like = data.Like
          })
        },
        hate:function(article){
            news.hate(article.ID,function(err,data){
               article.Hate = data.Hate
            }) 
        },

        infiniteHandler: function ($state) {

          if (this.articles.length > 500) {
            $state.complete();
            // console.log("dont...");
          } else {

            //  setTimeout(function () {
              // console.log("dododododo...");

              var api = '/hot?limit=5&offset='+this.page*5;
              var site = this
              news.getNew(api,function(err,data){
                if(data.length>0){
                  for(var t=0;t<data.length;t++){
                    site.articles.push(data[t])
                  }
                  $state.loaded();
                  site.page ++
                }else{
                  $state.complete();
                }
              })
            // }.bind(this), 1000);
          }
        },

    },


    watch:{
      articles:function(){
        localStorage.setItem("hot_articles",JSON.stringify(this.articles))
        localStorage.setItem("hot_page",this.page)
      },
      }
}
</script>



