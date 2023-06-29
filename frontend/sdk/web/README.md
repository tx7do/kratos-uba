# Web SDK

## 初始化

初始化方法内部已处理为单例

```html

<script type="text/javascript" src="report_sdk.js"></script>
<script type="text/javascript">

    const params = {"appid": "********", "appkey": "********"}
    //第一个参数为 上报服务的启动地址
    //第二个参数为 appid   由BI后台创建应用后生成
    //第三个参数为 appkey  由BI后台创建应用后生成
    //第四个参数为 客户端debug类型 
    // 0 为正常入kafka 
    // 1 为测试模式（数据入库）
    // 2 为测试模式（数据不入库）
    //当debug为 1 or 2 并且该distinctid在BI系统的Debug模式功能中添加了,就可触发测试模式
    let eventReport = new EventReport("http://127.0.0.1:8091", params["appid"], params["appkey"], 0)

</script>
```

## 手动设置用户的distinctid

```js
eventReport.identify("23232")
```

## 注销当前用户

```js
eventReport.logout()
```

## 获取当前用户的distinctid

```js
eventReport.getDistinctId()
```

## 登录

由于distinctid为本地离线UUID,如想接入自己业务服务器用户id用此方法

```js
eventReport.login()
```

## 获取设备注册时间戳

```js
eventReport.getRegTime()
```

## 设置全局属性

例如用户名，年龄这些属性需要设置在每个事件中，就可以调用此方法

```js
eventReport.setSuperProperties({name: "张三", age: 18})
```

## 删除全局属性值

```js
// 删除name属性
eventReport.unsetSuperProperties("name")
```

## 清空全局属性

```js
eventReport.clearSuperProperties()
```

## 上报用户属性

```js
// 链式调用
eventReport.userSetOnce({"address": "井湾子街道"}) //如果您要上传的用户属性只要设置一次，则可以调用 userSetOnce 来进行设置，当该属性之前已经有值的时候，将会忽略这条信息
    .userSet({name_tset: "张三", age: 18})//对于一般的用户属性，您可以调用 userSet 来进行设置，如果之前存在该用户属性将会覆盖原有的属性值，如果之前不存在该用户属性，则会新建该用户属性。
    .userUnset("name_tset")//当您要清空用户的某个用户属性值时，您可以调用 userUnset 来对指定属性进行清空操作
    .userAdd({age: 1})//当您要上传数值型的属性时，您可以调用 userAdd 来对该属性进行累加操作，如果该属性还未被设置，则会赋值 0 后再进行计算。如果传入负值，等同于减法操作
    .trackUserData()//最终上报
```

## 事件上报

```js
//进行事件上报 
//第一个参数为 事件名
//第二个参数为 事件附加属性 类型为对象
eventReport.track("访问网站", {feeling: "Hello World!"})
```

## 完整示例

```html

<script type="text/javascript" src="report_sdk.js"></script>
<script type="text/javascript">
    const params = {"appid": "********", "appkey": "********"}
    const eventReport = new EventReport("http://127.0.0.1:8091", params["appid"], params["appkey"], 0)
    eventReport.track("访问网站", {feeling: "Hello World!"})
</script>
```
