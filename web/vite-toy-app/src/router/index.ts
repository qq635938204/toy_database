import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import ToyList from '../views/ToyList.vue'
import ToyDetail from '../views/ToyDetail.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: ToyList },
  { path: '/toydetail/:id', component: ToyDetail, name: 'ToyDetail' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router