import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 系统管理 band：组织人员/权限/租户/站内消息/日志审计/系统设置/开发者七组收拢。
// 各分组沿用原模块路径（/opm、/permission…），子页路径全部保持旧 URL 不变；
// 「系统设置」原占用 /system 根，现收为 band 下的 settings 组，子页用绝对路径。
const AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const system: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: BasicLayout,
    redirect: '/opm/org-units',
    meta: {
      order: 900,
      icon: 'lucide:settings',
      title: $t('menu.system.moduleName'),
    },
    children: [
      {
        path: '/opm',
        name: 'OrganizationalPersonnelManagement',
        redirect: '/opm/users',
        meta: {
          order: 900,
          icon: 'lucide:users',
          title: $t('menu.opm.moduleName'),
          keepAlive: true,
          authority: AUTH,
        },
        children: [
          {
            path: 'org-units',
            name: 'OrgUnitManagement',
            meta: {
              order: 1,
              icon: 'lucide:layers',
              title: $t('menu.opm.orgUnit'),
              authority: AUTH,
            },
            component: () => import('#/views/app/opm/org_unit/index.vue'),
          },
          {
            path: 'positions',
            name: 'PositionManagement',
            meta: {
              order: 2,
              icon: 'lucide:briefcase',
              title: $t('menu.opm.position'),
              authority: AUTH,
            },
            component: () => import('#/views/app/opm/position/index.vue'),
          },
          {
            path: 'users',
            name: 'UserManagement',
            meta: {
              order: 3,
              icon: 'lucide:user',
              title: $t('menu.opm.user'),
              authority: AUTH,
            },
            component: () => import('#/views/app/opm/user/index.vue'),
          },
          {
            path: 'users/detail/:id',
            name: 'UserDetail',
            meta: {
              hideInMenu: true,
              title: $t('menu.opm.userDetail'),
              authority: AUTH,
            },
            component: () => import('#/views/app/opm/user/detail/index.vue'),
          },
        ],
      },
      {
        path: '/permission',
        name: 'PermissionManagement',
        redirect: '/permission/codes',
        meta: {
          order: 910,
          icon: 'lucide:shield-check',
          title: $t('menu.permission.moduleName'),
          keepAlive: true,
          authority: AUTH,
        },
        children: [
          {
            path: 'codes',
            name: 'PermissionPointManagement',
            meta: {
              order: 1,
              icon: 'lucide:shield-ellipsis',
              title: $t('menu.permission.permission'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/permission/permission/index.vue'),
          },
          {
            path: 'roles',
            name: 'RoleManagement',
            meta: {
              order: 2,
              icon: 'lucide:shield-user',
              title: $t('menu.permission.role'),
              authority: AUTH,
            },
            component: () => import('#/views/app/permission/role/index.vue'),
          },
        ],
      },
      {
        path: '/tenant',
        name: 'TenantManagement',
        meta: {
          order: 920,
          icon: 'lucide:building-2',
          title: $t('menu.tenant.moduleName'),
          authority: ['sys:platform_admin'],
        },
        component: () => import('#/views/app/tenant/tenant/index.vue'),
      },
      {
        path: '/internal-message',
        name: 'InternalMessageManagement',
        redirect: '/internal-message/messages',
        meta: {
          order: 930,
          icon: 'lucide:mail',
          title: $t('menu.internalMessage.moduleName'),
          keepAlive: true,
          authority: AUTH,
        },
        children: [
          {
            path: 'messages',
            name: 'InternalMessageList',
            meta: {
              order: 1,
              icon: 'lucide:message-circle-more',
              title: $t('menu.internalMessage.internalMessage'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/internal_message/message/index.vue'),
          },
          {
            path: 'categories',
            name: 'InternalMessageCategoryManagement',
            meta: {
              order: 2,
              icon: 'lucide:calendar-check',
              title: $t('menu.internalMessage.internalMessageCategory'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/internal_message/category/index.vue'),
          },
        ],
      },
      {
        path: '/log',
        name: 'LogAuditManagement',
        redirect: '/log/login-audit-logs',
        meta: {
          order: 940,
          icon: 'lucide:logs',
          title: $t('menu.log.moduleName'),
          keepAlive: true,
          authority: AUTH,
        },
        children: [
          {
            path: 'login-audit-logs',
            name: 'LoginAuditLog',
            meta: {
              icon: 'lucide:user-lock',
              title: $t('menu.log.loginAuditLog'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/log/login_audit_log/index.vue'),
          },
          {
            path: 'api-audit-logs',
            name: 'ApiAuditLog',
            meta: {
              icon: 'lucide:file-clock',
              title: $t('menu.log.apiAuditLog'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/log/api_audit_log/index.vue'),
          },
          {
            path: 'operation-audit-logs',
            name: 'OperationAuditLog',
            meta: {
              icon: 'lucide:shield-ellipsis',
              title: $t('menu.log.operationAuditLog'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/log/operation_audit_log/index.vue'),
          },
          {
            path: 'data-access-audit-logs',
            name: 'DataAccessAuditLog',
            meta: {
              icon: 'lucide:shield-check',
              title: $t('menu.log.dataAccessAuditLog'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/log/data_access_audit_log/index.vue'),
          },
          {
            path: 'permission-audit-logs',
            name: 'PermissionAuditLog',
            meta: {
              icon: 'lucide:shield-alert',
              title: $t('menu.log.permissionAuditLog'),
              authority: AUTH,
            },
            component: () =>
              import('#/views/app/log/permission_audit_log/index.vue'),
          },
        ],
      },
      {
        path: 'settings',
        name: 'SystemSettings',
        redirect: '/system/menus',
        meta: {
          order: 950,
          icon: 'lucide:sliders-horizontal',
          title: $t('menu.system.settings'),
          keepAlive: true,
          authority: AUTH,
        },
        children: [
          {
            path: '/system/menus',
            name: 'MenuManagement',
            meta: {
              order: 1,
              icon: 'lucide:square-menu',
              title: $t('menu.system.menu'),
              authority: ['sys:platform_admin'],
            },
            component: () => import('#/views/app/system/menu/index.vue'),
          },
          {
            path: '/system/files',
            name: 'FileManagement',
            meta: {
              order: 2,
              icon: 'lucide:file-search',
              title: $t('menu.system.file'),
              authority: AUTH,
            },
            component: () => import('#/views/app/system/file/index.vue'),
          },
          {
            path: '/system/login-policies',
            name: 'LoginPolicyManagement',
            meta: {
              order: 3,
              icon: 'lucide:shield-x',
              title: $t('menu.system.loginPolicy'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/system/login_policy/index.vue'),
          },
        ],
      },
      {
        path: '/developer',
        name: 'Developer',
        redirect: '/developer/apis',
        meta: {
          order: 960,
          icon: 'lucide:terminal',
          title: $t('menu.developer.moduleName'),
          keepAlive: true,
          authority: ['sys:platform_admin'],
        },
        children: [
          {
            path: 'apis',
            name: 'APIManagement',
            meta: {
              order: 1,
              icon: 'lucide:route',
              title: $t('menu.developer.api'),
              authority: ['sys:platform_admin'],
            },
            component: () => import('#/views/app/system/api/index.vue'),
          },
          {
            path: 'dict',
            name: 'DictManagement',
            meta: {
              order: 2,
              icon: 'lucide:library-big',
              title: $t('menu.developer.dict'),
              authority: ['sys:platform_admin'],
            },
            component: () => import('#/views/app/system/dict/index.vue'),
          },
          {
            path: 'tasks',
            name: 'TaskManagement',
            meta: {
              order: 3,
              icon: 'lucide:list-todo',
              title: $t('menu.developer.task'),
              authority: AUTH,
            },
            component: () => import('#/views/app/system/task/index.vue'),
          },
          {
            path: 'languages',
            name: 'LanguageManagement',
            meta: {
              order: 4,
              icon: 'lucide:globe',
              title: $t('menu.developer.language'),
              authority: ['sys:platform_admin'],
            },
            component: () => import('#/views/app/system/language/index.vue'),
          },
          {
            path: 'event-schemas',
            name: 'EventSchemaManagement',
            meta: {
              order: 5,
              icon: 'lucide:braces',
              title: $t('menu.developer.eventSchema'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/system/event-schema/index.vue'),
          },
        ],
      },
    ],
  },
];

export default system;
