package utils

import (
	pb "abouroumine.com/stc/ss_server/ss_proto"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

const (
	DBAddress = "localhost:50053"
	DbCert    = "./cert/db_server.crt"
	HOSTNAME  = "localhost"
)

func (s *Server) LandingRequest(in *wrappers.Int32Value) (*pb.Command, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(DBAddress, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStation, err := c.RequestLand(ctx, in)
	if err != nil {
		return nil, err
	}
	return newStation, nil
}

func (s *Server) TheLanding(in *wrappers.Int32Value) (*emptypb.Empty, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(DBAddress, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStation, err := c.Landing(ctx, in)
	if err != nil {
		return nil, err
	}
	return newStation, nil
}
