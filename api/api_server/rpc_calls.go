package api_server

import (
	pb "abouroumine.com/stc/api/api_proto"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"time"
)

const (
	AuthCert = "./cert/auth_server.crt"
	DbCert   = "./cert/db_server.crt"
	CCCert   = "./cert/cc_server.crt"
	SSCert   = "./cert/ss_server.crt"
)

func (s *Server) CheckLogIn(userInfo *pb.UserAuth) (*pb.JWTResponse, error) {
	creds, err := credentials.NewClientTLSFromFile(AuthCert, os.Getenv("AUTH_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("AUTH_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewAuthenticationInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	token, err := c.Login(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Server) AddNewUser(user *pb.UserAuth) (*bool, error) {
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

	c := pb.NewAuthenticationInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	isAdded, err := c.SignUp(ctx, user)
	if err != nil {
		return nil, err
	}
	return &isAdded.Value, nil
}

func (s *Server) RegisterStation(station *pb.Station) (*pb.Station, error) {
	creds, err := credentials.NewClientTLSFromFile(CCCert, os.Getenv("CC_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("CC_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStation, err := c.StationRegister(ctx, station)
	if err != nil {
		return nil, err
	}
	return newStation, nil
}

func (s *Server) GetAllStations(role, shipId *string) (*pb.Stations, error) {
	creds, err := credentials.NewClientTLSFromFile(CCCert, os.Getenv("CC_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("CC_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.AllStations(ctx, &pb.AllStationMsg{IdShip: *shipId, Role: *role})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) error {
	creds, err := credentials.NewClientTLSFromFile(CCCert, os.Getenv("CC_SERVICE_HOSTNAME"))
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("CC_SERVICE_ADDR"), opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = c.ShipRegister(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetAllShips() (*pb.Ships, error) {
	creds, err := credentials.NewClientTLSFromFile(CCCert, os.Getenv("CC_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("CC_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.AllShips(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) LandingRequest(value *pb.RequestDemand) (*pb.Command, error) {
	creds, err := credentials.NewClientTLSFromFile(SSCert, os.Getenv("SS_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(os.Getenv("SS_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.RequestLand(ctx, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) TheLanding(value *pb.RequestDemand) (*emptypb.Empty, error) {
	creds, err := credentials.NewClientTLSFromFile(SSCert, os.Getenv("SS_SERVICE_HOSTNAME"))
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(os.Getenv("SS_SERVICE_ADDR"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewShippingStationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.Landing(ctx, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}
