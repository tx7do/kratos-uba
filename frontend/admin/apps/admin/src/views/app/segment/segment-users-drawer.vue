<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { ubaservicev1_SegmentUser as SegmentUser } from '#/generated/api/admin/service/v1';

import { ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { fetchListSegmentUsers } from '#/api';
import { $t } from '#/locales';

const segmentId = ref<number>(0);

const gridOptions: VxeGridProps<SegmentUser> = {
  height: 'auto',
  stripe: true,
  autoResize: true,
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }) => {
        return await fetchListSegmentUsers({
          segmentId: segmentId.value,
          page: page.currentPage,
          pageSize: page.pageSize,
        } as any);
      },
    },
  },
  columns: [
    {
      title: $t('page.segment.userId'),
      field: 'userId',
      minWidth: 140,
    },
    {
      title: $t('page.segment.joinedAt'),
      field: 'joinedAt',
      formatter: 'formatDateTime',
      minWidth: 180,
    },
  ],
};

const gridEvents: VxeGridListeners<SegmentUser> = {};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  onOpenChange(isOpen) {
    if (isOpen) {
      const data = drawerApi.getData<Record<string, any>>();
      segmentId.value = data?.row?.id || 0;
      gridApi.query();
    }
  },
});

function getTitle() {
  return $t('page.segment.usersTitle');
}
</script>

<template>
  <Drawer :title="getTitle" class="w-[720px]">
    <Grid />
  </Drawer>
</template>
