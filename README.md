# cybertect

[![Build Status](https://github.com/vicanso/cybertect/workflows/Test/badge.svg)](https://github.com/vicanso/cybertect/actions)

基于`elton`的脚手架，实现了数据校验、行为统计等功能。

## 特性

### 简单易用的应用配置

- 应用配置通过加载default.yml + 当前GO_ENV所对应的yml组合生成，简化配置
- 支持配置参数的校验，保证应用启动时的参数准确性
- 支持优先从ENV中获取配置参数，若获取失败再使用yml配置



### 简单易用的缓存模块

缓存模块支持三类缓存：LRU+TTL的高速缓存，redis+ttl的共有缓存以及redis+lru+ttl的多级缓存

```go
// redis缓存中简化的struct读取与保存
err := cache.GetRedisCache().SetStruct(context.Background(), "key", &map[string]string{
  "name": "my name",
}, time.Minute)
data := make(map[string]string)
err = cache.GetRedisCache().GetStruct(context.Background(), "key", &data)
```
### 详尽的性能指标

性能指标中包括以下的相关指标：

- `goMaxProcs` 程序使用的最大CPU数量
- `threadCount` 程序当前线程数
- `memSys` 系统申请内存
- `memHeapSys` 系统申请heap
- `memHeapInuse` 使用中的heap
- `memFrees` heap对象释放的数量
- `routineCount` goroutine的数量
- `cpuUsage` CPU使用率
- `lastGC` 上一次GC的时间
- `numGC` GC的次数
- `recentPause` 最近一次GC暂停的时长
- `pauseTotal` GC暂停的总时长
- `connProcessing` 当前处理中的连接数（http.Server)
- `connProcessedCount` 处理过的连接总数
- `connAlive` 当前活跃的连接数
- `connCreatedCount` 被创建的连接总数
- `concurrency` 请求并发数
- `requestProcessedTotal` 处理过的请求总数

### Redis

redis模块记录了当前并发请求以及pipeline请求量，可以设置最大并发请求量，提供简单的熔断处理。


## commit

feat：新功能（feature）

fix：修补bug

docs：文档（documentation）

style： 格式（不影响代码运行的变动）

refactor：重构（即不是新增功能，也不是修改bug的代码变动）

test：增加测试

chore：构建过程或辅助工具的变动

## 启动数据库

### postgres

```
docker pull postgres:alpine

docker run -d --restart=always \
  -v $PWD/cybertect:/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD=A123456 \
  -p 5432:5432 \
  --name=cybertect \
  postgres:alpine

docker exec -it cybertect sh

psql -c "CREATE DATABASE cybertect;" -U postgres
psql -c "CREATE USER vicanso WITH PASSWORD 'A123456';" -U postgres
psql -c "GRANT ALL PRIVILEGES ON DATABASE cybertect to vicanso;" -U postgres
```

## redis

```
docker pull redis:alpine

docker run -d --restart=always \
  -p 6379:6379 \
  --name=redis \
  redis:alpine
```

## 规范

- 所有自定义的error都必须为hes.Error
