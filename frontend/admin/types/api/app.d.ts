//应用服务
//Service: ApplicationService
//查询应用列表
type ListApplication = (params: pagination.PagingRequest) => Promise<ListApplicationResponse>;

//查询应用详情
type GetApplication = (params: GetApplicationRequest) => Promise<Application>;

//创建应用
type CreateApplication = (params: CreateApplicationRequest) => Promise<Application>;

//更新应用
type UpdateApplication = (params: UpdateApplicationRequest) => Promise<Application>;

//删除应用
type DeleteApplication = (params: DeleteApplicationRequest) => Promise<{}>;

//查询应用详情
type GetApplicationByAppId = (params: GetApplicationByAppIdRequest) => Promise<Application>;

//应用
interface Application {
  id?: number;
  name?: string;
  appId?: string;
  appKey?: string;
  status?: string;
  creatorId?: number;
  ownerId?: number;
  remark?: string;
  keepMonth?: number;
  createTime?: string;
  updateTime?: string;
  deleteTime?: string;
}

//获取应用列表 - 答复
interface ListApplicationResponse {
  items?: Application[];
  total?: number;
}

//获取应用数据 - 请求
interface GetApplicationRequest {
  id?: number;
}

//创建应用 - 请求
interface CreateApplicationRequest {
  app?: Application;
  operatorId?: number;
}

//更新应用 - 请求
interface UpdateApplicationRequest {
  id?: number;
  app?: Application;
  operatorId?: number;
}

//删除应用 - 请求
interface DeleteApplicationRequest {
  id?: number;
  operatorId?: number;
}

interface GetApplicationByAppIdRequest {
  appId?: string;
}
