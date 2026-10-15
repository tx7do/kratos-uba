import type {
  ubaservicev1_GetReplaySessionResponse,
  ubaservicev1_ListReplaySessionsRequest,
  ubaservicev1_ListReplaySessionsResponse,
} from '#/generated/api/admin/service/v1';

import { queryClient } from '#/plugins/vue-query';
import { apiClient } from '#/api/client';

// ==============================
// Replay 会话回放
// ==============================

export async function fetchListReplaySessions(
  req: ubaservicev1_ListReplaySessionsRequest,
): Promise<ubaservicev1_ListReplaySessionsResponse> {
  return queryClient.fetchQuery({
    queryKey: ['listReplaySessions', req],
    queryFn: () =>
      apiClient.replayService.ListSessions(
        req,
      ) as Promise<ubaservicev1_ListReplaySessionsResponse>,
    staleTime: 0,
    retry: 0,
  });
}

export async function fetchReplaySessionEvents(sessionId: string) {
  return queryClient.fetchQuery({
    queryKey: ['replaySessionEvents', sessionId],
    queryFn: () =>
      apiClient.replayService.GetSessionEvents({
        sessionId,
      }) as Promise<ubaservicev1_GetReplaySessionResponse>,
    staleTime: 0,
    retry: 0,
  });
}

export async function deleteReplaySession(sessionId: string) {
  return apiClient.replayService.DeleteSession({ sessionId });
}
