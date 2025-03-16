declare module 'vue-router' {
    import { Router, RouteRecordRaw, RouterOptions } from 'vue-router'
    export function createRouter(options: RouterOptions): Router
    export function createWebHistory(base?: string): RouterOptions['history']
    export { RouteRecordRaw }
  }