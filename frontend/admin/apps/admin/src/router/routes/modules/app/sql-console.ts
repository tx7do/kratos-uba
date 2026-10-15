import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const sqlConsole: RouteRecordRaw[] = [
  {
    path: '/sql-console',
    name: 'SqlConsoleManagement',
    component: BasicLayout,
    redirect: '/sql-console/query',
    meta: {
      order: 420,
      icon: 'lucide:database-zap',
      title: $t('menu.sqlConsole.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'query',
        name: 'SqlConsoleQuery',
        meta: {
          order: 1,
          icon: 'lucide:square-terminal',
          title: $t('menu.sqlConsole.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/sql-console/index.vue'),
      },
    ],
  },
];

export default sqlConsole;
