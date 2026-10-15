<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_AlertRule as AlertRule } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListAlertRules,
  PaginationQuery,
  useDeleteAlertRule,
} from '#/api';
import { $t } from '#/locales';

import AlertRuleDrawer from './alert-rule-drawer.vue';

const { mutateAsync: deleteAlertRule } = useDeleteAlertRule();

const aggregatorText: Record<string, string> = {
  COUNT: 'page.alertRule.aggregator.count',
  ACTIVE_USERS: 'page.alertRule.aggregator.activeUsers',
  SUM_AMOUNT: 'page.alertRule.aggregator.sumAmount',
  AVG_AMOUNT: 'page.alertRule.aggregator.avgAmount',
};

const channelText: Record<string, string> = {
  INTERNAL_MESSAGE: 'page.alertRule.channel.internalMessage',
  WEBHOOK: 'page.alertRule.channel.webhook',
};

const comparisonSymbol: Record<string, string> = {
  GT: '>',
  GTE: '>=',
  LT: '<',
  LTE: '<=',
};

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.alertRule.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'eventName',
      label: $t('page.alertRule.eventName'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<AlertRule> = {
  height: 'auto',
  stripe: true,
  autoResize: true,
  toolbarConfig: {
    custom: true,
    export: true,
    refresh: true,
    zoom: true,
  },
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
    resizable: true,
  },
  tooltipConfig: {
    showAll: true,
    enterable: true,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await fetchListAlertRules(
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
      title: $t('page.alertRule.name'),
      field: 'name',
      minWidth: 150,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.alertRule.eventName'),
      field: 'eventName',
      minWidth: 140,
      align: 'left',
      formatter: ({ cellValue }) => cellValue || $t('page.alertRule.allEvents'),
    },
    {
      title: $t('page.alertRule.aggregator.label'),
      field: 'aggregator',
      minWidth: 110,
      formatter: ({ cellValue }) =>
        aggregatorText[cellValue] ? $t(aggregatorText[cellValue]) : cellValue,
    },
    {
      title: $t('page.alertRule.rule'),
      field: 'threshold',
      minWidth: 140,
      formatter: ({ row }) =>
        `${comparisonSymbol[row.comparison as string] || '>'} ${row.threshold}`,
    },
    {
      title: $t('page.alertRule.windowMinutes'),
      field: 'windowMinutes',
      minWidth: 100,
      formatter: ({ cellValue }) => `${cellValue} min`,
    },
    {
      title: $t('page.alertRule.channels'),
      field: 'channels',
      minWidth: 150,
      slots: { default: 'channels' },
    },
    {
      title: $t('page.alertRule.enabled'),
      field: 'enabled',
      minWidth: 90,
      slots: { default: 'enabled' },
    },
    {
      title: $t('page.alertRule.lastTriggeredAt'),
      field: 'lastTriggeredAt',
      formatter: 'formatDateTime',
      minWidth: 160,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      minWidth: 160,
    },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      minWidth: 120,
    },
  ],
};

const gridEvents: VxeGridListeners<AlertRule> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: AlertRuleDrawer,
  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

function openDrawer(create: boolean, row?: any) {
  drawerApi.setData({ create, row });
  drawerApi.open();
}

function handleCreate() {
  openDrawer(true);
}

function handleEdit(row: any) {
  openDrawer(false, row);
}

async function handleDelete(row: any) {
  try {
    await deleteAlertRule(row.id);
    notification.success({
      message: $t('ui.notification.delete_success'),
    });
    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('ui.notification.delete_failed'),
    });
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.alert.rule')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.alertRule.button.create') }}
        </a-button>
      </template>
      <template #channels="{ row }">
        <a-tag
          v-for="channel in row.channels || []"
          :key="channel"
          class="mr-1"
        >
          {{ channelText[channel] ? $t(channelText[channel]) : channel }}
        </a-tag>
      </template>
      <template #enabled="{ row }">
        <a-tag :color="enableBoolToColor(row.enabled)">
          {{ enableBoolToName(row.enabled) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click.stop="handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('menu.alert.rule'),
            })
          "
          @confirm="handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>

<style scoped></style>
