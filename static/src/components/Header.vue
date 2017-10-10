<template>
<div>

<!-- 
  <nav class="navbar navbar-expand navbar-dark bg-dark">
    <a class="navbar-brand" href="#">Rf</a>

    <div class="collapse navbar-collapse">
      <div class="navbar-nav">
            <router-link class="nav-item nav-link"  to="/">Hot</router-link>
            <router-link class="nav-item nav-link"  to="/new">New</router-link>
            <router-link class="nav-item nav-link" to="/post">Post</router-link>
      </div>
    </div>
    </nav> -->

    <div class="container">
	<div class="row justify-content-center">


<div class="col-md-8 col-sm-9">

<ul class="nav nav-tabs">
  <li class="nav-item">
     <router-link class="nav-item nav-link"  to="/">微信热文</router-link>
  </li>
  <li class="nav-item">
    <router-link class="nav-item nav-link"  to="/new">最新收录</router-link>
  </li>
  <!-- <li class="nav-item">
    <router-link class="nav-item nav-link" to="/post">提交</router-link>
  </li> -->
</ul>
</div>
</div>
</div>
</div>
</template>
<style>
  .nav-bar {
    width: 100%;
    height: 44px;
    box-shadow: 0 2px 10px 0 rgba(0,0,0,0.1);
    background-color: #fff;
    position: relative;
    z-index: 1000;
  }
  .nav-bar .btn-back {
    height: 100%;
    position: absolute;
    top: 0;
    left: 10px;
  }
  .nav-bar .arrow-back {
    width: 8px;
    height: 16px;
    float: left;
    margin-top: 14px;
    margin-right: 15px;
  }
  .nav-bar .title {
    width: 100%;
    text-align: center;
    font-size: 16px;
    line-height: 46px;
    color: #666;
  }
</style>

<script>
  export default {
    props: {
      title: {
        type: String
      },
      backTo: {
        type: String,
        default: '/'
      }
    },
    methods: {
      back () {
        this.$router.push(this.backTo)
      }
    }
  }
</script>