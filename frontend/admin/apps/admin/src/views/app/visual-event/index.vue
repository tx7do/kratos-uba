<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_VisualEvent as VisualEvent } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListVisualEvents,
  PaginationQuery,
  useDeleteVisualEvent,
} from '#/api';
import VisualEventDrawer from './visual-event-drawer.vue';

const { mutateAsync: deleteEvent } = useDeleteVisualEvent();

const matchTypeText: Record<string, string> = {
  XPATH_EXACT: 'page.visualEvent.match.exact',
  XPATH_PREFIX: 'page.visualEvent.match.prefix',
};

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'eventName',
      label: $t('page.visualEvent.eventName'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<VisualEvent> = {
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
        return await fetchListVisualEvents(
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
      title: $t('page.visualEvent.eventName'),
      field: 'eventName',
      minWidth: 160,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.visualEvent.displayName'),
      field: 'displayName',
      minWidth: 140,
      align: 'left',
    },
    {
      title: $t('page.visualEvent.appId'),
      field: 'appId',
      minWidth: 130,
    },
    {
      title: $t('page.visualEvent.pageUrl'),
      field: 'pageUrl',
      minWidth: 220,
      align: 'left',
      formatter: ({ cellValue }) => cellValue || $t('page.visualEvent.allPages'),
    },
    {
      title: $t('page.visualEvent.elementXpath'),
      field: 'elementXpath',
      minWidth: 300,
      align: 'left',
      slots: { default: 'xpath' },
    },
    {
      title: $t('page.visualEvent.matchType'),
      field: 'matchType',
      minWidth: 110,
      formatter: ({ cellValue }) =>
        matchTypeText[cellValue] ? $t(matchTypeText[cellValue]) : cellValue,
    },
    {
      title: $t('page.visualEvent.enabled'),
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
      minWidth: 120,
    },
  ],
};

const gridEvents: VxeGridListeners<VisualEvent> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: VisualEventDrawer,
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
    await deleteEvent(row.id);
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
    <Grid :table-title="$t('menu.visualEvent.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.visualEvent.button.create') }}
        </a-button>
      </template>
      <template #xpath="{ row }">
        <span class="font-mono text-xs">{{ row.elementXpath }}</span>
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
              moduleName: $t('menu.visualEvent.moduleName'),
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
