import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const userProfile: RouteRecordRaw[] = [
  {
    path: '/user-360',
    name: 'UserProfileManagement',
    component: BasicLayout,
    redirect: '/user-360/profile',
    meta: {
      order: 1008,
      icon: 'lucide:user-round-search',
      title: $t('menu.userProfile.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'profile',
        name: 'UserProfile360',
        meta: {
          order: 1,
          icon: 'lucide:id-card',
          title: $t('menu.userProfile.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/user-profile/index.vue'),
      },
    ],
  },
];

export default userProfile;
