<script lang="ts" setup>
import { reactive } from 'vue';

import { ColPage } from '@vben/common-ui';

import OrgList from '#/views/app/opm/user/org-list.vue';
import UserList from '#/views/app/opm/user/user-list.vue';

const props = reactive({
  leftCollapsedWidth: 0,
  leftCollapsible: false,
  // min/max 反转会把拖拽范围锁死（min>max 恒不满足），左栏无法拉伸
  leftMaxWidth: 40,
  leftMinWidth: 20,
  leftWidth: 25,
  rightWidth: 75,
  resizable: true,
  splitHandle: true,
  splitLine: true,
});
</script>

<template>
  <ColPage auto-content-height v-bind="props">
    <template #left>
      <OrgList />
    </template>
    <UserList />
  </ColPage>
</template>
