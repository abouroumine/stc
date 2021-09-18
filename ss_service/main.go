package main

import (
	pb "abouroumine.com/stc/ss_server/ss_proto"
	"abouroumine.com/stc/ss_server/utils"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const (
	SSAddress = ":50052"
	SSCert    = "./cert/ss_server.crt"
	SSKey     = "./cert/ss_server.key"
)

func main() {
	cert, err := tls.LoadX509KeyPair(SSCert, SSKey)
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", SSAddress)
	if err != nil {
		return
	}
	pb.RegisterShippingStationServer(s, &utils.Server{})
	if err := s.Serve(lis); err != nil {
		return
	}
}
