package api_server

import (
	pb "abouroumine.com/stc/api/api_proto"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	s *http.ServeMux
}

type CredClaim struct {
	Username string `json:"username"`
	Userid   string `json:"userid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (s *Server) Initialize() {
	s.s = http.NewServeMux()

	// Authentication Service
	s.s.HandleFunc("/user/signup", s.VerifyMethod(s.SignUP, "POST"))
	s.s.HandleFunc("/auth/login", s.VerifyMethod(s.LogIn, "POST"))

	// Command Central Service
	s.s.HandleFunc("/centcom/station/register", s.VerifyMethod(s.StationRegister, "POST"))
	s.s.HandleFunc("/centcom/station/all", s.VerifyMethod(s.AllStations, "GET"))
	s.s.HandleFunc("/centcom/ship/register", s.VerifyMethod(s.ShipRegister, "POST"))
	s.s.HandleFunc("/centcom/ship/all", s.VerifyMethod(s.AllShips, "GET"))

	// Shipping Station Service
	s.s.HandleFunc("/shipping-station/request-landing", s.VerifyMethod(s.RequestLanding, "POST"))
	s.s.HandleFunc("/shipping-station/land", s.VerifyMethod(s.Landing, "POST"))

	http.ListenAndServe(":8080", s.s)
}

var jwtKey = []byte("Signing-Key")

func VerifyJWT(r *http.Request) (*string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}
	token := c.Value
	claims := &CredClaim{}

	t, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, nil
	}
	return &claims.Role, nil
}

func (s *Server) VerifyMethod(next http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next(w, r)
		} else {
			PrepareResponse(w, http.StatusBadRequest, "")
			return
		}
	}
}

func (s *Server) SignUP(w http.ResponseWriter, r *http.Request) {
	var user pb.UserAuth
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		PrepareResponse(w, http.StatusUnauthorized, "")
		return
	}
	switch user.Role {
	case string(SHIP):
		_, err = s.AddNewUser(&user)
		if err != nil {
			PrepareResponse(w, http.StatusUnauthorized, "")
			return
		}
		PrepareResponse(w, http.StatusOK, "ok")
		return
	case string(STATION), string(COMMAND):
		role, err := VerifyJWT(r)
		if err != nil {
			PrepareResponse(w, http.StatusUnauthorized, "")
			return
		}
		if role == nil {
			PrepareResponse(w, http.StatusUnauthorized, "")
			return
		}
		if *role != string(COMMAND) {
			PrepareResponse(w, http.StatusUnauthorized, "")
			return
		}
		_, err = s.AddNewUser(&user)
		if err != nil {
			PrepareResponse(w, http.StatusUnauthorized, "")
			return
		}
		PrepareResponse(w, http.StatusOK, "ok")
		return
	default:
		PrepareResponse(w, http.StatusUnauthorized, "")
		return
	}
}

func (s *Server) LogIn(w http.ResponseWriter, r *http.Request) {
	var user pb.UserAuth
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		PrepareResponse(w, http.StatusUnauthorized, "")
		return
	}
	token, err := s.CheckLogIn(&user)
	if err != nil || token == nil {
		PrepareResponse(w, http.StatusUnauthorized, "")
		return
	}
	t, err := strconv.Atoi(token.GetExp())
	if err != nil {
		PrepareResponse(w, http.StatusUnauthorized, "")
		return
	}
	tm := time.Unix(int64(t), 0)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.GetToken(),
		Expires: tm,
	})
	PrepareResponse(w, http.StatusOK, "Welcome")
}

func (s *Server) StationRegister(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || *role != string(SHIP) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var station pb.Station
	err = json.NewDecoder(r.Body).Decode(&station)
	if err != nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	createdStation, err := s.RegisterStation(&station)
	if err != nil || createdStation == nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, *createdStation)
}

func (s *Server) AllStations(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || (*role != string(SHIP) && *role != string(COMMAND)) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var shipId struct{ ShipId string }
	err = json.NewDecoder(r.Body).Decode(&shipId)
	if err != nil {
		theNil := ""
		result, er := s.GetAllStations(role, &theNil)
		if er != nil || result == nil {
			PrepareResponse(w, http.StatusBadRequest, "Bad Request")
			return
		}
		PrepareResponse(w, http.StatusOK, result.Stations)
		return
	}
	result, err := s.GetAllStations(role, &shipId.ShipId)
	if err != nil || result == nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, result.Stations)
}

func (s *Server) ShipRegister(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || *role != string(SHIP) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var weight struct{ Weight float32 }
	err = json.NewDecoder(r.Body).Decode(&weight)
	if err != nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	err = s.RegisterShip(&wrappers.FloatValue{Value: weight.Weight})
	if err != nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, "Success!")
}

func (s *Server) AllShips(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || *role != string(COMMAND) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := s.GetAllShips()
	if err != nil || result == nil {
		fmt.Println(err)
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, result.Ships)
}

func (s *Server) RequestLanding(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || *role != string(SHIP) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var info pb.RequestDemand
	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := s.LandingRequest(&info)
	if err != nil || result == nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, *result)
}

func (s *Server) Landing(w http.ResponseWriter, r *http.Request) {
	role, err := VerifyJWT(r)
	if err != nil || role == nil || *role != string(SHIP) {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var info pb.RequestDemand
	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := s.TheLanding(&info)
	if err != nil || result == nil {
		PrepareResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}
	PrepareResponse(w, http.StatusOK, *result)
}
