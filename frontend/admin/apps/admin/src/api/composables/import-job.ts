import type {
  ubaservicev1_ImportJob,
  ubaservicev1_ListImportJobResponse,
} from '#/generated/api/admin/service/v1';

import {
  useMutation,
  type UseMutationOptions,
  useQuery,
  type UseQueryOptions,
} from '@tanstack/vue-query';

import { apiClient } from '#/api/client';
import { queryClient } from '#/plugins/vue-query';
import type { PaginationQuery } from '#/transport/rest';

// ==============================
// ImportJob 历史数据导入
// ==============================

export function useListImportJobs(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListImportJobResponse, Error>,
) {
  return useQuery({
    queryKey: ['listImportJobs', query],
    queryFn: () => apiClient.importService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListImportJobs(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listImportJobs', params],
    queryFn: () => apiClient.importService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCreateImportJob(
  options?: UseMutationOptions<
    ubaservicev1_ImportJob,
    Error,
    { name: string; appId?: string; eventsJson: string }
  >,
) {
  return useMutation({
    mutationFn: (req) =>
      apiClient.importService.Create({
        name: req.name,
        appId: req.appId,
        eventsJson: req.eventsJson,
      } as any),
    ...options,
  });
}

export function useCancelImportJob(
  options?: UseMutationOptions<ubaservicev1_ImportJob, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.importService.Cancel({ id }),
    ...options,
  });
}

export function useDeleteImportJob(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.importService.Delete({ id }),
    ...options,
  });
}
