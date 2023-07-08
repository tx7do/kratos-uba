//用户服务
//Service: UserService
//查询用户列表
type ListUser = (params: PagingRequest) => Promise<ListUserResponse>;

//查询用户详情
type GetUser = (params: GetUserRequest) => Promise<User>;

//创建用户
type CreateUser = (params: CreateUserRequest) => Promise<User>;

//更新用户
type UpdateUser = (params: UpdateUserRequest) => Promise<User>;

//删除用户
type DeleteUser = (params: DeleteUserRequest) => Promise<{}>;

//验证密码
type VerifyPassword = (params: VerifyPasswordRequest) => Promise<VerifyPasswordResponse>;

//用户是否存在
type UserExists = (params: UserExistsRequest) => Promise<UserExistsResponse>;

//查询用户详情
type GetUserByUserName = (params: GetUserByUserNameRequest) => Promise<User>;

enum VerifyPasswordResult {
  SUCCESS = 0, //验证成功
  ACCOUNT_NOT_EXISTS = 1, //账号不存在
  WRONG_PASSWORD = 2, //密码错误
  FREEZE = 3, //已冻结
  DELETED = 4, //已删除
}

//用户
interface User {
  id?: number; //用户ID
  userName?: string; //用户登陆名，使用手机号
  realName?: string; //真实姓名
  email?: string; //电子邮箱
  avatar?: string; //头像
  roleId?: number;
  creatorId?: number;
  authority?: string;
  description?: string;
  password?: string;
  createTime?: string; //创建时间
  updateTime?: string; //更新时间
  deleteTime?: string; //删除时间
  status?: string;
}

//获取用户列表 - 答复
interface ListUserResponse {
  items?: User[];
  total?: number;
}

//获取用户数据 - 请求
interface GetUserRequest {
  id?: number;
}

//创建用户 - 请求
interface CreateUserRequest {
  user?: User;
  operatorId?: number;
}

//更新用户 - 请求
interface UpdateUserRequest {
  id?: number;
  user?: User;
  operatorId?: number;
}

//删除用户 - 请求
interface DeleteUserRequest {
  id?: number;
  operatorId?: number;
}

//验证密码 - 请求
interface VerifyPasswordRequest {
  userName?: string;
  password?: string;
}

//验证密码 - 答复
interface VerifyPasswordResponse {
  result?: VerifyPasswordResult;
}

//用户是否存在 - 请求
interface UserExistsRequest {
  userName?: string;
}

//用户是否存在 - 答复
interface UserExistsResponse {
  exist?: boolean;
}

interface GetUserByUserNameRequest {
  userName?: string;
}
