<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_FeatureFlag as FeatureFlag } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import {
  LucideFilePenLine,
  LucideFlaskConical,
  LucideTrash2,
} from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListFeatureFlags,
  PaginationQuery,
  useDeleteFeatureFlag,
} from '#/api';
import FeatureFlagDrawer from './feature-flag-drawer.vue';
import FlagResultDrawer from './flag-result-drawer.vue';

const { mutateAsync: deleteFlag } = useDeleteFeatureFlag();

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'key',
      label: $t('page.featureFlag.key'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<FeatureFlag> = {
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
        return await fetchListFeatureFlags(
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
      title: $t('page.featureFlag.key'),
      field: 'key',
      minWidth: 180,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.featureFlag.name'),
      field: 'name',
      minWidth: 150,
      align: 'left',
    },
    {
      title: $t('page.featureFlag.variants'),
      field: 'variants',
      minWidth: 240,
      slots: { default: 'variants' },
    },
    {
      title: $t('page.featureFlag.rolloutPercentage'),
      field: 'rolloutPercentage',
      minWidth: 110,
      formatter: ({ cellValue }) => `${cellValue ?? 100}%`,
    },
    {
      title: $t('page.featureFlag.goalEventName'),
      field: 'goalEventName',
      minWidth: 140,
    },
    {
      title: $t('page.featureFlag.enabled'),
      field: 'enabled',
      minWidth: 90,
      slots: { default: 'enabled' },
    },
    {
      title: $t('page.featureFlag.startedAt'),
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
      minWidth: 160,
    },
  ],
};

const gridEvents: VxeGridListeners<FeatureFlag> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: FeatureFlagDrawer,
  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

const [ResultDrawer, resultDrawerApi] = useVbenDrawer({
  connectedComponent: FlagResultDrawer,
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

function handleViewResult(row: any) {
  resultDrawerApi.setData({ row });
  resultDrawerApi.open();
}

async function handleDelete(row: any) {
  try {
    await deleteFlag(row.id);
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
    <Grid :table-title="$t('menu.featureFlag.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.featureFlag.button.create') }}
        </a-button>
      </template>
      <template #variants="{ row }">
        <a-tag
          v-for="variant in row.variants || []"
          :key="variant.key"
          class="mr-1"
        >
          {{ variant.key }} · {{ Math.round((variant.weight / 10) * 10) / 10 }}%
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
          :icon="h(LucideFlaskConical)"
          :title="$t('page.featureFlag.viewResult')"
          @click.stop="handleViewResult(row)"
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
              moduleName: $t('menu.featureFlag.moduleName'),
            })
          "
          @confirm="handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
    <ResultDrawer />
  </Page>
</template>

<style scoped></style>
