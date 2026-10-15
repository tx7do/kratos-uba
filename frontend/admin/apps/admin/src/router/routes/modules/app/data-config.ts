import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 数据配置：采集与建模的数据基建（应用/埋点/导入/SQL/对象/标签）。
// 单页模块坍缩为叶子；标签管理有三个子页，保留为三级分组。
// 子路由用绝对路径保持旧 URL 不变。
const AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const dataConfig: RouteRecordRaw[] = [
  {
    path: '/data-config',
    name: 'DataConfiguration',
    component: BasicLayout,
    redirect: '/app',
    meta: {
      order: 400,
      icon: 'lucide:database',
      title: $t('menu.dataConfig.moduleName'),
    },
    children: [
      {
        path: '/app',
        name: 'ApplicationManagement',
        meta: {
          order: 400,
          icon: 'lucide:square-stack',
          title: $t('menu.application.moduleName'),
          authority: AUTH,
        },
        component: () =>
          import('#/views/app/application/application/index.vue'),
      },
      {
        path: '/visual-events',
        name: 'VisualEventManagement',
        meta: {
          order: 410,
          icon: 'lucide:mouse-pointer-click',
          title: $t('menu.visualEvent.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/visual-event/index.vue'),
      },
      {
        path: '/import-jobs',
        name: 'ImportJobManagement',
        meta: {
          order: 415,
          icon: 'lucide:file-up',
          title: $t('menu.importJob.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/import-job/index.vue'),
      },
      {
        path: '/sql-console',
        name: 'SqlConsoleManagement',
        meta: {
          order: 420,
          icon: 'lucide:database-zap',
          title: $t('menu.sqlConsole.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/sql-console/index.vue'),
      },
      {
        path: '/object',
        name: 'ObjectManagement',
        meta: {
          order: 430,
          icon: 'lucide:layers',
          title: $t('menu.object.moduleName'),
          authority: AUTH,
        },
        component: () => import('#/views/app/object/object/index.vue'),
      },
      {
        path: '/tag',
        name: 'TagManagement',
        redirect: '/tag/tags',
        meta: {
          order: 440,
          icon: 'lucide:tags',
          title: $t('menu.tag.moduleName'),
          authority: AUTH,
        },
        children: [
          {
            path: 'tags',
            name: 'TagDefinitionManagement',
            meta: {
              order: 1,
              icon: 'lucide:badge-check',
              title: $t('menu.tag.tags'),
              authority: AUTH,
            },
            component: () => import('#/views/app/tag/tags/index.vue'),
          },
          {
            path: 'user-tags',
            name: 'UserTagManagement',
            meta: {
              order: 2,
              icon: 'lucide:bookmark',
              title: $t('menu.tag.userTags'),
              authority: AUTH,
            },
            component: () => import('#/views/app/tag/user-tags/index.vue'),
          },
          {
            path: 'ids',
            name: 'IDMappingManagement',
            meta: {
              order: 3,
              icon: 'lucide:link',
              title: $t('menu.tag.ids'),
              authority: AUTH,
            },
            component: () => import('#/views/app/tag/ids/index.vue'),
          },
        ],
      },
    ],
  },
];

export default dataConfig;
