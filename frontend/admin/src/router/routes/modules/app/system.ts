import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const system: AppRouteModule = {
  path: '/system',
  name: 'System',
  component: LAYOUT,
  meta: {
    orderNo: 2000,
    icon: 'ant-design:setting-outline',
    title: t('routes.app.system.moduleName'),
  },
  children: [
    {
      path: 'users',
      name: 'UserManagement',
      meta: {
        title: t('routes.app.system.userManage'),
        ignoreKeepAlive: false,
        icon: 'ant-design:user-outlined',
      },
      component: () => import('/@/views/app/system/users/index.vue'),
    },

    {
      path: 'apps',
      name: 'AppManagement',
      meta: {
        title: t('routes.app.system.appManage'),
        ignoreKeepAlive: false,
        icon: 'ant-design:appstore-outlined',
      },
      component: () => import('/@/views/app/system/apps/index.vue'),
    },
  ],
};

export default system;
