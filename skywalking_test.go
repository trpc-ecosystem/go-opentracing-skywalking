//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 Tencent.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

package skywalking

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/plugin"
)

const configInfo = `
plugins:
 tracing:
   skywalking:                                     # skywalking 配置
     server: echo.hello                            # 服务名
     service: trpc.weiling.test.Hello              # 服务实例名 建议与服务名一致 这边是测试所以不一致
     address: localhost:11800                      # skywalking 服务器地址 ip：port
     check_interval: 10s                           # 健康检查时间 默认 20s
     max_send_queue_size:  100                     # 最大发送队列 默认 30000
     auth: tester								   # 作者
     props:                                        # 元属性
       app_id: test
     component_id: 0                              # 组件 ID 主要是图标 23 代表 grpc 服务
`

// TestSkyWalkingPlugin_Setup 测试 skyWalkingPlugin Setup 方法。
func TestSkyWalkingPlugin_Setup(t *testing.T) {
	p := &skyWalkingPlugin{}
	cfg := trpc.Config{}
	_ = yaml.Unmarshal([]byte(configInfo), &cfg)
	log.Debugf("cfg:%+v", cfg)
	conf := cfg.Plugins["tracing"]["skywalking"]

	Convey("Setup 测试用例 1: ", t, func() {
		err := p.Setup("test", &plugin.YamlNodeDecoder{Node: &conf})
		So(err, ShouldEqual, nil)
	})
}
