package models

type Users struct {
	tableName struct{} `pg:"users,alias:users"`
	Id        int      `pg:"id,pk"`
	Username  string   `pg:",unique"`
	Password  string
	Role      string
}
