package main

import (
	pb "abouroumine.com/stc/cc_server/cc_proto"
	"abouroumine.com/stc/cc_server/utils"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const (
	CCAddress = "localhost:50051"
	CCCert    = "./cert/cc_server.crt"
	CCKey     = "./cert/cc_server.key"
)

func main() {
	cert, err := tls.LoadX509KeyPair(CCCert, CCKey)
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", CCAddress)
	if err != nil {
		return
	}
	pb.RegisterCCServiceServer(s, &utils.Server{})
	if err := s.Serve(lis); err != nil {
		return
	}
}
