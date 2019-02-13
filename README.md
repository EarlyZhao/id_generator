
# 分布式ID生成器
产生整型ID，依赖关系型数据库产生ID，通过多缓冲区来实现高性能。支持Go 1.9 及以上版本。
- 高性能，单机100000➕
- 简单，不依赖Web框架，只依赖少量第三方包
- ID连续递增(可能是弊端)
- 易扩展



# 工作原理
![多buffer](https://github.com/EarlyZhao/id_generator/blob/master/images/buffer.jpg?raw=true)

工作原理如上图，每个业务类型有一个buffer数组，数量可配置。ID在buffer中产生，buffer中有一个`Current`值，每生成一个ID，`Current`的值会加1，当`Current`等于`ended_at`时，说明当前buffer中缓冲的ID已经用完，这时buffer指针会指向下一个buffer，同时会创建一个gorouting去填满耗尽的buffer(一条UPDATE和SELECT语句)。ID的创建都是基于内存的`Current++`操作，性能很棒。

如果buffer很小，同时buffer的数量也很小，当所有buffer都耗尽而填满buffer的gorouting还没有结束时，程序会产生等待。 **所以在实际使用时，一定要根据并发量设置合适的buffer容量和buffer数量**。如果每秒的ID消耗在10万，那么建议每个buffer的大小就设置为10万，同时其数量不少于2个为宜，调整buffer大小和数量能适应并发峰值就行。

buffer的大小由`interval`的值决定，每个业务类型可以自由配置这个值。具体的配置说明可见下文。

# 如何使用
## 1.下载代码：
```go
go get github.com/EarlyZhao/id_generator
```
项目使用[dep](https://github.com/golang/dep)来作为版本依赖管理，[安装dep](https://github.com/golang/dep#Installation) 。

## 2.然后安装依赖包：
```go
dep ensure -v 
```
在启动服务时可能会遇到`undefined: proto.ProtoPackageIsVersion3`问题，可参考[这个解决方案](https://github.com/golang/protobuf/issues/763#issuecomment-449856852)，手动下载合适的proto包到vendor目录。
## 3.创建配置文件
配置文件用来做以下配置：
- 数据库
- buffer数量
- 修改业务配置的凭证

文件内容如下，见[config.yml.example](https://github.com/EarlyZhao/id_generator/blob/master/config.yml.example)：
```yml
secret: "this-is-a-secret-token-for-setup-data" # 修改业务配置的凭证
buffers: 2    # buffer数量
database: "mysql" # 暂时只支持MySQL
mysql:
  username: "root"
  password: "password"
  host: "localhost"
  port: 3306
  database: "id_generator"
```
配置文件需要是.yml文件，当以上数据配置完毕后，接下来就可以启动服务了。

## 4.启动服务

HTTP 服务：
```go
// go run main.go -h
go run main.go  -addr 127.0.0.1 -p 8080 -config_path your_file_path.yml
```
gRPC:
```go
// -run 默认为http
go run main.go  -addr 127.0.0.1 -p 8080 -run grpc -config_path your_file_path.yml
```

首次启动服务时，会通过[gorm](http://gorm.io/docs/migration.html)的Auto-Migrate功能创建核心的业务数据表：
```sql
lists | CREATE TABLE `lists` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `business_type` varchar(255) NOT NULL,
  `business_desc` varchar(255) DEFAULT NULL,
  `interval` bigint(20) DEFAULT '10000',
  `started_at` bigint(20) DEFAULT '1',
  `ended_at` bigint(20) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `enable` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `business_type` (`business_type`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8
```

`ended_at`为bigint类型，可以保证ID充足。

如果想关闭访问日志，可以加上`-close_log`选项来关闭，提升服务性能：
```go
go run main.go  -addr 127.0.0.1 -p 8080 -close_log -config_path your_file_path.yml
```
配置调度器数量，默认为逻辑CPU数量的一半：
```go
go run main.go  -addr 127.0.0.1 -p 8080 -max_procs 10 -config_path your_file_path.yml
```
## 5. 配置业务类型（只能通过HTTP）

### **创建业务类型**：
```bash
curl -XPOST http://127.0.0.1:8080/lists -d '{
    "access_token": "this-is-a-secret-token-for-setup-data",
    "business_type": "test3",
    "business_desc": "这是日志数据的ID配置"
    "interval": 10000,
    "start_at":2235435678
}'
```
上面的请求会在数据库中插入一条数据。
- `business_type`为配置的业务类型的唯一标识，有唯一索引保证唯一。
- `interval`用来配置buffer的大小
- `start_at`配置ID的起点

### **查看业务类型**：
```bash
curl http://127.0.0.1:8080/lists
```
### **修改业务类型**：
修改前需要先关闭该业务类型，避免造成脏数据。
#### *关闭该业务类型**：
```bash
curl -XPUT http://127.0.0.1:8080/lists -d '{
    "access_token": "this-is-a-secret-token-for-setup-data",
    "business_type": "test3",
    "enable": "0"
}'
```
#### *更新业务类型**:
```bash
curl -XPUT http://127.0.0.1:8080/lists -d '{
    "access_token": "this-is-a-secret-token-for-setup-data",
    "business_type": "test3",
    "business_desc": "日志数据的ID配置2",
    "start_at": "2235435678",
    "enable": "1",
    "interval": 5000
}'
```
### 创建ID
#### HTTP
```bash
# GET /unique_ids/:business_type
curl http://127.0.0.1:8080/unique_ids/test3
```
结果：
```json
{
    "status": 1,
    "data": {
        "business": "test",
        "id": 1322895678
    }
}
```
#### gRPC
- gRPC的[.proto文件](https://github.com/EarlyZhao/id_generator/blob/master/grpc/id_rpc/unique_id.proto)
- gRPC获取ID请见[gRPC示例](https://github.com/EarlyZhao/id_generator/blob/master/rpc_client_example.go)
# 部署
当只部署一个应用时，程序能提供绝对递增的数据。如果追求高并发(数十万)，需要部署多个应用，但此时只能保证ID是趋势递增的。无论出于性能要求还是高可用，都建议部署多个应用。

![示例架构图](https://github.com/EarlyZhao/id_generator/blob/master/images/id_schema.jpg?raw=true)

当业务类型过多时，可以通过business_type对数据库分片。

由于Golang本身对fork的支持不够好，程序是以前台进程的方式启动，在实际使用时可以通过[supervisor](http://supervisord.org/)来管理。或者使用nohup：
```shell
go build -o id_generator main.go
nohup  ./id_generator -config_path ~/id_generator.yml >id_generator.log &
```

