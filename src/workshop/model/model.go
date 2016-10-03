package model

type SquadBu struct {
	Bu_name    string `json:"bu_name" bson:"bu_name"`
	Squad_name string `json:"squad_name" bson:"squad_name"`
}

type DevSquad struct {
	Squad_name string `json:"squad_name" bson:"squad_name"`
	Dev_name   string `json:"dev_name" bson:"dev_name"`
}

type SquadDeactive struct {
	Squad_name string `json:"squad_name" bson:"squad_name"`
}

type BuDeactive struct {
	Bu_name string `json:"bu_name" bson:"bu_name"`
}