package models

type Dock struct {
	tableName       struct{} `pg:"dock"`
	Id              int      `pg:"id,pk"`
	NumDockingPorts int
	Occupied        int
	Weight          float32
	StationId       int
}
