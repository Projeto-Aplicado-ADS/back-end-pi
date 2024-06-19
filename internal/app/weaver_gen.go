// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "projeto-api/internal/app/Component",
		Iface: reflect.TypeOf((*Component)(nil)).Elem(),
		Impl:  reflect.TypeOf(implapp{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return component_local_stub{impl: impl.(Component), tracer: tracer, allUsersMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "AllUsers", Remote: false, Generated: true}), createUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "CreateUser", Remote: false, Generated: true}), getOneUserByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "GetOneUserById", Remote: false, Generated: true}), getUserByEmailAndPasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "GetUserByEmailAndPassword", Remote: false, Generated: true}), listUserByEmailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "ListUserByEmail", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return component_client_stub{stub: stub, allUsersMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "AllUsers", Remote: true, Generated: true}), createUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "CreateUser", Remote: true, Generated: true}), getOneUserByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "GetOneUserById", Remote: true, Generated: true}), getUserByEmailAndPasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "GetUserByEmailAndPassword", Remote: true, Generated: true}), listUserByEmailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "projeto-api/internal/app/Component", Method: "ListUserByEmail", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return component_server_stub{impl: impl.(Component), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return component_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Component] = (*implapp)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*implapp)(nil)

// Local stub implementations.

type component_local_stub struct {
	impl                             Component
	tracer                           trace.Tracer
	allUsersMetrics                  *codegen.MethodMetrics
	createUserMetrics                *codegen.MethodMetrics
	getOneUserByIdMetrics            *codegen.MethodMetrics
	getUserByEmailAndPasswordMetrics *codegen.MethodMetrics
	listUserByEmailMetrics           *codegen.MethodMetrics
}

// Check that component_local_stub implements the Component interface.
var _ Component = (*component_local_stub)(nil)

