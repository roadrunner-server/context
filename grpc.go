package context

var (
	// GrpcTranscoderSpanKey is gRPC transcoder span context key
	GrpcTranscoderSpanKey = &ContextKey{name: "grpc_transcoder_span"} //nolint:gochecknoglobals

	// GrpcMethodDescKey is gRPC method descriptor context key
	GrpcMethodDescKey = &ContextKey{name: "grpc_method_desc"} //nolint:gochecknoglobals

	// GrpcRequestMessageKey is gGRPC request message context key
	GrpcRequestMessageKey = &ContextKey{name: "grpc_request_message"} //nolint:gochecknoglobals
)
