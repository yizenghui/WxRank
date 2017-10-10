<template>

    <div class="container">
	<div class="row justify-content-center">


<div class="col-md-8 col-sm-9">  <br>
<form @submit.prevent="submit">
  <div class="form-group">
    <label for="exampleInputEmail1">微信文章地址</label>
    <input  v-model="url" type="url" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter URL">
    <small id="emailHelp" class="form-text text-muted">如果您想要与小伙伴们分享美文，请粘贴微信文章地址.</small>
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>
</div>
</div>
</div>
</template>

<script>
import news from '../api/news.js';

  export default {
    name: 'app',
    data(){
        return{
           url:""
        }
    },
    methods: {
       submit: function(event) {
            
              news.post(this.url,function(err,data){
                if(data == "0"){
                  alert("收录失败")
                }else{
                  alert("收录成功")
                }
              });

           
        }
        
    },
    components: {
      
    }
  }
</script>
