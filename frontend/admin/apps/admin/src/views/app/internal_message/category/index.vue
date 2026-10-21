<script lang="ts" setup>const { mutateAsync: deleteInternalMessageCategory } = useDeleteInternalMessageCategory();
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { type internal_messageservicev1_InternalMessageCategory as InternalMessageCategory } from '#/generated/api/admin/service/v1';
import { $t } from '#/locales';
import { PaginationQuery, enableBoolToColor, enableBoolToName, fetchListInternalMessageCategories, useDeleteInternalMessageCategory } from '#/api';
import TableActions from '#/components/TableActions.vue';
import type { TableAction } from '#/components/TableActions.vue';
import TableAddButton from '#/components/TableAddButton.vue';
import TableExportButton from '#/components/TableExportButton.vue';

import InternalMessageCategoryDrawer from './internal-message-category-drawer.vue';
const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.internalMessageCategory.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: $t('page.internalMessageCategory.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<InternalMessageCategory> = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  height: 'auto',

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await fetchListInternalMessageCategories(new PaginationQuery({ paging: { page: page.currentPage, pageSize: page.pageSize }, formValues: formValues }));
      },
    },
  },

  columns: [
    {
      title: $t('page.internalMessageCategory.name'),
      field: 'name',
    },
    {
      title: $t('page.internalMessageCategory.code'),
      field: 'code',
    },
    { title: $t('ui.table.sortOrder'), field: 'sortOrder', width: 70 },
    {
      title: $t('ui.table.status'),
      field: 'isEnabled',
      slots: { default: 'isEnabled' },
      width: 95,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: $t('ui.table.remark'), field: 'remark' },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 140,
    },
  ],
};

const exportFetcher = (page: number, pageSize: number) =>
  fetchListInternalMessageCategories(new PaginationQuery({ paging: { page, pageSize } }));

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [Drawer, drawerApi] = useVbenDrawer({
  // 连接抽离的组件
  connectedComponent: InternalMessageCategoryDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      // 关闭时，重载表格数据
      gridApi.reload();
    }
  },
});

/* 打开模态窗口 */
function openDrawer(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });

  drawerApi.open();
}

/* 创建 */
function handleCreate() {

  openDrawer(true);
}

/* 编辑 */
function handleEdit(row: any) {
  openDrawer(false, row);
}

/* 删除 */
async function handleDelete(row: any) {

  try {
    await deleteInternalMessageCategory({ id: row.id });

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

/* 操作列描述符：编辑行内、删除 Popconfirm 确认（≤3 个全行内，定式由 TableActions 统一执行） */
function buildRowActions(): TableAction[] {
  return [
    {
      key: 'edit',
      label: $t('common.edit'),
      icon: h(LucideFilePenLine),
      onClick: handleEdit,
    },
    {
      key: 'delete',
      label: $t('common.delete'),
      icon: h(LucideTrash2),
      danger: true,
      confirm: {
        title: $t('ui.text.do_you_want_delete', {
          moduleName: $t('page.internalMessageCategory.moduleName'),
        }),
      },
      onClick: handleDelete,
    },
  ];
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.internalMessage.internalMessageCategory')">
      <template #toolbar-tools>
        <TableExportButton :fetcher="exportFetcher" :columns="gridOptions.columns" filename="im-categories" />
        <TableAddButton @click="handleCreate" />
      </template>
      <template #isEnabled="{ row }">
        <a-tag :color="enableBoolToColor(row.isEnabled)">
          {{ enableBoolToName(row.isEnabled) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <TableActions :row="row" :actions="buildRowActions()" />
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
