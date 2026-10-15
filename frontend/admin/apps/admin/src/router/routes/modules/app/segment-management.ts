import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const segment: RouteRecordRaw[] = [
  {
    path: '/segments',
    name: 'SegmentManagement',
    component: BasicLayout,
    redirect: '/segments/list',
    meta: {
      order: 200,
      icon: 'lucide:users-round',
      title: $t('menu.segment.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'SegmentList',
        meta: {
          order: 1,
          icon: 'lucide:list-filter',
          title: $t('menu.segment.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/segment/index.vue'),
      },
    ],
  },
];

export default segment;
