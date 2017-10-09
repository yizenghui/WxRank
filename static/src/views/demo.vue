<template>
  <div>     
    <nav-bar title="Refresh & Infinite"></nav-bar>

    <scroller style="padding-top: 44px; max-width:480px;"
              :on-refresh="refresh"
              :on-infinite="infinite"
              width="480px"
              ref="my_scroller">

      <div v-for="(item, index) in items" @click="onItemClick(index, item)"
          class="row" :class="{'grey-bg': index % 2 == 0}">
        {{ item }}
      </div>
    </scroller>
  </div>
</template>
<style>
/*
html, body {
  margin: 0;
}

* {
  box-sizing: border-box;
}

.row {
  width: 100%;
  height: 50px;
  padding: 10px 0;
  font-size: 16px;
  line-height: 30px;
  text-align: center;
  color: #444;
  background-color: #fff;
}

.grey-bg {
  background-color: #eee;
}

.header {
  position: fixed;
  top: 0;
  left: 0;
  height: 44px;
  width: 100%;
  box-shadow: 0 2px 10px 0 rgba(0,0,0,0.1);
  background-color: #fff;
  z-index: 1000;
  color: #666;
}

.header > .title {
  font-size: 16px;
  line-height: 44px;
  text-align: center;
  margin: 0 auto;
}
*/
</style>

<script>
import NavBar from '../components/NavBar.vue';

  export default {

    components: {
      NavBar
    },
    data () {
      return {
        items: []
      }
    },  
    mounted:function() {
      for (let i = 1; i <= 20; i++) {
        this.items.push(i + ' - keep walking, be 2 with you.')
      }
      this.top = 1
      this.bottom = 20
    },
    methods: {
      refresh(done) {
        setTimeout(() => {
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
    }
  }
</script>