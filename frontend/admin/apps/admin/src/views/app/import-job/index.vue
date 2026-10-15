<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_ImportJob as ImportJob } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  fetchListImportJobs,
  PaginationQuery,
  useCancelImportJob,
  useDeleteImportJob,
} from '#/api';
import ImportDrawer from './import-drawer.vue';

const { mutateAsync: deleteJob } = useDeleteImportJob();
const { mutateAsync: cancelJob } = useCancelImportJob();

const statusColor: Record<string, string> = {
  PENDING: 'processing',
  RUNNING: 'processing',
  COMPLETED: 'success',
  PARTIAL: 'warning',
  FAILED: 'error',
  CANCELLED: 'default',
};

const statusText: Record<string, string> = {
  PENDING: 'page.importJob.status.pending',
  RUNNING: 'page.importJob.status.running',
  COMPLETED: 'page.importJob.status.completed',
  PARTIAL: 'page.importJob.status.partial',
  FAILED: 'page.importJob.status.failed',
  CANCELLED: 'page.importJob.status.cancelled',
};

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.importJob.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<ImportJob> = {
  height: 'auto',
  stripe: true,
  autoResize: true,
  toolbarConfig: {
    custom: true,
    refresh: true,
    zoom: true,
  },
  pagerConfig: {},
  rowConfig: {
    isHover: true,
    resizable: true,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await fetchListImportJobs(
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
      title: $t('page.importJob.name'),
      field: 'name',
      minWidth: 180,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.importJob.appId'),
      field: 'appId',
      minWidth: 140,
    },
    {
      title: $t('ui.table.status'),
      field: 'status',
      minWidth: 120,
      slots: { default: 'status' },
    },
    {
      title: $t('page.importJob.progress'),
      field: 'totalCount',
      minWidth: 200,
      formatter: ({ row }) =>
        `${row.importedCount ?? 0}/${row.totalCount ?? 0}${
          (row.failedCount ?? 0) > 0 ? ` (${row.failedCount} ${$t('page.importJob.failedUnit')})` : ''
        }`,
    },
    {
      title: $t('page.importJob.duration'),
      field: 'durationMs',
      minWidth: 110,
      formatter: ({ cellValue }) =>
        cellValue ? `${(cellValue / 1000).toFixed(1)}s` : '-',
    },
    {
      title: $t('page.importJob.startedAt'),
      field: 'startedAt',
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
      minWidth: 140,
    },
  ],
};

const gridEvents: VxeGridListeners<ImportJob> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: ImportDrawer,
  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

async function handleCancel(row: any) {
  try {
    await cancelJob(row.id);
    notification.success({ message: $t('page.importJob.cancelSuccess') });
    await gridApi.reload();
  } catch {
    notification.error({ message: $t('page.importJob.cancelFailed') });
  }
}

async function handleDelete(row: any) {
  try {
    await deleteJob(row.id);
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
    <Grid :table-title="$t('menu.importJob.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="drawerApi.setData({}); drawerApi.open()">
          {{ $t('page.importJob.button.create') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-tag :color="statusColor[row.status as string] || 'default'">
          {{ (row.status && statusText[row.status as string]) ? $t(statusText[row.status as string]!) : (row.status ?? '') }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          v-if="['PENDING', 'RUNNING'].includes(row.status)"
          size="small"
          type="link"
          danger
          @click.stop="handleCancel(row)"
        >
          {{ $t('page.importJob.cancel') }}
        </a-button>
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="$t('ui.text.do_you_want_delete', { moduleName: $t('menu.importJob.moduleName') })"
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
