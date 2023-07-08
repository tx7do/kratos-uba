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
    icon: 'ant-design:radar-chart-outlined',
    title: t('routes.app.behaviorAnalysis.moduleName'),
  },
  children: [
    {
      path: 'event',
      name: 'EventAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/event/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.event'),
      },
    },
    {
      path: 'retain',
      name: 'RetainAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/retain/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.retain'),
      },
    },
    {
      path: 'funnel',
      name: 'FunnelAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/funnel/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.funnel'),
      },
    },
    {
      path: 'path',
      name: 'PathAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/path/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.path'),
      },
    },
  ],
};

export default behavior;
