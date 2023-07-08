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
      name: 'MetaData',
      component: () => import('/@/views/app/meta_data/meta/index.vue'),
      meta: {
        icon: 'ant-design:control-outlined',
        title: t('routes.app.metaData.meta'),
      },
    },
    {
      path: 'event',
      name: 'EventData',
      component: () => import('/@/views/app/meta_data/event/index.vue'),
      meta: {
        icon: 'ant-design:database-outlined',
        title: t('routes.app.metaData.event'),
      },
    },
    {
      path: 'user',
      name: 'UserData',
      component: () => import('/@/views/app/meta_data/user/index.vue'),
      meta: {
        icon: 'ant-design:contacts-outlined',
        title: t('routes.app.metaData.user'),
      },
    },
  ],
};

export default metaData;
