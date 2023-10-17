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

import (
	"context"
	"time"

	"github.com/SkyAPM/go2sky"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
)

// SkyFilter skywalking 过滤器实现
type SkyFilter struct {
	tracer *go2sky.Tracer
	config *config
}

// NewSkyWalking 实例化过滤器实现
func NewSkyWalking(tracer *go2sky.Tracer, config *config) *SkyFilter {
	return &SkyFilter{tracer: tracer, config: config}
}

// InterceptClient is the client filter for tracing RPC.
func (sky SkyFilter) InterceptClient(ctx context.Context, req, rsp interface{}, f filter.ClientHandleFunc) error {
	msg := codec.Message(ctx)
	span, err := sky.tracer.CreateExitSpan(ctx, msg.ClientRPCName(), sky.config.Server, func(key, value string) error {
		md := msg.ClientMetaData()
		if md == nil {
			md = codec.MetaData{}
		}
		md[key] = []byte(value)
		msg.WithClientMetaData(md)
		return nil
	})
	span.SetComponent(sky.config.ComponentID)
	span.SetSpanLayer(agentv3.SpanLayer_RPCFramework)

	if err != nil {
		log.Errorf("span init error, trace information: %v", err)
	}
	err = f(ctx, req, rsp)
	if err != nil {
		span.Error(time.Now(), err.Error())
	}
	span.End()
	return err
}

// InterceptServer is the server filter for tracing  RPC.
func (sky SkyFilter) InterceptServer(
	ctx context.Context,
	req interface{},
	f filter.ServerHandleFunc,
) (interface{}, error) {
	msg := codec.Message(ctx)
	span, ctx, err := sky.tracer.CreateEntrySpan(ctx, msg.ServerRPCName(), func(key string) (string, error) {
		md := msg.ServerMetaData()
		parent := md[key]
		ctx = msg.Context()
		return string(parent), nil
	})
	span.SetComponent(sky.config.ComponentID)
	span.SetSpanLayer(agentv3.SpanLayer_RPCFramework)
	if err != nil {
		log.Errorf("span init error, trace information: %v", err)
	}
	rsp, err := f(ctx, req)
	if err != nil {
		span.Error(time.Now(), err.Error())
	}
	span.End()
	return rsp, err
}

// ClientFilter 客户端 RPC 调用分布式追踪过滤器
// Deprecated: Use InterceptClient instead.
func (sky SkyFilter) ClientFilter() filter.ClientFilter {
	return sky.InterceptClient
}

// ServerFilter 服务端 RPC 调用分布式追踪过滤器
// Deprecated: Use InterceptServer instead.
func (sky SkyFilter) ServerFilter() filter.ClientFilter {
	return func(ctx context.Context, req, rsp interface{}, f filter.ClientHandleFunc) error {
		_, err := sky.InterceptServer(ctx, req, func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, f(ctx, req, rsp)
		})
		return err
	}
}
