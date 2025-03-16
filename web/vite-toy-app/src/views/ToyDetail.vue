<template>
  <div class="toy-detail">
    <h1>玩具详情</h1>
    <div v-if="toy">
      <img :src="getImageUrl(toy.image_path)" alt="Toy Image" />
      <h2>{{ toy.name }}</h2>
      <p>MB编号: {{ toy.mb_num }}</p>
      <p>购买时间: {{ toy.buy_date }}</p>
      <p>简介: {{ toy.description }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ToyDetail',
  data() {
    return {
      toy: null
    }
  },
  created() {
    this.fetchToyDetail()
  },
  methods: {
    async fetchToyDetail() {
      const toyId = this.$route.params.id
      try {
        const response = await axios.get(`http://127.0.0.1:9988/v1/toy/${toyId}`)
        if (response.data.code === 0) {
          this.toy = response.data.data
        } else {
          console.error('Error fetching toy detail:', response.data.message)
        }
      } catch (error) {
        console.error('Error fetching toy detail:', error)
      }
    },
    getImageUrl(imagePath) {
      return `http://127.0.0.1:9988/${imagePath}`
    }
  }
}
</script>

<style scoped>
.toy-detail {
  padding: 16px;
  background-color: #1c1c1c;
  color: white;
}
.toy-detail img {
  width: 300px;
  height: 200px;
  object-fit: cover;
}
</style>