<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  enableBoolToColor,
  enableBoolToName,
  fetchListAiProviders,
  PaginationQuery,
  useDeleteAiProvider,
} from '#/api';
import { $t } from '#/locales';

import AiProviderDrawer from './provider-drawer.vue';

defineOptions({ name: 'AiProviderManagement' });

const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name__contains',
      label: $t('page.ai.provider.name'),
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
        return await fetchListAiProviders(
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
      title: $t('page.ai.provider.name'),
      field: 'name',
      fixed: 'left',
      minWidth: 140,
    },
    {
      title: $t('page.ai.provider.modelType'),
      field: 'modelType',
      slots: { default: 'modelType' },
      minWidth: 90,
    },
    {
      title: $t('page.ai.provider.modelName'),
      field: 'modelName',
      minWidth: 160,
    },
    {
      title: $t('page.ai.provider.baseUrl'),
      field: 'baseUrl',
      minWidth: 220,
    },
    {
      title: $t('page.ai.provider.apiKeyHint'),
      field: 'apiKeyHint',
      minWidth: 140,
    },
    {
      title: $t('page.ai.provider.isDefault'),
      field: 'isDefault',
      slots: { default: 'isDefault' },
      minWidth: 90,
    },
    {
      title: $t('page.ai.provider.isEnabled'),
      field: 'isEnabled',
      slots: { default: 'isEnabled' },
      minWidth: 90,
    },
    {
      title: $t('ui.table.remark'),
      field: 'remark',
      minWidth: 160,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      minWidth: 140,
    },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      minWidth: 90,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: AiProviderDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

function openDrawer(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });
  drawerApi.open();
}

function handleCreate() {
  openDrawer(true);
}

function handleEdit(row: any) {
  openDrawer(false, row);
}

const { mutateAsync: deleteProvider } = useDeleteAiProvider();

async function handleDelete(row: any) {
  try {
    await deleteProvider({ id: row.id });
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

function modelTypeName(v: number | undefined) {
  // 与 proto ai.service.v1.AiProvider.ModelType 对齐（1=LOCAL，2=CLOUD）
  if (v === 1) return $t('page.ai.provider.modelTypeLocal');
  if (v === 2) return $t('page.ai.provider.modelTypeCloud');
  return '-';
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.ai.provider')">
      <template #toolbar-tools>
        <a-button type="primary" class="mr-2" @click="handleCreate">
          {{ $t('page.ai.provider.button.create') }}
        </a-button>
      </template>
      <template #modelType="{ row }">
        <Tag :color="row.modelType === 2 ? 'blue' : 'green'">
          {{ modelTypeName(row.modelType) }}
        </Tag>
      </template>
      <template #isDefault="{ row }">
        <Tag v-if="row.isDefault" color="gold">
          {{ $t('page.ai.provider.isDefaultYes') }}
        </Tag>
        <span v-else>-</span>
      </template>
      <template #isEnabled="{ row }">
        <Tag :color="enableBoolToColor(row.isEnabled)">
          {{ enableBoolToName(row.isEnabled) }}
        </Tag>
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
              moduleName: row.name || $t('page.ai.provider.moduleName'),
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
