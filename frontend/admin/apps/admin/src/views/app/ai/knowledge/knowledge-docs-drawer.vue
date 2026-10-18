<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Button,
  message,
  notification,
  Select,
  Table,
  Textarea,
  Upload,
} from 'ant-design-vue';

import {
  type AiDoc,
  fetchListAiDocs,
  useDeleteAiDoc,
  useSearchAiKnowledge,
  useUploadAiDoc,
  useUploadAiDocFile,
} from '#/api';

defineOptions({ name: 'KnowledgeDocsDrawer' });

const row = ref();
const docs = ref<AiDoc[]>([]);
const loading = ref(false);

// 上传文件
const uploadingFile = ref(false);
const { mutateAsync: uploadDocFile } = useUploadAiDocFile();

// 粘贴文本
const pasteOpen = ref(false);
const pasteName = ref('');
const pasteContent = ref('');
const uploadingText = ref(false);
const { mutateAsync: uploadDocText } = useUploadAiDoc();

// 检索测试
const searchQuery = ref('');
const searchTopK = ref<number | undefined>(3);
const searching = ref(false);
const searchHits = ref<
  Array<{
    chunkIndex?: number;
    content: string;
    docName: string;
    score: number;
  }>
>([]);
const { mutateAsync: searchKnowledge } = useSearchAiKnowledge();

const { mutateAsync: deleteDoc } = useDeleteAiDoc();

const docColumns = computed(() => [
  { title: $t('page.ai.docs.name'), dataIndex: 'name', key: 'name' },
  {
    title: $t('page.ai.docs.chunkCount'),
    dataIndex: 'chunkCount',
    key: 'chunkCount',
    width: 90,
  },
  {
    title: $t('page.ai.docs.status'),
    dataIndex: 'status',
    key: 'status',
    width: 100,
  },
  {
    title: $t('ui.table.action'),
    key: 'action',
    width: 90,
  },
]);

const title = computed(() => {
  const baseName = row.value?.row?.name ?? '';
  return `${$t('page.ai.docs.manage')}${baseName ? ` · ${baseName}` : ''}`;
});

async function reloadDocs() {
  if (!row.value?.row?.id) return;
  loading.value = true;
  try {
    const resp = await fetchListAiDocs({ baseId: row.value.row.id });
    docs.value = resp.items ?? [];
  } finally {
    loading.value = false;
  }
}

// 文件入库：beforeUpload 返回 false 阻止自动上传，转 base64 后走 UploadDocFile
async function handleFile(file: File) {
  if (!row.value?.row?.id) return false;
  if (file.size > 10 * 1024 * 1024) {
    message.error($t('page.ai.docs.fileTooLarge'));
    return false;
  }
  uploadingFile.value = true;
  try {
    const base64 = await fileToBase64(file);
    const resp = await uploadDocFile({
      baseId: row.value.row.id,
      fileName: file.name,
      contentBase64: base64,
    });
    notification.success({
      message: $t('page.ai.docs.uploadSuccess', {
        count: resp.chunkCount ?? 0,
      }),
    });
    await reloadDocs();
  } catch {
    notification.error({ message: $t('page.ai.docs.uploadFailed') });
  } finally {
    uploadingFile.value = false;
  }
  return false;
}

async function handlePasteSubmit() {
  if (!pasteName.value.trim() || !pasteContent.value.trim()) {
    message.warning($t('page.ai.docs.pasteRequired'));
    return;
  }
  uploadingText.value = true;
  try {
    const resp = await uploadDocText({
      baseId: row.value.row.id,
      name: pasteName.value.trim(),
      content: pasteContent.value,
    });
    notification.success({
      message: $t('page.ai.docs.uploadSuccess', {
        count: resp.chunkCount ?? 0,
      }),
    });
    pasteOpen.value = false;
    pasteName.value = '';
    pasteContent.value = '';
    await reloadDocs();
  } catch {
    notification.error({ message: $t('page.ai.docs.uploadFailed') });
  } finally {
    uploadingText.value = false;
  }
}

async function handleDeleteDoc(record: AiDoc) {
  try {
    await deleteDoc({ baseId: row.value.row.id, id: record.id as number });
    await reloadDocs();
  } catch {
    notification.error({ message: $t('ui.notification.delete_failed') });
  }
}

async function handleSearch() {
  if (!searchQuery.value.trim()) {
    message.warning($t('page.ai.docs.searchRequired'));
    return;
  }
  searching.value = true;
  searchHits.value = [];
  try {
    const resp = await searchKnowledge({
      baseId: row.value.row.id,
      query: searchQuery.value,
      topK: searchTopK.value,
    });
    const hits = resp.hits ?? [];
    searchHits.value = hits.map((hit) => ({
      content: hit.content ?? '',
      score: hit.score ?? 0,
      docName: docNameOf(hit.doc?.id),
      chunkIndex: hit.chunkIndex,
    }));
  } catch {
    message.error($t('page.ai.docs.searchFailed'));
  } finally {
    searching.value = false;
  }
}

