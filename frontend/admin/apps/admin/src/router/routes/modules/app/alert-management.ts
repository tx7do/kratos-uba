import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const alert: RouteRecordRaw[] = [
  {
    path: '/alert',
    name: 'AlertManagement',
    component: BasicLayout,
    redirect: '/alert/rules',
    meta: {
      order: 1005,
      icon: 'lucide:bell-ring',
      title: $t('menu.alert.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'rules',
        name: 'AlertRuleManagement',
        meta: {
          order: 1,
          icon: 'lucide:bell-plus',
          title: $t('menu.alert.rule'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/alert/rule/index.vue'),
      },
    ],
  },
];

export default alert;
