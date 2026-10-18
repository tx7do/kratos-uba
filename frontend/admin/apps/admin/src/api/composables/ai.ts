import type {
  aiservicev1_AiConversation,
  aiservicev1_AiDoc,
  aiservicev1_AiKnowledgeBase,
  aiservicev1_AiMessage,
  aiservicev1_AiProvider,
  aiservicev1_AskAiQueryRequest,
  aiservicev1_AskAiQueryResponse,
  aiservicev1_ChatRequest,
  aiservicev1_ChatResponse,
  aiservicev1_CreateAiProviderRequest,
  aiservicev1_DeleteAiConversationRequest,
  aiservicev1_DeleteAiDocRequest,
  aiservicev1_DeleteAiMessageRequest,
  aiservicev1_DeleteAiProviderRequest,
  aiservicev1_GetAiKnowledgeBaseRequest,
  aiservicev1_GetAiProviderRequest,
  aiservicev1_ListAiConversationResponse,
  aiservicev1_ListAiDocsRequest,
  aiservicev1_ListAiKnowledgeBaseResponse,
  aiservicev1_ListAiMessageResponse,
  aiservicev1_ListAiProviderResponse,
  aiservicev1_ListAiUsageLogResponse,
  aiservicev1_ReindexDocsRequest,
  aiservicev1_SyncOpsKnowledgeRequest,
  aiservicev1_SyncOpsKnowledgeResponse,
  aiservicev1_SearchAiKnowledgeRequest,
  aiservicev1_SearchAiKnowledgeResponse,
  aiservicev1_UpdateAiProviderRequest,
  aiservicev1_UploadAiDocFileRequest,
  aiservicev1_UploadAiDocRequest,
  aiservicev1_UploadAiDocResponse,
  aiservicev1_UsageSummaryResponse,
} from '#/generated/api/admin/service/v1';

import {
  useMutation,
  type UseMutationOptions,
  useQuery,
  type UseQueryOptions,
} from '@tanstack/vue-query';

import { apiClient } from '#/api/client';
import { queryClient } from '#/plugins/vue-query';
import { makeUpdateMask, type PaginationQuery } from '#/transport/rest';

// ==============================
// AI 模型提供商
// ==============================
export function useListAiProviders(
  query: PaginationQuery,
  options?: UseQueryOptions<aiservicev1_ListAiProviderResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAiProviders', query],
    queryFn: () => apiClient.aiProviderService.List(query.toRawParams()),
    ...options,
  });
}

export function useGetAiProvider(
  req: aiservicev1_GetAiProviderRequest,
  options?: UseQueryOptions<aiservicev1_AiProvider, Error>,
) {
  return useQuery({
    queryKey: ['getAiProvider', req],
    queryFn: () => apiClient.aiProviderService.Get(req),
    ...options,
  });
}

export async function fetchListAiProviders(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAiProviders', params],
    queryFn: () => apiClient.aiProviderService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCreateAiProvider(
  options?: UseMutationOptions<
    object,
    Error,
    aiservicev1_CreateAiProviderRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiProviderService.Create(req),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiProviders'] }),
    ...options,
  });
}

export function useUpdateAiProvider(
  options?: UseMutationOptions<
    object,
    Error,
    aiservicev1_UpdateAiProviderRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiProviderService.Update(req),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiProviders'] }),
    ...options,
  });
}

export function useDeleteAiProvider(
  options?: UseMutationOptions<object, Error, aiservicev1_DeleteAiProviderRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiProviderService.Delete(req),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiProviders'] }),
    ...options,
  });
}

// ==============================
// AI 对话会话 / 消息
// ==============================
export function useListAiConversations(
  query: PaginationQuery,
  options?: UseQueryOptions<aiservicev1_ListAiConversationResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAiConversations', query],
    queryFn: () => apiClient.aiConversationService.List(query.toRawParams()),
    ...options,
  });
}

export function fetchListAiConversations(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAiConversations', params],
    queryFn: () => apiClient.aiConversationService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useUpdateAiConversation(
  options?: UseMutationOptions<object, Error, { id: number; data: aiservicev1_AiConversation }>,
) {
  return useMutation({
    mutationFn: ({ id, data }) =>
      apiClient.aiConversationService.Update({
        id,
        data,
        updateMask: makeUpdateMask(['title']),
      }),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiConversations'] }),
    ...options,
  });
}

export function useDeleteAiConversation(
  options?: UseMutationOptions<object, Error, aiservicev1_DeleteAiConversationRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiConversationService.Delete(req),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiConversations'] }),
    ...options,
  });
}

export function useListAiMessages(
  query: PaginationQuery,
  options?: UseQueryOptions<aiservicev1_ListAiMessageResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAiMessages', query],
    queryFn: () => apiClient.aiMessageService.List(query.toRawParams()),
    ...options,
  });
}

export function fetchListAiMessages(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAiMessages', params],
    queryFn: () => apiClient.aiMessageService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useDeleteAiMessage(
  options?: UseMutationOptions<object, Error, aiservicev1_DeleteAiMessageRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiMessageService.Delete(req),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiMessages'] }),
    ...options,
  });
}

// ==============================
// AI 对话（POST 发起；增量经 SSE `ai_chat_chunk` / `ai_chat_tool` 推送）
// ==============================
export function useSendChat(
  options?: UseMutationOptions<aiservicev1_ChatResponse, Error, aiservicev1_ChatRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiChatService.Chat(req),
    ...options,
  });
}

