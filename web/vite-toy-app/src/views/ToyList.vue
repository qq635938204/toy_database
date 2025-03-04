<!-- filepath: /e:/Code/toy_database/web/vite-toy-app/src/views/ToyList.vue -->
<template>
    <div class="toy-list">
      <ToyCard v-for="toy in toys" :key="toy.id" :toy="toy" />
    </div>
  </template>
  
  <script>
  import ToyCard from '../components/ToyCard.vue'
  
  export default {
    name: 'ToyList',
    components: {
      ToyCard
    },
    data() {
      return {
        toys: [
          { id: 1, image: 'path/to/image1.jpg', brand: 'Brand A', purchaseDate: '2021-01-01' },
          { id: 2, image: 'path/to/image2.jpg', brand: 'Brand B', purchaseDate: '2021-02-01' },
          // 添加更多玩具信息
        ]
      }
    }
  }
  </script>
  
  <style scoped>
  .toy-list {
    display: flex;
    flex-wrap: wrap;
  }
  </style>