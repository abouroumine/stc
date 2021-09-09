package utils

import (
	m "abouroumine.com/stc/db_service/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const (
	user         = "ayoubbouroumine"
	passwordAuth = ""
	passwordCC   = ""
	host         = "localhost:5432"
	dbNameAuth   = "stc_auth"
	dbNameCC     = "stc_cc"
)

func (s *Server) ConnectPostgresSQLToAuthDB() {
	s.DB = pg.Connect(&pg.Options{
		User:     user,
		Password: passwordAuth,
		Database: dbNameAuth,
		Addr:     host,
	})
	s.CreateUserAuthTables()
}

func (s *Server) ConnectPostgresSQLToCCDB() {
	s.DB = pg.Connect(&pg.Options{
		User:     user,
		Password: passwordCC,
		Database: dbNameCC,
		Addr:     host,
	})
	s.CreateCCTables()
}

func (s *Server) CreateTable(model interface{}) error {
	opts := &orm.CreateTableOptions{}
	err := s.DB.Model(model).CreateTable(opts)
	return err
}

func (s *Server) CreateUserAuthTables() {
	err := s.CreateTable(&m.Users{})
	if err != nil {
		return
	} else {
		newUser := new(m.Users)
		newUser.Username = "ayoub"
		newUser.Password = "ayoub1111"
		newUser.Role = "Command"
		_, _ = s.DB.Model(newUser).Insert()
	}
}

func (s *Server) CreateCCTables() {
	_ = s.CreateTable(&m.Command{})
	_ = s.CreateTable(&m.Dock{})
	_ = s.CreateTable(&m.Station{})
	_ = s.CreateTable(&m.Ship{})
}
