<!-- filepath: /e:/Code/toy_database/web/vite-toy-app/src/components/ToyCard.vue -->
<template>
  <div class="toy-card">
    <img :src="getImageUrl(toy.image_path)" alt="Toy Image" />
    <div class="toy-header">
      <h2>{{ toy.year_on_base }} {{ toy.name }}</h2>
    </div>
    <div class="toy-details">
      <p>MB编号: {{ toy.mb_num }}</p>
      <p>购买时间: {{ toy.buy_date }}</p>
      <p class="toy-description">简介: {{ toy.description }}</p>
    </div>
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
  text-align: left; /* 左对齐 */
}
.toy-card img {
  width: 300px;
  height: 200px;
  object-fit: cover; /* 保持图片比例 */
}
.toy-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.toy-header p {
  margin: 0;
}
.toy-header h2 {
  margin: 8px 0;
}
.toy-details {
  display: flex;
  flex-direction: column; /* 垂直排列 */
}
.toy-details p {
  margin: 4px 0;
}
.toy-description {
  width: 300px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3; /* 限制显示3行 */
  line-clamp: 3; /* 限制显示3行 */
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: normal;
  word-break: break-all; /* 强制换行 */
  color: #888;/* 设置颜色淡一点 */
}
</style>