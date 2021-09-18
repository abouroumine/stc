package utils

import (
	pb "abouroumine.com/stc/ss_server/ss_proto"
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

func (s *Server) LandingRequest(in *pb.RequestDemand) (*pb.Command, error) {
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

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Here we start
	stationId, err := strconv.Atoi(in.IdStation)
	if err != nil {
		return nil, err
	}
	station, err := c.StationInfo(ctx, &wrappers.Int32Value{Value: int32(stationId)})
	if err != nil {
		return nil, err
	}

	shipId, err := strconv.Atoi(in.IdShip)
	if err != nil {
		return nil, err
	}
	ship, err := c.ShipInfo(ctx, &wrappers.Int32Value{Value: int32(shipId)})
	if err != nil {
		return nil, err
	}

	if station.Capacity < ship.Weight {
		return nil, nil
	}

	needNotToWait := true
	if station.Capacity-station.UsedCapacity >= ship.Weight {
		return s.ToDock(station.Docks, &needNotToWait)
	} else {
		needNotToWait = false
		return s.ToDock(station.Docks, &needNotToWait)
	}
}

func (s *Server) TheLanding(in *pb.RequestDemand) (*emptypb.Empty, error) {
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

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Here we start
	stationId, err := strconv.Atoi(in.IdStation)
	if err != nil {
		return nil, err
	}
	station, err := c.StationInfo(ctx, &wrappers.Int32Value{Value: int32(stationId)})
	if err != nil {
		return nil, err
	}

	shipId, err := strconv.Atoi(in.IdShip)
	if err != nil {
		return nil, err
	}
	ship, err := c.ShipInfo(ctx, &wrappers.Int32Value{Value: int32(shipId)})
	if err != nil {
		return nil, err
	}
	return s.ToLand(c, ctx, station, ship, int(in.Time))
}
