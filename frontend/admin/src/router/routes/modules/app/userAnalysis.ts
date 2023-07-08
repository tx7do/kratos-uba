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
    hideChildrenInMenu: true,
    icon: 'ant-design:pie-chart-outlined',
    title: t('routes.app.user.moduleName'),
  },
  children: [
    {
      path: 'index',
      name: 'UserAnalysisPage',
      component: () => import('/@/views/app/user_analysis/index.vue'),
      meta: {
        icon: 'ant-design:pie-chart-outlined',
        title: t('routes.app.user.moduleName'),
        hideMenu: true,
      },
    },
  ],
};

export default kanban;
