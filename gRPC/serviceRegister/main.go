package serviceRegister

import (
	"context"
	"google.golang.org/grpc/grpclog"
	"reflect"
)

/*
 服务注册
*/

//Server服务端
type Server struct {
	m map[string]*service //根据服务名称获取服务信息
}

type service struct {
	server interface{}
	md     map[string]*MethodDesc
	sd     map[string]*StreamDesc
	mdata  interface{}
}

type MethodDesc struct {
	MethodName string
	Handler    methodHandler
}

type methodHandler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor UnaryServerInterceptor) (interface{}, error)

type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

type UnaryServerInfo struct {
	// Server is the service implementation the user provides. This is read-only.
	Server interface{}
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod string
}

// UnaryServerInfo consists of various information about a unary RPC on
// server side. All per-rpc information may be mutated by the interceptor.
type ServiceDesc struct {
	ServiceName string
	//HandlerType: 用于检查接口是否被实现
	HandlerType interface{}
	Methods     []MethodDesc
	Streams     []StreamDesc
	Metadata    interface{}
}

// UnaryHandler defines the handler invoked by UnaryServerInterceptor to complete the normal
// execution of a unary RPC. If a UnaryHandler returns an error, it should be produced by the
// status package, or else gRPC will use codes.Unknown as the status code and err.Error() as
// the status message of the RPC.
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)

// StreamDesc represents a streaming RPC service's method specification.
type StreamDesc struct {
	StreamName string
	Handler    StreamHandler

	// At least one of these is true.
	ServerStreams bool
	ClientStreams bool
}

// StreamHandler defines the handler called by gRPC server to complete the
// execution of a streaming RPC. If a StreamHandler returns an error, it
// should be produced by the status package, or else gRPC will use
// codes.Unknown as the status code and err.Error() as the status message
// of the RPC.
type StreamHandler func(srv interface{}, stream ServerStream) error

// ServerStream defines the server-side behavior of a streaming RPC.
//
// All errors returned from ServerStream methods are compatible with the
// status package.
type ServerStream interface {
	// SetHeader sets the header metadata. It may be called multiple times.
	// When call multiple times, all the provided metadata will be merged.
	// All the metadata will be sent out when one of the following happens:
	//  - ServerStream.SendHeader() is called;
	//  - The first response is sent out;
	//  - An RPC status is sent out (error or success).
	SetHeader(MD) error
	// SendHeader sends the header metadata.
	// The provided md and headers set by SetHeader() will be sent.
	// It fails if called multiple times.
	SendHeader(MD) error
	// SetTrailer sets the trailer metadata which will be sent with the RPC status.
	// When called more than once, all the provided metadata will be merged.
	SetTrailer(MD)
	// Context returns the context for this stream.
	Context() context.Context
	// SendMsg sends a message. On error, SendMsg aborts the stream and the
	// error is returned directly.
	//
	// SendMsg blocks until:
	//   - There is sufficient flow control to schedule m with the transport, or
	//   - The stream is done, or
	//   - The stream breaks.
	//
	// SendMsg does not wait until the message is received by the client. An
	// untimely stream closure may result in lost messages.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not safe
	// to call SendMsg on the same stream in different goroutines.
	SendMsg(m interface{}) error
	// RecvMsg blocks until it receives a message into m or the stream is
	// done. It returns io.EOF when the client has performed a CloseSend. On
	// any non-EOF error, the stream is aborted and the error contains the
	// RPC status.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not
	// safe to call RecvMsg on the same stream in different goroutines.
	RecvMsg(m interface{}) error
}

// MD is a mapping from metadata keys to values. Users should use the following
// two convenience functions New and Pairs to generate MD.
type MD map[string][]string

/*
	1 参数说明
		sd: 服务描述
        ss：服务实例
    2 功能说明
		判断 实例是否实现了服务接口，如果实现了就调用注册方法进行注册，否则退出当前进程
*/
func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}) {
	ht := reflect.TypeOf(sd.HandlerType).Elem()
	st := reflect.TypeOf(ss)
	if !st.Implements(ht) {
		grpclog.Fatalf("grpc: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht)
	}

}

func (s *Server) register(sd *ServiceDesc, ss interface{}) {
	//检验服务是否已经注册过
	if _, ok := s.m[sd.ServiceName]; ok {
		grpclog.Fatalf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
	}

	//创建服务实例
	srv := &service{
		server: ss,
		md:     make(map[string]*MethodDesc),
		sd:     make(map[string]*StreamDesc),
		mdata:  sd.Metadata,
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		srv.md[d.MethodName] = d
	}
	for i := range sd.Streams {
		d := &sd.Streams[i]
		srv.sd[d.StreamName] = d
	}
	s.m[sd.ServiceName] = srv
}
