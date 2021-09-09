package main

import (
	pb "abouroumine.com/stc/auth_service/auth_proto"
	"abouroumine.com/stc/auth_service/utils"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const (
	AuthPort = ":50050"
	AuthCrt  = "./cert/auth_server.crt"
	AuthKey  = "./cert/auth_server.key"
)

func main() {
	cert, err := tls.LoadX509KeyPair(AuthCrt, AuthKey)
	if err != nil {
		return
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	s := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", AuthPort)
	if err != nil {
		return
	}
	pb.RegisterAuthenticationInfoServer(s, &utils.Server{})
	if err := s.Serve(lis); err != nil {
		return
	}
}
