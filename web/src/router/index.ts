import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

export enum RouteName {
  Index = 'index',
  About = 'about',
  PostList = 'pages',
  Contact = 'contact'
}

export const constructPostDetailPath = (id: string) => `/posts/${id}` as const
export const constructPostListPath = (pageNumber: string | number) =>
  `/pages/${pageNumber}` as const

const MainPage = () => import('/@/views/MainPage.vue')
const AboutPage = () => import('/@/views/AboutPage.vue')
const PostListPage = () => import('/@/views/PostListPage.vue')

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: RouteName.Index,
    component: MainPage
  },
  {
    path: '/about',
    name: RouteName.About,
    component: AboutPage
  },
  {
    path: constructPostListPath(':pageNumber(\\d+)'),
    name: RouteName.PostList,
    component: PostListPage
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
