<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import {
  LucideDatabaseZap,
  LucideFilePenLine,
  LucideFolderOpen,
  LucideTrash2,
} from '@vben/icons';

import { message, notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  fetchListAiKnowledgeBases,
  PaginationQuery,
  useDeleteAiKnowledgeBase,
  useReindexAiDocs,
  useSyncOpsKnowledge,
} from '#/api';
import { $t } from '#/locales';

import KnowledgeDocsDrawer from './knowledge-docs-drawer.vue';
import AiKnowledgeDrawer from './knowledge-drawer.vue';

defineOptions({ name: 'AiKnowledgeManagement' });

const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name__contains',
      label: $t('page.ai.knowledge.name'),
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
        return await fetchListAiKnowledgeBases(
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
      title: $t('page.ai.knowledge.name'),
      field: 'name',
      fixed: 'left',
      minWidth: 160,
    },
    {
      title: $t('page.ai.knowledge.description'),
      field: 'description',
      minWidth: 220,
    },
    {
      title: $t('page.ai.knowledge.embeddingModel'),
      field: 'embeddingModel',
      minWidth: 180,
    },
    {
      title: $t('page.ai.knowledge.docCount'),
      field: 'docCount',
      minWidth: 90,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      minWidth: 140,
    },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      minWidth: 130,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [EditDrawer, editDrawerApi] = useVbenDrawer({
  connectedComponent: AiKnowledgeDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      gridApi.reload();
    }
  },
});

const [DocsDrawer, docsDrawerApi] = useVbenDrawer({
  connectedComponent: KnowledgeDocsDrawer,
});

const { mutateAsync: deleteBase } = useDeleteAiKnowledgeBase();
const { mutateAsync: reindexDocs } = useReindexAiDocs();
const { mutateAsync: syncOpsKnowledge, isPending: syncingOps } =
  useSyncOpsKnowledge();

function openEditDrawer(create: boolean, row?: any) {
  editDrawerApi.setData({ create, row });
  editDrawerApi.open();
}

function handleCreate() {
  openEditDrawer(true);
}

function handleEdit(row: any) {
  openEditDrawer(false, row);
}

function handleDocs(row: any) {
  docsDrawerApi.setData({ row });
  docsDrawerApi.open();
}

async function handleDelete(row: any) {
  try {
    await deleteBase({ id: row.id });
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

async function handleReindex(baseId?: number) {
  try {
    await reindexDocs({ baseId });
    message.success($t('page.ai.knowledge.reindexEnqueued'));
  } catch {
    message.error($t('page.ai.knowledge.reindexFailed'));
  }
}

async function handleSyncOps() {
  try {
    const resp = await syncOpsKnowledge({});
    notification.success({
      message: $t('page.ai.knowledge.syncOpsSuccess', {
        docs: resp.docCount ?? 0,
        chunks: resp.chunkCount ?? 0,
      }),
    });
    await gridApi.reload();
  } catch {
    notification.error({ message: $t('page.ai.knowledge.syncOpsFailed') });
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.ai.knowledge')">
      <template #toolbar-tools>
        <a-button class="mr-2" :loading="syncingOps" @click="handleSyncOps">
          {{ $t('page.ai.knowledge.syncOps') }}
        </a-button>
        <a-button class="mr-2" @click="handleReindex()">
          <template #icon>
            <LucideDatabaseZap />
          </template>
          {{ $t('page.ai.knowledge.reindexAll') }}
        </a-button>
        <a-button type="primary" class="mr-2" @click="handleCreate">
          {{ $t('page.ai.knowledge.button.create') }}
        </a-button>
      </template>
      <template #docCount="{ row }">
        {{ row.docCount ?? 0 }}
      </template>
      <template #action="{ row }">
        <a-tooltip :title="$t('page.ai.knowledge.docs.manage')">
          <a-button
            type="link"
            :icon="h(LucideFolderOpen)"
            @click.stop="handleDocs(row)"
          />
        </a-tooltip>
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click.stop="handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: row.name || $t('page.ai.knowledge.moduleName'),
            })
          "
          @confirm="handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <EditDrawer />
    <DocsDrawer />
  </Page>
</template>
