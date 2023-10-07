// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.

// Package skywalking implements skywalking plugin.
package skywalking

import (
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/plugin"
)

const (
	// pluginName 插件名 skywalking
	pluginName = "skywalking"
	// pluginType 插件类别 tracing
	pluginType         = "tracing"
	defaultComponentID = 23
)

var globalTracer *go2sky.Tracer

// GlobalTracer 全局的 tracer
func GlobalTracer() *go2sky.Tracer {
	return globalTracer
}

func init() {
	plugin.Register(pluginName, &skyWalkingPlugin{})
}

// skyWalkingPlugin trpc 插件实现
type skyWalkingPlugin struct{}

// Type skyWalkingPlugin trpc 插件名字
func (p *skyWalkingPlugin) Type() string {
	return pluginType
}

// Setup skyWalkingPlugin 初始化
func (p *skyWalkingPlugin) Setup(name string, decoder plugin.Decoder) error {
	config := newConfig()
	err := decoder.Decode(config)
	if err != nil {
		return err
	}
	opts, err := adaptReport(config)
	if err != nil {
		return err
	}
	report, err := reporter.NewGRPCReporter(config.Address, opts...)

	if err != nil {
		return err
	}
	tracerOption := []go2sky.TracerOption{go2sky.WithReporter(report), go2sky.WithInstance(config.Service)}
	if config.Sampler != 0 {
		tracerOption = append(tracerOption, go2sky.WithSampler(config.Sampler))
	}
	tracer, err := go2sky.NewTracer(config.Server, tracerOption...)
	if err != nil {
		return err
	}
	// 将 tracer 暴露出来，便于在其他地方调用 tracer.CreateLocalSpan 等 tracer 方法进行上报 span
	// 新版本 go2sky 可以使用 go2sky.SetGlobalTracer(tracer) https://github.com/SkyAPM/go2sky#global-tracer
	globalTracer = tracer
	sky := NewSkyWalking(tracer, config)
	filter.Register(pluginName, sky.InterceptServer, sky.InterceptClient)
	return nil
}

// adaptReport 自定义配置转换为 skyWalking 配置
func adaptReport(config *config) (opts []reporter.GRPCReporterOption, err error) {
	result := make([]reporter.GRPCReporterOption, 0, 4)
	if config.CheckInterval != "" {
		t, err := time.ParseDuration(config.CheckInterval)
		if err != nil {
			return result, err
		}
		result = append(result, reporter.WithCheckInterval(t))
	}

	result = append(result, reporter.WithInstanceProps(config.InstanceProps))
	if config.MaxSendQueueSize != 0 {
		result = append(result, reporter.WithMaxSendQueueSize(config.MaxSendQueueSize))
	}
	if config.Auth != "" {
		result = append(result, reporter.WithAuthentication(config.Auth))
	}
	if config.ComponentID == 0 {
		config.ComponentID = defaultComponentID
	}
	return result, nil
}
