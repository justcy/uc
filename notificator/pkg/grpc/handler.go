package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/justcy/uc/notificator/pkg/endpoint"
	pb "github.com/justcy/uc/notificator/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeSendSmsHandler creates the handler logic
func makeSendSmsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendSmsEndpoint, decodeSendSmsRequest, encodeSendSmsResponse, options...)
}

// decodeSendSmsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendSms request.
// TODO implement the decoder
func decodeSendSmsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeSendSmsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendSmsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
func (g *grpcServer) SendSms(ctx context1.Context, req *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	_, rep, err := g.sendSms.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendSmsReply), nil
}

// makeCheckSmsHandler creates the handler logic
func makeCheckSmsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CheckSmsEndpoint, decodeCheckSmsRequest, encodeCheckSmsResponse, options...)
}

// decodeCheckSmsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CheckSms request.
// TODO implement the decoder
func decodeCheckSmsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeCheckSmsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCheckSmsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
func (g *grpcServer) CheckSms(ctx context1.Context, req *pb.CheckSmsRequest) (*pb.CheckSmsReply, error) {
	_, rep, err := g.checkSms.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CheckSmsReply), nil
}

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
// TODO implement the decoder
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
func (g *grpcServer) SendEmail(ctx context1.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	_, rep, err := g.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailReply), nil
}
