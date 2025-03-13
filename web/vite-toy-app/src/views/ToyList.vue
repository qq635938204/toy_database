<template>
  <div class="toy-list">
    <div class="carousel">
      <img :src="images[currentImageIndex]" alt="Carousel Image" />
      <div class="carousel-indicators">
        <span v-for="(image, index) in images" :key="index" :class="{ active: index === currentImageIndex }"></span>
      </div>
    </div>
    <div class="search-container">
      <input type="text" v-model="searchQuery" placeholder="搜索.." />
      <button @click="searchToys"><i class="fas fa-search"></i></button>
      <span class="fixed-text">mbx.zhang3woshi.cn</span>
    </div>
    <div class="toy-cards">
      <ToyCard v-for="toy in filteredToys" :key="toy.id" :toy="toy" />
    </div>
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
    },
    searchToys() {
      // 搜索玩具的逻辑
      console.log('搜索玩具:', this.searchQuery)
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
  flex-direction: column;
  align-items: center;
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
.carousel-indicators {
  position: absolute;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 5px;
}
.carousel-indicators span {
  display: block;
  width: 10px;
  height: 10px;
  background-color: #ccc;
  border-radius: 50%;
}
.carousel-indicators span.active {
  background-color: #333;
}
.search-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 16px;
}
.search-container input {
  padding: 8px;
  width: 300px; /* 固定宽度 */
  box-sizing: border-box;
}
.search-container button {
  padding: 8px;
  background-color: #333;
  color: white;
  border: none;
  cursor: pointer;
}
.search-container button i {
  font-size: 16px;
}
.fixed-text {
  margin-left: auto;
  padding-left: 16px;
  padding-right: 16px; /* 增加右侧距离 */
  font-size: 25px;
  font-weight: bold; /* 加粗显示 */
  color: #000; /* 正常颜色 */
}
.toy-cards {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start; /* 从左到右显示 */
  gap: 16px;
  width: 100%; /* 确保卡片列表占满可用宽度 */
}
</style>