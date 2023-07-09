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
      path: 'retention',
      name: 'RetainAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/retain/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.retention'),
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
      path: 'interval',
      name: 'IntervalAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/interval/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.interval'),
      },
    },
    {
      path: 'distribution',
      name: 'DistributionAnalysisPage',
      component: () => import('/@/views/app/behavior_analysis/distribution/index.vue'),
      meta: {
        icon: 'ant-design:radar-chart-outlined',
        title: t('routes.app.behaviorAnalysis.distribution'),
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
