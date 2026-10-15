import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// 风控与通知：风险事件处置 + 通知触达基建（通知管理仅平台管理员可见）。
// 子路由用绝对/相对路径保持旧 URL 不变。
const riskNotification: RouteRecordRaw[] = [
  {
    path: '/risk-notification',
    name: 'RiskNotification',
    component: BasicLayout,
    redirect: '/risk/events',
    meta: {
      order: 500,
      icon: 'lucide:shield',
      title: $t('menu.riskNotification.moduleName'),
    },
    children: [
      {
        path: '/risk',
        name: 'RiskManagement',
        redirect: '/risk/events',
        meta: {
          order: 500,
          icon: 'lucide:shield-alert',
          title: $t('menu.risk.moduleName'),
          authority: ['sys:platform_admin', 'sys:tenant_manager'],
        },
        children: [
          {
            path: 'events',
            name: 'RiskEventManagement',
            meta: {
              order: 1,
              icon: 'lucide:triangle-alert',
              title: $t('menu.risk.event'),
              authority: ['sys:platform_admin', 'sys:tenant_manager'],
            },
            component: () => import('#/views/app/risk/event/index.vue'),
          },
          {
            path: 'rules',
            name: 'RiskRuleManagement',
            meta: {
              order: 2,
              icon: 'lucide:file-check',
              title: $t('menu.risk.rule'),
              authority: ['sys:platform_admin', 'sys:tenant_manager'],
            },
            component: () => import('#/views/app/risk/rule/index.vue'),
          },
          {
            path: 'webhooks',
            name: 'WebhookManagement',
            meta: {
              order: 3,
              icon: 'lucide:webhook',
              title: $t('menu.risk.webhook'),
              authority: ['sys:platform_admin', 'sys:tenant_manager'],
            },
            component: () => import('#/views/app/risk/webhook/index.vue'),
          },
        ],
      },
      {
        path: '/notification',
        name: 'NotificationManagement',
        redirect: '/notification/deliveries',
        meta: {
          order: 510,
          icon: 'lucide:mail-warning',
          title: $t('menu.notification.moduleName'),
          authority: ['sys:platform_admin'],
        },
        children: [
          {
            path: 'channels',
            name: 'NotificationChannelManagement',
            meta: {
              order: 1,
              icon: 'lucide:radio-tower',
              title: $t('menu.notification.channels'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/notification/channel/index.vue'),
          },
          {
            path: 'rules',
            name: 'NotificationRuleManagement',
            meta: {
              order: 2,
              icon: 'lucide:route',
              title: $t('menu.notification.rules'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/notification/rule/index.vue'),
          },
          {
            path: 'templates',
            name: 'NotificationTemplateManagement',
            meta: {
              order: 3,
              icon: 'lucide:layout-template',
              title: $t('menu.notification.templates'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/notification/template/index.vue'),
          },
          {
            path: 'deliveries',
            name: 'NotificationDeliveryManagement',
            meta: {
              order: 4,
              icon: 'lucide:mail-check',
              title: $t('menu.notification.deliveries'),
              authority: ['sys:platform_admin'],
            },
            component: () =>
              import('#/views/app/notification/delivery/index.vue'),
          },
        ],
      },
    ],
  },
];

export default riskNotification;
