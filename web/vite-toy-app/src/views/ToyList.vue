<!-- filepath: e:\Code\toy_database\web\vite-toy-app\src\views\ToyList.vue -->
<template>
  <div class="toy-list">
    <div class="carousel">
      <img :src="images[currentImageIndex]" alt="Carousel Image" />
    </div>
    <input type="text" v-model="searchQuery" placeholder="搜索玩具..." />
    <ToyCard v-for="toy in filteredToys" :key="toy.id" :toy="toy" />
  </div>
</template>

<script>
import axios from 'axios'
import ToyCard from '../components/ToyCard.vue'

export default {
  name: 'ToyList',
  components: {
    ToyCard
  },
  data() {
    return {
      toys: [],
      searchQuery: '',
      images: [
        'http://127.0.0.1:9988/image/banner1.JPG',
        'http://127.0.0.1:9988/image/banner2.JPG',
        'http://127.0.0.1:9988/image/banner3.JPG'
      ],
      currentImageIndex: 0
    }
  },
  mounted() {
    this.fetchToys()
    this.startCarousel()
  },
  methods: {
    async fetchToys() {
      try {
        const response = await axios.get('http://127.0.0.1:9988/v1/toy/list')
        if (response.data.code === 0) {
          this.toys = response.data.data
        } else {
          console.error('Error fetching toys:', response.data.message)
        }
      } catch (error) {
        console.error('Error fetching toys:', error)
      }
    },
    startCarousel() {
      setInterval(() => {
        this.currentImageIndex = (this.currentImageIndex + 1) % this.images.length
      }, 3000) // 每3秒切换一次图片
    }
  },
  computed: {
    filteredToys() {
      return this.toys.filter(toy => 
        toy.name.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    }
  }
}
</script>

<style scoped>
.toy-list {
  display: flex;
  flex-wrap: wrap;
}
.carousel {
  width: 100%;
  height: 200px;
  margin-bottom: 16px;
  overflow: hidden;
  position: relative;
}
.carousel img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
input {
  margin-bottom: 16px;
  padding: 8px;
  width: 100%;
  box-sizing: border-box;
}
</style>