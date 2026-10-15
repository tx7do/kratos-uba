import type {
  ubaservicev1_CountDashboardResponse,
  ubaservicev1_GetDashboardRequest,
  ubaservicev1_ListDashboardResponse,
  ubaservicev1_Dashboard,
} from '#/generated/api/admin/service/v1';

import { useMutation, type UseMutationOptions, useQuery, type UseQueryOptions } from '@tanstack/vue-query';
import { unref, type MaybeRef } from 'vue';

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
  req: MaybeRef<ubaservicev1_GetDashboardRequest>,
  options?: UseQueryOptions<ubaservicev1_Dashboard, Error>,
) {
  return useQuery({
    queryKey: ['getDashboard', req],
    // 请求参数可能是 computed（路由参数变化时自动换 key 重查），调用前先解包
    queryFn: () => apiClient.dashboardService.Get(unref(req)),
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
    // data 必须带上 id：updateMask 里有 id 而 data.id 为空时，
    // go-crud 会把 id 列 SET NULL，触发数据库 not-null 约束
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.dashboardService.Update({
        id,
        data: { id, ...values } as any,
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
