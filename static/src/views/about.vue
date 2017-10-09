<template>
  <div style="max-width:480px;margin:0 auto;">
     <p class="example-list-item" v-for="item in list" v-text="item"></p>
    <infinite-loading @infinite="infiniteHandler" :distance="distance" ref="infiniteLoading"></infinite-loading>
  </div>
</template>

<script>
  import http from 'axios'
  import InfiniteLoading from 'vue-infinite-loading'
  export default {
    name: 'app',
    data(){
        return{
            distance: 100,
            list: []
        }
    },
    methods: {
      
        infiniteHandler: function ($state) {
          console.log("dododododo...");
          if (this.list.length > 50) {
            $state.complete();
          } else {
            setTimeout(function () {
              var temp = [];
              for (var i = this.list.length; i <= this.list.length + 10; i++) {
                temp.push(i);
              }
              this.list = this.list.concat(temp);
              $state.loaded();
            }.bind(this), 1000);
          }
        }
    },
    components: {
      InfiniteLoading
    }
  }
</script>

<style>
   body{
      margin: 0;
    }
    .example-list-item{
      margin: 0;
      padding: 0 10px;
      font-size: 14px;
      line-height: 40px;
      color: #666;
      background-color: #fafafa;
      border-top: 1px solid #fff;
      border-bottom: 1px solid #e3e3e3;
    }
    .example-list-item::before{
      content: 'Line: ';
    }
</style>