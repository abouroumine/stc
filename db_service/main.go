package main

import (
	pb "abouroumine.com/stc/db_service/db_proto"
	"abouroumine.com/stc/db_service/utils"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const (
	DBAddress = ":50053"
	DbCrt     = "./cert/db_server.crt"
	DbKey     = "./cert/db_server.key"
)

func main() {
	cert, err := tls.LoadX509KeyPair(DbCrt, DbKey)
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", DBAddress)
	if err != nil {
		return
	}
	pb.RegisterAuthenticationInfoServer(s, &utils.Server{})
	pb.RegisterCCServiceServer(s, &utils.Server{})
	pb.RegisterShippingStationServer(s, &utils.Server{})
	if err := s.Serve(lis); err != nil {
		return
	}
}
