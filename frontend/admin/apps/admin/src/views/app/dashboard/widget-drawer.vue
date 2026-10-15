<script lang="ts" setup>
import type { ubaservicev1_DashboardWidget as DashboardWidget } from '#/generated/api/admin/service/v1';

import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { LucidePencil, LucidePlus, LucideTrash2 } from '@vben/icons';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { CodeEditor } from '#/adapter/component/Editor';
import { useUpdateDashboard } from '#/api';

import {
  defaultParamsJson,
  isKnownModel,
  modelLabelKey,
  WIDGET_MODELS,
} from './widget-models';

const { mutateAsync: updateDashboard } = useUpdateDashboard();

const dashboardId = ref(0);
const widgets = ref<DashboardWidget[]>([]);

// 组件编辑区是否展开；editingIndex = -1 表示新增，>=0 表示修改对应下标
const editingOpen = ref(false);
const editingIndex = ref(-1);
// 待确认删除的行下标（行内二次确认，不依赖弹层）
const confirmDeleteIndex = ref(-1);
const formTitle = ref('');
const formModel = ref<string>();
const formSpan = ref<number>(12);
const formParams = ref('');

const getTitle = computed(() => $t('page.dashboard.widget.manage'));

function resetEditor() {
  editingOpen.value = false;
  editingIndex.value = -1;
  confirmDeleteIndex.value = -1;
  formTitle.value = '';
  formModel.value = undefined;
  formSpan.value = 12;
  formParams.value = '';
}

function openCreate() {
  resetEditor();
  editingOpen.value = true;
  formModel.value = WIDGET_MODELS[0]?.value;
  formParams.value = defaultParamsJson(formModel.value || 'event-trend');
}

function openEdit(widget: DashboardWidget, index: number) {
  resetEditor();
  editingOpen.value = true;
  editingIndex.value = index;
  formTitle.value = widget.title || '';
  formModel.value = widget.model;
  formSpan.value = Number(widget.span ?? 12);
  formParams.value = widget.paramsJson || defaultParamsJson(widget.model || '');
}

/** 校验并把编辑区内容写回组件列表；失败时返回 false */
function commitEditor(): boolean {
  const title = formTitle.value.trim();
  if (!title) {
    notification.error({ message: $t('page.dashboard.widget.titleRequired') });
    return false;
  }
  if (!formModel.value) {
    notification.error({ message: $t('page.dashboard.widget.modelRequired') });
    return false;
  }
  let params: any;
  try {
    params = JSON.parse(formParams.value || '{}');
  } catch {
    notification.error({ message: $t('page.dashboard.widget.paramsInvalid') });
    return false;
  }
  if (typeof params !== 'object' || params === null || Array.isArray(params)) {
    notification.error({ message: $t('page.dashboard.widget.paramsInvalid') });
    return false;
  }

  const span = Math.min(Math.max(Number(formSpan.value) || 12, 1), 24);
  const widget: DashboardWidget = {
    id:
      editingIndex.value >= 0
        ? (widgets.value[editingIndex.value]?.id ?? nextWidgetId())
        : nextWidgetId(),
    model: formModel.value,
    paramsJson: JSON.stringify(params, null, 2),
    span,
    title,
  };

  if (editingIndex.value >= 0) {
    widgets.value.splice(editingIndex.value, 1, widget);
  } else {
    widgets.value.push(widget);
  }
  resetEditor();
  return true;
}

/** 生成看板内唯一的组件 ID（w1、w2 …） */
function nextWidgetId(): string {
  let max = 0;
  for (const w of widgets.value) {
    const matched = /^w(\d+)$/.exec(w.id || '');
    if (matched) {
      max = Math.max(max, Number(matched[1]));
    }
  }
  return `w${max + 1}`;
}

function removeWidget(index: number) {
  widgets.value.splice(index, 1);
  confirmDeleteIndex.value = -1;
  if (editingIndex.value === index) {
    resetEditor();
  }
}

/** 切换模型时，若参数还没动过则同步换成新模型的默认模板 */
function handleModelChange() {
  let current: any;
  try {
    current = JSON.parse(formParams.value || '{}');
  } catch {
    return;
  }
  const untouched =
    current === null ||
    typeof current !== 'object' ||
    Object.keys(current).length === 0 ||
    (current.lastDays !== undefined && Object.keys(current).length <= 4);
  if (untouched) {
    formParams.value = defaultParamsJson(formModel.value || '');
  }
}

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {
    if (editingOpen.value && !commitEditor()) {
      return;
    }

    drawerApi.setState({ loading: true });
    try {
      await updateDashboard({
        id: dashboardId.value,
        values: { widgets: widgets.value.map((w) => ({ ...w })) },
      });
      notification.success({
        message: $t('ui.notification.update_success'),
      });
      drawerApi.close();
    } catch {
      notification.error({
        message: $t('ui.notification.update_failed'),
      });
    } finally {
      drawerApi.setState({ loading: false });
    }
  },

  onOpenChange(isOpen) {
    if (isOpen) {
      const data = drawerApi.getData<{
        dashboardId: number;
        widgets: DashboardWidget[];
      }>();
      dashboardId.value = data?.dashboardId ?? 0;
      widgets.value = (data?.widgets || []).map((w) => ({ ...w }));
      resetEditor();
    }
  },
});
</script>

