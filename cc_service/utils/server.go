package utils

import (
	pb "abouroumine.com/stc/cc_server/cc_proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	pb.UnimplementedCCServiceServer
}

func (s *Server) StationRegister(ctx context.Context, in *pb.Station) (*pb.Station, error) {
	return s.RegisterStation(in)
}

func (s *Server) AllStations(ctx context.Context, in *pb.AllStationMsg) (*pb.Stations, error) {
	return s.GetAllStations(in)
}

func (s *Server) ShipRegister(ctx context.Context, in *wrapperspb.FloatValue) (*emptypb.Empty, error) {
	return s.RegisterShip(in)
}

func (s *Server) AllShips(ctx context.Context, in *emptypb.Empty) (*pb.Ships, error) {
	return s.GetAllShips(in)
}
