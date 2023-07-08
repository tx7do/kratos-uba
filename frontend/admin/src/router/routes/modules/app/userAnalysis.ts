import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const kanban: AppRouteModule = {
  path: '/user-analysis',
  name: 'UserAnalysis',
  component: LAYOUT,
  redirect: '/user-analysis/index',
  meta: {
    orderNo: 3,
    icon: 'ant-design:pie-chart-outlined',
    title: t('routes.app.userAnalysis.moduleName'),
  },
  children: [
    {
      path: 'attribute',
      name: 'UserAttributeAnalysisPage',
      component: () => import('/@/views/app/user_analysis/attribute/index.vue'),
      meta: {
        icon: 'ant-design:pie-chart-outlined',
        title: t('routes.app.userAnalysis.attribute'),
      },
    },
    {
      path: 'group',
      name: 'UserGroupPage',
      component: () => import('/@/views/app/user_analysis/group/index.vue'),
      meta: {
        icon: 'ant-design:pie-chart-outlined',
        title: t('routes.app.userAnalysis.group'),
      },
    },
    {
      path: 'list',
      name: 'UserListPage',
      component: () => import('/@/views/app/user_analysis/list/index.vue'),
      meta: {
        icon: 'ant-design:pie-chart-outlined',
        title: t('routes.app.userAnalysis.list'),
      },
    },
    {
      path: 'event',
      name: 'UserEventDetail',
      component: () => import('/@/views/app/user_analysis/event/index.vue'),
      meta: {
        icon: 'ant-design:pie-chart-outlined',
        title: t('routes.app.userAnalysis.event'),
      },
    },
  ],
};

export default kanban;
