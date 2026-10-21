<script lang="ts" setup>
import type { NotificationItem } from './types';

import { Bell, Check } from '@vben/icons';
import { $t } from '@vben/locales';
import {
  VbenButton,
  VbenIconButton,
  VbenPopover,
  VbenTooltip,
  VbenScrollbar,
} from '@vben-core/shadcn-ui';

import { useToggle } from '@vueuse/core';

interface Props {
  /**
   * 显示圆点
   */
  dot?: boolean;
  /**
   * 消息列表（最近 5 条，已读+未读混合）
   */
  notifications?: NotificationItem[];
  /**
   * 未读数（红底数字角标；>0 时显示，封顶 99+）
   */
  count?: number;
  /**
   * 「全部已读」请求进行中
   */
  makeAllLoading?: boolean;
}

defineOptions({ name: 'NotificationPopup' });

withDefaults(defineProps<Props>(), {
  count: 0,
  dot: false,
  makeAllLoading: false,
  notifications: () => [],
});

const emit = defineEmits<{
  makeAll: [];
  viewAll: [];
}>();

const [open, toggle] = useToggle();

function close() {
  open.value = false;
}

function handleMakeAll() {
  emit('makeAll');
}

// 点击条目与底部链接同义：进入收件箱页（对齐 react 端行为）
function handleClick() {
  emit('viewAll');
  close();
}
</script>
<template>
  <VbenPopover
    v-model:open="open"
    content-class="relative right-2 w-[336px] p-0"
  >
    <template #trigger>
      <div class="flex-center mr-2 h-full" @click.stop="toggle()">
        <VbenTooltip side="bottom">
          <template #trigger>
            <VbenIconButton class="bell-button text-foreground relative">
              <span
                v-if="dot"
                class="bg-primary absolute right-0.5 top-0.5 h-2 w-2 rounded"
              ></span>
              <span
                v-if="count > 0"
                class="bg-destructive text-destructive-foreground absolute -right-1 -top-1 h-4 min-w-4 rounded-full px-1 text-[10px] font-semibold leading-4"
              >
                {{ count > 99 ? '99+' : count }}
              </span>
              <Bell class="size-4" />
            </VbenIconButton>
          </template>
          {{ $t('ui.widgets.notifications') }}
        </VbenTooltip>
      </div>
    </template>

    <!-- 面板对齐 react 端 notification-popover：336 宽 + 10/10/6 内边距 -->
    <div class="relative p-2.5 pb-1.5">
      <!-- 头部：标题 + 未读数 + 全部已读 -->
      <div
        class="border-border mb-1 flex items-center justify-between border-b pb-2.5 pt-0.5"
      >
        <div class="text-foreground flex items-center gap-2 text-sm font-semibold">
          {{ $t('ui.widgets.inboxTitle') }}
          <span
            v-if="count > 0"
            class="bg-primary/10 text-primary rounded-full px-2 text-xs font-normal leading-[18px]"
          >
            {{ $t('ui.widgets.unreadCount', { count }) }}
          </span>
        </div>
        <VbenButton
          :disabled="count <= 0"
          :loading="makeAllLoading"
          class="text-primary shrink-0"
          size="sm"
          variant="ghost"
          @click="handleMakeAll"
        >
          <Check class="size-3.5" />
          {{ $t('ui.widgets.markAllAsRead') }}
        </VbenButton>
      </div>

      <!-- 消息列表：未读带圆点加粗 -->
      <VbenScrollbar v-if="notifications.length > 0">
        <div class="flex w-full flex-col">
          <div
            v-for="item in notifications"
            :key="item.id"
            class="hover:bg-accent flex w-full cursor-pointer items-center gap-2.5 rounded-lg px-2 py-[9px]"
            @click="handleClick()"
          >
            <span
              class="mt-1.5 size-1.5 shrink-0 rounded-full"
              :class="item.isRead ? 'bg-transparent' : 'bg-primary'"
            ></span>
            <div class="min-w-0 flex-1">
              <p
                class="truncate text-[13px] leading-5"
                :class="
                  item.isRead
                    ? 'text-muted-foreground'
                    : 'text-foreground font-semibold'
                "
              >
                {{ item.title || '-' }}
              </p>
              <p
                v-if="item.message"
                class="text-muted-foreground mt-0.5 truncate text-xs"
              >
                {{ item.message }}
              </p>
            </div>
            <span class="text-muted-foreground shrink-0 text-xs">
              {{ item.date }}
            </span>
          </div>
        </div>
      </VbenScrollbar>

      <div
        v-else
        class="flex-center text-muted-foreground min-h-[150px] w-full text-xs"
      >
        {{ $t('ui.widgets.noMessages') }}
      </div>

      <!-- 底部：进入收件箱 -->
      <div class="border-border mt-1 border-t pt-2 text-center">
        <VbenButton
          class="text-primary"
          size="sm"
          variant="ghost"
          @click="handleClick()"
        >
          {{ $t('ui.widgets.viewAll') }}
        </VbenButton>
      </div>
    </div>
  </VbenPopover>
</template>

<style scoped>
:deep(.bell-button) {
  &:hover {
    svg {
      animation: bell-ring 1s both;
    }
  }
}

@keyframes bell-ring {
  0%,
  100% {
    transform-origin: top;
  }

  15% {
    transform: rotateZ(10deg);
  }

  30% {
    transform: rotateZ(-10deg);
  }

  45% {
    transform: rotateZ(5deg);
  }

  60% {
    transform: rotateZ(-5deg);
  }

  75% {
    transform: rotateZ(2deg);
  }
}
</style>
