<script lang="ts" setup>
import type {
  aiservicev1_AiQueryHistoryItem,
  aiservicev1_AskAiQueryResponse,
} from '#/generated/api/admin/service/v1';

import { computed, h, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { LucidePlay, LucideTrash2 } from '@vben/icons';

import { Card, message, Switch, Table } from 'ant-design-vue';

import { useAskAiQuery } from '#/api';
import { $t } from '#/locales';

defineOptions({ name: 'AiQuery' });

interface HistoryTurn {
  question: string;
  sql: string;
  resultSummary: string;
}

const question = ref('');
const withAnswer = ref(true);
const running = ref(false);
const result = ref<aiservicev1_AskAiQueryResponse>();
const history = ref<HistoryTurn[]>([]);

const { mutateAsync: ask } = useAskAiQuery();

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
  const q = question.value.trim();
  if (!q) {
    message.warning($t('page.ai.query.emptyQuestion'));
    return;
  }
  running.value = true;
  try {
    const historyItems: aiservicev1_AiQueryHistoryItem[] = history.value
      .slice(-5)
      .map((h) => ({
        question: h.question,
        sql: h.sql,
        resultSummary: h.resultSummary,
      }));
    const resp = await ask({
      question: q,
      lang: 'zh-CN',
      withAnswer: withAnswer.value,
      history: historyItems,
    });
    result.value = resp;
    if (resp.errorMessage) {
      message.warning($t('page.ai.query.executeFailed'));
    }
    // 记入追问历史（供模型消解指代）
    history.value.push({
      question: q,
      sql: resp.sql ?? '',
      resultSummary: (resp.rows || [])
        .slice(0, 3)
        .map((r) => (r.values || []).join(' | '))
        .join('；'),
    });
    if (history.value.length > 10) {
      history.value = history.value.slice(-10);
    }
  } catch (error: any) {
    result.value = undefined;
    message.error(error?.response?.data?.message || String(error));
  } finally {
    running.value = false;
  }
}

function clearHistory() {
  history.value = [];
  result.value = undefined;
  question.value = '';
}
</script>

<template>
  <Page :title="$t('menu.ai.query')">
    <div class="flex flex-col gap-4">
      <!-- 提问区 -->
      <Card size="small">
        <div class="flex items-end gap-2">
          <a-textarea
            v-model:value="question"
            :placeholder="$t('page.ai.query.placeholder')"
            :rows="2"
            class="flex-1"
            @keydown.enter.exact.prevent="run"
          />
        </div>
        <div class="mt-3 flex flex-wrap items-center gap-3">
          <a-button type="primary" :loading="running" @click="run">
            <template #icon>
              <LucidePlay />
            </template>
            {{ $t('page.ai.query.ask') }}
          </a-button>
          <div class="flex items-center gap-1 text-sm">
            <Switch v-model:checked="withAnswer" size="small" />
            {{ $t('page.ai.query.withAnswer') }}
          </div>
          <a-button :icon="h(LucideTrash2)" @click="clearHistory">
            {{ $t('page.ai.query.clear') }}
          </a-button>
          <span class="text-xs opacity-60">
            {{ $t('page.ai.query.safetyNote') }}
          </span>
        </div>
      </Card>

      <!-- 自然语言结论 -->
      <Card
        v-if="result?.answer"
        size="small"
        :title="$t('page.ai.query.answer')"
      >
        <div class="whitespace-pre-wrap text-sm">{{ result.answer }}</div>
      </Card>

      <!-- 生成的 SQL -->
      <Card v-if="result?.sql" size="small" :title="$t('page.ai.query.sql')">
        <pre
          class="overflow-x-auto rounded bg-gray-50 p-3 text-xs dark:bg-gray-900"
          >{{ result.sql }}
        </pre>
        <div v-if="result.errorMessage" class="mt-2 text-sm text-red-500">
          {{ result.errorMessage }}
        </div>
      </Card>

      <!-- 结果表 -->
      <Card v-if="result && !result.errorMessage" size="small">
        <template #title>
          <div class="flex items-center gap-2">
            {{ $t('page.ai.query.result') }}
            <span class="text-xs opacity-50">
              {{
                $t('page.ai.query.rowCount', { count: result.rowCount ?? 0 })
              }}
            </span>
          </div>
        </template>
        <Table
          :columns="tableColumns"
          :data-source="tableData"
          :pagination="{ pageSize: 20, showSizeChanger: false }"
          :scroll="{ x: 'max-content' }"
          row-key="__row"
          size="small"
        />
      </Card>
    </div>
  </Page>
</template>
