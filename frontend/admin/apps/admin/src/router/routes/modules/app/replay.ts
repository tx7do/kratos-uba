import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const replay: RouteRecordRaw[] = [
  {
    path: '/replays',
    name: 'ReplayManagement',
    component: BasicLayout,
    redirect: '/replays/list',
    meta: {
      order: 310,
      icon: 'lucide:monitor-play',
      title: $t('menu.replay.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'ReplaySessionList',
        meta: {
          order: 1,
          icon: 'lucide:history',
          title: $t('menu.replay.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/replay/index.vue'),
      },
    ],
  },
];

export default replay;
