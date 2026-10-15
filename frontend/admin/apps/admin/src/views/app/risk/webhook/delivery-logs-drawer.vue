<script lang="ts" setup>
import type { ubaservicev1_WebhookDeliveryLog as DeliveryLog } from '#/generated/api/admin/service/v1';

import { ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';

import { apiClient } from '#/api/client';
import { $t } from '#/locales';

const logs = ref<DeliveryLog[]>([]);
const loading = ref(false);

async function load(webhookId: number) {
  loading.value = true;
  try {
    const resp = await apiClient.webhookService.ListDeliveryLogs({
      webhookId,
      limit: 100,
    } as any);
    logs.value = (resp.items || []) as DeliveryLog[];
  } finally {
    loading.value = false;
  }
}

const [Drawer, drawerApi] = useVbenDrawer({
  onOpenChange(isOpen) {
    if (isOpen) {
      const data = drawerApi.getData<Record<string, any>>();
      load(data?.row?.id || 0);
    }
  },
});
</script>

<template>
  <Drawer :title="$t('page.webhook.deliveryLogs')" class="w-[860px]">
    <a-spin :spinning="loading">
      <a-table
        :data-source="logs"
        :pagination="{ pageSize: 10 }"
        row-key="id"
        size="small"
      >
        <a-table-column
          :title="$t('page.webhook.log.eventType')"
          data-index="eventType"
          :width="130"
        />
        <a-table-column
          :title="$t('page.webhook.log.success')"
          data-index="success"
          :width="90"
        >
          <template #default="{ text }">
            <a-tag :color="text ? 'green' : 'red'">
              {{
                text ? $t('page.webhook.log.ok') : $t('page.webhook.log.failed')
              }}
            </a-tag>
          </template>
        </a-table-column>
        <a-table-column
          :title="$t('page.webhook.log.statusCode')"
          data-index="statusCode"
          :width="100"
        />
        <a-table-column
          :title="$t('page.webhook.log.retryCount')"
          data-index="retryCount"
          :width="90"
        />
        <a-table-column
          :title="$t('page.webhook.log.errorMessage')"
          data-index="errorMessage"
          :width="200"
        >
          <template #default="{ text }">
            {{ text || '-' }}
          </template>
        </a-table-column>
        <a-table-column
          :title="$t('page.webhook.log.sentAt')"
          data-index="sentAt"
          :width="170"
        >
          <template #default="{ text }">
            {{ text ? new Date(text).toLocaleString() : '-' }}
          </template>
        </a-table-column>
      </a-table>
    </a-spin>
  </Drawer>
</template>
