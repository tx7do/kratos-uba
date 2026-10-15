import type {
  ubaservicev1_CountReportSubscriptionResponse,
  ubaservicev1_GetReportSubscriptionRequest,
  ubaservicev1_ListReportSubscriptionResponse,
  ubaservicev1_ReportSubscription,
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
// ReportSubscription 定时报告
// ==============================

export function useListReportSubscriptions(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListReportSubscriptionResponse, Error>,
) {
  return useQuery({
    queryKey: ['listReportSubscriptions', query],
    queryFn: () => apiClient.reportSubscriptionService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListReportSubscriptions(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listReportSubscriptions', params],
    queryFn: () =>
      apiClient.reportSubscriptionService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountReportSubscriptions(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountReportSubscriptionResponse, Error>,
) {
  return useQuery({
    queryKey: ['countReportSubscriptions', query],
    queryFn: () => apiClient.reportSubscriptionService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetReportSubscription(
  req: ubaservicev1_GetReportSubscriptionRequest,
  options?: UseQueryOptions<ubaservicev1_ReportSubscription, Error>,
) {
  return useQuery({
    queryKey: ['getReportSubscription', req],
    queryFn: () => apiClient.reportSubscriptionService.Get(req),
    ...options,
  });
}

export function useCreateReportSubscription(
  options?: UseMutationOptions<
    ubaservicev1_ReportSubscription,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.reportSubscriptionService.Create({
        data: { ...values } as any,
      }),
    ...options,
  });
}

export function useUpdateReportSubscription(
  options?: UseMutationOptions<
    ubaservicev1_ReportSubscription,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.reportSubscriptionService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteReportSubscription(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.reportSubscriptionService.Delete({ id }),
    ...options,
  });
}
