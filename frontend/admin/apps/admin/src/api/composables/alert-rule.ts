import type {
  ubaservicev1_CountAlertRuleResponse,
  ubaservicev1_GetAlertRuleRequest,
  ubaservicev1_ListAlertRuleResponse,
  ubaservicev1_AlertRule,
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
// AlertRule 指标阈值告警
// ==============================

export function useListAlertRules(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListAlertRuleResponse, Error>,
) {
  return useQuery({
    queryKey: ['listAlertRules', query],
    queryFn: () => apiClient.alertRuleService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListAlertRules(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listAlertRules', params],
    queryFn: () => apiClient.alertRuleService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountAlertRules(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountAlertRuleResponse, Error>,
) {
  return useQuery({
    queryKey: ['countAlertRules', query],
    queryFn: () => apiClient.alertRuleService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetAlertRule(
  req: ubaservicev1_GetAlertRuleRequest,
  options?: UseQueryOptions<ubaservicev1_AlertRule, Error>,
) {
  return useQuery({
    queryKey: ['getAlertRule', req],
    queryFn: () => apiClient.alertRuleService.Get(req),
    ...options,
  });
}

export function useCreateAlertRule(
  options?: UseMutationOptions<
    ubaservicev1_AlertRule,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.alertRuleService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateAlertRule(
  options?: UseMutationOptions<
    ubaservicev1_AlertRule,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.alertRuleService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteAlertRule(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.alertRuleService.Delete({ id }),
    ...options,
  });
}
