<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { computed, ref } from 'vue';

import { Page, type VbenFormProps } from '@vben/common-ui';

import {
  Card,
  InputNumber,
  message,
  notification,
  Progress,
  Statistic,
} from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  fetchListAiUsageLogs,
  PaginationQuery,
  useAiUsageSummary,
  useGetTenantQuota,
  useSetTenantQuota,
} from '#/api';
import { $t } from '#/locales';

defineOptions({ name: 'AiUsage' });

const summary = useAiUsageSummary();

// 配额进度（0~100，封顶；超限变红由 status 控制）
const quotaPercent = computed(() => {
  const used = summary.data.value?.monthTokens ?? 0;
  const limit = Number(summary.data.value?.quotaLimit ?? 0);
  if (limit <= 0) return 0;
  return Math.min(100, Math.round((used / limit) * 100));
});
const quotaExceeded = computed(
  () =>
    (summary.data.value?.monthTokens ?? 0) >=
    Number(summary.data.value?.quotaLimit ?? 0),
);

// 租户配额查看 / 设置（平台管理员）
const quotaTenantId = ref<number | undefined>();
const quotaLimitInput = ref<number | undefined>();
const quotaInfo = ref<{
  configured?: boolean;
  monthlyTokenLimit?: number | string;
}>();
const { mutateAsync: getTenantQuota, isPending: queryingQuota } =
  useGetTenantQuota();
const { mutateAsync: setTenantQuota, isPending: savingQuota } =
  useSetTenantQuota();

async function handleQueryQuota() {
  if (!quotaTenantId.value) {
    message.warning($t('page.ai.usage.tenantRequired'));
    return;
  }
  try {
    const resp = await getTenantQuota({ tenantId: quotaTenantId.value });
    quotaInfo.value = resp;
    quotaLimitInput.value = Number(resp.monthlyTokenLimit ?? 0);
  } catch {
    notification.error({ message: $t('page.ai.usage.quotaQueryFailed') });
  }
}

async function handleSaveQuota() {
  if (!quotaTenantId.value) {
    message.warning($t('page.ai.usage.tenantRequired'));
    return;
  }
  try {
    await setTenantQuota({
      tenantId: quotaTenantId.value,
      monthlyTokenLimit: quotaLimitInput.value ?? 0,
    } as never);
    notification.success({ message: $t('page.ai.usage.quotaSaveSuccess') });
    await summary.refetch();
  } catch (error: any) {
    notification.error({
      message: $t('page.ai.usage.quotaSaveFailed'),
      description: error?.response?.data?.message || String(error),
    });
  }
}

const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'modelName__contains',
      label: $t('page.ai.usage.modelName'),
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
        return await fetchListAiUsageLogs(
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
      title: $t('page.ai.usage.modelName'),
      field: 'modelName',
      fixed: 'left',
      minWidth: 180,
    },
    {
      title: $t('page.ai.usage.promptTokens'),
      field: 'promptTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.completionTokens'),
      field: 'completionTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.totalTokens'),
      field: 'totalTokens',
      minWidth: 110,
    },
    {
      title: $t('page.ai.usage.durationMs'),
      field: 'durationMs',
      minWidth: 100,
    },
    {
      title: $t('page.ai.usage.conversationId'),
      field: 'conversationId',
      minWidth: 110,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      minWidth: 140,
    },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <div class="mb-4 grid grid-cols-2 gap-4 lg:grid-cols-4">
      <Card size="small">
        <Statistic
          :title="$t('page.ai.usage.monthTokens')"
          :value="summary.data.value?.monthTokens ?? 0"
        />
      </Card>
      <Card size="small">
        <Statistic
          :title="$t('page.ai.usage.monthCalls')"
          :value="summary.data.value?.monthCalls ?? 0"
        />
      </Card>
      <!-- 配额进度（本租户配置了月度配额时显示） -->
      <Card v-if="summary.data.value?.quotaConfigured" size="small">
        <div class="flex flex-col gap-1">
          <span class="text-sm opacity-60">
            {{ $t('page.ai.usage.quotaTitle') }}
          </span>
          <Progress
            :percent="quotaPercent"
            :status="quotaExceeded ? 'exception' : 'normal'"
            size="small"
          />
          <span class="text-xs opacity-60">
            {{ summary.data.value?.monthTokens ?? 0 }} /
            {{ summary.data.value?.quotaLimit ?? 0 }}
          </span>
        </div>
      </Card>
    </div>

    <!-- 租户配额管理（仅平台管理员可保存成功；租户用户查询会被强制指向本租户） -->
    <Card size="small" class="mb-4" :title="$t('page.ai.usage.quotaManage')">
      <div class="flex flex-wrap items-center gap-3">
        <InputNumber
          v-model:value="quotaTenantId"
          :min="1"
          :placeholder="$t('page.ai.usage.tenantId')"
          class="w-40"
        />
        <a-button :loading="queryingQuota" @click="handleQueryQuota">
          {{ $t('page.ai.usage.quotaQuery') }}
        </a-button>
        <span v-if="quotaInfo" class="text-sm">
          {{
            quotaInfo.configured
              ? $t('page.ai.usage.quotaConfigured', {
                  limit: String(quotaInfo.monthlyTokenLimit ?? 0),
                })
              : $t('page.ai.usage.quotaUnlimited')
          }}
        </span>
        <div class="flex items-center gap-2">
          <InputNumber
            v-model:value="quotaLimitInput"
            :min="0"
            :placeholder="$t('page.ai.usage.monthlyLimit')"
            class="w-52"
          />
          <a-button
            type="primary"
            :loading="savingQuota"
            @click="handleSaveQuota"
          >
            {{ $t('page.ai.usage.quotaSave') }}
          </a-button>
          <span class="text-xs opacity-60">
            {{ $t('page.ai.usage.quotaHint') }}
          </span>
        </div>
      </div>
    </Card>

    <Grid :table-title="$t('menu.ai.usage')" />
  </Page>
</template>
