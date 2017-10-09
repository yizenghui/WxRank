<template>
  <div id="app">
        <header-tab></header-tab>
    <router-view/>
  </div>
</template>

<script>
import HeaderTab from '@/components/Header';
export default {
  name: 'app',
  components:{
    HeaderTab,
  }
}
</script>
