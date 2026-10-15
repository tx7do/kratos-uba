import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 用户洞察：以个体视角串起用户 360 与会话回放。
// 单页模块坍缩为叶子，子路由用绝对路径保持旧 URL 不变。
const AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const insight: RouteRecordRaw[] = [
  {
    path: '/insight',
    name: 'UserInsights',
    component: BasicLayout,
    redirect: '/user-360',
    meta: {
      order: 300,
      icon: 'lucide:telescope',
      title: $t('menu.insight.moduleName'),
    },
    children: [
      {
        path: '/user-360',
        name: 'UserProfileManagement',
        meta: {
          order: 300,
          icon: 'lucide:user-round-search',
          title: $t('menu.userProfile.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/user-profile/index.vue'),
      },
      {
        path: '/replays',
        name: 'ReplayManagement',
        meta: {
          order: 310,
          icon: 'lucide:monitor-play',
          title: $t('menu.replay.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/replay/index.vue'),
      },
    ],
  },
];

export default insight;
