<script lang="ts" setup>
import type { VbenFormProps } from '@vben-core/form-ui';
import type {
  VxeGridDefines,
  VxeGridInstance,
  VxeGridListeners,
  VxeGridPropTypes,
  VxeGridProps as VxeTableGridProps,
} from 'vxe-table';

import type { ExtendedVxeGridApi, VxeGridProps } from './types';

import {
  computed,
  h,
  nextTick,
  onMounted,
  onUnmounted,
  toRaw,
  useSlots,
  useTemplateRef,
  watch,
} from 'vue';

import { usePriorityValues } from '@vben/hooks';
import {
  EmptyIcon,
  LucideDownload,
  LucideRefreshCw,
  LucideRotateCcw,
  Maximize,
  Search,
  SlidersHorizontal,
} from '@vben/icons';
import { $t } from '@vben/locales';
import { usePreferences } from '@vben/preferences';
import { cloneDeep, cn, mergeWithArrayOverride } from '@vben/utils';
import {
  VbenHelpTooltip,
  VbenIconButton,
  VbenLoading,
  VbenTooltip,
} from '@vben-core/shadcn-ui';

import { VxeGrid, VxeUI } from 'vxe-table';

import { extendProxyOptions } from './extends';
import { useTableForm } from './init';

import 'vxe-table/styles/cssvar.scss';
import 'vxe-pc-ui/styles/cssvar.scss';
import './style.css';

interface Props extends VxeGridProps {
  api: ExtendedVxeGridApi;
}

const props = withDefaults(defineProps<Props>(), {});

const FORM_SLOT_PREFIX = 'form-';

const TOOLBAR_ACTIONS = 'toolbar-actions';
const TOOLBAR_TOOLS = 'toolbar-tools';

const gridRef = useTemplateRef<VxeGridInstance>('gridRef');

const state = props.api?.useStore?.();

const {
  gridOptions,
  class: className,
  gridClass,
  gridEvents,
  formOptions,
  tableTitle,
  tableTitleHelp,
  showSearchForm,
} = usePriorityValues(props, state);

const { isMobile } = usePreferences();

/**
 * 列表工具栏默认工具（下载/刷新/全屏/列设置）：vben 化按钮 + VbenTooltip 样式化提示，
 * 替代 vxe 内建钮的原生 title；下载/全屏/列设置经 commitProxy 与原内建行为一致，
 * 刷新走 api.query()（保留搜索表单值）。
 * 沿用 vxe 内建 toolbarConfig 开关语义：对应键（export/refresh/zoom/custom）显式
 * false 即隐藏该钮；缺省全显示，存量页面不受影响。
 */
const defaultToolbarTools = computed(() => {
  const flags = gridOptions.value?.toolbarConfig as
    | Record<string, boolean | undefined>
    | undefined;
  return [
    {
      code: 'export',
      icon: LucideDownload,
      label: $t('common.download'),
      onClick: () => gridRef.value?.commitProxy('export'),
    },
    {
      code: 'refresh',
      icon: LucideRotateCcw,
      label: $t('common.refresh'),
      onClick: () => props.api?.query(),
    },
    {
      code: 'zoom',
      icon: Maximize,
      label: $t('common.fullscreen'),
      onClick: () => gridRef.value?.commitProxy('zoom'),
    },
    {
      code: 'custom',
      icon: SlidersHorizontal,
      label: $t('common.columnSetting'),
      onClick: () => gridRef.value?.commitProxy('custom'),
    },
  ].filter((tool) => flags?.[tool.code] !== false);
});

const slots = useSlots();

const [Form, formApi] = useTableForm({
  handleSubmit: async () => {
    const formValues = formApi.form.values;
    formApi.setLatestSubmissionValues(toRaw(formValues));
    props.api.reload(formValues);
  },
  handleReset: async () => {
    await formApi.resetForm();
    const formValues = formApi.form.values;
    formApi.setLatestSubmissionValues(formValues);
    props.api.reload(formValues);
  },
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
  showCollapseButton: true,
  // 搜索栏按钮统一（三端对齐 ele）：搜索（实心蓝+放大镜）在前、重置（文字样式+循环箭头）在后；
  // icon 尺寸须钉住（lucide-vue-next 默认 width=24）；间距与y轴对齐由全局 .ant-btn flex 基座负责
  // （framework/styles/antd/index.css），此处不再手写 svg 对齐类
  actionButtonsReverse: true,
  submitButtonOptions: {
    class: 'whitespace-nowrap [&>svg]:size-3.5',
    content: $t('common.search'),
    icon: h(Search),
  },
  resetButtonOptions: {
    class: 'whitespace-nowrap [&>svg]:size-3.5',
    icon: h(LucideRefreshCw),
    type: 'text',
  },
  wrapperClass: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3',
});

const showTableTitle = computed(() => {
  return !!slots.tableTitle?.() || tableTitle.value;
});

const showToolbar = computed(() => {
  return (
    !!slots[TOOLBAR_ACTIONS]?.() ||
    !!slots[TOOLBAR_TOOLS]?.() ||
    showTableTitle.value
  );
});

