import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const behavior: AppRouteModule = {
  path: '/behavior-analysis',
  name: 'BehaviorAnalysis',
  component: LAYOUT,
  redirect: '/behavior-analysis/index',
  meta: {
    orderNo: 2,
    hideChildrenInMenu: true,
    icon: 'ant-design:radar-chart-outlined',
    title: t('routes.app.behavior.moduleName'),
  },
  children: [
    {
      path: 'index',
      name: 'BehaviorAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behavior.moduleName'),
        hideMenu: true,
      },
    },
  ],
};

export default behavior;
