import type {
  ubaservicev1_CountFeatureFlagResponse,
  ubaservicev1_FlagResultResponse,
  ubaservicev1_FeatureFlag,
  ubaservicev1_GetFeatureFlagRequest,
  ubaservicev1_GetFlagResultRequest,
  ubaservicev1_ListFeatureFlagResponse,
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
// FeatureFlag 功能开关 / A/B 实验
// ==============================

export function useListFeatureFlags(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListFeatureFlagResponse, Error>,
) {
  return useQuery({
    queryKey: ['listFeatureFlags', query],
    queryFn: () => apiClient.featureFlagService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListFeatureFlags(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listFeatureFlags', params],
    queryFn: () => apiClient.featureFlagService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountFeatureFlags(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountFeatureFlagResponse, Error>,
) {
  return useQuery({
    queryKey: ['countFeatureFlags', query],
    queryFn: () => apiClient.featureFlagService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetFeatureFlag(
  req: ubaservicev1_GetFeatureFlagRequest,
  options?: UseQueryOptions<ubaservicev1_FeatureFlag, Error>,
) {
  return useQuery({
    queryKey: ['getFeatureFlag', req],
    queryFn: () => apiClient.featureFlagService.Get(req),
    ...options,
  });
}

export function useCreateFeatureFlag(
  options?: UseMutationOptions<
    ubaservicev1_FeatureFlag,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.featureFlagService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateFeatureFlag(
  options?: UseMutationOptions<
    ubaservicev1_FeatureFlag,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.featureFlagService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteFeatureFlag(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.featureFlagService.Delete({ id }),
    ...options,
  });
}

export async function fetchFlagResult(req: ubaservicev1_GetFlagResultRequest) {
  return queryClient.fetchQuery({
    queryKey: ['flagResult', req],
    queryFn: () =>
      apiClient.featureFlagService.GetFlagResult(
        req,
      ) as Promise<ubaservicev1_FlagResultResponse>,
    staleTime: 0,
    retry: 0,
  });
}
