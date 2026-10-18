import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

// AI 助手：模型提供商 / 对话 / 知识库（RAG）/ 智能问数 / 用量。
// 二次管理页与工具页同级；对话面向日常使用，均限管理角色。
const AUTH = ['sys:platform_admin', 'sys:tenant_manager'];

const ai: RouteRecordRaw[] = [
  {
    path: '/ai',
    name: 'AiAssistant',
    component: BasicLayout,
    redirect: '/ai/chat',
    meta: {
      order: 450,
      icon: 'lucide:sparkles',
      title: $t('menu.ai.moduleName'),
    },
    children: [
      {
        path: 'chat',
        name: 'AiChatManagement',
        meta: {
          order: 450,
          icon: 'lucide:bot-message-square',
          title: $t('menu.ai.chat'),
          authority: AUTH,
        },
        component: () => import('#/views/app/ai/chat/index.vue'),
      },
      {
        path: 'knowledge',
        name: 'AiKnowledgeManagement',
        meta: {
          order: 455,
          icon: 'lucide:book-open-text',
          title: $t('menu.ai.knowledge'),
          authority: AUTH,
        },
        component: () => import('#/views/app/ai/knowledge/index.vue'),
      },
      {
        path: 'query',
        name: 'AiQueryManagement',
        meta: {
          order: 460,
          icon: 'lucide:message-circle-question',
          title: $t('menu.ai.query'),
          authority: AUTH,
        },
        component: () => import('#/views/app/ai/query/index.vue'),
      },
      {
        path: 'providers',
        name: 'AiProviderManagement',
        meta: {
          order: 465,
          icon: 'lucide:cpu',
          title: $t('menu.ai.provider'),
          authority: AUTH,
        },
        component: () => import('#/views/app/ai/provider/index.vue'),
      },
      {
        path: 'usage',
        name: 'AiUsageManagement',
        meta: {
          order: 470,
          icon: 'lucide:gauge',
          title: $t('menu.ai.usage'),
          authority: AUTH,
        },
        component: () => import('#/views/app/ai/usage/index.vue'),
      },
    ],
  },
];

export default ai;
