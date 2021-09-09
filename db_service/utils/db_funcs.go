package utils

import (
	pb "abouroumine.com/stc/db_service/db_proto"
	m "abouroumine.com/stc/db_service/models"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *Server) CheckLoginInfo(user *pb.UserAuth) (*pb.UserAuth, error) {
	// Select all users.
	us := new(m.Users)
	s.ConnectPostgresSQLToAuthDB()
	defer s.DB.Close()
	err := s.DB.Model(us).
		Where("username = ?", user.GetUsername()).
		Where("password = ?", user.GetPassword()).
		Limit(1).Select()
	if err != nil {
		return nil, err
	}
	newUser := pb.UserAuth{
		Username: us.Username,
		Password: us.Password,
		Role:     us.Role,
		Userid:   strconv.Itoa(us.Id),
	}
	return &newUser, nil
}

func (s *Server) AddNewUser(user *pb.UserAuth) (*wrappers.BoolValue, error) {
	us := new(m.Users)
	us.Username = user.Username
	us.Role = user.Role
	us.Password = user.Password
	s.ConnectPostgresSQLToAuthDB()
	defer s.DB.Close()
	_, err := s.DB.Model(us).Insert()
	if err != nil {
		return nil, err
	}
	return &wrappers.BoolValue{Value: true}, nil
}

func (s *Server) RegisterStation(station *pb.Station) (*pb.Station, error) {
	sta := new(m.Station)
	sta.Capacity = station.Capacity

	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	_, err := s.DB.Model(sta).Insert()
	if err != nil {
		return nil, err
	}
	var docks []*m.Dock
	for _, v := range station.Docks {
		l := m.Dock{NumDockingPorts: int(v.NumDockingPorts), StationId: sta.Id}
		docks = append(docks, &l)
	}
	_, err = s.DB.Model(&docks).Insert()
	if err != nil {
		return nil, err
	}

	var docksResult []*pb.Dock
	for _, v := range docks {
		l := pb.Dock{Id: strconv.Itoa(v.Id), NumDockingPorts: int32(v.NumDockingPorts)}
		docksResult = append(docksResult, &l)
	}

	result := pb.Station{
		Id:    strconv.Itoa(sta.Id),
		Docks: docksResult,
	}
	return &result, nil
}

func (s *Server) GetAllStationsForCommand() (*pb.Stations, error) {
	var stations []m.Station
	var resultStations []*pb.Station
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	err := s.DB.Model(&stations).Select()
	if err != nil {
		return nil, err
	}
	for _, v1 := range stations {
		var docks []m.Dock
		var pbDocks []*pb.Dock
		err = s.DB.Model(&docks).Where("station_id = ?", v1.Id).Select()
		if err != nil {
			return nil, err
		}
		for _, v2 := range docks {
			newDock := pb.Dock{
				Id:              strconv.Itoa(v2.Id),
				NumDockingPorts: int32(v2.NumDockingPorts),
				Occupied:        int32(v2.Occupied),
				Weight:          v2.Weight,
			}
			pbDocks = append(pbDocks, &newDock)
		}
		newStation := pb.Station{
			Id:           strconv.Itoa(v1.Id),
			Capacity:     v1.Capacity,
			UsedCapacity: v1.UserCapacity,
			Docks:        pbDocks,
		}
		resultStations = append(resultStations, &newStation)
	}
	var result pb.Stations
	result.Stations = resultStations
	return &result, nil
}

func (s *Server) GetAllStationsForShip(in *string) (*pb.Stations, error) {

	return nil, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) (*emptypb.Empty, error) {
	ship := new(m.Ship)
	ship.Weight = in.Value
	ship.Status = DOCKED
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	_, err := s.DB.Model(ship).Insert()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) GetAllShips(in *emptypb.Empty) (*pb.Ships, error) {
	var ships []m.Ship
	var resultShips []*pb.Ship
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	err := s.DB.Model(&ships).Select()
	if err != nil {
		return nil, err
	}
	for _, v2 := range ships {
		newShip := pb.Ship{
			Id:     strconv.Itoa(v2.Id),
			Status: v2.Status,
			Weight: v2.Weight,
		}
		resultShips = append(resultShips, &newShip)
	}
	return &pb.Ships{Ships: resultShips}, nil
}

func (s *Server) LandingRequest(in *wrappers.Int32Value) (*pb.Command, error) {
	// time := in.Value
	//var stations []m.Station
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	//err := s.DB.Model(&stations).Select()
	/*if err != nil {
		return nil, err
	}*/
	command := &pb.Command{
		Command:        "land",
		DockingStation: "1",
	}
	return command, nil
}

func (s *Server) TheLanding(in *wrappers.Int32Value) (*emptypb.Empty, error) {
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	// _, err := s.DB.Model(ship).Insert()
	//if err != nil {
	//	return nil, err
	//}
	return &emptypb.Empty{}, nil
}
