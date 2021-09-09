package models

type Ship struct {
	tableName struct{} `pg:"ship,alias:ship"`
	Id        int      `pg:"id,pk"`
	Status    string
	Weight    float32
	Time      int32
}
