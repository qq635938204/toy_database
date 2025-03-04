<!-- filepath: /e:/Code/toy_database/web/vite-toy-app/src/App.vue -->
<template>
    <div id="app">
      <ToyList />
    </div>
  </template>
  
  <script>
  import ToyList from './views/ToyList.vue'
  
  export default {
    name: 'App',
    components: {
      ToyList
    }
  }
  </script>
  
  <style>
  /* 添加全局样式 */
  </style>