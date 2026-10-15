<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_ReportSubscription as ReportSubscription } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListReportSubscriptions,
  PaginationQuery,
  useDeleteReportSubscription,
} from '#/api';
import ReportSubscriptionDrawer from './report-subscription-drawer.vue';

const { mutateAsync: deleteSubscription } = useDeleteReportSubscription();

const modelText: Record<string, string> = {
  'event-trend': 'page.report.model.eventTrend',
  'active-users': 'page.report.model.activeUsers',
  'group-by': 'page.report.model.groupBy',
  'funnel': 'page.report.model.funnel',
  'retention': 'page.report.model.retention',
};

const channelText: Record<string, string> = {
  REPORT_CHANNEL_INTERNAL_MESSAGE: 'page.report.channel.internalMessage',
  REPORT_CHANNEL_WEBHOOK: 'page.report.channel.webhook',
};

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.report.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<ReportSubscription> = {
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
        return await fetchListReportSubscriptions(
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
      title: $t('page.report.name'),
      field: 'name',
      minWidth: 160,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.report.model.label'),
      field: 'model',
      minWidth: 130,
      formatter: ({ cellValue }) =>
        modelText[cellValue] ? $t(modelText[cellValue]) : cellValue,
    },
    {
      title: $t('page.report.intervalMinutes'),
      field: 'intervalMinutes',
      minWidth: 110,
      formatter: ({ cellValue }) => `${cellValue} min`,
    },
    {
      title: $t('page.report.channels'),
      field: 'channels',
      minWidth: 160,
      slots: { default: 'channels' },
    },
    {
      title: $t('page.report.enabled'),
      field: 'enabled',
      minWidth: 90,
      slots: { default: 'enabled' },
    },
    {
      title: $t('page.report.lastRunAt'),
      field: 'lastRunAt',
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

const gridEvents: VxeGridListeners<ReportSubscription> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: ReportSubscriptionDrawer,
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
    await deleteSubscription(row.id);
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
    <Grid :table-title="$t('menu.report.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.report.button.create') }}
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
              moduleName: $t('menu.report.moduleName'),
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
