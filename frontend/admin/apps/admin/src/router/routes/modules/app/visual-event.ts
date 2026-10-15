import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const visualEvent: RouteRecordRaw[] = [
  {
    path: '/visual-events',
    name: 'VisualEventManagement',
    component: BasicLayout,
    redirect: '/visual-events/list',
    meta: {
      order: 1013,
      icon: 'lucide:mouse-pointer-click',
      title: $t('menu.visualEvent.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'VisualEventList',
        meta: {
          order: 1,
          icon: 'lucide:crosshair',
          title: $t('menu.visualEvent.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/visual-event/index.vue'),
      },
    ],
  },
];

export default visualEvent;
