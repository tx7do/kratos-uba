<script lang="ts" setup>
import type { ubaservicev1_DashboardWidget as DashboardWidgetType } from '#/generated/api/admin/service/v1';

import { computed, watch } from 'vue';
import { useRoute } from 'vue-router';

import { Page } from '@vben/common-ui';

import { useGetDashboard } from '#/api';

import DashboardWidget from './dashboard-widget.vue';

const route = useRoute();

const dashboardId = computed(() => Number(route.params.id) || 0);

const req = computed(() => ({ queryBy: { id: dashboardId.value } }));

// composable 的 options 参数覆盖了内部 queryKey，这里补齐并按需启用
const { data: dashboard, refetch } = useGetDashboard(
  req as any,
  {
    queryKey: ['getDashboardView', req],
    enabled: () => dashboardId.value > 0,
  } as any,
);

const widgets = computed(
  () => (dashboard.value?.widgets || []) as DashboardWidgetType[],
);

const title = computed(() => dashboard.value?.name || '');

watch(dashboardId, () => refetch());
</script>

<template>
  <Page :title="title" auto-content-height>
    <a-row v-if="widgets.length > 0" :gutter="[16, 16]">
      <a-col
        v-for="widget in widgets"
        :key="widget.id"
        :md="24"
        :lg="Math.min(Math.max(widget.span || 12, 6), 24)"
        :xl="Math.min(Math.max(widget.span || 12, 6), 24)"
      >
        <DashboardWidget :widget="widget" />
      </a-col>
    </a-row>
    <a-empty v-else :description="$t('page.dashboard.empty')" />
  </Page>
</template>
