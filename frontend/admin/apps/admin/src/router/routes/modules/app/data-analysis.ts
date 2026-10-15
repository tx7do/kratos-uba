import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 数据分析：26 个分析页按 README 三分类（通用/深度洞察/游戏专项）收成二级分组。
// 子路由用绝对路径保持旧 URL 不变（/data-analysis/event-trend 等），分组节点仅是菜单层级。
const BASIC_AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const dataAnalysis: RouteRecordRaw[] = [
  {
    path: '/data-analysis',
    name: 'DataAnalysis',
    component: BasicLayout,
    redirect: '/data-analysis/event-trend',
    meta: {
      order: 100,
      icon: 'lucide:chart-bar',
      title: $t('menu.dataAnalysis.moduleName'),
      authority: BASIC_AUTH,
    },
    children: [
      {
        path: 'basic',
        name: 'DataAnalysisBasic',
        redirect: '/data-analysis/event-trend',
        meta: {
          order: 1,
          icon: 'lucide:chart-column',
          title: $t('menu.dataAnalysis.basic'),
          authority: BASIC_AUTH,
        },
        children: [
          {
            path: '/data-analysis/event-trend',
            name: 'EventTrendAnalysis',
            meta: {
              order: 1,
              icon: 'lucide:trending-up',
              title: $t('menu.dataAnalysis.eventTrend'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/event-trend/index.vue'),
          },
          {
            path: '/data-analysis/funnel',
            name: 'FunnelAnalysis',
            meta: {
              order: 2,
              icon: 'lucide:filter',
              title: $t('menu.dataAnalysis.funnel'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/funnel/index.vue'),
          },
          {
            path: '/data-analysis/retention',
            name: 'RetentionAnalysis',
            meta: {
              order: 3,
              icon: 'lucide:repeat',
              title: $t('menu.dataAnalysis.retention'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/retention/index.vue'),
          },
          {
            path: '/data-analysis/dimension-compare',
            name: 'DimensionCompare',
            meta: {
              order: 4,
              icon: 'lucide:bar-chart-3',
              title: $t('menu.dataAnalysis.dimensionCompare'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/dimension-compare/index.vue'),
          },
          {
            path: '/data-analysis/attribution',
            name: 'AttributionAnalysis',
            meta: {
              order: 5,
              icon: 'lucide:target',
              title: $t('menu.dataAnalysis.attribution'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/attribution/index.vue'),
          },
          {
            path: '/data-analysis/distribution',
            name: 'DistributionAnalysis',
            meta: {
              order: 6,
              icon: 'lucide:bar-chart-big',
              title: $t('menu.dataAnalysis.distribution'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/distribution/index.vue'),
          },
          {
            path: '/data-analysis/behavior-sequence',
            name: 'BehaviorSequence',
            meta: {
              order: 7,
              icon: 'lucide:git-commit-horizontal',
              title: $t('menu.dataAnalysis.behaviorSequence'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/behavior-sequence/index.vue'),
          },
          {
            path: '/data-analysis/click',
            name: 'ClickHeatmap',
            meta: {
              order: 8,
              icon: 'lucide:mouse-pointer-click',
              title: $t('menu.dataAnalysis.click'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/click/index.vue'),
          },
          {
            path: '/data-analysis/path-sankey',
            name: 'PathSankey',
            meta: {
              order: 9,
              icon: 'lucide:share-2',
              title: $t('menu.dataAnalysis.pathSankey'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/path-sankey/index.vue'),
          },
          {
            path: '/data-analysis/segmentation',
            name: 'Segmentation',
            meta: {
              order: 10,
              icon: 'lucide:users',
              title: $t('menu.dataAnalysis.segmentation'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/segmentation/index.vue'),
          },
        ],
      },
      {
        path: 'insight',
        name: 'DataAnalysisInsight',
        redirect: '/data-analysis/lifecycle',
        meta: {
          order: 2,
          icon: 'lucide:lightbulb',
          title: $t('menu.dataAnalysis.insight'),
          authority: BASIC_AUTH,
        },
        children: [
          {
            path: '/data-analysis/lifecycle',
            name: 'Lifecycle',
            meta: {
              order: 1,
              icon: 'lucide:git-fork',
              title: $t('menu.dataAnalysis.lifecycle'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/lifecycle/index.vue'),
          },
          {
            path: '/data-analysis/churn',
            name: 'Churn',
            meta: {
              order: 2,
              icon: 'lucide:user-minus',
              title: $t('menu.dataAnalysis.churn'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/churn/index.vue'),
          },
          {
            path: '/data-analysis/interval',
            name: 'Interval',
            meta: {
              order: 3,
              icon: 'lucide:timer',
              title: $t('menu.dataAnalysis.interval'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/interval/index.vue'),
          },
          {
            path: '/data-analysis/matrix',
            name: 'Matrix',
            meta: {
              order: 4,
              icon: 'lucide:layout-grid',
              title: $t('menu.dataAnalysis.matrix'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/matrix/index.vue'),
          },
          {
            path: '/data-analysis/revenue',
            name: 'Revenue',
            meta: {
              order: 5,
              icon: 'lucide:dollar-sign',
              title: $t('menu.dataAnalysis.revenue'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/revenue/index.vue'),
          },
          {
            path: '/data-analysis/session-analysis',
            name: 'SessionAnalysis',
            meta: {
              order: 6,
              icon: 'lucide:waypoints',
              title: $t('menu.dataAnalysis.sessionAnalysis'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/session-analysis/index.vue'),
          },
          {
            path: '/data-analysis/sessions',
            name: 'SessionManagement',
            meta: {
              order: 7,
              icon: 'lucide:clock',
              title: $t('menu.dataAnalysis.session'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/session/index.vue'),
          },
          {
            path: '/data-analysis/anomaly',
            name: 'Anomaly',
            meta: {
              order: 8,
              icon: 'lucide:activity',
              title: $t('menu.dataAnalysis.anomaly'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/anomaly/index.vue'),
          },
          {
            path: '/data-analysis/new-vs-old',
            name: 'NewVsOld',
            meta: {
              order: 9,
              icon: 'lucide:users-round',
              title: $t('menu.dataAnalysis.newVsOld'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/new-vs-old/index.vue'),
          },
        ],
      },
      {
        path: 'game',
        name: 'DataAnalysisGame',
        redirect: '/data-analysis/level-analysis',
        meta: {
          order: 3,
          icon: 'lucide:gamepad-2',
          title: $t('menu.dataAnalysis.game'),
          authority: BASIC_AUTH,
        },
        children: [
          {
            path: '/data-analysis/level-analysis',
            name: 'LevelAnalysis',
            meta: {
              order: 1,
              icon: 'lucide:triangle-right',
              title: $t('menu.dataAnalysis.levelAnalysis'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/level-analysis/index.vue'),
          },
          {
            path: '/data-analysis/whale-tier',
            name: 'WhaleTier',
            meta: {
              order: 2,
              icon: 'lucide:gem',
              title: $t('menu.dataAnalysis.whaleTier'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/whale-tier/index.vue'),
          },
          {
            path: '/data-analysis/ltv',
            name: 'LTV',
            meta: {
              order: 3,
              icon: 'lucide:trending-up',
              title: $t('menu.dataAnalysis.ltv'),
              authority: BASIC_AUTH,
            },
            component: () => import('#/views/app/data-analysis/ltv/index.vue'),
          },
          {
            path: '/data-analysis/server-retention',
            name: 'ServerRetention',
            meta: {
              order: 4,
              icon: 'lucide:server',
              title: $t('menu.dataAnalysis.serverRetention'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/server-retention/index.vue'),
          },
          {
            path: '/data-analysis/online-stats',
            name: 'OnlineStats',
            meta: {
              order: 5,
              icon: 'lucide:radio',
              title: $t('menu.dataAnalysis.onlineStats'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/online-stats/index.vue'),
          },
          {
            path: '/data-analysis/economy',
            name: 'Economy',
            meta: {
              order: 6,
              icon: 'lucide:coins',
              title: $t('menu.dataAnalysis.economy'),
              authority: BASIC_AUTH,
            },
            component: () =>
              import('#/views/app/data-analysis/economy/index.vue'),
          },
        ],
      },
      {
        path: '/data-analysis/realtime-screen',
        name: 'RealtimeScreen',
        meta: {
          order: 4,
          icon: 'lucide:radio',
          title: $t('menu.dataAnalysis.realtimeScreen'),
          authority: BASIC_AUTH,
        },
        component: () =>
          import('#/views/app/data-analysis/realtime-screen/index.vue'),
      },
    ],
  },
];

export default dataAnalysis;
