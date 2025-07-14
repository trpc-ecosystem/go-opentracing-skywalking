English| [中文](README.zh_CN.md) 

# OpenTracing SkyWalking Plugin

[![Go Reference](https://pkg.go.dev/badge/github.com/trpc-ecosystem/go-opentracing-skywalking.svg)](https://pkg.go.dev/github.com/trpc-ecosystem/go-opentracing-skywalking)
[![Go Report Card](https://goreportcard.com/badge/trpc.group/trpc-go/trpc-opentracing-skywalking)](https://goreportcard.com/report/trpc.group/trpc-go/trpc-opentracing-skywalking)
[![LICENSE](https://img.shields.io/badge/license-Apache--2.0-green.svg)](https://github.com/trpc-ecosystem/go-opentracing-skywalking/blob/main/LICENSE)
[![Releases](https://img.shields.io/github/release/trpc-ecosystem/go-opentracing-skywalking.svg?style=flat-square)](https://github.com/trpc-ecosystem/go-opentracing-skywalking/releases)
[![Tests](https://github.com/trpc-ecosystem/go-opentracing-skywalking/actions/workflows/prc.yml/badge.svg)](https://github.com/trpc-ecosystem/go-opentracing-skywalking/actions/workflows/prc.yml)
[![Coverage](https://codecov.io/gh/trpc-ecosystem/go-opentracing-skywalking/branch/main/graph/badge.svg)](https://app.codecov.io/gh/trpc-ecosystem/go-opentracing-skywalking/tree/main)

## Install Skywalking

1. Install docker and docker-compose
2. Run docker-compose.yml

```shell
docker-compose up -d
```

## Import Package

```go
    _ "trpc.group/trpc-go/trpc-opentracing-skywalking"
```

## Configure Plugin

```yaml
plugins:
  tracing:
    skywalking:                                     # Skywalking configuration
      server: echo.hello                            # Service name
      service: trpc.weiling.test.Hello              # Service instance name, recommended to be the same as the service name, but not the same here for testing purposes
      address: localhost:11800                      # Skywalking server address ip:port
      check_interval: 10s                           # Health check interval, default 20s
      max_send_queue_size:  100                     # Maximum send queue size, default 30000
      props:                                        # Meta properties
        app_id: test
      component_id: 23                              # Component ID, mainly for the icon, 23 represents the grpc service
      auth: 1260.4526611xxx                         # Authentication token, can be left blank if not available
      sampler : 1                                   # Sampling rate, default to all samples, sampler[0-1] float type
```

For component_id configuration, click [here](https://github.com/apache/skywalking//blob/master/apm-protocol/apm-network/src/main/java/org/apache/skywalking/apm/network/trace/component/ComponentsDefine.java)

## Configure Server

```yaml
server:  # Server configuration
  app: echo  # Business application name
  server: hello  # Process service name
  bin_path: /usr/local/trpc/bin/  # Path to the binary executable and framework configuration files
  conf_path: /usr/local/trpc/conf/  # Path to the business configuration files
  data_path: /usr/local/trpc/data/  # Path to the business data files
  service:  # Business services provided, can have multiple
    - name: trpc.weiling.test.Hello  # Service routing name
      ip: 127.0.0.1  # Service listening IP address, can use placeholder ${ip}, IP and NIC are optional, IP has priority
      nic: eth0  # Network card address for service listening, not needed if IP is configured
      port: 8021  # Service listening port, can use placeholder ${port}
      network: tcp  # Network listening type, tcp or udp
      protocol: trpc  # Application layer protocol, trpc or http
      timeout: 1000  # Maximum request processing time, in milliseconds
      idletime: 3000  # Connection idle time, in milliseconds
      filter:
        - skywalking
```

## Configure Client

```yaml
client:  # Client backend configuration
  timeout: 100000  # Maximum request processing time for all backends
  namespace: Development  # Environment for all backends
  service:  # Configuration for individual backend
    - callee: trpc.weiling.test.Hello  # Backend service protocol file's service name, if callee
```

You can also pass this `client.Option`, `client.WithFilter(filter.GetClient("skywalking"))`, when using the trpc client proxy, so you don't need to configure the client in trpc_go.yaml.

## View Skywalking Distributed Tracing Data Locally

http://localhost:8080

## Copyright

The copyright notice pertaining to the Tencent code in this repo was previously in the name of “THL A29 Limited.”  That entity has now been de-registered.  You should treat all previously distributed copies of the code as if the copyright notice was in the name of “Tencent.”
