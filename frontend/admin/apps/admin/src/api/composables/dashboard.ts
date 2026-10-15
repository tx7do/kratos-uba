import type {
  ubaservicev1_CountDashboardResponse,
  ubaservicev1_GetDashboardRequest,
  ubaservicev1_ListDashboardResponse,
  ubaservicev1_Dashboard,
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
// Dashboard 自定义看板
// ==============================

export function useListDashboards(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListDashboardResponse, Error>,
) {
  return useQuery({
    queryKey: ['listDashboards', query],
    queryFn: () => apiClient.dashboardService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListDashboards(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listDashboards', params],
    queryFn: () => apiClient.dashboardService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountDashboards(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountDashboardResponse, Error>,
) {
  return useQuery({
    queryKey: ['countDashboards', query],
    queryFn: () => apiClient.dashboardService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetDashboard(
  req: ubaservicev1_GetDashboardRequest,
  options?: UseQueryOptions<ubaservicev1_Dashboard, Error>,
) {
  return useQuery({
    queryKey: ['getDashboard', req],
    queryFn: () => apiClient.dashboardService.Get(req),
    ...options,
  });
}

export function useCreateDashboard(
  options?: UseMutationOptions<
    ubaservicev1_Dashboard,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.dashboardService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateDashboard(
  options?: UseMutationOptions<
    ubaservicev1_Dashboard,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.dashboardService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteDashboard(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.dashboardService.Delete({ id }),
    ...options,
  });
}
