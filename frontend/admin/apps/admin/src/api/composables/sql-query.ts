import type {
  ubaservicev1_ExecuteSqlQueryRequest,
  ubaservicev1_ExecuteSqlQueryResponse,
} from '#/generated/api/admin/service/v1';

import { queryClient } from '#/plugins/vue-query';
import { apiClient } from '#/api/client';

// ==============================
// SqlQuery 自助只读查询
// ==============================

export async function fetchExecuteSqlQuery(
  req: ubaservicev1_ExecuteSqlQueryRequest,
): Promise<ubaservicev1_ExecuteSqlQueryResponse> {
  return queryClient.fetchQuery({
    queryKey: ['executeSqlQuery', req],
    queryFn: () =>
      apiClient.sqlQueryService.ExecuteQuery(
        req,
      ) as Promise<ubaservicev1_ExecuteSqlQueryResponse>,
    staleTime: 0,
    retry: 0,
  });
}
