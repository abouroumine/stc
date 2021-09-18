package utils

import (
	pb "abouroumine.com/stc/db_service/db_proto"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	DB *pg.DB
	pb.UnimplementedAuthenticationInfoServer
	pb.UnimplementedCCServiceServer
	pb.UnimplementedShippingStationServer
}

func (s *Server) CheckInfoDB(ctx context.Context, in *pb.UserAuth) (*pb.UserAuth, error) {
	return s.CheckLoginInfo(in)
}

func (s *Server) SignUp(ctx context.Context, in *pb.UserAuth) (*wrapperspb.BoolValue, error) {
	return s.AddNewUser(in)
}

func (s *Server) StationRegister(ctx context.Context, in *pb.Station) (*pb.Station, error) {
	return s.RegisterStation(in)
}

func (s *Server) ShipCCInfo(ctx context.Context, in *wrappers.Int32Value) (*pb.Ship, error) {
	return s.GetShipInfo(in)
}

func (s *Server) AllStationsNoCondition(ctx context.Context, in *emptypb.Empty) (*pb.Stations, error) {
	return s.GetAllStationsForCommand()
}

func (s *Server) AllStationsWithCondition(ctx context.Context, in *wrappers.FloatValue) (*pb.Stations, error) {
	return s.GetAllStationsForShip(in)
}

func (s *Server) ShipRegister(ctx context.Context, in *wrappers.FloatValue) (*emptypb.Empty, error) {
	return s.RegisterShip(in)
}

func (s *Server) AllShips(ctx context.Context, in *emptypb.Empty) (*pb.Ships, error) {
	return s.GetAllShips(in)
}

func (s *Server) ShipInfo(ctx context.Context, in *wrappers.Int32Value) (*pb.Ship, error) {
	return s.GetShipInfo(in)
}

func (s *Server) StationInfo(ctx context.Context, in *wrappers.Int32Value) (*pb.Station, error) {
	return s.GetStationInfo(in)
}

func (s *Server) UpdateTheLandData(ctx context.Context, in *pb.UpdateLandData) (*emptypb.Empty, error) {
	return s.TheLandUpdate(in)
}
