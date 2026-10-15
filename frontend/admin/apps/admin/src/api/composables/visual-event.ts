import type {
  ubaservicev1_CountVisualEventResponse,
  ubaservicev1_GetVisualEventRequest,
  ubaservicev1_ListVisualEventResponse,
  ubaservicev1_VisualEvent,
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
// VisualEvent 可视化埋点
// ==============================

export function useListVisualEvents(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListVisualEventResponse, Error>,
) {
  return useQuery({
    queryKey: ['listVisualEvents', query],
    queryFn: () => apiClient.visualEventService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListVisualEvents(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listVisualEvents', params],
    queryFn: () => apiClient.visualEventService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountVisualEvents(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountVisualEventResponse, Error>,
) {
  return useQuery({
    queryKey: ['countVisualEvents', query],
    queryFn: () => apiClient.visualEventService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetVisualEvent(
  req: ubaservicev1_GetVisualEventRequest,
  options?: UseQueryOptions<ubaservicev1_VisualEvent, Error>,
) {
  return useQuery({
    queryKey: ['getVisualEvent', req],
    queryFn: () => apiClient.visualEventService.Get(req),
    ...options,
  });
}

export function useCreateVisualEvent(
  options?: UseMutationOptions<
    ubaservicev1_VisualEvent,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.visualEventService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateVisualEvent(
  options?: UseMutationOptions<
    ubaservicev1_VisualEvent,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.visualEventService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteVisualEvent(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.visualEventService.Delete({ id }),
    ...options,
  });
}
