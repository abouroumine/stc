package utils

import (
	pb "abouroumine.com/stc/ss_server/ss_proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	pb.UnimplementedShippingStationServer
}

func (s *Server) RequestLand(ctx context.Context, in *wrapperspb.Int32Value) (*pb.Command, error) {
	return s.LandingRequest(in)
}

func (s *Server) Landing(ctx context.Context, in *wrapperspb.Int32Value) (*emptypb.Empty, error) {
	return s.TheLanding(in)
}
