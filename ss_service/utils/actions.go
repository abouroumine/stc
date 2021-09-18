package utils

import (
	pb "abouroumine.com/stc/ss_server/ss_proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"sort"
	"time"
)

func (s *Server) ToDock(docks []*pb.Dock, isNotWaiting *bool) (*pb.Command, error) {
	times := make([]int, 0)
	for _, v := range docks {
		if v.NumDockingPorts > v.Occupied && *isNotWaiting {
			// finalize by a return
			return &pb.Command{Command: "land", DockingStation: v.Id}, nil
		}
		for _, v2 := range v.Ships {
			times = append(times, int(v2.StartTime+v2.Time)-time.Now().Second())
		}
	}
	// All Ports are Occupied Can not Dock Now we Need to Wait.
	sort.Ints(times)
	if len(times) <= 0 {
		return nil, nil
	}
	timeToWait := times[0]
	return &pb.Command{Command: "wait", Duration: int32(timeToWait)}, nil
}

func (s *Server) ToLand(c pb.ShippingStationClient, ctx context.Context, station *pb.Station, ship *pb.Ship, t int) (*emptypb.Empty, error) {
	docks := station.Docks
	for _, v := range docks {
		if v.NumDockingPorts > v.Occupied {
			func() {
				_, _ = c.UpdateTheLandData(ctx, &pb.UpdateLandData{
					IdShip:    ship.Id,
					IdDock:    v.Id,
					IdStation: station.Id,
					Weight:    ship.Weight,
					Time:      int32(t),
				})
			}()
			return &emptypb.Empty{}, nil
		}
	}
	return nil, nil
}
