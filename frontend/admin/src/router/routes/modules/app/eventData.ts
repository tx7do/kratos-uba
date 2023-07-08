import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const eventData: AppRouteModule = {
  path: '/event-manage',
  name: 'EventManage',
  component: LAYOUT,
  meta: {
    orderNo: 4,
    icon: 'ant-design:api-outlined',
    title: t('routes.app.eventData.moduleName'),
  },
  children: [
    {
      path: 'report',
      name: 'Report',
      component: () => import('/@/views/app/event_data/report/index.vue'),
      meta: {
        icon: 'ant-design:node-index-outlined',
        title: t('routes.app.eventData.report'),
      },
    },
    {
      path: 'realtime',
      name: 'RealtimeData',
      component: () => import('/@/views/app/event_data/realtime/index.vue'),
      meta: {
        icon: 'ant-design:build-outlined',
        title: t('routes.app.eventData.realtimeData'),
      },
    },
    {
      path: 'debug',
      name: 'DebugManage',
      component: () => import('/@/views/app/event_data/debug/index.vue'),
      meta: {
        icon: 'ant-design:bug-outlined',
        title: t('routes.app.eventData.debug'),
      },
    },
  ],
};

export default eventData;
