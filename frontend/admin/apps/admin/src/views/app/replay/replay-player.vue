<script lang="ts" setup>
import { onBeforeUnmount, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';

import { $t } from '#/locales';
import { fetchReplaySessionEvents } from '#/api';

const sessionId = ref('');
const loading = ref(false);
const containerRef = ref<HTMLDivElement>();
const playerInstance = ref<any>(null);

async function load(sessionIdValue: string) {
  loading.value = true;
  destroyPlayer();
  try {
    const resp = await fetchReplaySessionEvents(sessionIdValue);
    let events: unknown[] = [];
    try {
      events = JSON.parse(resp.eventsJson || '[]');
    } catch {
      events = [];
    }

    // rrweb-player 动态导入（按需加载，构建器自动 code-split）
    const [{ default: RrwebPlayer }, cssModule] = await Promise.all([
      import('rrweb-player'),
      import('rrweb-player/dist/style.css'),
    ]);
    void cssModule;

    if (containerRef.value) {
      containerRef.value.innerHTML = '';
      playerInstance.value = new RrwebPlayer({
        target: containerRef.value,
        props: {
          events: events as any,
          showController: true,
          autoPlay: false,
        } as any,
      });
    }
  } finally {
    loading.value = false;
  }
}

function destroyPlayer() {
  playerInstance.value?.destroy();
  playerInstance.value = null;
  if (containerRef.value) {
    containerRef.value.innerHTML = '';
  }
}

const [Drawer, drawerApi] = useVbenDrawer({
  onOpenChange(isOpen) {
    if (isOpen) {
      const data = drawerApi.getData<Record<string, any>>();
      if (data?.sessionId) {
        sessionId.value = data.sessionId;
        load(data.sessionId);
      }
    } else {
      destroyPlayer();
    }
  },
});

onBeforeUnmount(destroyPlayer);
</script>

<template>
  <Drawer :title="$t('page.replay.playerTitle')" class="w-[820px]">
    <a-spin :spinning="loading">
      <p v-if="sessionId" class="text-muted-foreground mb-2 text-xs">
        {{ $t('page.replay.sessionLabel') }}:
        <span class="font-mono">{{ sessionId }}</span>
      </p>
      <div
        ref="containerRef"
        :style="{
          width: '100%',
          maxHeight: '60vh',
          overflow: 'hidden',
          borderRadius: '4px',
          border: '1px solid #e5e7eb',
        }"
      />
    </a-spin>
  </Drawer>
</template>
