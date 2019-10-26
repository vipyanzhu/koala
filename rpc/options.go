package rpc

import (
	"time"
)

type RpcOptions struct {
	ConnTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	ServiceName  string
	//注册中心名字
	RegisterName string
	//注册中心地址
	RegisterAddr string
	//注册中心路径
	RegisterPath string
}

type RpcOptionFunc func(opts *RpcOptions)

func WithConnTimeout(timeout time.Duration) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.ConnTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.WriteTimeout = timeout
	}
}

func WithReadTimeout(timeout time.Duration) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.ReadTimeout = timeout
	}
}

func WithServiceName(serviceName string) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.ServiceName = serviceName
	}
}

func WithRegisterName(name string) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.RegisterName = name
	}
}

func WithRegisterAddr(addr string) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.RegisterAddr = addr
	}
}

func WithRegisterPath(path string) RpcOptionFunc {
	return func(opts *RpcOptions) {
		opts.RegisterPath = path
	}
}