// ==============================
// AI 知识库（RAG）
// ==============================
export function useListAiKnowledgeBases(
  query: PaginationQuery,
  options?: UseQueryOptions<aiservicev1_ListAiKnowledgeBaseResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAiKnowledgeBases', query],
    queryFn: () => apiClient.aiKnowledgeBaseService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListAiKnowledgeBases(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAiKnowledgeBases', params],
    queryFn: () => apiClient.aiKnowledgeBaseService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useGetAiKnowledgeBase(
  req: aiservicev1_GetAiKnowledgeBaseRequest,
  options?: UseQueryOptions<aiservicev1_AiKnowledgeBase, Error>,
) {
  return useQuery({
    queryKey: ['getAiKnowledgeBase', req],
    queryFn: () => apiClient.aiKnowledgeBaseService.Get(req),
    ...options,
  });
}

export function useCreateAiKnowledgeBase(
  options?: UseMutationOptions<object, Error, { data: aiservicev1_AiKnowledgeBase }>,
) {
  return useMutation({
    mutationFn: ({ data }) => apiClient.aiKnowledgeBaseService.Create({ data }),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] }),
    ...options,
  });
}

export function useUpdateAiKnowledgeBase(
  options?: UseMutationOptions<object, Error, { id: number; data: aiservicev1_AiKnowledgeBase }>,
) {
  return useMutation({
    mutationFn: ({ id, data }) =>
      apiClient.aiKnowledgeBaseService.Update({
        id,
        data,
        updateMask: makeUpdateMask(['name', 'description', 'providerId', 'embeddingModel']),
      }),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] }),
    ...options,
  });
}

export function useDeleteAiKnowledgeBase(
  options?: UseMutationOptions<object, Error, { id: number }>,
) {
  return useMutation({
    mutationFn: ({ id }) => apiClient.aiKnowledgeBaseService.Delete({ id }),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] }),
    ...options,
  });
}

export function fetchListAiDocs(params: aiservicev1_ListAiDocsRequest) {
  return queryClient.fetchQuery({
    queryKey: ['listAiDocs', params.baseId],
    queryFn: () => apiClient.aiKnowledgeBaseService.ListDocs(params),
    staleTime: 0,
    retry: 0,
  });
}

export function useUploadAiDoc(
  options?: UseMutationOptions<aiservicev1_UploadAiDocResponse, Error, aiservicev1_UploadAiDocRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.UploadDoc(req),
    onSuccess: (_data, req) => {
      queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] });
      queryClient.invalidateQueries({ queryKey: ['listAiDocs', req.baseId] });
    },
    ...options,
  });
}

export function useUploadAiDocFile(
  options?: UseMutationOptions<
    aiservicev1_UploadAiDocResponse,
    Error,
    aiservicev1_UploadAiDocFileRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.UploadDocFile(req),
    onSuccess: (_data, req) => {
      queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] });
      queryClient.invalidateQueries({ queryKey: ['listAiDocs', req.baseId] });
    },
    ...options,
  });
}

export function useDeleteAiDoc(
  options?: UseMutationOptions<object, Error, aiservicev1_DeleteAiDocRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.DeleteDoc(req),
    onSuccess: (_data, req) => {
      queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] });
      queryClient.invalidateQueries({ queryKey: ['listAiDocs', req.baseId] });
    },
    ...options,
  });
}

export function useSearchAiKnowledge(
  options?: UseMutationOptions<
    aiservicev1_SearchAiKnowledgeResponse,
    Error,
    aiservicev1_SearchAiKnowledgeRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.Search(req),
    ...options,
  });
}

export function useReindexAiDocs(
  options?: UseMutationOptions<object, Error, aiservicev1_ReindexDocsRequest>,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.ReindexDocs(req),
    ...options,
  });
}

export function useSyncOpsKnowledge(
  options?: UseMutationOptions<
    aiservicev1_SyncOpsKnowledgeResponse,
    Error,
    aiservicev1_SyncOpsKnowledgeRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiKnowledgeBaseService.SyncOpsKnowledge(req),
    onSuccess: () =>
      queryClient.invalidateQueries({ queryKey: ['listAiKnowledgeBases'] }),
    ...options,
  });
}

// ==============================
// AI 用量流水
// ==============================
export function useListAiUsageLogs(
  query: PaginationQuery,
  options?: UseQueryOptions<aiservicev1_ListAiUsageLogResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAiUsageLogs', query],
    queryFn: () => apiClient.aiUsageLogService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListAiUsageLogs(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAiUsageLogs', params],
    queryFn: () => apiClient.aiUsageLogService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useAiUsageSummary(
  options?: UseQueryOptions<aiservicev1_UsageSummaryResponse, Error>,
) {
  return useQuery({
    queryKey: ['aiUsageSummary'],
    queryFn: () => apiClient.aiUsageLogService.GetUsageSummary({}),
    ...options,
  });
}

// ==============================
// 智能问数（NL → 只读 SQL → 结构化结果）
// ==============================
export function useAskAiQuery(
  options?: UseMutationOptions<
    aiservicev1_AskAiQueryResponse,
    Error,
    aiservicev1_AskAiQueryRequest
  >,
) {
  return useMutation({
    mutationFn: (req) => apiClient.aiQueryService.Ask(req),
    ...options,
  });
}

// ==============================
// 共享小件
// ==============================

/** AiDoc 文档行（列表用）。 */
export type AiDoc = aiservicev1_AiDoc;

/** AiMessage 消息行。 */
export type AiMessage = aiservicev1_AiMessage;
