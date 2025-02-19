import type { RouteRecordRaw } from 'vue-router'

export const SettingsRouteName = {
  EditPost: 'posts',
  ManagePost: 'dashboard',
  UserProfile: 'profile'
} as const

type RouteName = (typeof SettingsRouteName)[keyof typeof SettingsRouteName]
type StaticRoutes = Omit<typeof SettingsRouteName, 'EditPost'>
type StaticRouteName = StaticRoutes[keyof StaticRoutes]
type DynamicRouteName = Exclude<RouteName, StaticRouteName>

const constructSettingsStaticPath = (routeName: StaticRouteName) =>
  `/settings/${routeName}` as const
const constructSettingsDynamicPath = (
  routeName: DynamicRouteName,
  id: string
) => `/settings/${routeName}/${id}` as const

const EditPostPage = () => import('/@/views/Settings/EditPostPage.vue')
const ManagePostPage = () => import('/@/views/Settings/ManagePostPage.vue')
const UserProfilePage = () => import('/@/views/Settings/UserProfilePage.vue')

export const settingsRoutes: RouteRecordRaw[] = [
  {
    path: constructSettingsDynamicPath('posts', ':postId(\\d+)'),
    name: SettingsRouteName.EditPost,
    component: EditPostPage
  },
  {
    path: constructSettingsStaticPath('dashboard'),
    name: SettingsRouteName.ManagePost,
    component: ManagePostPage
  },
  {
    path: constructSettingsStaticPath('profile'),
    name: SettingsRouteName.UserProfile,
    component: UserProfilePage
  }
]
