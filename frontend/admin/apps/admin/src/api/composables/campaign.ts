import type {
  ubaservicev1_Campaign,
  ubaservicev1_CountCampaignResponse,
  ubaservicev1_GetCampaignRequest,
  ubaservicev1_ListCampaignResponse,
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
// Campaign 营销触达
// ==============================

export function useListCampaigns(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_ListCampaignResponse, Error>,
) {
  return useQuery({
    queryKey: ['listCampaigns', query],
    queryFn: () => apiClient.campaignService.List(query.toRawParams()),
    ...options,
  });
}

export async function fetchListCampaigns(params: PaginationQuery) {
  return queryClient.fetchQuery({
    queryKey: ['listCampaigns', params],
    queryFn: () => apiClient.campaignService.List(params.toRawParams()),
    staleTime: 0,
    retry: 0,
  });
}

export function useCountCampaigns(
  query: PaginationQuery,
  options?: UseQueryOptions<ubaservicev1_CountCampaignResponse, Error>,
) {
  return useQuery({
    queryKey: ['countCampaigns', query],
    queryFn: () => apiClient.campaignService.Count(query.toRawParams()),
    ...options,
  });
}

export function useGetCampaign(
  req: ubaservicev1_GetCampaignRequest,
  options?: UseQueryOptions<ubaservicev1_Campaign, Error>,
) {
  return useQuery({
    queryKey: ['getCampaign', req],
    queryFn: () => apiClient.campaignService.Get(req),
    ...options,
  });
}

export function useCreateCampaign(
  options?: UseMutationOptions<
    ubaservicev1_Campaign,
    Error,
    Record<string, any>
  >,
) {
  return useMutation({
    mutationFn: (values) =>
      apiClient.campaignService.Create({ data: { ...values } as any }),
    ...options,
  });
}

export function useUpdateCampaign(
  options?: UseMutationOptions<
    ubaservicev1_Campaign,
    Error,
    { id: number; values: Record<string, any> }
  >,
) {
  return useMutation({
    mutationFn: ({ id, values }: { id: number; values: Record<string, any> }) =>
      apiClient.campaignService.Update({
        id,
        data: { ...values } as any,
        updateMask: makeUpdateMask(Object.keys(values ?? {})),
      }),
    ...options,
  });
}

export function useDeleteCampaign(
  options?: UseMutationOptions<object, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.campaignService.Delete({ id }),
    ...options,
  });
}

export function useSendCampaign(
  options?: UseMutationOptions<ubaservicev1_Campaign, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.campaignService.Send({ id }),
    ...options,
  });
}

export function useCancelCampaign(
  options?: UseMutationOptions<ubaservicev1_Campaign, Error, number>,
) {
  return useMutation({
    mutationFn: (id) => apiClient.campaignService.Cancel({ id }),
    ...options,
  });
}
