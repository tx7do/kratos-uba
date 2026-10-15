import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const report: RouteRecordRaw[] = [
  {
    path: '/report-subscriptions',
    name: 'ReportManagement',
    component: BasicLayout,
    redirect: '/report-subscriptions/list',
    meta: {
      order: 240,
      icon: 'lucide:calendar-clock',
      title: $t('menu.report.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'ReportSubscriptionList',
        meta: {
          order: 1,
          icon: 'lucide:mail-check',
          title: $t('menu.report.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/report/index.vue'),
      },
    ],
  },
];

export default report;
