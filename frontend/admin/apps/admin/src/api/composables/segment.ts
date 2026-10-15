import type {
  ubaservicev1_CountSegmentResponse,
  ubaservicev1_GetSegmentRequest,
  ubaservicev1_ListSegmentResponse,
  ubaservicev1_ListSegmentUsersRequest,
  ubaservicev1_ListSegmentUsersResponse,
  ubaservicev1_RefreshSegmentRequest,
  ubaservicev1_Segment,
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
// Segment 用户分群
// ==============================

export function useListSegments(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListSegmentResponse, Error>,
) {
  return useQuery({
    queryKey: ['listSegments', query],
    queryFn: () => apiClient.segmentService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListSegments(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listSegments', params],
    queryFn: () => apiClient.segmentService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountSegments(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountSegmentResponse, Error>,
) {
  return useQuery({
    queryKey: ['countSegments', query],
    queryFn: () => apiClient.segmentService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetSegment(
  req: ubaservicev1_GetSegmentRequest,
  options?: UseQueryOptions<ubaservicev1_Segment, Error>,
) {
  return useQuery({
    queryKey: ['getSegment', req],
    queryFn: () => apiClient.segmentService.Get(req),
    ...options,
  });
}

export function useCreateSegment(
  options?: UseMutationOptions<
    ubaservicev1_Segment,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.segmentService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateSegment(
  options?: UseMutationOptions<
    ubaservicev1_Segment,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.segmentService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteSegment(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.segmentService.Delete({ id }),
    ...options,
  });
}

export function useRefreshSegment(
  options?: UseMutationOptions<ubaservicev1_Segment, Error, number>,
) {
  return useMutation({
    mutationFn: (id) =>
      apiClient.segmentService.Refresh({ id } as ubaservicev1_RefreshSegmentRequest),
    ...options,
  });
}

export function fetchListSegmentUsers(
  req: ubaservicev1_ListSegmentUsersRequest,
) {
  return queryClient.fetchQuery({
    queryKey: ['listSegmentUsers', req],
    queryFn: () =>
      apiClient.segmentService.ListUsers(
        req as ubaservicev1_ListSegmentUsersRequest,
      ) as Promise<ubaservicev1_ListSegmentUsersResponse>,
    staleTime: 0,
    retry: 0,
  });
}
