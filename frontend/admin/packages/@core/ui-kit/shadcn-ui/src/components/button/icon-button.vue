<script setup lang="ts">
import type { ButtonVariants } from '../../ui';
import type { VbenButtonProps } from './button';

import { computed, useSlots } from 'vue';

import { cn } from '@vben-core/shared/utils';

import { VbenTooltip } from '../tooltip';
import VbenButton from './button.vue';

interface Props extends VbenButtonProps {
  class?: any;
  disabled?: boolean;
  // radix asChild（TooltipTrigger 等）会把自身 handler 与外部 @click 合并成数组下发
  onClick?: ((...args: any[]) => void) | ((...args: any[]) => void)[];
  tooltip?: string;
  tooltipDelayDuration?: number;
  tooltipSide?: 'bottom' | 'left' | 'right' | 'top';
  variant?: ButtonVariants;
}

const props = withDefaults(defineProps<Props>(), {
  disabled: false,
  onClick: () => {},
  tooltipDelayDuration: 200,
  tooltipSide: 'bottom',
  variant: 'icon',
});

const slots = useSlots();

const showTooltip = computed(() => !!slots.tooltip || !!props.tooltip);

// radix asChild 合并的 handler 数组逐个调用，语义同原生元素上的数组 onClick
function handleClick(e: MouseEvent) {
  if (Array.isArray(props.onClick)) {
    props.onClick.forEach((fn) => fn(e));
  } else {
    props.onClick(e);
  }
}
</script>

<template>
  <VbenButton
    v-if="!showTooltip"
    :class="cn('rounded-full', props.class)"
    :disabled="disabled"
    :variant="variant"
    size="icon"
    @click="handleClick"
  >
    <slot></slot>
  </VbenButton>

  <VbenTooltip
    v-else
    :delay-duration="tooltipDelayDuration"
    :side="tooltipSide"
  >
    <template #trigger>
      <VbenButton
        :class="cn('rounded-full', props.class)"
        :disabled="disabled"
        :variant="variant"
        size="icon"
        @click="handleClick"
      >
        <slot></slot>
      </VbenButton>
    </template>
    <slot v-if="slots.tooltip" name="tooltip"> </slot>
    <template v-else>
      {{ tooltip }}
    </template>
  </VbenTooltip>
</template>
