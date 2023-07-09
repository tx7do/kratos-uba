import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const metaData: AppRouteModule = {
  path: '/meta-manage',
  name: 'MetaManage',
  component: LAYOUT,
  meta: {
    orderNo: 5,
    hideChildrenInMenu: false,
    icon: 'ant-design:project-outlined',
    title: t('routes.app.metaData.moduleName'),
  },
  children: [
    {
      path: 'meta',
      name: 'MetaEvent',
      component: () => import('/@/views/app/meta_data/meta/index.vue'),
      meta: {
        icon: 'ant-design:control-outlined',
        title: t('routes.app.metaData.meta'),
      },
    },
    {
      path: 'event',
      name: 'EventProperty',
      component: () => import('/@/views/app/meta_data/event/index.vue'),
      meta: {
        icon: 'ant-design:database-outlined',
        title: t('routes.app.metaData.event'),
      },
    },
    {
      path: 'user',
      name: 'UserProperty',
      component: () => import('/@/views/app/meta_data/user/index.vue'),
      meta: {
        icon: 'ant-design:contacts-outlined',
        title: t('routes.app.metaData.user'),
      },
    },
    {
      path: 'dimension',
      name: 'DimensionData',
      component: () => import('/@/views/app/meta_data/dimension/index.vue'),
      meta: {
        icon: 'ant-design:contacts-outlined',
        title: t('routes.app.metaData.dimension'),
      },
    },
    {
      path: 'virtual',
      name: 'VirtualProperty',
      component: () => import('/@/views/app/meta_data/virtual/index.vue'),
      meta: {
        icon: 'ant-design:contacts-outlined',
        title: t('routes.app.metaData.virtual'),
      },
    },
    {
      path: 'virtual-event',
      name: 'VirtualEvent',
      component: () => import('/@/views/app/meta_data/virtual_event/index.vue'),
      meta: {
        icon: 'ant-design:contacts-outlined',
        title: t('routes.app.metaData.virtualEvent'),
      },
    },
  ],
};

export default metaData;
