package utils

type Role string

const (
	SHIP     Role = "Ship"
	STATION  Role = "Station"
	COMMAND  Role = "Command"
	DOCKED        = "docked"
	INFLIGHT      = "in-flight"
)
