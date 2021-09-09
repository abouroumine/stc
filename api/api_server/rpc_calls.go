package api_server

import (
	pb "abouroumine.com/stc/api/api_proto"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

const (
	AuthAddress = "localhost:50050"
	CCAddress   = "localhost:50051"
	SSAddress   = "localhost:50052"
	DBAddress   = "localhost:50053"
	AuthCert    = "./cert/auth_server.crt"
	DbCert      = "./cert/db_server.crt"
	CCCert      = "./cert/cc_server.crt"
	SSCert      = "./cert/ss_server.crt"
	HOSTNAME    = "localhost"
)

func (s *Server) CheckLogIn(userInfo *pb.UserAuth) (*pb.JWTResponse, error) {
	creds, err := credentials.NewClientTLSFromFile(AuthCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(AuthAddress, opts...)
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
	creds, err := credentials.NewClientTLSFromFile(CCCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(CCAddress, opts...)
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

func (s *Server) GetAllStations(role *string) (*pb.Stations, error) {
	creds, err := credentials.NewClientTLSFromFile(CCCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(CCAddress, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewCCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.AllStations(ctx, &wrappers.StringValue{Value: *role})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) error {
	creds, err := credentials.NewClientTLSFromFile(CCCert, HOSTNAME)
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(CCAddress, opts...)
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
	creds, err := credentials.NewClientTLSFromFile(CCCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(CCAddress, opts...)
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

func (s *Server) LandingRequest(value *wrappers.Int32Value) (*pb.Command, error) {
	creds, err := credentials.NewClientTLSFromFile(SSCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(SSAddress, opts...)
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

func (s *Server) TheLanding(value *wrappers.Int32Value) (*emptypb.Empty, error) {
	creds, err := credentials.NewClientTLSFromFile(SSCert, HOSTNAME)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(SSAddress, opts...)
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
