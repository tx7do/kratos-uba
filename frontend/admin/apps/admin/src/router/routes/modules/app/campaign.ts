import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const campaign: RouteRecordRaw[] = [
  {
    path: '/campaigns',
    name: 'CampaignManagement',
    component: BasicLayout,
    redirect: '/campaigns/list',
    meta: {
      order: 1012,
      icon: 'lucide:megaphone',
      title: $t('menu.campaign.moduleName'),
      authority: ['sys:platform_admin', 'sys:tenant_manager'],
    },
    children: [
      {
        path: 'list',
        name: 'CampaignList',
        meta: {
          order: 1,
          icon: 'lucide:target',
          title: $t('menu.campaign.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        component: () => import('#/views/app/campaign/index.vue'),
      },
    ],
  },
];

export default campaign;
