import type { RouteRecordRaw } from 'vue-router'
import { RouteName } from '.'

export const SettingsRouteName = {
  EditPost: 'posts',
  ManagePost: 'dashboard',
  UserProfile: 'profile'
} as const

type RouteName = (typeof SettingsRouteName)[keyof typeof SettingsRouteName]

const constructSettingsPath = (name: RouteName) => `/settings/${name}` as const

const EditPostPage = () => import('/@/views/Settings/EditPostPage.vue')
const ManagePostPage = () => import('/@/views/Settings/ManagePostPage.vue')
const UserProfilePage = () => import('/@/views/Settings/UserProfilePage.vue')

export const settingsRoutes: RouteRecordRaw[] = [
  {
    path: `${constructSettingsPath('posts')}/:postId(\\d+)`,
    name: SettingsRouteName.EditPost,
    component: EditPostPage
  },
  {
    path: constructSettingsPath('dashboard'),
    name: SettingsRouteName.ManagePost,
    component: ManagePostPage
  },
  {
    path: constructSettingsPath('profile'),
    name: SettingsRouteName.UserProfile,
    component: UserProfilePage
  }
]
