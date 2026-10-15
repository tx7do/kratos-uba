<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_Segment as Segment } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import {
  LucideFilePenLine,
  LucideRefreshCw,
  LucideTrash2,
  LucideUsers,
} from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListSegments,
  PaginationQuery,
  useDeleteSegment,
  useRefreshSegment,
} from '#/api';
import { $t } from '#/locales';

import SegmentDrawer from './segment-drawer.vue';
import SegmentUsersDrawer from './segment-users-drawer.vue';

const { mutateAsync: deleteSegment } = useDeleteSegment();
const { mutateAsync: refreshSegment } = useRefreshSegment();

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.segment.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<Segment> = {
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
        return await fetchListSegments(
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
      title: $t('page.segment.name'),
      field: 'name',
      minWidth: 160,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.segment.description'),
      field: 'description',
      minWidth: 200,
      align: 'left',
    },
    {
      title: $t('page.segment.userCount'),
      field: 'userCount',
      minWidth: 110,
    },
    {
      title: $t('page.segment.refreshIntervalMinutes'),
      field: 'refreshIntervalMinutes',
      minWidth: 120,
      formatter: ({ cellValue }) =>
        cellValue > 0 ? `${cellValue} min` : $t('page.segment.manualRefresh'),
    },
    {
      title: $t('page.segment.lastRefreshedAt'),
      field: 'lastRefreshedAt',
      formatter: 'formatDateTime',
      minWidth: 160,
    },
    {
      title: $t('page.segment.enabled'),
      field: 'enabled',
      minWidth: 90,
      slots: { default: 'enabled' },
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
      minWidth: 180,
    },
  ],
};

const gridEvents: VxeGridListeners<Segment> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: SegmentDrawer,
  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

const [UsersDrawer, usersDrawerApi] = useVbenDrawer({
  connectedComponent: SegmentUsersDrawer,
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

function handleViewUsers(row: any) {
  usersDrawerApi.setData({ row });
  usersDrawerApi.open();
}

async function handleRefresh(row: any) {
  try {
    await refreshSegment(row.id);
    notification.success({
      message: $t('page.segment.refreshSuccess'),
    });
    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('page.segment.refreshFailed'),
    });
  }
}

async function handleDelete(row: any) {
  try {
    await deleteSegment(row.id);
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
    <Grid :table-title="$t('menu.segment.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.segment.button.create') }}
        </a-button>
      </template>
      <template #userCount="{ row }">
        <a-badge
          :count="row.userCount"
          :number-style="{ backgroundColor: '#1677ff' }"
          :max-count="9999999"
        />
      </template>
      <template #enabled="{ row }">
        <a-tag :color="enableBoolToColor(row.enabled)">
          {{ enableBoolToName(row.enabled) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideUsers)"
          :title="$t('page.segment.viewUsers')"
          @click.stop="handleViewUsers(row)"
        />
        <a-button
          type="link"
          :icon="h(LucideRefreshCw)"
          :title="$t('page.segment.refresh')"
          @click.stop="handleRefresh(row)"
        />
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
              moduleName: $t('menu.segment.moduleName'),
            })
          "
          @confirm="handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
    <UsersDrawer />
  </Page>
</template>

<style scoped></style>