function docNameOf(docId?: number) {
  if (docId === undefined || docId === null) return '-';
  const doc = docs.value.find((d: AiDoc) => d.id === docId);
  return doc?.name ?? `#${docId}`;
}

function fileToBase64(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.addEventListener('load', () => {
      const result = String(reader.result ?? '');
      // data:...;base64,xxxx → xxxx
      const idx = result.indexOf('base64,');
      resolve(idx === -1 ? result : result.slice(idx + 'base64,'.length));
    });
    reader.addEventListener('error', () => reject(reader.error));
    reader.readAsDataURL(file);
  });
}

const [Drawer, drawerApi] = useVbenDrawer({
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      row.value = drawerApi.getData<Record<string, any>>();
      searchHits.value = [];
      searchQuery.value = '';
      void reloadDocs();
    }
  },
});
</script>

<template>
  <Drawer class="w-[880px]" :title="title">
    <div class="flex flex-col gap-4">
      <!-- 上传 -->
      <div class="flex flex-wrap items-center gap-2">
        <Upload
          :multiple="false"
          :before-upload="handleFile"
          :show-upload-list="false"
          accept=".txt,.md,.docx,.pdf"
        >
          <Button type="primary" :loading="uploadingFile">
            {{ $t('page.ai.docs.uploadFile') }}
          </Button>
        </Upload>
        <Button :loading="uploadingText" @click="pasteOpen = !pasteOpen">
          {{ $t('page.ai.docs.pasteText') }}
        </Button>
        <span class="text-xs opacity-60">
          {{ $t('page.ai.docs.uploadHint') }}
        </span>
      </div>

      <!-- 粘贴文本面板 -->
      <div v-if="pasteOpen" class="rounded border p-3">
        <div class="mb-2 flex flex-col gap-2">
          <Textarea
            v-model:value="pasteName"
            :placeholder="$t('page.ai.docs.name')"
            :rows="1"
          />
          <Textarea
            v-model:value="pasteContent"
            :placeholder="$t('page.ai.docs.pastePlaceholder')"
            :rows="6"
          />
          <div class="flex justify-end gap-2">
            <Button @click="pasteOpen = false">
              {{ $t('ui.button.cancel') }}
            </Button>
            <Button
              type="primary"
              :loading="uploadingText"
              @click="handlePasteSubmit"
            >
              {{ $t('ui.button.confirm') }}
            </Button>
          </div>
        </div>
      </div>

      <!-- 文档列表 -->
      <Table
        :columns="docColumns"
        :data-source="docs"
        :loading="loading"
        :pagination="false"
        row-key="id"
        size="small"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'status'">
            <a-tag :color="record.status === 'READY' ? 'success' : 'error'">
              {{ record.status }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-popconfirm
              :cancel-text="$t('ui.button.cancel')"
              :ok-text="$t('ui.button.ok')"
              :title="
                $t('ui.text.do_you_want_delete', { moduleName: record.name })
              "
              @confirm="handleDeleteDoc(record as AiDoc)"
            >
              <a-button danger type="link" size="small">
                {{ $t('ui.button.delete') }}
              </a-button>
            </a-popconfirm>
          </template>
        </template>
      </Table>

      <!-- 检索测试 -->
      <div class="rounded border p-3">
        <div class="mb-2 font-medium">{{ $t('page.ai.docs.searchTitle') }}</div>
        <div class="flex flex-wrap items-center gap-2">
          <Textarea
            v-model:value="searchQuery"
            :placeholder="$t('page.ai.docs.searchPlaceholder')"
            :rows="2"
            class="min-w-[320px] flex-1"
          />
          <Select
            v-model:value="searchTopK"
            class="w-24"
            :options="[
              { label: 'Top 3', value: 3 },
              { label: 'Top 5', value: 5 },
              { label: 'Top 10', value: 10 },
            ]"
          />
          <Button type="primary" :loading="searching" @click="handleSearch">
            {{ $t('page.ai.docs.searchRun') }}
          </Button>
        </div>
        <div v-if="searchHits.length > 0" class="mt-3 flex flex-col gap-2">
          <div
            v-for="(hit, index) in searchHits"
            :key="index"
            class="rounded bg-gray-50 p-2 text-xs dark:bg-gray-800"
          >
            <div class="mb-1 flex items-center gap-2 opacity-70">
              <span>#{{ index + 1 }} {{ hit.docName }}</span>
              <span v-if="hit.chunkIndex !== undefined">
                {{ $t('page.ai.docs.chunk') }}{{ hit.chunkIndex }}
              </span>
              <span class="ml-auto">score: {{ hit.score.toFixed(4) }}</span>
            </div>
            <div class="whitespace-pre-wrap">{{ hit.content }}</div>
          </div>
        </div>
      </div>
    </div>
  </Drawer>
</template>
