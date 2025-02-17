import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

export enum RouteName {
  Index = 'index',
  About = 'about',
  PostList = 'page',
  Contact = 'contact'
}

export const constructPostDetailPagePath = (id: string) =>
  `/posts/${id}` as const

const MainPage = () => import('/@/views/MainPage.vue')

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: RouteName.Index,
    component: MainPage
  }
]

const routerHistory = createWebHistory(import.meta.env.BASE_URL)

const router = createRouter({
  history: routerHistory,
  routes
})

router.beforeEach(to => {
  // trailing slashを消す
  if (to.path !== '/' && to.path.endsWith('/')) {
    return to.path.slice(0, -1)
  }
  return true
})

export default router
