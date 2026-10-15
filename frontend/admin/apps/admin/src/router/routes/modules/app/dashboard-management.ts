import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const dashboard: RouteRecordRaw[] = [
  {
    path: '/custom-dashboards',
    name: 'CustomDashboardManagement',
    component: BasicLayout,
    redirect: '/custom-dashboards/list',
    meta: {
      order: 230,
      icon: 'lucide:layout-dashboard',
      title: $t('menu.dashboard.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'CustomDashboardList',
        meta: {
          order: 1,
          icon: 'lucide:gallery-horizontal-end',
          title: $t('menu.dashboard.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/dashboard/index.vue'),
      },
      {
        path: 'view/:id',
        name: 'CustomDashboardView',
        meta: {
          order: 2,
          hideInMenu: true,
          icon: 'lucide:chart-line',
          title: $t('menu.dashboard.view'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/dashboard/view.vue'),
      },
    ],
  },
];

export default dashboard;
