package models

type Command struct {
	tableName      struct{} `pg:"command,alias:command"`
	Id             int      `pg:"id,pk"`
	Command        string
	Duration       int32
	DockingStation string
}
