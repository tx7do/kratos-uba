import { defHttp } from '/@/utils/http/axios';

const route = {
  Users: '/users',
  UserWithId: (id) => `/users/${id}`,
};

// 获取列表
export const ListUser: ListUser = (params) => {
  return defHttp.get<ListUserResponse>({
    url: route.Users,
    params,
  });
};

// 获取
export const GetUser: GetUser = (params) => {
  return defHttp.get<User>({
    url: route.UserWithId(params.id),
  });
};

// 创建
export const CreateUser: CreateUser = (params) => {
  return defHttp.post<User>(
    { url: route.Users, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateUser: UpdateUser = (params) => {
  return defHttp.put<User>({
    url: route.UserWithId(params.id),
    data: params.user,
  });
};

// 删除
export const DeleteUser: DeleteUser = (params) => {
  return defHttp.delete({
    url: route.UserWithId(params.id),
  });
};
