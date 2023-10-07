# tRPC OpenTracing SkyWalking 插件
## 1.Skywalking 安装
  1. 安装 docker docker-compose
  2. 运行 docker-compose.yml 
    ```shell
        docker-compose up -d
    ```

## 2.在业务代码上导入包
```go
   _ "trpc.group/trpc-go/trpc-opentracing-skywalking"
```

## 3.插件配置
```yaml
plugins:
 tracing:
   skywalking:                                     # skywalking 配置
     server: echo.hello                            # 服务名
     service: trpc.weiling.test.Hello              # 服务实例名 建议与服务名一致 这边是测试所以不一致
     address: localhost:11800                      # skywalking 服务器地址 ip：port
     check_interval: 10s                           # 健康检查时间 默认 20s
     max_send_queue_size:  100                     # 最大发送队列 默认 30000
     props:                                        # 元属性
       app_id: test
     component_id: 23                              # 组件 ID 主要是图标 23 代表 grpc 服务
     auth: 1260.4526611xxx                         # 认证用的 token，如果没有可以不填
     sampler : 1                                   # 采样率，默认全部采样，sampler[0-1] float 类型

```
component_id 配置点击[这里](https://github.com/apache/skywalking//blob/master/apm-protocol/apm-network/src/main/java/org/apache/skywalking/apm/network/trace/component/ComponentsDefine.java)
## 4.客户端配置
```yaml
client:  # 客户端调用的后端配置
  timeout: 100000  # 针对所有后端的请求最长处理时间
  namespace: Development  # 针对所有后端的环境
  service:  # 针对单个后端的配置
    - callee: trpc.weiling.test.Hello  # 后端服务协议文件的 service name, 如何 callee 和下面的 name 一样，那只需要配置一个即可
      name: trpc.weiling.test.Hello  # 后端服务名字路由的 service name，有注册到名字服务的话，下面 target 可以不用配置
      target: ip://127.0.0.1:8021  # 后端服务地址
      network: tcp  # 后端服务的网络类型 tcp udp
      protocol: trpc  # 应用层协议 trpc http
      timeout: 1000  # 请求最长处理时间
      serialization: 0  # 序列化方式 0-pb 1-jce 2-json 3-flatbuffer，默认不要配置
      filter:
      - skywalking

```
这里也可以在使用 trpc client proxy 时传入这个 `client.Option,client.WithFilter(filter.GetClient("skywalking")), `
从而不在 trpc_go.yaml 里面进行 client 的配置


## 5.业务服务使用 skywalking
```yaml
server:  # 服务端配置
  app: echo  # 业务的应用名
  server: hello  # 进程服务名
  bin_path: /usr/local/trpc/bin/  # 二进制可执行文件和框架配置文件所在路径
  conf_path: /usr/local/trpc/conf/  # 业务配置文件所在路径
  data_path: /usr/local/trpc/data/  # 业务数据文件所在路径
  service:  # 业务服务提供的 service，可以有多个
    - name: trpc.weiling.test.Hello  # service 的路由名称
      ip: 127.0.0.1  # 服务监听 ip 地址 可使用占位符 ${ip},ip 和 nic 二选一，优先 ip
      nic: eth0  # 服务监听的网卡地址 有 ip 就不需要配置
      port: 8021  # 服务监听端口 可使用占位符 ${port}
      network: tcp  # 网络监听类型  tcp udp
      protocol: trpc  # 应用层协议 trpc http
      timeout: 1000  # 请求最长处理时间 单位 毫秒
      idletime: 3000  # 连接空闲时间 单位 毫秒
      filter:
        - skywalking
```

## 6.本地查看 skywalking 分布式追踪数据
http://localhost:8080