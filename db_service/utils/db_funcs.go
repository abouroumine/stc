package utils

import (
	pb "abouroumine.com/stc/db_service/db_proto"
	m "abouroumine.com/stc/db_service/models"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"sync"
	"time"
)

var mutex = sync.Mutex{}

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
	mutex.Lock()
	s.ConnectPostgresSQLToAuthDB()
	defer s.DB.Close()
	_, err := s.DB.Model(us).Insert()
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	mutex.Unlock()
	return &wrappers.BoolValue{Value: true}, nil
}

func (s *Server) RegisterStation(station *pb.Station) (*pb.Station, error) {
	sta := new(m.Station)
	sta.Capacity = station.Capacity
	sta.IsRegistered = true
	mutex.Lock()
	s.ConnectPostgresSQLToCCDB()
	tx, err := s.DB.Begin()
	defer tx.Close()
	_, err = tx.Model(sta).Insert()
	if err != nil {
		tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	var docks []*m.Dock
	for _, v := range station.Docks {
		l := m.Dock{NumDockingPorts: int(v.NumDockingPorts), StationId: sta.Id}
		docks = append(docks, &l)
	}
	_, err = s.DB.Model(&docks).Insert()
	if err != nil {
		tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	tx.Commit()
	mutex.Unlock()
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
			UsedCapacity: v1.UsedCapacity,
			Docks:        pbDocks,
		}
		resultStations = append(resultStations, &newStation)
	}
	var result pb.Stations
	result.Stations = resultStations
	return &result, nil
}

func (s *Server) GetAllStationsForShip(in *wrappers.FloatValue) (*pb.Stations, error) {
	var stations []m.Station
	var resultStations []*pb.Station
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	err := s.DB.Model(&stations).Select()
	if err != nil {
		return nil, err
	}
	for _, v1 := range stations {
		if (v1.Capacity - v1.UsedCapacity) >= in.Value {
			var docks []m.Dock
			var pbDocks []*pb.Dock
			err = s.DB.Model(&docks).Where("station_id = ?", v1.Id).Select() // .Where("num_docking_ports > occupied")
			if err != nil {
				return nil, err
			}
			for _, v2 := range docks {
				if v2.NumDockingPorts > v2.Occupied {
					newDock := pb.Dock{
						Id:              strconv.Itoa(v2.Id),
						NumDockingPorts: int32(v2.NumDockingPorts),
						Occupied:        int32(v2.Occupied),
						Weight:          v2.Weight,
					}
					pbDocks = append(pbDocks, &newDock)
				}
			}
			if len(docks) > 0 {
				newStation := pb.Station{
					Id:           strconv.Itoa(v1.Id),
					Capacity:     v1.Capacity,
					UsedCapacity: v1.UsedCapacity,
					Docks:        pbDocks,
				}
				resultStations = append(resultStations, &newStation)
			}
		}
	}
	var result pb.Stations
	result.Stations = resultStations
	return &result, nil
}

