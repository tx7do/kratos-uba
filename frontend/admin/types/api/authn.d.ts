//用户后台登陆认证服务
//Service: AuthenticationService
//后台登陆
type Login = (params: LoginRequest) => Promise<LoginResponse>;

//后台登出
type Logout = (params: LogoutRequest) => Promise<{}>;

//后台获取已经登陆的用户的数据
type GetMe = () => Promise<user.service.v1.User>;

//用户后台登陆 - 请求
interface LoginRequest {
  userName?: string;
  password?: string;
}

//用户后台登陆 - 回应
interface LoginResponse {
  id?: number;
  userName?: string;
  token?: string;
  refreshToken?: string;
}

//用户后台登出 - 请求
interface LogoutRequest {
  id?: number;
}

//获取当前用户信息 - 请求
interface GetMeRequest {
  id?: number;
}
