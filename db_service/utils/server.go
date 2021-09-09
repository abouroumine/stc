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

func (s *Server) AllStations(ctx context.Context, in *wrappers.StringValue) (*pb.Stations, error) {
	role := in.Value
	if role == string(COMMAND) {
		return s.GetAllStationsForCommand()
	} else {
		return s.GetAllStationsForShip(&role)
	}
}

func (s *Server) ShipRegister(ctx context.Context, in *wrappers.FloatValue) (*emptypb.Empty, error) {
	return s.RegisterShip(in)
}

func (s *Server) AllShips(ctx context.Context, in *emptypb.Empty) (*pb.Ships, error) {
	return s.GetAllShips(in)
}

func (s *Server) RequestLand(ctx context.Context, in *wrapperspb.Int32Value) (*pb.Command, error) {
	return s.LandingRequest(in)
}

func (s *Server) Landing(ctx context.Context, in *wrapperspb.Int32Value) (*emptypb.Empty, error) {
	return s.TheLanding(in)
}