func (s *Server) RegisterShip(in *wrappers.FloatValue) (*emptypb.Empty, error) {
	ship := new(m.Ship)
	ship.Weight = in.Value
	ship.Status = INFLIGHT
	mutex.Lock()
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	_, err := s.DB.Model(ship).Insert()
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	mutex.Unlock()
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

func (s *Server) GetShipInfo(in *wrappers.Int32Value) (*pb.Ship, error) {
	ship := new(m.Ship)
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	err := s.DB.Model(ship).Where("id = ?", in.Value).Select()
	if err != nil {
		return nil, err
	}
	shipInfo := pb.Ship{
		Id:        strconv.Itoa(ship.Id),
		Status:    ship.Status,
		Weight:    ship.Weight,
		Time:      0,
		StartTime: 0,
	}
	return &shipInfo, nil
}

func (s *Server) GetStationInfo(in *wrappers.Int32Value) (*pb.Station, error) {
	station := new(m.Station)
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()

	err := s.DB.Model(station).Where("id = ?", in.Value).Select()
	if err != nil {
		return nil, err
	}
	var docks []m.Dock
	var pbDocks []*pb.Dock
	err = s.DB.Model(&docks).Where("station_id = ?", station.Id).Select()
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
	stationInfo := &pb.Station{
		Id:           strconv.Itoa(station.Id),
		Capacity:     station.Capacity,
		UsedCapacity: station.UsedCapacity,
		Docks:        pbDocks,
		IsRegistered: station.IsRegistered,
	}
	return stationInfo, nil
}

func (s *Server) TheLandUpdate(in *pb.UpdateLandData) (*emptypb.Empty, error) {
	s.ConnectPostgresSQLToCCDB()
	defer s.DB.Close()
	stationId, err := strconv.Atoi(in.IdStation)
	if err != nil {
		return nil, err
	}
	station := new(m.Station)
	err = s.DB.Model(station).Where("id = ?", stationId).Select()
	if err != nil {
		return nil, err
	}
	dock := new(m.Dock)
	dockId, err := strconv.Atoi(in.IdDock)
	if err != nil {
		return nil, err
	}
	err = s.DB.Model(dock).Where("id = ?", dockId).Select()
	if err != nil {
		return nil, err
	}
	ship := new(m.Ship)
	shipId, err := strconv.Atoi(in.IdShip)
	if err != nil {
		return nil, err
	}
	err = s.DB.Model(ship).Where("id = ?", shipId).Select()
	if err != nil {
		return nil, err
	}
	mutex.Lock()

	tx, err := s.DB.Begin()
	defer tx.Close()

	// Add Weight
	_, err = tx.Model(station).Set("used_capacity = ?", station.UsedCapacity+in.Weight).Where("id = ?id").Update()
	if err != nil {
		_ = tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	_, err = tx.Model(dock).Set("occupied = ?", dock.Occupied+1).Where("id = ?id").Update()
	if err != nil {
		_ = tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	_, err = tx.Model(dock).Set("weight = ?", dock.Weight+in.Weight).Where("id = ?id").Update()
	if err != nil {
		_ = tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	_, err = tx.Model(ship).Set("status = ?", DOCKED).Where("id = ?id").Update()
	if err != nil {
		_ = tx.Rollback()
		mutex.Unlock()
		return nil, err
	}
	_, err = tx.Model(ship).Set("dock_id = ?", in.IdDock).Where("id = ?id").Update()
	if err != nil {
		_ = tx.Rollback()
		mutex.Unlock()
		return nil, err
	}

	tx.Commit()
	mutex.Unlock()
	go func() {

		// Remove Weight
		time.Sleep(time.Second * time.Duration(in.Time))
		s.ConnectPostgresSQLToCCDB()
		defer s.DB.Close()
		tx1, err := s.DB.Begin()
		defer tx1.Close()
		stationId, err = strconv.Atoi(in.IdStation)
		if err != nil {
			_ = tx1.Rollback()
		}
		station := new(m.Station)
		err = tx1.Model(station).Where("id = ?", stationId).Select()
		if err != nil {
			_ = tx1.Rollback()
		}
		dock := new(m.Dock)
		dockId, err := strconv.Atoi(in.IdDock)
		if err != nil {
			_ = tx1.Rollback()
		}
		err = s.DB.Model(dock).Where("id = ?", dockId).Select()
		if err != nil {
			_ = tx1.Rollback()
		}
		ship := new(m.Ship)
		shipId, err := strconv.Atoi(in.IdShip)
		if err != nil {
			_ = tx1.Rollback()
		}
		err = s.DB.Model(ship).Where("id = ?", shipId).Select()
		if err != nil {
			_ = tx1.Rollback()
		}

		mutex.Lock()
		_, err = tx1.Model(station).Set("used_capacity = ?", station.UsedCapacity-in.Weight).Where("id = ?id").Update()
		if err != nil {
			_ = tx1.Rollback()
			mutex.Unlock()
		}
		_, err = tx1.Model(dock).Set("occupied = ?", dock.Occupied-1).Where("id = ?id").Update()
		if err != nil {
			_ = tx1.Rollback()
			mutex.Unlock()
		}
		_, err = tx1.Model(dock).Set("weight = ?", dock.Weight-in.Weight).Where("id = ?id").Update()
		if err != nil {
			_ = tx1.Rollback()
			mutex.Unlock()
		}
		_, err = tx1.Model(ship).Set("status = ?", INFLIGHT).Where("id = ?id").Update()
		if err != nil {
			_ = tx1.Rollback()
			mutex.Unlock()
		}
		_, err = tx1.Model(ship).Set("dock_id = ?", nil).Where("id = ?id").Update()
		if err != nil {
			_ = tx1.Rollback()
			mutex.Unlock()
		}
		tx1.Commit()
		mutex.Unlock()
	}()

	return &emptypb.Empty{}, nil
}
