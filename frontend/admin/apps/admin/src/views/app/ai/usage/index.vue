<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, type VbenFormProps } from '@vben/common-ui';

import { Card, Statistic } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  fetchListAiUsageLogs,
  PaginationQuery,
  useAiUsageSummary,
} from '#/api';
import { $t } from '#/locales';

defineOptions({ name: 'AiUsage' });

const summary = useAiUsageSummary();

const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'modelName__contains',
      label: $t('page.ai.usage.modelName'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps = {
  toolbarConfig: {
    custom: true,
    export: true,
    refresh: true,
    zoom: true,
  },
  height: 'auto',
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  stripe: true,

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await fetchListAiUsageLogs(
          new PaginationQuery({
            paging: { page: page.currentPage, pageSize: page.pageSize },
            formValues,
          }),
        );
      },
    },
  },

  columns: [
    {
      title: $t('page.ai.usage.modelName'),
      field: 'modelName',
      fixed: 'left',
      minWidth: 180,
    },
    {
      title: $t('page.ai.usage.promptTokens'),
      field: 'promptTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.completionTokens'),
      field: 'completionTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.totalTokens'),
      field: 'totalTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.durationMs'),
      field: 'durationMs',
      minWidth: 100,
    },
    {
      title: $t('page.ai.usage.conversationId'),
      field: 'conversationId',
      minWidth: 110,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      minWidth: 140,
    },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <div class="mb-4 grid grid-cols-2 gap-4 lg:grid-cols-4">
      <Card size="small">
        <Statistic
          :title="$t('page.ai.usage.monthTokens')"
          :value="summary.data.value?.monthTokens ?? 0"
        />
      </Card>
      <Card size="small">
        <Statistic
          :title="$t('page.ai.usage.monthCalls')"
          :value="summary.data.value?.monthCalls ?? 0"
        />
      </Card>
    </div>
    <Grid :table-title="$t('menu.ai.usage')" />
  </Page>
</template>
