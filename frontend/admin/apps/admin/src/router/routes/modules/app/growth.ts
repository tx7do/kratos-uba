import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 增长运营：分群 → 触达 → 实验 → 看板 → 报告 → 告警的增长闭环。
// 单页模块坍缩为叶子（原模块根路径/name/title 保留，组件即原列表页）；
// 子路由用绝对路径保持旧 URL 不变。
const AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const growth: RouteRecordRaw[] = [
  {
    path: '/growth',
    name: 'GrowthOperations',
    component: BasicLayout,
    redirect: '/segments',
    meta: {
      order: 200,
      icon: 'lucide:trending-up',
      title: $t('menu.growth.moduleName'),
    },
    children: [
      {
        path: '/segments',
        name: 'SegmentManagement',
        meta: {
          order: 200,
          icon: 'lucide:users-round',
          title: $t('menu.segment.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/segment/index.vue'),
      },
      {
        path: '/campaigns',
        name: 'CampaignManagement',
        meta: {
          order: 210,
          icon: 'lucide:megaphone',
          title: $t('menu.campaign.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/campaign/index.vue'),
      },
      {
        path: '/feature-flags',
        name: 'FeatureFlagManagement',
        meta: {
          order: 220,
          icon: 'lucide:toggle-right',
          title: $t('menu.featureFlag.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/feature-flag/index.vue'),
      },
      {
        path: '/custom-dashboards',
        name: 'CustomDashboardManagement',
        meta: {
          order: 230,
          icon: 'lucide:layout-dashboard',
          title: $t('menu.dashboard.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/dashboard/index.vue'),
        children: [
          {
            path: 'view/:id',
            name: 'CustomDashboardView',
            meta: {
              order: 1,
              hideInMenu: true,
              icon: 'lucide:chart-line',
              title: $t('menu.dashboard.view'),
              authority: AUTH,
            },
            component: () => import('#/views/app/dashboard/view.vue'),
          },
        ],
      },
      {
        path: '/report-subscriptions',
        name: 'ReportManagement',
        meta: {
          order: 240,
          icon: 'lucide:calendar-clock',
          title: $t('menu.report.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/report/index.vue'),
      },
      {
        path: '/alert',
        name: 'AlertManagement',
        meta: {
          order: 250,
          icon: 'lucide:bell-ring',
          title: $t('menu.alert.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/alert/rule/index.vue'),
      },
    ],
  },
];

export default growth;
