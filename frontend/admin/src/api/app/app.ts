import { defHttp } from '/@/utils/http/axios';

const route = {
  Applications: '/apps',
  ApplicationWithId: (id) => `/apps/${id}`,
};

// 获取列表
export const ListApplication: ListApplication = (params) => {
  return defHttp.get<ListApplicationResponse>({
    url: route.Applications,
    params,
  });
};

// 获取
export const GetApplication: GetApplication = (params) => {
  return defHttp.get<Application>({
    url: route.ApplicationWithId(params.id),
  });
};

// 创建
export const CreateApplication: CreateApplication = (params) => {
  return defHttp.post<Application>(
    { url: route.Applications, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateApplication: UpdateApplication = (params) => {
  return defHttp.put<Application>({
    url: route.ApplicationWithId(params.id),
    data: params.app,
  });
};

// 删除
export const DeleteApplication: DeleteApplication = (params) => {
  return defHttp.delete({
    url: route.ApplicationWithId(params.id),
  });
};
