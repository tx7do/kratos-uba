<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_Dashboard as Dashboard } from '#/generated/api/admin/service/v1';

import { h } from 'vue';
import { useRouter } from 'vue-router';

import { Page, useVbenDrawer } from '@vben/common-ui';
import {
  LucideExternalLink,
  LucideFilePenLine,
  LucideTrash2,
} from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListDashboards,
  PaginationQuery,
  useDeleteDashboard,
} from '#/api';
import { $t } from '#/locales';

import DashboardDrawer from './dashboard-drawer.vue';

const router = useRouter();

const { mutateAsync: deleteDashboard } = useDeleteDashboard();

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.dashboard.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<Dashboard> = {
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
  tooltipConfig: {
    showAll: true,
    enterable: true,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await fetchListDashboards(
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
      title: $t('page.dashboard.name'),
      field: 'name',
      minWidth: 180,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.dashboard.description'),
      field: 'description',
      minWidth: 240,
      align: 'left',
    },
    {
      title: $t('page.dashboard.widgetCount'),
      field: 'widgets',
      minWidth: 100,
      formatter: ({ cellValue }) => `${(cellValue || []).length}`,
    },
    {
      title: $t('page.dashboard.enabled'),
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
      minWidth: 160,
    },
  ],
};

const gridEvents: VxeGridListeners<Dashboard> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: DashboardDrawer,
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

function handleView(row: any) {
  router.push({ path: `/dashboard/custom/view/${row.id}` });
}

async function handleDelete(row: any) {
  try {
    await deleteDashboard(row.id);
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
    <Grid :table-title="$t('menu.dashboard.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.dashboard.button.create') }}
        </a-button>
      </template>
      <template #enabled="{ row }">
        <a-tag :color="enableBoolToColor(row.enabled)">
          {{ enableBoolToName(row.enabled) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideExternalLink)"
          :title="$t('page.dashboard.view')"
          @click.stop="handleView(row)"
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
              moduleName: $t('menu.dashboard.moduleName'),
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