const toolbarOptions = computed(() => {
  const slotActions = slots[TOOLBAR_ACTIONS]?.();
  const slotTools = slots[TOOLBAR_TOOLS]?.();

  const toolbarConfig: VxeGridPropTypes.ToolbarConfig = {
    // 内建工具钮（下载/刷新/全屏/列设置）全部禁用：由本组件模板以 vben 化按钮
    // （VbenIconButton + VbenTooltip 样式化提示）在 toolbar-tools 插槽前部渲染，
    // 行为经 commitProxy 与原内建钮一致
    refresh: false,
    export: false,
    zoom: false,
    custom: false,
  };

  if (!showToolbar.value) {
    return { toolbarConfig };
  }

  // if (gridOptions.value?.toolbarConfig?.search) {
  // }
  // 强制使用固定的toolbar配置，不允许用户自定义
  // 减少配置的复杂度，以及后续维护的成本
  toolbarConfig.slots = {
    ...(slotActions || showTableTitle.value
      ? { buttons: TOOLBAR_ACTIONS }
      : {}),
    ...(slotTools ? { tools: TOOLBAR_TOOLS } : {}),
  };

  return { toolbarConfig };
});

const options = computed(() => {
  const globalGridConfig = VxeUI?.getConfig()?.grid ?? {};

  const mergedOptions: VxeTableGridProps = cloneDeep(
    mergeWithArrayOverride(
      {},
      toolbarOptions.value,
      toRaw(gridOptions.value),
      globalGridConfig,
    ),
  );

  if (mergedOptions.proxyConfig) {
    const { ajax } = mergedOptions.proxyConfig;
    mergedOptions.proxyConfig.enabled = !!ajax;
    // 不自动加载数据, 由组件控制
    mergedOptions.proxyConfig.autoLoad = false;
  }

  if (mergedOptions.pagerConfig) {
    const mobileLayouts = [
      'PrevJump',
      'PrevPage',
      'Number',
      'NextPage',
      'NextJump',
    ] as any;
    const layouts = [
      'Total',
      'Sizes',
      'Home',
      ...mobileLayouts,
      'End',
    ] as readonly string[];
    mergedOptions.pagerConfig = mergeWithArrayOverride(
      {},
      mergedOptions.pagerConfig,
      {
        pageSize: 20,
        background: true,
        pageSizes: [10, 20, 30, 50, 100, 200],
        className: 'mt-2 w-full',
        layouts: isMobile.value ? mobileLayouts : layouts,
        size: 'mini' as const,
      },
    );
  }
  if (mergedOptions.formConfig) {
    mergedOptions.formConfig.enabled = false;
  }
  return mergedOptions;
});

function onToolbarToolClick(event: VxeGridDefines.ToolbarToolClickEventParams) {
  if (event.code === 'search') {
    props.api?.toggleSearchForm?.();
  }
  (
    gridEvents.value?.toolbarToolClick as VxeGridListeners['toolbarToolClick']
  )?.(event);
}

const events = computed(() => {
  return {
    ...gridEvents.value,
    toolbarToolClick: onToolbarToolClick,
  };
});

const delegatedSlots = computed(() => {
  const resultSlots: string[] = [];

  for (const key of Object.keys(slots)) {
    if (!['empty', 'form', 'loading', TOOLBAR_ACTIONS].includes(key)) {
      resultSlots.push(key);
    }
  }
  return resultSlots;
});

const delegatedFormSlots = computed(() => {
  const resultSlots: string[] = [];

  for (const key of Object.keys(slots)) {
    if (key.startsWith(FORM_SLOT_PREFIX)) {
      resultSlots.push(key);
    }
  }
  return resultSlots.map((key) => key.replace(FORM_SLOT_PREFIX, ''));
});

async function init() {
  await nextTick();
  const globalGridConfig = VxeUI?.getConfig()?.grid ?? {};
  const defaultGridOptions: VxeTableGridProps = mergeWithArrayOverride(
    {},
    toRaw(gridOptions.value),
    toRaw(globalGridConfig),
  );
  // 内部主动加载数据，防止form的默认值影响
  const autoLoad = defaultGridOptions.proxyConfig?.autoLoad;
  const enableProxyConfig = options.value.proxyConfig?.enabled;
  if (enableProxyConfig && autoLoad) {
    props.api.reload(formApi.form?.values ?? {});
  }

  // form 由 vben-form代替，所以不适配formConfig，这里给出警告
  const formConfig = gridOptions.value?.formConfig;
  // 处理某个页面加载多个Table时，第2个之后的Table初始化报出警告
  // 因为第一次初始化之后会把defaultGridOptions和gridOptions合并后缓存进State
  if (formConfig && formConfig.enabled) {
    console.warn(
      '[Vben Vxe Table]: The formConfig in the grid is not supported, please use the `formOptions` props',
    );
  }
  props.api?.setState?.({ gridOptions: defaultGridOptions });
  // form 由 vben-form 代替，所以需要保证query相关事件可以拿到参数
  extendProxyOptions(props.api, defaultGridOptions, () =>
    formApi.getLatestSubmissionValues(),
  );
}

