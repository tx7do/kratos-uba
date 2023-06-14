# Kratos 用户行为分析系统 - 后端

## 后端组成

1. Admin Service，系统后台的BFF；
2. Agent Service，埋点上报的BFF；
3. Core Service，核心服务；
4. Logger Service，消费Kafka，入库到ClickHouse（如果使用ClickHouse，它可以不需要，因为CH具有Kafka Engine）；
5. Report Service，查询日志，并输出为报表。

## 后端依赖的第三方组件

- PostgreSQL
- ClickHouse
- Redis
- ZooKeeper
- Kafka
- Consul
- Jaeger
- MinIO

请使用`docker-compose`的方式安装以上第三方依赖组件。

## 参考资料

- [Business Intelligence in Microservices: Improving Performance](https://dzone.com/articles/business-intelligence-in-microservices-improving-p)
- [基于 ClickHouse 高性能引擎集群构建数据湖](https://toutiao.io/posts/pklw5vz/preview)
- [ClickHouse整合Kafka](https://learn-bigdata.incubator.edurt.io/docs/ClickHouse/Action/engine-kafka/)
- [Apply CDC from MySQL to ClickHouse](https://medium.com/@hoptical/apply-cdc-from-mysql-to-clickhouse-d660873311c7)
- [ClickHouse 在实时场景的应用和优化](https://mp.weixin.qq.com/s/hqUCFSr8cu3x3u8HCA6WYg)
- [ClickHouse基础&实践&调优全视角解析](https://xie.infoq.cn/article/37886f3baca09057580bdd5aa)
- [从维护几百张表到只需维护一张表，一个UEI模型就够了](https://zhuanlan.zhihu.com/p/623182999)
- [BI花5天完成的分析，UBA只需30秒](https://zhuanlan.zhihu.com/p/629574865)
