package utils

import (
	pb "abouroumine.com/stc/auth_service/auth_proto"
	"context"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"strconv"
	"time"
)

const (
	DbCrt = "./cert/db_server.crt"
)

type Server struct {
	pb.UnimplementedAuthenticationInfoServer
}

var jwtKey = []byte("Signing-Key")

type CredClaim struct {
	Username string `json:"username"`
	Userid   string `json:"userid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GetToken(user *pb.UserAuth) (*string, *time.Time, error) {
	expire := time.Now().Add(15 * time.Minute)
	claims := &CredClaim{
		Username: user.Username,
		Userid:   user.Userid,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, nil, err
	}
	return &tokenString, &expire, nil
}

func (s *Server) Login(ctx context.Context, in *pb.UserAuth) (*pb.JWTResponse, error) {
	user, err := s.CheckLogIn(in)
	if err != nil || user == nil {
		return nil, err
	}

	// We Add JWT Here
	tokenString, exp, err := GetToken(user)
	if err != nil {
		return nil, err
	}

	theJWT := pb.JWTResponse{
		Token: *tokenString,
		Exp:   strconv.FormatInt(exp.Unix(), 10),
	}

	return &theJWT, err
}

func (s *Server) CheckLogIn(userInfo *pb.UserAuth) (*pb.UserAuth, error) {
	creds, err := credentials.NewClientTLSFromFile(DbCrt, os.Getenv("DB_SERVICE_HOSTNAME"))
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

	user, err := c.CheckInfoDB(ctx, userInfo)
	if err != nil || user == nil {
		return nil, err
	}

	return user, nil
}
