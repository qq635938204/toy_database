<!-- filepath: /e:/Code/toy_database/web/vite-toy-app/src/components/ToyCard.vue -->
<template>
  <div class="toy-card">
    <img :src="getImageUrl(toy.image_path)" alt="Toy Image" />
    <h2>{{ toy.name }}</h2>
    <p>MB编号: {{ toy.mb_num }}</p>
    <p>购买时间: {{ toy.buy_date }}</p>
    <p>价格: {{ toy.price }}</p>
    <p>库存: {{ toy.stock }}</p>
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
  },
  methods: {
    getImageUrl(imagePath) {
      return `http://127.0.0.1:9988/${imagePath}`
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
  width: 300px;
  height: 200px;
  object-fit: cover; /* 保持图片比例 */
}
.toy-card h2 {
  margin: 8px 0;
}
.toy-card p {
  margin: 4px 0;
}
</style>