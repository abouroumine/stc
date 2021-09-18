package utils

import (
	pb "abouroumine.com/stc/cc_server/cc_proto"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"strconv"
	"time"
)

const (
	DbCert = "./cert/db_server.crt"
)

func (s *Server) RegisterStation(in *pb.Station) (*pb.Station, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCert, os.Getenv("DB_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("DB_SERVICE_ADDR"), opts...)
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

func (s *Server) GetAllStations(in *pb.AllStationMsg) (*pb.Stations, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCert, os.Getenv("DB_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("DB_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if in.Role == string(COMMAND) {
		result, err := c.AllStationsNoCondition(ctx, &emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	id, err := strconv.Atoi(in.IdShip)
	if err != nil {
		return nil, err
	}
	ship, err := c.ShipCCInfo(ctx, &wrappers.Int32Value{Value: int32(id)})
	if err != nil {
		return nil, err
	}
	result, err := c.AllStationsWithCondition(ctx, &wrappers.FloatValue{Value: ship.Weight})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) (*emptypb.Empty, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCert, os.Getenv("DB_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("DB_SERVICE_ADDR"), opts...)
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
	creds, err := credentials.NewClientTLSFromFile(DbCert, os.Getenv("DB_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("DB_SERVICE_ADDR"), opts...)
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
