package models

type Station struct {
	tableName    struct{} `pg:"station,alias:station"`
	Id           int      `pg:"id,pk"`
	Capacity     float32
	UserCapacity float32
	Docks        []*Dock `pg:"rel:has-many"`
	IsRegistered bool
}
