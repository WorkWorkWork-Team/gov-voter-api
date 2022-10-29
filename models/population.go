package model

import "time"

type UserInfo struct {
	CitizenID   int       `db:"CitizenID"`
	LazerId     string    `db:"LazerID"`
	Name        string    `db:"Name"`
	Lastname    string    `db:"Lastname"`
	Birthday    time.Time `db:"BirthDay"`
	Nationality string    `db:"Nationality"`
	DistricID   string    `db:"DistrictID"`
}
