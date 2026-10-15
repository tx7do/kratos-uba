import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const featureFlag: RouteRecordRaw[] = [
  {
    path: '/feature-flags',
    name: 'FeatureFlagManagement',
    component: BasicLayout,
    redirect: '/feature-flags/list',
    meta: {
      order: 1004,
      icon: 'lucide:toggle-right',
      title: $t('menu.featureFlag.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'FeatureFlagList',
        meta: {
          order: 1,
          icon: 'lucide:split',
          title: $t('menu.featureFlag.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/feature-flag/index.vue'),
      },
    ],
  },
];

export default featureFlag;