<template>
  <Drawer :title="getTitle" class="w-[640px]">
    <div class="mx-4 flex flex-col gap-3">
      <div class="flex items-center justify-between">
        <span class="font-medium">
          {{ $t('page.dashboard.widget.list') }}
          <a-tag class="ml-1">{{ widgets.length }}</a-tag>
        </span>
        <a-button size="small" type="primary" @click="openCreate">
          <LucidePlus class="mr-1 size-4" />
          {{ $t('page.dashboard.widget.add') }}
        </a-button>
      </div>

      <div
        v-if="widgets.length === 0"
        class="text-muted-foreground py-2 text-sm"
      >
        {{ $t('page.dashboard.widget.empty') }}
      </div>
      <template v-else>
        <div
          v-for="(widget, index) in widgets"
          :key="widget.id || index"
          class="flex items-center justify-between rounded-md border px-3 py-2"
        >
          <div class="flex min-w-0 flex-col gap-1">
            <span class="truncate font-medium">{{ widget.title }}</span>
            <div class="flex items-center gap-1">
              <a-tag v-if="isKnownModel(widget.model)" color="processing">
                {{ $t(modelLabelKey(widget.model)) }}
              </a-tag>
              <a-tag v-else color="warning">{{ widget.model }}</a-tag>
              <a-tag>
                {{ $t('page.dashboard.widget.span') }} {{ widget.span ?? 12 }}
              </a-tag>
            </div>
          </div>
          <div class="flex shrink-0 items-center gap-1">
            <template v-if="confirmDeleteIndex === index">
              <span class="text-destructive text-xs">
                {{ $t('page.dashboard.widget.confirmDelete') }}
              </span>
              <a-button size="small" @click="confirmDeleteIndex = -1">
                {{ $t('ui.button.cancel') }}
              </a-button>
              <a-button
                danger
                size="small"
                type="primary"
                @click="removeWidget(index)"
              >
                {{ $t('ui.button.ok') }}
              </a-button>
            </template>
            <template v-else>
              <a-button
                size="small"
                type="text"
                @click="openEdit(widget, index)"
              >
                <LucidePencil class="size-4" />
              </a-button>
              <a-button
                danger
                size="small"
                type="text"
                @click="confirmDeleteIndex = index"
              >
                <LucideTrash2 class="size-4" />
              </a-button>
            </template>
          </div>
        </div>
      </template>

      <a-divider class="my-1" />

      <div v-if="editingOpen" class="flex flex-col gap-3">
        <div class="flex items-center gap-3">
          <span class="w-20 shrink-0 text-right text-sm">
            {{ $t('page.dashboard.widget.title') }}
          </span>
          <a-input
            v-model:value="formTitle"
            :placeholder="$t('ui.placeholder.input')"
            allow-clear
            class="flex-1"
          />
        </div>
        <div class="flex items-center gap-3">
          <span class="w-20 shrink-0 text-right text-sm">
            {{ $t('page.dashboard.widget.model') }}
          </span>
          <a-select
            v-model:value="formModel"
            :placeholder="$t('ui.placeholder.select')"
            class="flex-1"
            @change="handleModelChange"
          >
            <a-select-option
              v-for="model in WIDGET_MODELS"
              :key="model.value"
              :value="model.value"
            >
              {{ $t(model.labelKey) }}
            </a-select-option>
          </a-select>
        </div>
        <div class="flex items-center gap-3">
          <span class="w-20 shrink-0 text-right text-sm">
            {{ $t('page.dashboard.widget.span') }}
          </span>
          <a-input-number
            v-model:value="formSpan"
            :max="24"
            :min="1"
            class="flex-1"
          />
          <span class="text-muted-foreground shrink-0 text-xs">
            {{ $t('page.dashboard.widget.spanHint') }}
          </span>
        </div>
        <div class="flex flex-col gap-1">
          <span class="w-20 shrink-0 text-right text-sm">
            {{ $t('page.dashboard.widget.params') }}
          </span>
          <CodeEditor
            v-model="formParams"
            :options="{
              language: 'json',
              lineNumbers: true,
              minimap: false,
              wordWrap: 'on',
            }"
            height="240px"
          />
          <span class="text-muted-foreground text-xs">
            {{ $t('page.dashboard.widget.paramsHint') }}
          </span>
        </div>
        <div class="flex justify-end gap-2">
          <a-button @click="resetEditor">
            {{ $t('page.dashboard.widget.cancelEdit') }}
          </a-button>
          <a-button type="primary" @click="commitEditor">
            {{ $t('page.dashboard.widget.save') }}
          </a-button>
        </div>
      </div>
    </div>
  </Drawer>
</template>
