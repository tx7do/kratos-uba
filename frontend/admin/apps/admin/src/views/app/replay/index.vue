<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_ReplaySessionSummary as ReplaySession } from '#/generated/api/admin/service/v1';

import { h } from 'vue';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { LucideMonitorPlay, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  deleteReplaySession,
  fetchListReplaySessions,
} from '#/api';
import ReplayPlayer from './replay-player.vue';

async function handleDelete(sessionId: string) {
  await deleteReplaySession(sessionId);
  notification.success({
    message: $t('ui.notification.delete_success'),
  });
  gridApi.reload();
}

const gridOptions: VxeGridProps<ReplaySession> = {
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
        // 服务端聚合查询不支持 go-crud 分页协议，全量拉最近 200 条后前端分页
        const resp = await fetchListReplaySessions({
          sessionId: formValues?.sessionId || undefined,
          userId: formValues?.userId ? Number(formValues.userId) : undefined,
        } as any);
        const items = resp.items || [];
        const pageNo = page.currentPage || 1;
        const size = page.pageSize || 20;
        return {
          total: items.length,
          items: items.slice((pageNo - 1) * size, pageNo * size),
        };
      },
    },
  },
  columns: [
    {
      title: $t('page.replay.sessionId'),
      field: 'sessionId',
      minWidth: 260,
      fixed: 'left',
      align: 'left',
      slots: { default: 'sessionId' },
    },
    {
      title: $t('page.replay.userId'),
      field: 'userId',
      minWidth: 100,
    },
    {
      title: $t('page.replay.pageUrl'),
      field: 'pageUrl',
      minWidth: 240,
      align: 'left',
    },
    {
      title: $t('page.replay.chunkCount'),
      field: 'chunkCount',
      minWidth: 100,
    },
    {
      title: $t('page.replay.eventCount'),
      field: 'eventCount',
      minWidth: 100,
    },
    {
      title: $t('page.replay.firstSeenAt'),
      field: 'firstSeenAt',
      formatter: 'formatDateTime',
      minWidth: 160,
    },
    {
      title: $t('page.replay.lastSeenAt'),
      field: 'lastSeenAt',
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

const gridEvents: VxeGridListeners<ReplaySession> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents,
});

const [PlayerDrawer, playerDrawerApi] = useVbenDrawer({
  connectedComponent: ReplayPlayer,
});

function handlePlay(row: any) {
  playerDrawerApi.setData({ sessionId: row.sessionId });
  playerDrawerApi.open();
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.replay.moduleName')">
      <template #sessionId="{ row }">
        <span class="font-mono text-xs">{{ row.sessionId }}</span>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideMonitorPlay)"
          :title="$t('page.replay.play')"
          @click.stop="handlePlay(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="$t('page.replay.deleteConfirm')"
          @confirm="handleDelete(row.sessionId)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <PlayerDrawer />
  </Page>
</template>

<style scoped></style>
