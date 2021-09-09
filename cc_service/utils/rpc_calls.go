package utils

import (
	pb "abouroumine.com/stc/cc_server/cc_proto"
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

func (s *Server) RegisterStation(in *pb.Station) (*pb.Station, error) {
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

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStation, err := c.StationRegister(ctx, in)
	if err != nil {
		return nil, err
	}
	return newStation, nil
}

func (s *Server) GetAllStations(role *wrappers.StringValue) (*pb.Stations, error) {
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

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.AllStations(ctx, role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) (*emptypb.Empty, error) {
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

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStation, err := c.ShipRegister(ctx, in)
	if err != nil {
		return nil, err
	}
	return newStation, nil
}

func (s *Server) GetAllShips(in *emptypb.Empty) (*pb.Ships, error) {
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

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.AllShips(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}
