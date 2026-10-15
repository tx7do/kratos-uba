<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_Campaign as Campaign } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import {
  LucideFilePenLine,
  LucideRocket,
  LucideSend,
  LucideTrash2,
} from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  fetchListCampaigns,
  PaginationQuery,
  useCancelCampaign,
  useDeleteCampaign,
  useSendCampaign,
} from '#/api';
import CampaignDrawer from './campaign-drawer.vue';

const { mutateAsync: deleteCampaign } = useDeleteCampaign();
const { mutateAsync: sendCampaign } = useSendCampaign();
const { mutateAsync: cancelCampaign } = useCancelCampaign();

const statusColor: Record<string, string> = {
  DRAFT: 'default',
  SCHEDULED: 'processing',
  SENDING: 'processing',
  SENT: 'success',
  PARTIAL: 'warning',
  FAILED: 'error',
  CANCELLED: 'default',
};

const statusText: Record<string, string> = {
  DRAFT: 'page.campaign.status.draft',
  SCHEDULED: 'page.campaign.status.scheduled',
  SENDING: 'page.campaign.status.sending',
  SENT: 'page.campaign.status.sent',
  PARTIAL: 'page.campaign.status.partial',
  FAILED: 'page.campaign.status.failed',
  CANCELLED: 'page.campaign.status.cancelled',
};

const formOptions = {
  collapsed: false,
  showCollapseButton: true,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.campaign.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<Campaign> = {
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
        return await fetchListCampaigns(
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
      title: $t('page.campaign.name'),
      field: 'name',
      minWidth: 160,
      fixed: 'left',
      align: 'left',
    },
    {
      title: $t('page.campaign.segmentId'),
      field: 'segmentId',
      minWidth: 100,
    },
    {
      title: $t('page.campaign.title'),
      field: 'title',
      minWidth: 180,
      align: 'left',
    },
    {
      title: $t('page.campaign.statusLabel'),
      field: 'status',
      minWidth: 110,
      slots: { default: 'status' },
    },
    {
      title: $t('page.campaign.progress'),
      field: 'totalCount',
      minWidth: 180,
      formatter: ({ row }) =>
        `${row.sentCount ?? 0}/${row.totalCount ?? 0}${
          (row.failedCount ?? 0) > 0 ? ` (${row.failedCount} 失败)` : ''
        }`,
    },
    {
      title: $t('page.campaign.scheduledAt'),
      field: 'scheduledAt',
      formatter: 'formatDateTime',
      minWidth: 160,
    },
    {
      title: $t('page.campaign.lastSentAt'),
      field: 'lastSentAt',
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
      minWidth: 200,
    },
  ],
};

const gridEvents: VxeGridListeners<Campaign> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: CampaignDrawer,
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

async function handleSend(row: any) {
  try {
    await sendCampaign(row.id);
    notification.success({
      message: $t('page.campaign.sendSuccess'),
    });
    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('page.campaign.sendFailed'),
    });
  }
}

async function handleCancel(row: any) {
  try {
    await cancelCampaign(row.id);
    notification.success({
      message: $t('page.campaign.cancelSuccess'),
    });
    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('page.campaign.cancelFailed'),
    });
  }
}

function canSend(row: Campaign) {
  return ['DRAFT', 'CANCELLED', 'FAILED'].includes(row.status || '');
}

function canCancel(row: Campaign) {
  return ['DRAFT', 'SCHEDULED'].includes(row.status || '');
}

async function handleDelete(row: any) {
  try {
    await deleteCampaign(row.id);
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
    <Grid :table-title="$t('menu.campaign.moduleName')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.campaign.button.create') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-tag :color="statusColor[row.status] || 'default'">
          {{
            row.status && statusText[row.status]
              ? $t(statusText[row.status] as string)
              : row.status
          }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          v-if="canSend(row)"
          type="link"
          :icon="h(LucideSend)"
          :title="$t('page.campaign.send')"
          @click.stop="handleSend(row)"
        />
        <a-button
          v-if="canCancel(row)"
          type="link"
          :icon="h(LucideRocket)"
          :title="$t('page.campaign.cancel')"
          @click.stop="handleCancel(row)"
        />
        <a-button
          v-if="['DRAFT', 'CANCELLED'].includes(row.status || '')"
          type="link"
          :icon="h(LucideFilePenLine)"
          @click.stop="handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('menu.campaign.moduleName'),
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
