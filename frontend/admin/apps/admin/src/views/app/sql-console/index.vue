<script lang="ts" setup>
import type { ubaservicev1_ExecuteSqlQueryResponse } from '#/generated/api/admin/service/v1';

import { computed, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { LucidePlay, LucideTable2 } from '@vben/icons';

import { Card, message } from 'ant-design-vue';

import { fetchExecuteSqlQuery } from '#/api';
import { $t } from '#/locales';
import { exportCsv } from '#/utils/csv-export';

const query = ref(
  "SELECT event_name, COUNT(*) AS cnt, COUNT(DISTINCT user_id) AS users\nFROM events_fact\nGROUP BY event_name\nORDER BY cnt DESC\nLIMIT 50",
);
const maxRows = ref<number>(200);
const running = ref(false);
const result = ref<ubaservicev1_ExecuteSqlQueryResponse>();

const tableColumns = computed(() =>
  (result.value?.columns || []).map((name) => ({
    title: name,
    dataIndex: name,
    key: name,
    ellipsis: true,
  })),
);

const tableData = computed(() => {
  const columns = result.value?.columns || [];
  return (result.value?.rows || []).map((row, index) => {
    const item: Record<string, any> = { __row: index };
    (row.values || []).forEach((value, i) => {
      item[columns[i] ?? String(i)] = value;
    });
    return item;
  });
});

async function run() {
  if (!query.value.trim()) {
    message.warning($t('page.sqlConsole.emptyQuery'));
    return;
  }
  running.value = true;
  try {
    result.value = await fetchExecuteSqlQuery({
      query: query.value,
      maxRows: maxRows.value,
    } as any);
  } catch (error: any) {
    result.value = undefined;
    message.error(error?.response?.data?.message || String(error));
  } finally {
    running.value = false;
  }
}

function handleExport() {
  const resp = result.value;
  if (!resp) return;
  exportCsv(
    'sql-query',
    resp.columns || [],
    (resp.rows || []).map((row) =>
      (row.values || []).map(
        (value) =>
          typeof value === 'string' || typeof value === 'number'
            ? value
            : JSON.stringify(value ?? null),
      ),
    ),
  );
}
</script>

<template>
  <Page auto-content-height>
    <Card size="small" class="mb-4">
      <Textarea
        v-model:value="query"
        :rows="8"
        class="font-mono"
        :placeholder="$t('page.sqlConsole.placeholder')"
        @keydown.ctrl.enter="run"
        @keydown.meta.enter="run"
      />
      <div class="mt-3 flex flex-wrap items-center gap-3">
        <a-button type="primary" :loading="running" @click="run">
          <template #icon>
            <LucidePlay />
          </template>
          {{ $t('page.sqlConsole.run') }}
          <span class="ml-1 text-xs opacity-60">(Ctrl+Enter)</span>
        </a-button>
        <span class="text-muted-foreground text-sm">
          {{ $t('page.sqlConsole.maxRows') }}
        </span>
        <a-select v-model:value="maxRows" style="width: 120px">
          <a-select-option :value="200">200</a-select-option>
          <a-select-option :value="500">500</a-select-option>
          <a-select-option :value="1000">1000</a-select-option>
        </a-select>
        <span class="text-muted-foreground text-xs">
          {{ $t('page.sqlConsole.readOnlyHint') }}
        </span>
      </div>
    </Card>

    <Card size="small">
      <template #title>
        <div class="flex items-center gap-2">
          <LucideTable2 />
          {{ $t('page.sqlConsole.results') }}
          <template v-if="result">
            <a-tag>
              {{ result.rowCount }} {{ $t('page.sqlConsole.rowsUnit') }}
            </a-tag>
            <a-tag v-if="result.truncated" color="orange">
              {{ $t('page.sqlConsole.truncated') }}
            </a-tag>
            <a-tag>{{ result.durationMs }} ms</a-tag>
          </template>
        </div>
      </template>
      <template #extra>
        <a-button
          v-if="result"
          size="small"
          @click="handleExport"
        >
          {{ $t('page.analytics.exportCsv') }}
        </a-button>
      </template>
      <a-table
        v-if="result"
        :data-source="tableData"
        :columns="tableColumns"
        :pagination="{ pageSize: 20, showSizeChanger: false }"
        :scroll="{ x: 'max-content' }"
        row-key="__row"
        size="small"
      />
      <a-empty v-else :description="$t('page.sqlConsole.emptyResult')" />
    </Card>
  </Page>
</template>

<style scoped></style>
