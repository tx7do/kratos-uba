<script lang="ts" setup>
import { LucidePlus } from '@vben/icons';

/**
 * 列表工具栏「新增」按钮（与 TableExportButton 配对）。
 * icon↔文本间距与 y 轴居中由全局 .ant-btn flex 基座统一负责
 * （framework/styles/antd/index.css），此处只保留工具栏间距与防换行；
 * 默认文案 ui.button.add、加号图标；@click、disabled 等属性透传内部
 * a-button；text 传最终文案（需自行翻译）。
 */
defineProps<{ text?: string }>();
</script>

<template>
  <a-button type="primary" ghost class="mr-2 whitespace-nowrap">
    <template #icon>
      <LucidePlus />
    </template>
    {{ text ?? $t('ui.button.add') }}
  </a-button>
</template>
