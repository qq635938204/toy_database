<!-- filepath: /e:/Code/toy_database/web/vite-toy-app/src/components/ToyCard.vue -->
<template>
    <div class="toy-card">
      <img :src="toy.image" alt="Toy Image" />
      <h2>{{ toy.brand }}</h2>
      <p>购买时间: {{ toy.purchaseDate }}</p>
    </div>
  </template>
  
  <script>
  export default {
    name: 'ToyCard',
    props: {
      toy: {
        type: Object,
        required: true
      }
    }
  }
  </script>
  
  <style scoped>
  .toy-card {
    border: 1px solid #ccc;
    padding: 16px;
    margin: 16px;
    text-align: center;
  }
  .toy-card img {
    max-width: 100%;
    height: auto;
  }
  </style>