func (s component_local_stub) AllUsers(ctx context.Context) (r0 AllUsersOut, err error) {
	// Update metrics.
	begin := s.allUsersMetrics.Begin()
	defer func() { s.allUsersMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.Component.AllUsers", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AllUsers(ctx)
}

func (s component_local_stub) CreateUser(ctx context.Context, a0 UserIn) (r0 bool, err error) {
	// Update metrics.
	begin := s.createUserMetrics.Begin()
	defer func() { s.createUserMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.Component.CreateUser", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateUser(ctx, a0)
}

func (s component_local_stub) GetOneUserById(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	// Update metrics.
	begin := s.getOneUserByIdMetrics.Begin()
	defer func() { s.getOneUserByIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.Component.GetOneUserById", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetOneUserById(ctx, a0)
}

func (s component_local_stub) GetUserByEmailAndPassword(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	begin := s.getUserByEmailAndPasswordMetrics.Begin()
	defer func() { s.getUserByEmailAndPasswordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.Component.GetUserByEmailAndPassword", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetUserByEmailAndPassword(ctx, a0, a1)
}

func (s component_local_stub) ListUserByEmail(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	// Update metrics.
	begin := s.listUserByEmailMetrics.Begin()
	defer func() { s.listUserByEmailMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "app.Component.ListUserByEmail", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ListUserByEmail(ctx, a0)
}

// Client stub implementations.

type component_client_stub struct {
	stub                             codegen.Stub
	allUsersMetrics                  *codegen.MethodMetrics
	createUserMetrics                *codegen.MethodMetrics
	getOneUserByIdMetrics            *codegen.MethodMetrics
	getUserByEmailAndPasswordMetrics *codegen.MethodMetrics
	listUserByEmailMetrics           *codegen.MethodMetrics
}

// Check that component_client_stub implements the Component interface.
var _ Component = (*component_client_stub)(nil)

func (s component_client_stub) AllUsers(ctx context.Context) (r0 AllUsersOut, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.allUsersMetrics.Begin()
	defer func() { s.allUsersMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.Component.AllUsers", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	*(*[]UsersOut)(&r0) = serviceweaver_dec_slice_UsersOut_50a126a0(dec)
	err = dec.Error()
	return
}

func (s component_client_stub) CreateUser(ctx context.Context, a0 UserIn) (r0 bool, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.createUserMetrics.Begin()
	defer func() { s.createUserMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.Component.CreateUser", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

func (s component_client_stub) GetOneUserById(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getOneUserByIdMetrics.Begin()
	defer func() { s.getOneUserByIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.Component.GetOneUserById", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s component_client_stub) GetUserByEmailAndPassword(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getUserByEmailAndPasswordMetrics.Begin()
	defer func() { s.getUserByEmailAndPasswordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.Component.GetUserByEmailAndPassword", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s component_client_stub) ListUserByEmail(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.listUserByEmailMetrics.Begin()
	defer func() { s.listUserByEmailMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "app.Component.ListUserByEmail", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 4, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.2 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type component_server_stub struct {
	impl    Component
	addLoad func(key uint64, load float64)
}

// Check that component_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*component_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s component_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "AllUsers":
		return s.allUsers
	case "CreateUser":
		return s.createUser
	case "GetOneUserById":
		return s.getOneUserById
	case "GetUserByEmailAndPassword":
		return s.getUserByEmailAndPassword
	case "ListUserByEmail":
		return s.listUserByEmail
	default:
		return nil
	}
}

func (s component_server_stub) allUsers(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.AllUsers(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_UsersOut_50a126a0(enc, ([]UsersOut)(r0))
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s component_server_stub) createUser(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 UserIn
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.CreateUser(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s component_server_stub) getOneUserById(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetOneUserById(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s component_server_stub) getUserByEmailAndPassword(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.GetUserByEmailAndPassword(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s component_server_stub) listUserByEmail(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.ListUserByEmail(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type component_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that component_reflect_stub implements the Component interface.
var _ Component = (*component_reflect_stub)(nil)

func (s component_reflect_stub) AllUsers(ctx context.Context) (r0 AllUsersOut, err error) {
	err = s.caller("AllUsers", ctx, []any{}, []any{&r0})
	return
}

func (s component_reflect_stub) CreateUser(ctx context.Context, a0 UserIn) (r0 bool, err error) {
	err = s.caller("CreateUser", ctx, []any{a0}, []any{&r0})
	return
}

func (s component_reflect_stub) GetOneUserById(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	err = s.caller("GetOneUserById", ctx, []any{a0}, []any{&r0})
	return
}

func (s component_reflect_stub) GetUserByEmailAndPassword(ctx context.Context, a0 string, a1 string) (err error) {
	err = s.caller("GetUserByEmailAndPassword", ctx, []any{a0, a1}, []any{})
	return
}

func (s component_reflect_stub) ListUserByEmail(ctx context.Context, a0 string) (r0 UsersOut, err error) {
	err = s.caller("ListUserByEmail", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*UserIn)(nil)

type __is_UserIn[T ~struct {
	weaver.AutoMarshal
	FullName string    "json:\"full_name\""
	Email    string    "json:\"email\""
	Phone    string    "json:\"phone,omitempty\""
	Password string    "json:\"password\""
	Birthday time.Time "json:\"birthday,omitempty\""
}] struct{}

var _ __is_UserIn[UserIn]

func (x *UserIn) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("UserIn.WeaverMarshal: nil receiver"))
	}
	enc.String(x.FullName)
	enc.String(x.Email)
	enc.String(x.Phone)
	enc.String(x.Password)
	enc.EncodeBinaryMarshaler(&x.Birthday)
}

func (x *UserIn) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("UserIn.WeaverUnmarshal: nil receiver"))
	}
	x.FullName = dec.String()
	x.Email = dec.String()
	x.Phone = dec.String()
	x.Password = dec.String()
	dec.DecodeBinaryUnmarshaler(&x.Birthday)
}

var _ codegen.AutoMarshal = (*UsersOut)(nil)

type __is_UsersOut[T ~struct {
	weaver.AutoMarshal
	ID        string    "json:\"id\""
	FullName  string    "json:\"full_name\""
	Email     string    "json:\"email\""
	Phone     string    "json:\"phone,omitempty\""
	IsAdmin   bool      "json:\"is_admin,omitempty\""
	Birthday  time.Time "json:\"birthday,omitempty\""
	CreatedAt time.Time "json:\"created_at,omitempty\""
}] struct{}

var _ __is_UsersOut[UsersOut]

func (x *UsersOut) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("UsersOut.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ID)
	enc.String(x.FullName)
	enc.String(x.Email)
	enc.String(x.Phone)
	enc.Bool(x.IsAdmin)
	enc.EncodeBinaryMarshaler(&x.Birthday)
	enc.EncodeBinaryMarshaler(&x.CreatedAt)
}

func (x *UsersOut) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("UsersOut.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.String()
	x.FullName = dec.String()
	x.Email = dec.String()
	x.Phone = dec.String()
	x.IsAdmin = dec.Bool()
	dec.DecodeBinaryUnmarshaler(&x.Birthday)
	dec.DecodeBinaryUnmarshaler(&x.CreatedAt)
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_UsersOut_50a126a0(enc *codegen.Encoder, arg []UsersOut) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_UsersOut_50a126a0(dec *codegen.Decoder) []UsersOut {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]UsersOut, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
