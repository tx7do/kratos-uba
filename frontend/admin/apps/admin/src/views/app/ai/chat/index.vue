<script lang="ts" setup>
import type { AiMessage } from '#/api';
import type {
  aiservicev1_AiRole,
  aiservicev1_ChatCitation,
} from '#/generated/api/admin/service/v1';

import { computed, h, nextTick, onBeforeUnmount, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import {
  LucidePencil,
  LucidePlus,
  LucideSend,
  LucideTrash2,
  LucideWrench,
} from '@vben/icons';

import { message, notification, Popconfirm, Tag } from 'ant-design-vue';
import { marked } from 'marked';

import {
  fetchListAiConversations,
  fetchListAiKnowledgeBases,
  fetchListAiMessages,
  fetchListAiProviders,
  useDeleteAiConversation,
  useSendChat,
  useUpdateAiConversation,
} from '#/api';
import { $t } from '#/locales';
import { PaginationQuery } from '#/transport/rest';
import { globalSSEClient } from '#/transport/sse';

defineOptions({ name: 'AiChat' });

interface ChatItem {
  id?: number;
  role: aiservicev1_AiRole;
  content: string;
  modelName?: string;
  promptTokens?: number;
  completionTokens?: number;
  durationMs?: number;
  errorMessage?: string;
  streaming?: boolean;
  citations?: aiservicev1_ChatCitation[];
  toolEvents?: Array<{ arguments: string; name: string; result: string }>;
}

const conversations = ref<Array<{ id: number; title?: string }>>([]);
const activeConversationId = ref<null | number>(null);
const messages = ref<ChatItem[]>([]);
const loadingMessages = ref(false);

// 历史消息分页：后端按 -created_at（最新在前）返回，页面内反转成时间升序展示；
// 「加载更早消息」追加下一页并前插。
const messagesRef = ref<HTMLElement>();
const CHAT_PAGE_SIZE = 50;
const messagesPage = ref(1);
const hasMoreMessages = ref(false);
const loadingOlder = ref(false);
const sending = ref(false);
const input = ref('');

// 供应商 / 知识库选择（0=默认 / 不注入）
const providers = ref<Array<{ id: number; modelName?: string; name?: string }>>(
  [],
);
const knowledgeBases = ref<Array<{ id: number; name?: string }>>([]);
const selectedProviderId = ref<number | undefined>();
const selectedKnowledgeBaseId = ref<number | undefined>();

// 流式状态：chunk 经 SSE 到达时挂到最后一条 streaming 消息
const streamingConvId = ref<null | number>(null);

// 会话重命名（行内编辑）
const editingId = ref<null | number>(null);
const editingTitle = ref('');

const { mutateAsync: sendChat } = useSendChat();
const { mutateAsync: updateConversation } = useUpdateAiConversation();
const { mutateAsync: deleteConversation } = useDeleteAiConversation();

const canSend = computed(() => !sending.value && input.value.trim().length > 0);

onMounted(async () => {
  globalSSEClient.on('ai_chat_chunk', handleChunkEvent);
  globalSSEClient.on('ai_chat_tool', handleToolEvent);
  await Promise.all([reloadConversations(), reloadOptions()]);
});

onBeforeUnmount(() => {
  globalSSEClient.off('ai_chat_chunk', handleChunkEvent);
  globalSSEClient.off('ai_chat_tool', handleToolEvent);
});

// ── SSE 增量 ───────────────────────────────────────────────

function streamingItem(): ChatItem | undefined {
  return [...messages.value].reverse().find((m) => m.streaming);
}

function handleChunkEvent(data: unknown) {
  const chunk = data as { conversationId?: number; delta?: string };
  if (
    streamingConvId.value !== null &&
    chunk.conversationId !== streamingConvId.value
  ) {
    return; // 不是当前会话的流（其他标签页发的）
  }
  if (streamingConvId.value === null && chunk.conversationId) {
    streamingConvId.value = chunk.conversationId;
  }
  const item = streamingItem();
  if (item && chunk.delta) {
    item.content += chunk.delta;
    void scrollToBottom();
  }
}

function handleToolEvent(data: unknown) {
  const tool = data as {
    arguments?: string;
    conversationId?: number;
    name?: string;
    result?: string;
  };
  if (
    streamingConvId.value !== null &&
    tool.conversationId !== streamingConvId.value
  ) {
    return;
  }
  const item = streamingItem();
  if (item) {
    item.toolEvents = item.toolEvents ?? [];
    item.toolEvents.push({
      name: tool.name ?? '',
      arguments: tool.arguments ?? '',
      result: tool.result ?? '',
    });
    void scrollToBottom();
  }
}

// ── 会话列表 ───────────────────────────────────────────────

async function reloadConversations() {
  try {
    const resp = await fetchListAiConversations(
      new PaginationQuery({ paging: { pageSize: 100 } }),
    );
    conversations.value = (resp.items ?? []).map((c) => ({
      id: c.id as number,
      title: c.title ?? '',
    }));
    if (
      activeConversationId.value !== null &&
      !conversations.value.some((c) => c.id === activeConversationId.value)
    ) {
      activeConversationId.value = null;
      messages.value = [];
    }
  } catch {
    // 列表加载失败不阻断新对话
  }
}

async function reloadOptions() {
  try {
    const [providerResp, kbResp] = await Promise.all([
      fetchListAiProviders(new PaginationQuery({ paging: { pageSize: 100 } })),
      fetchListAiKnowledgeBases(
        new PaginationQuery({ paging: { pageSize: 100 } }),
      ),
    ]);
    providers.value = (providerResp.items ?? []).map((p) => ({
      id: p.id as number,
      name: p.name ?? '',
      modelName: p.modelName ?? '',
    }));
    knowledgeBases.value = (kbResp.items ?? []).map((k) => ({
      id: k.id as number,
      name: k.name ?? '',
    }));
  } catch {
    // 选项加载失败：后端仍有默认提供商兜底
  }
}

async function selectConversation(id: number) {
  if (sending.value) return;
  activeConversationId.value = id;
  loadingMessages.value = true;
  messagesPage.value = 1;
  try {
    const resp = await fetchListAiMessages(
      new PaginationQuery({
        paging: { page: 1, pageSize: CHAT_PAGE_SIZE },
        formValues: { conversationId: id },
      }),
    );
    const items = resp.items ?? [];
    hasMoreMessages.value = items.length >= CHAT_PAGE_SIZE;
    // 最新在前 → 反转为时间升序展示
    messages.value = items.reverse().map((m) => entityToItem(m));
    await scrollToBottom();
  } finally {
    loadingMessages.value = false;
  }
}

/** 加载更早的一页历史（前插到当前消息上方）。 */
async function loadOlderMessages() {
  if (
    !activeConversationId.value ||
    loadingOlder.value ||
    !hasMoreMessages.value
  ) {
    return;
  }
  loadingOlder.value = true;
  const el = messagesRef.value;
  const prevHeight = el?.scrollHeight ?? 0;
  try {
    const nextPage = messagesPage.value + 1;
    const resp = await fetchListAiMessages(
      new PaginationQuery({
        paging: { page: nextPage, pageSize: CHAT_PAGE_SIZE },
        formValues: { conversationId: activeConversationId.value },
      }),
    );
    const items = resp.items ?? [];
    hasMoreMessages.value = items.length >= CHAT_PAGE_SIZE;
    messagesPage.value = nextPage;
    const older = items.reverse().map((m) => entityToItem(m));
    messages.value = [...older, ...messages.value];
    await nextTick();
    // 保持视口停留在原内容位置，避免前插后跳动
    if (el) {
      el.scrollTop = el.scrollHeight - prevHeight;
    }
  } finally {
    loadingOlder.value = false;
  }
}

function entityToItem(m: AiMessage): ChatItem {
  return {
    id: m.id,
    role: m.role ?? 'USER',
    content: m.content ?? '',
    modelName: m.modelName ?? undefined,
    promptTokens: m.promptTokens ?? undefined,
    completionTokens: m.completionTokens ?? undefined,
    durationMs: m.durationMs ?? undefined,
    errorMessage: m.errorMessage ?? undefined,
  };
}

function handleNewConversation() {
  if (sending.value) return;
  activeConversationId.value = null;
  streamingConvId.value = null;
  messages.value = [];
  input.value = '';
}

// ── 发送 ───────────────────────────────────────────────────

async function handleSend() {
  const content = input.value.trim();
  if (!content || sending.value) return;

  sending.value = true;
  input.value = '';
  streamingConvId.value = activeConversationId.value;

  messages.value.push({ role: 'USER', content });
  const placeholder: ChatItem = {
    role: 'ASSISTANT',
    content: '',
    streaming: true,
  };
  messages.value.push(placeholder);
  void scrollToBottom();

  try {
    const resp = await sendChat({
      content,
      conversationId: Number(activeConversationId.value ?? 0),
      providerId: Number(selectedProviderId.value ?? 0),
      knowledgeBaseId: Number(selectedKnowledgeBaseId.value ?? 0),
    });

    // done 帧为准：替换流式占位（SSE 断线时这是唯一回复来源）
    const done = resp.message;
    placeholder.streaming = false;
    placeholder.content = done?.content ?? placeholder.content;
    placeholder.id = done?.id;
    placeholder.modelName = done?.modelName ?? undefined;
    placeholder.promptTokens = done?.promptTokens ?? undefined;
    placeholder.completionTokens = done?.completionTokens ?? undefined;
    placeholder.durationMs = done?.durationMs ?? undefined;
    if (done?.errorMessage) {
      placeholder.errorMessage = done.errorMessage;
    }
    placeholder.citations = resp.citations;

    const convId = resp.conversation?.id;
    if (convId && convId !== activeConversationId.value) {
      activeConversationId.value = convId;
      await reloadConversations();
    } else {
      await reloadConversations();
    }
  } catch (error: any) {
    placeholder.streaming = false;
    placeholder.errorMessage = error?.response?.data?.message || String(error);
    notification.error({ message: $t('page.ai.chat.sendFailed') });
  } finally {
    sending.value = false;
    streamingConvId.value = null;
    await scrollToBottom();
  }
}

// ── 重命名 / 删除 ──────────────────────────────────────────

function startRename(conv: { id: number; title?: string }) {
  editingId.value = conv.id;
  editingTitle.value = conv.title ?? '';
}

async function saveRename() {
  if (editingId.value === null) return;
  const id = editingId.value;
  const title = editingTitle.value.trim();
  editingId.value = null;
  if (!title) return;
  try {
    await updateConversation({ id, data: { title } as never });
    await reloadConversations();
  } catch {
    message.error($t('page.ai.chat.renameFailed'));
  }
}

async function handleDeleteConversation(conv: { id: number }) {
  try {
    await deleteConversation({ id: conv.id });
    if (activeConversationId.value === conv.id) {
      activeConversationId.value = null;
      messages.value = [];
    }
    await reloadConversations();
  } catch {
    message.error($t('ui.notification.delete_failed'));
  }
}

// ── 渲染辅助 ───────────────────────────────────────────────

const rendered = (content: string) => marked.parse(content, { async: false });

async function scrollToBottom() {
  await nextTick();
  const el = messagesRef.value;
  if (el) {
    el.scrollTop = el.scrollHeight;
  }
}
</script>

<template>
  <Page :title="$t('menu.ai.chat')">
    <div class="flex h-[calc(100vh-160px)] gap-3">
      <!-- 会话列表 -->
      <div class="bg-card flex w-60 flex-col rounded-lg border">
        <div class="p-2">
          <a-button type="primary" block @click="handleNewConversation">
            <template #icon>
              <LucidePlus />
            </template>
            {{ $t('page.ai.chat.newConversation') }}
          </a-button>
        </div>
        <div class="flex-1 overflow-y-auto px-2 pb-2">
          <div
            v-for="conv in conversations"
            :key="conv.id"
            class="hover:bg-accent group mb-1 flex cursor-pointer items-center gap-1 rounded px-2 py-2 text-sm"
            :class="{
              'bg-accent': conv.id === activeConversationId,
            }"
            @click="selectConversation(conv.id)"
          >
            <template v-if="editingId === conv.id">
              <input
                v-model="editingTitle"
                class="w-full rounded border px-1 text-sm"
                @click.stop
                @keyup.enter="saveRename"
                @keyup.esc="editingId = null"
                @blur="saveRename"
              />
            </template>
            <template v-else>
              <span class="flex-1 truncate" :title="conv.title">
                {{ conv.title || $t('page.ai.chat.newConversation') }}
              </span>
              <span class="hidden shrink-0 group-hover:inline-flex">
                <a-button
                  type="link"
                  size="small"
                  :icon="h(LucidePencil)"
                  @click.stop="startRename(conv)"
                />
                <Popconfirm
                  :cancel-text="$t('ui.button.cancel')"
                  :ok-text="$t('ui.button.ok')"
                  :title="
                    $t('ui.text.do_you_want_delete', {
                      moduleName: conv.title || '',
                    })
                  "
                  @confirm="handleDeleteConversation(conv)"
                >
                  <a-button
                    danger
                    type="link"
                    size="small"
                    :icon="h(LucideTrash2)"
                    @click.stop
                  />
                </Popconfirm>
              </span>
            </template>
          </div>
        </div>
      </div>

      <!-- 消息区 -->
      <div class="bg-card flex flex-1 flex-col rounded-lg border">
        <div ref="messagesRef" class="flex-1 overflow-y-auto p-4">
          <div v-if="hasMoreMessages" class="mb-3 text-center">
            <a-button
              size="small"
              type="link"
              :loading="loadingOlder"
              @click="loadOlderMessages"
            >
              {{ $t('page.ai.chat.loadOlder') }}
            </a-button>
          </div>
          <div
            v-if="messages.length === 0 && !loadingMessages"
            class="flex h-full items-center justify-center text-sm opacity-50"
          >
            {{ $t('page.ai.chat.empty') }}
          </div>
          <div
            v-for="(item, index) in messages"
            :key="item.id ?? `local-${index}`"
            class="mb-4 flex gap-2"
            :class="item.role === 'USER' ? 'flex-row-reverse' : ''"
          >
            <div
              class="flex size-8 shrink-0 items-center justify-center rounded-full text-xs font-medium"
              :class="
                item.role === 'USER'
                  ? 'bg-primary text-primary-foreground'
                  : 'bg-muted text-muted-foreground'
              "
            >
              {{
                item.role === 'USER'
                  ? $t('page.ai.chat.roleUser')
                  : $t('page.ai.chat.roleAssistant')
              }}
            </div>
            <div class="max-w-[78%]">
              <!-- 工具调用可见化 -->
              <div
                v-if="item.toolEvents && item.toolEvents.length > 0"
                class="mb-1 flex flex-col gap-1"
              >
                <div
                  v-for="(tool, ti) in item.toolEvents"
                  :key="ti"
                  class="rounded border px-2 py-1 text-xs opacity-80"
                >
                  <Tag color="purple">
                    <LucideWrench class="mr-1 inline size-3" />
                    {{ tool.name }}
                  </Tag>
                  <span v-if="tool.arguments" class="opacity-60">
                    {{ tool.arguments }}
                  </span>
                </div>
              </div>
              <div
                class="rounded-lg px-3 py-2 text-sm"
                :class="
                  item.role === 'USER'
                    ? 'bg-primary text-primary-foreground whitespace-pre-wrap'
                    : 'bg-muted'
                "
              >
                <span
                  v-if="item.role === 'ASSISTANT'"
                  v-html="rendered(item.content)"
                ></span>
                <template v-else>{{ item.content }}</template>
                <span
                  v-if="item.streaming && !item.content"
                  class="ml-1 inline-block animate-pulse"
                >
                  …
                </span>
              </div>
              <div v-if="item.errorMessage" class="mt-1 text-xs text-red-500">
                {{ item.errorMessage }}
              </div>
              <!-- RAG 引用来源 -->
              <div
                v-if="
                  item.citations && item.citations.length > 0 && !item.streaming
                "
                class="mt-1 rounded border px-2 py-1 text-xs opacity-80"
              >
                <div class="mb-1 font-medium">
                  {{ $t('page.ai.chat.citations') }}
                </div>
                <div
                  v-for="(cite, ci) in item.citations"
                  :key="ci"
                  class="truncate"
                  :title="cite.snippet"
                >
                  [{{ ci + 1 }}] {{ cite.docName }}#{{ cite.chunkIndex }}
                  <span class="opacity-50">
                    ({{ (cite.score ?? 0).toFixed(3) }})
                  </span>
                </div>
              </div>
              <div
                v-if="
                  item.role === 'ASSISTANT' && !item.streaming && item.modelName
                "
                class="mt-1 flex gap-2 text-xs opacity-50"
              >
                <span>{{ item.modelName }}</span>
                <span v-if="item.durationMs"> {{ item.durationMs }}ms </span>
                <span v-if="item.promptTokens">
                  {{ item.promptTokens }}+{{ item.completionTokens }} tokens
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- 输入区 -->
        <div class="border-t p-3">
          <div class="mb-2 flex flex-wrap gap-2">
            <a-select
              v-model:value="selectedProviderId"
              class="min-w-[220px]"
              allow-clear
              :placeholder="$t('page.ai.chat.defaultProvider')"
              :options="
                providers.map((p) => ({
                  label: `${p.name}（${p.modelName}）`,
                  value: p.id,
                }))
              "
            />
            <a-select
              v-model:value="selectedKnowledgeBaseId"
              class="min-w-[200px]"
              allow-clear
              :placeholder="$t('page.ai.chat.knowledgeBase')"
              :options="
                knowledgeBases.map((k) => ({
                  label: k.name,
                  value: k.id,
                }))
              "
            />
          </div>
          <div class="flex items-end gap-2">
            <a-textarea
              v-model:value="input"
              :placeholder="$t('page.ai.chat.inputPlaceholder')"
              :rows="2"
              :disabled="sending"
              @keydown.enter.exact.prevent="handleSend"
            />
            <a-button
              type="primary"
              :loading="sending"
              :disabled="!canSend"
              @click="handleSend"
            >
              <template #icon>
                <LucideSend />
              </template>
              {{ $t('page.ai.chat.send') }}
            </a-button>
          </div>
        </div>
      </div>
    </div>
  </Page>
</template>
