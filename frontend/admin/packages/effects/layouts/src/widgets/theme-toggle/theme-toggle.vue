<script lang="ts" setup>
import { $t } from '@vben/locales';
import { usePreferences } from '@vben/preferences';
import { VbenTooltip } from '@vben-core/shadcn-ui';
import { updatePreferences } from '@vben/preferences';

import ThemeButton from './theme-button.vue';

defineOptions({
  name: 'ThemeToggle',
});

withDefaults(defineProps<{ shouldOnHover?: boolean }>(), {
  shouldOnHover: false,
});

const { isDark } = usePreferences();

function handleChange(isDark: boolean) {
  updatePreferences({
    theme: { mode: isDark ? 'dark' : 'light' },
  });
}
</script>
<template>
  <div>
    <VbenTooltip :disabled="false" side="bottom">
      <template #trigger>
        <ThemeButton
          :model-value="isDark"
          type="icon"
          @update:model-value="handleChange"
        />
      </template>
      {{ $t('preferences.theme.title') }}
    </VbenTooltip>
  </div>
</template>
