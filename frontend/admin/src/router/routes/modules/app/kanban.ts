import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const kanban: AppRouteModule = {
  path: '/kanban',
  name: 'Kanban',
  component: LAYOUT,
  redirect: '/kanban/index',
  meta: {
    orderNo: 1,
    hideChildrenInMenu: true,
    icon: 'ant-design:dashboard-outlined',
    title: t('routes.app.kanban.moduleName'),
  },
  children: [
    {
      path: 'index',
      name: 'KanbanPage',
      component: () => import('/@/views/app/kanban/index.vue'),
      meta: {
        icon: 'ant-design:dashboard-outlined',
        title: t('routes.app.kanban.moduleName'),
        hideMenu: true,
      },
    },
  ],
};

export default kanban;
