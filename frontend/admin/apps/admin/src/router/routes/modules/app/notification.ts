import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const notification: RouteRecordRaw[] = [
  {
    path: '/notification',
    name: 'NotificationManagement',
    component: BasicLayout,
    redirect: '/notification/deliveries',
    meta: {
      order: 1011,
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
        component: () => import('#/views/app/notification/channel/index.vue'),
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
        component: () => import('#/views/app/notification/rule/index.vue'),
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
        component: () => import('#/views/app/notification/template/index.vue'),
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
        component: () => import('#/views/app/notification/delivery/index.vue'),
      },
    ],
  },
];

export default notification;