// formOptions支持响应式
watch(
  formOptions,
  () => {
    formApi.setState((prev) => {
      const finalFormOptions: VbenFormProps = mergeWithArrayOverride(
        {},
        formOptions.value,
        prev,
      );
      return {
        ...finalFormOptions,
        collapseTriggerResize: !!finalFormOptions.showCollapseButton,
      };
    });
  },
  {
    immediate: true,
  },
);

onMounted(() => {
  props.api?.mount?.(gridRef.value, formApi);
  init();
});

onUnmounted(() => {
  formApi?.unmount?.();
  props.api?.unmount?.();
});
</script>

<template>
  <div :class="cn('bg-card h-full rounded-md', className)">
    <VxeGrid
      ref="gridRef"
      :class="
        cn(
          'p-2 pt-0',
          {
            'pt-0': showToolbar && !formOptions,
          },
          gridClass,
        )
      "
      v-bind="options"
      v-on="events"
    >
      <!-- 左侧操作区域或者title -->
      <template v-if="showToolbar" #toolbar-actions="slotProps">
        <slot v-if="showTableTitle" name="table-title">
          <div class="mr-1 pl-1 text-[1rem]">
            {{ tableTitle }}
            <VbenHelpTooltip v-if="tableTitleHelp" trigger-class="pb-1">
              {{ tableTitleHelp }}
            </VbenHelpTooltip>
          </div>
        </slot>
        <slot name="toolbar-actions" v-bind="slotProps"> </slot>
      </template>

      <!-- 继承默认的slot -->
      <template
        v-for="slotName in delegatedSlots"
        :key="slotName"
        #[slotName]="slotProps"
      >
        <slot :name="slotName" v-bind="slotProps"></slot>
        <!-- 默认工具（下载/刷新/全屏/列设置）：vben 化按钮 + VbenTooltip 样式化提示，
             替代 vxe 内建钮的原生 title；点击经 commitProxy 与原内建行为一致 -->
        <div
          v-if="slotName === TOOLBAR_TOOLS"
          class="mr-1 flex items-center gap-x-1"
        >
          <VbenTooltip
            v-for="tool in defaultToolbarTools"
            :key="tool.code"
            side="bottom"
          >
            <template #trigger>
              <VbenIconButton @click="tool.onClick()">
                <component :is="tool.icon" class="size-4" />
              </VbenIconButton>
            </template>
            {{ tool.label }}
          </VbenTooltip>
        </div>
      </template>

      <!-- form表单 -->
      <template #form>
        <!-- px-1.5(6px)：把搜索表单与表格内容对齐到同一条竖线。网格卡只有 p-2(8px) 一层内缩，
             而 vxe 的表头/单元格文字还额外吃 --vxe-ui-table-cell-padding-small(6px)
             （size:'small'，src/adapter/vxe-table.ts:33），实测表单字段在 8px、单元格文字在 14px
             ——盒子同源、内容错开 6px。下方 pb-4(16px) 保留：与 §2.6「同级块之间 16px」同值。 -->
        <div
          v-if="formOptions"
          v-show="showSearchForm !== false"
          class="relative rounded py-3 px-1.5 pb-4"
        >
          <slot name="form">
            <Form>
              <template
                v-for="slotName in delegatedFormSlots"
                :key="slotName"
                #[slotName]="slotProps"
              >
                <slot
                  :name="`${FORM_SLOT_PREFIX}${slotName}`"
                  v-bind="slotProps"
                ></slot>
              </template>
              <template #reset-before="slotProps">
                <slot name="reset-before" v-bind="slotProps"></slot>
              </template>
              <template #submit-before="slotProps">
                <slot name="submit-before" v-bind="slotProps"></slot>
              </template>
              <template #expand-before="slotProps">
                <slot name="expand-before" v-bind="slotProps"></slot>
              </template>
              <template #expand-after="slotProps">
                <slot name="expand-after" v-bind="slotProps"></slot>
              </template>
            </Form>
          </slot>
          <!-- overflow-hidden 不能留：.gowind-grid-form-gap 的 ::before/::after
               圆角补块画在色带框外（上/下缘外侧各 6px），留着会被整体裁掉；
               空色带本身也没有可裁内容。补块画法见 style.css 同名类。 -->
          <div
            class="gowind-grid-form-gap bg-background-deep z-100 absolute -left-2 bottom-1 h-2 w-[calc(100%+1rem)] md:bottom-2 md:h-3"
          ></div>
        </div>
      </template>
      <!-- loading -->
      <template #loading>
        <slot name="loading">
          <VbenLoading :spinning="true" />
        </slot>
      </template>
      <!-- 统一控状态 -->
      <template #empty>
        <slot name="empty">
          <EmptyIcon class="mx-auto" />
          <div class="mt-2">{{ $t('common.noData') }}</div>
        </slot>
      </template>
    </VxeGrid>
  </div>
</template>
