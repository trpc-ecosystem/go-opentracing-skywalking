//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 THL A29 Limited, a Tencent company.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

package skywalking

// config skywalking 配置信息
type config struct {
	Server           string            `yaml:"server"`              // 服务名 默认为服务名
	Service          string            `yaml:"service"`             // 实例名 默认为服务名
	Address          string            `yaml:"address"`             // skywalking 服务器名称 ip：port
	CheckInterval    string            `yaml:"check_interval"`      // conn 健康检查时间间隔
	MaxSendQueueSize int               `yaml:"max_send_queue_size"` // 可发送消息队列
	Auth             string            `yaml:"auth"`                // skywalking 鉴权信息
	InstanceProps    map[string]string `yaml:"props"`               // 元属性
	ComponentID      int32             `yaml:"component_id"`        // 组件 id
	Sampler          float64           `yaml:"sampler"`             // 抽样率
}

// newConfig
func newConfig() *config {
	return &config{}
}
