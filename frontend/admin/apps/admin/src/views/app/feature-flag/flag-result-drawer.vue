<script lang="ts" setup>
import type {
  ubaservicev1_FlagVariantResult,
} from '#/generated/api/admin/service/v1';

import { ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';

import { $t } from '#/locales';
import { fetchFlagResult } from '#/api';

const flagId = ref<number>(0);
const variants = ref<ubaservicev1_FlagVariantResult[]>([]);
const goalEvent = ref<string>();
const loading = ref(false);

async function load() {
  loading.value = true;
  try {
    const resp = await fetchFlagResult({ id: flagId.value } as any);
    variants.value = (resp.variants || []) as ubaservicev1_FlagVariantResult[];
    goalEvent.value = resp.goalEventName;
  } finally {
    loading.value = false;
  }
}

const [Drawer, drawerApi] = useVbenDrawer({
  onOpenChange(isOpen) {
    if (isOpen) {
      const data = drawerApi.getData<Record<string, any>>();
      flagId.value = data?.row?.id || 0;
      load();
    }
  },
});
</script>

<template>
  <Drawer :title="$t('page.featureFlag.resultTitle')" class="w-[720px]">
    <a-spin :spinning="loading">
      <a-alert class="mb-4" type="info">
        {{ $t('page.featureFlag.resultHint') }}
        <template v-if="goalEvent">
          {{ $t('page.featureFlag.goalEvent') }}:
          <a-tag color="blue">{{ goalEvent }}</a-tag>
        </template>
      </a-alert>
      <a-table
        :data-source="variants"
        :pagination="false"
        row-key="variant"
        size="small"
      >
        <a-table-column
          :title="$t('page.featureFlag.variant')"
          data-index="variant"
          :width="160"
        >
          <template #default="{ text }">
            <a-tag :color="text === 'control' ? 'default' : 'blue'">
              {{ text }}
            </a-tag>
          </template>
        </a-table-column>
        <a-table-column
          :title="$t('page.featureFlag.exposedUsers')"
          data-index="exposedUsers"
          :width="140"
        />
        <a-table-column
          :title="$t('page.featureFlag.convertedUsers')"
          data-index="convertedUsers"
          :width="140"
        />
        <a-table-column
          :title="$t('page.featureFlag.conversionRate')"
          data-index="conversionRate"
          :width="160"
        >
          <template #default="{ text }">
            <a-progress
              :percent="Math.round((text || 0) * 10000) / 100"
              size="small"
            />
          </template>
        </a-table-column>
      </a-table>
    </a-spin>
  </Drawer>
</template>
