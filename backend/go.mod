module kratos-bi

go 1.19

require (
	entgo.io/contrib v0.3.5
	entgo.io/ent v0.11.10
	github.com/go-chassis/sc-client v0.7.0
	github.com/go-kratos/aegis v0.1.4
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20230201021226-239121155d0f
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20230201021226-239121155d0f
	github.com/go-kratos/kratos/contrib/config/kubernetes/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20230201021226-239121155d0f
	github.com/go-kratos/kratos/contrib/log/aliyun/v2 v2.0.0-20230222075022-bab67facad0d
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20230222075022-bab67facad0d
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20230222075022-bab67facad0d
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20230222075022-bab67facad0d
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20230222075022-bab67facad0d
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20230201021226-239121155d0f
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/registry/servicecomb/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20230213033822-0a076443cba1
	github.com/go-kratos/kratos/v2 v2.6.1
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.7.0
	github.com/go-zookeeper/zk v1.0.3
	github.com/google/gnostic v0.6.9
	github.com/google/subcommands v1.2.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.18.0
	github.com/jackc/pgx/v4 v4.17.2
	github.com/lib/pq v1.10.7
	github.com/minio/minio-go/v7 v7.0.49
	github.com/nacos-group/nacos-sdk-go v1.1.4
	github.com/olekukonko/tablewriter v0.0.5
	github.com/oschwald/geoip2-golang v1.8.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	github.com/tx7do/kratos-authn v0.0.8
	github.com/tx7do/kratos-authn/engine/jwt v0.0.0-20230304060223-c4ad57a481a3
	github.com/tx7do/kratos-authn/middleware v0.0.0-20230304060223-c4ad57a481a3
	github.com/tx7do/kratos-authz v0.0.4
	github.com/tx7do/kratos-authz/middleware v0.0.0-20230216134644-a8796a116ab5
	github.com/tx7do/kratos-transport v1.0.5
	github.com/tx7do/kratos-transport/broker/kafka v0.0.0-20230503101546-8cd8dcbefa36
	github.com/tx7do/kratos-transport/transport/kafka v0.0.0-20230503101546-8cd8dcbefa36
	go.etcd.io/etcd/client/v3 v3.5.7
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/exporters/zipkin v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.8.0
	golang.org/x/tools v0.6.1-0.20230222164832-25d2519c8696
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	k8s.io/client-go v0.26.2
)

require (
	ariga.io/atlas v0.9.2-0.20230303073438-03a4779a6338 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.166 // indirect
	github.com/aliyun/aliyun-log-go-sdk v0.1.43 // indirect
	github.com/apolloconfig/agollo/v4 v4.3.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-chassis/cari v0.9.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jhump/protoreflect v1.10.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/oschwald/maxminddb-golang v1.10.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0-beta.8 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	github.com/shirou/gopsutil/v3 v3.21.8 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.11.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.2 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/xdg/scram v1.0.5 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.26.2 // indirect
	k8s.io/apimachinery v0.26.2 // indirect
	k8s.io/klog/v2 v2.90.0 // indirect
	k8s.io/kube-openapi v0.0.0-20230210211930-4b0756abdef5 // indirect
	k8s.io/utils v0.0.0-20230209194617-a36077c30491 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
