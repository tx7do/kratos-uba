<script lang="ts" setup>
import { Maximize } from '@vben-core/icons';

import { useFullscreen } from '@vueuse/core';
import { $t } from '@vben/locales';

import { VbenIconButton } from '../button';
import { VbenTooltip } from '../tooltip';

defineOptions({ name: 'FullScreen' });

const { isFullscreen, toggle } = useFullscreen();

// 重新检查全屏状态
isFullscreen.value = !!(
  document.fullscreenElement ||
  // @ts-ignore
  document.webkitFullscreenElement ||
  // @ts-ignore
  document.mozFullScreenElement ||
  // @ts-ignore
  document.msFullscreenElement
);
</script>
<template>
  <VbenTooltip side="bottom">
    <template #trigger>
      <VbenIconButton @click="toggle">
        <!-- 退出全屏：与 react 顶栏 MinimizeBracketsIcon 同款路径（手抄混合版，
             下半两角外肘贴边）；不用 lucide Minimize——其下半两角是内肘，与 react 顶栏不一致 -->
        <svg
          v-if="isFullscreen"
          class="text-foreground size-4"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M8 3v3a2 2 0 0 1-2 2H3" />
          <path d="M21 8h-3a2 2 0 0 1-2-2V3" />
          <path d="M3 16v3a2 2 0 0 0 2 2h3" />
          <path d="M16 21h3a2 2 0 0 0 2-2v-3" />
        </svg>
        <Maximize v-else class="text-foreground size-4" />
      </VbenIconButton>
    </template>
    {{ isFullscreen ? $t('common.exitFullscreen') : $t('common.fullscreen') }}
  </VbenTooltip>
</template>
