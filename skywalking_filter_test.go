// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.

package skywalking

import (
	context2 "context"
	"testing"

	"trpc.group/trpc-go/trpc-go/codec"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
)

// TestSkyFilter_ClientFilter 测试 SkyFilter ClientFilter.
func TestSkyFilter_ClientFilter(t *testing.T) {
	r, _ := reporter.NewLogReporter()
	tra, _ := go2sky.NewTracer("test_server", go2sky.WithReporter(r))
	sky := SkyFilter{
		tracer: tra,
		config: &config{
			Server: "test_server",
		},
	}
	Convey("ClientFilter 测试用例", t, func() {
		ctx := context2.Background()
		ctx, newMsg := codec.WithNewMessage(ctx)
		newMsg.WithClientRPCName("test")
		ctx = context2.WithValue(ctx, codec.ContextKeyMessage, newMsg)
		ft := sky.ClientFilter()

		err := ft(ctx, "", "", func(ctx context2.Context, req interface{},
			rsp interface{}) (err error) {
			return nil
		})
		So(err, ShouldEqual, nil)
	})
}

// TestSkyFilter_ServerFilter 测试 SkyFilter ServerFilter 方法。
func TestSkyFilter_ServerFilter(t *testing.T) {
	r, _ := reporter.NewLogReporter()
	tra, _ := go2sky.NewTracer("test_server", go2sky.WithReporter(r))
	sky := SkyFilter{
		tracer: tra,
		config: &config{
			Server: "test_server",
		},
	}
	Convey("ServerFilter 测试用例", t, func() {
		ctx := context2.Background()
		ctx, newMsg := codec.WithNewMessage(ctx)
		newMsg.WithServerRPCName("test")
		ctx = context2.WithValue(ctx, codec.ContextKeyMessage, newMsg)
		ft := sky.ServerFilter()

		err := ft(ctx, "", "", func(ctx context2.Context, req interface{},
			rsp interface{}) (err error) {
			return nil
		})
		So(err, ShouldEqual, nil)
	})
}

// TestSkyFilter_ClientFilter 测试 SkyFilter ClientFilter.
func TestSkyFilter_ClientInterceptor(t *testing.T) {
	r, _ := reporter.NewLogReporter()
	tra, _ := go2sky.NewTracer("test_server", go2sky.WithReporter(r))
	sky := SkyFilter{
		tracer: tra,
		config: &config{
			Server: "test_server",
		},
	}

	t.Run("InterceptClient Ok", func(t *testing.T) {
		ctx := context2.Background()
		ctx, newMsg := codec.WithNewMessage(ctx)
		newMsg.WithClientRPCName("test")
		ctx = context2.WithValue(ctx, codec.ContextKeyMessage, newMsg)
		err := sky.InterceptClient(ctx, "", "", func(ctx context2.Context, req interface{},
			rsp interface{}) (err error) {
			return nil
		})
		require.Nil(t, err)
	})
}

// TestSkyFilter_ServerFilter 测试 SkyFilter ServerFilter 方法。
func TestSkyFilter_ServerInterceptor(t *testing.T) {
	r, _ := reporter.NewLogReporter()
	tra, _ := go2sky.NewTracer("test_server", go2sky.WithReporter(r))
	sky := SkyFilter{
		tracer: tra,
		config: &config{
			Server: "test_server",
		},
	}
	t.Run("ServerFilter Ok", func(t *testing.T) {
		ctx := context2.Background()
		ctx, newMsg := codec.WithNewMessage(ctx)
		newMsg.WithServerRPCName("test")
		ctx = context2.WithValue(ctx, codec.ContextKeyMessage, newMsg)
		_, err := sky.InterceptServer(ctx, "", func(ctx context2.Context, req interface{}) (rsp interface{}, err error) {
			return rsp, nil
		})
		require.Nil(t, err)
	})
}
