import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const importJob: RouteRecordRaw[] = [
  {
    path: '/import-jobs',
    name: 'ImportJobManagement',
    component: BasicLayout,
    redirect: '/import-jobs/list',
    meta: {
      order: 415,
      icon: 'lucide:file-up',
      title: $t('menu.importJob.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'ImportJobList',
        meta: {
          order: 1,
          icon: 'lucide:upload',
          title: $t('menu.importJob.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/import-job/index.vue'),
      },
    ],
  },
];

export default importJob;
