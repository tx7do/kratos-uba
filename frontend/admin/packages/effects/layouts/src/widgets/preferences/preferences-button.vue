<script lang="ts" setup>
import { SvgSettingIcon } from '@vben/icons';
import { $t } from '@vben/locales';
import { VbenIconButton, VbenTooltip } from '@vben-core/shadcn-ui';

import Preferences from './preferences.vue';

const emit = defineEmits<{ clearPreferencesAndLogout: [] }>();

function clearPreferencesAndLogout() {
  emit('clearPreferencesAndLogout');
}
</script>
<template>
  <Preferences @clear-preferences-and-logout="clearPreferencesAndLogout">
    <VbenTooltip side="bottom">
      <template #trigger>
        <VbenIconButton>
          <SvgSettingIcon class="text-foreground size-4" />
        </VbenIconButton>
      </template>
      {{ $t('preferences.title') }}
    </VbenTooltip>
  </Preferences>
</template>
